package transport

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
)

type Client struct {
	log         *ulog.Log
	url         string
	key         string
	secret      string
	proxy       *url.URL
	logUri      bool
	logResponse bool
	onHttpError func(err error, attempt int) bool
}

func NewClient() *Client {
	return &Client{
		log: ulog.Empty(),
		url: MainBaseUrl,
	}
}

func (o *Client) WithUrl(url string) *Client {
	o.url = url
	return o
}

func (o *Client) WithByTickUrl() *Client {
	o.url = MainBaseByTickUrl
	return o
}

func (o *Client) WithAuth(key, secret string) *Client {
	o.key = key
	o.secret = secret
	return o
}

func (o *Client) WithProxy(proxy string) *Client {
	var err error
	o.proxy, err = ParseProxy(proxy)
	if err != nil {
		panic(fmt.Sprintf("set proxy fail: %v", err))
	}
	return o
}

func (o *Client) WithLog(log *ulog.Log) *Client {
	o.log = log
	return o
}

func (o *Client) WithLogUri(logUri bool) *Client {
	o.logUri = logUri
	return o
}

func (o *Client) WithLogResponse(logResponse bool) *Client {
	o.logResponse = logResponse
	return o
}

func (o *Client) WithOnHttpError(onHttpError func(err error, attempt int) bool) *Client {
	o.onHttpError = onHttpError
	return o
}

func (o *Client) RestProxy() {
	o.proxy = nil
}

func (o *Client) Key() string {
	return o.key
}

func (o *Client) Secret() string {
	return o.secret
}

func (o *Client) Request(method string, path string, param any, ret any, sign bool) (err error) {
	var httpError bool
	var attempt int
	for {
		err, httpError = o.request(method, path, param, ret, sign)
		if httpError && o.onHttpError != nil {
			if o.onHttpError(err, attempt) {
				continue
			}
			attempt++
		}
		break
	}
	return err
}

func (o *Client) RequestPublic(method string, path string, param any, ret any) error {
	return o.Request(method, path, param, ret, false)
}

func (o *Client) RequestPrivate(method string, path string, param any, ret any) error {
	return o.Request(method, path, param, ret, true)
}

func (o *Client) GetPublic(path string, param any, ret any) error {
	return o.RequestPublic(http.MethodGet, path, param, ret)
}

func (o *Client) Get(path string, param any, ret any) error {
	return o.RequestPrivate(http.MethodGet, path, param, ret)
}

func (o *Client) Post(path string, param any, ret any) error {
	return o.RequestPrivate(http.MethodPost, path, param, ret)
}

func (o *Client) Delete(path string, param any, ret any) error {
	return o.RequestPrivate(http.MethodDelete, path, param, ret)
}

func (o *Client) request(method string, path string, param any, ret any, sign bool) (err error, httpError bool) {
	logf := func(format string, a ...any) {
		m := fmt.Sprintf(format, a...)
		if err == nil {
			o.log.Infof("%s[%s]: %s", method, path, m)
		} else {
			o.log.Errorf("%s[%s]: %s", method, path, m)
		}
	}
	timestamp := time.Now()
	u, err := url.Parse(o.url)
	if err != nil {
		logf("url fail: %v", err)
		return
	}
	u.Path = path
	p := NewParam().From(param)
	vals := p.Make()
	var signHeader func(http.Header)
	if sign {
		if p.HeaderSign {
			signHeader = o.signQueryHeader(vals)
		} else {
			vals = o.signQuery(vals)
		}
	}
	var reqbody []byte
	if p.IsJson {
		m := make(map[string]any)
		for name, list := range vals {
			if len(list) > 0 {
				m[name] = list[0]
			}
		}
		reqbody, _ = json.Marshal(m)
	} else {
		u.RawQuery = vals.Encode()
		u.RawQuery = strings.Replace(u.RawQuery, "%2C", ",", -1)
	}
	if o.logUri {
		o.log.Debug("uri:", u.String())
	}
	req, err := http.NewRequest(method, u.String(), bytes.NewReader(reqbody))
	if err != nil {
		logf("init request fail: %v", err)
		return
	}
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Referer", "GinArea")
		req.Header.Set("x-referer", "GinArea")
	}
	if signHeader != nil {
		signHeader(req.Header)
	}
	client := &http.Client{}
	if o.proxy != nil {
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(o.proxy),
		}
	}
	resp, err := client.Do(req)
	elapsedTime := time.Since(timestamp).Truncate(time.Millisecond)
	if err != nil {
		logf("request fail [%s]: %v", elapsedTime.String(), err)
		httpError = true
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	m := fmt.Sprintf("%s %s", resp.Status, elapsedTime.String())
	if len(body) >= 0 {
		m = fmt.Sprintf("%s %s", m, ufmt.ByteSizeDense(len(body)))
		if o.logResponse {
			o.log.Debug("response:", string(body))
		}
	}
	ok := resp.StatusCode == http.StatusOK
	if ok {
		if len(body) == 0 {
			err = errors.New("response body is empty")
			logf("%v", err)
			return
		}
		timestamp := time.Now()
		body = []byte(strings.Replace(string(body), `"result":{}`, `"result":null`, 1))
		err = json.Unmarshal(body, ret)
		if err == nil {
			elapsedTime := time.Since(timestamp).Truncate(time.Millisecond)
			if elapsedTime > time.Millisecond {
				m = fmt.Sprintf("%s json:%s", m, elapsedTime.String())
			}
			s := reflect.ValueOf(ret)
			s = s.Elem()
			var e Error
			e.Code = int(s.FieldByName("RetCode").Int())
			e.Text = s.FieldByName("RetMsg").String()
			if !e.Empty() {
				err = &e
				m = fmt.Sprintf("%s %v", m, err)
			}
		} else {
			m = fmt.Sprintf("%s json unmarshal fail: %v", m, err)
		}
	} else {
		err = errors.New(fmt.Sprintf("http status-code: %d", resp.StatusCode))
		logf("%v", err)
		httpError = true
		return
	}
	logf("%s", m)
	return
}

func (o *Client) signQuery(src url.Values) url.Values {
	i := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	ts := strconv.Itoa(i)
	if src == nil {
		src = url.Values{}
	}
	src.Add("api_key", o.key)
	src.Add("timestamp", ts)
	src.Add("sign", makeSignature(src, o.secret))
	return src
}

func (o *Client) signQueryHeader(src url.Values) func(http.Header) {
	i := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	ts := strconv.Itoa(i)
	sign := makeSignature(src, o.secret)
	return func(h http.Header) {
		h.Set("X-BAPI-API-KEY", o.key)
		h.Set("X-BAPI-TIMESTAMP", ts)
		h.Set("X-BAPI-SIGN", sign)
	}
}

func makeSignature(src url.Values, key string) string {
	keys := make([]string, len(src))
	i := 0
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	s := ""
	for _, k := range keys {
		s += k + "=" + src.Get(k) + "&"
	}
	s = s[0 : len(s)-1]
	h := hmac.New(sha256.New, []byte(key))
	_, err := io.WriteString(h, s)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
