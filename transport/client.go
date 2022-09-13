package transport

import (
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

	"github.com/msw-x/moon/ulog"
)

type Client struct {
	log    *ulog.Log
	url    string
	key    string
	secret string
	proxy  *url.URL
	logUri bool
}

func NewClient(url, key, secret string) *Client {
	c := Client{
		log:    ulog.New("client"),
		url:    url,
		key:    key,
		secret: secret,
	}
	c.log.Info("url:", c.url)
	return &c
}

func (this *Client) WithProxy(proxy string) *Client {
	var err error
	this.proxy, err = ParseProxy(proxy)
	if err != nil {
		panic(fmt.Sprintf("set proxy fail: %v", err))
	}
	return this
}

func (this *Client) WithLogUri(logUri bool) *Client {
	this.logUri = logUri
	return this
}

func (this *Client) Request(method string, path string, param any, ret any) (err error) {
	logf := func(format string, a ...any) {
		m := fmt.Sprintf(format, a...)
		if err == nil {
			this.log.Infof("%s[%s]: %s", method, path, m)
		} else {
			err = errors.New(m)
			this.log.Errorf("%s[%s]: %s", method, path, m)
		}
	}
	timestamp := time.Now()
	u, err := url.Parse(this.url)
	if err != nil {
		logf("url fail: %v", err)
		return
	}
	u.Path = path
	p := NewUrlParam().From(param).Make()
	p = this.signQuery(p)
	u.RawQuery = p.Encode()
	u.RawQuery = strings.Replace(u.RawQuery, "%2C", ",", -1)
	if this.logUri {
		ulog.Debug("uri:", u.String())
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		logf("init request fail: %v", err)
		return
	}
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Referer", "GinArea")
		req.Header.Set("x-referer", "GinArea")
	}
	client := &http.Client{}
	if this.proxy != nil {
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(this.proxy),
		}
	}
	resp, err := client.Do(req)
	elapsedTime := time.Since(timestamp).Truncate(time.Millisecond)
	if err != nil {
		logf("request fail [%s]: %v", elapsedTime.String(), err)
		return
	}
	defer resp.Body.Close()
	m := fmt.Sprintf("%s %s", resp.Status, elapsedTime.String())
	if resp.ContentLength >= 0 {
		//m = fmt.Sprintf("%s %s", m, ufmt.ByteSizeDense(resp.ContentLength))
	}
	ok := resp.StatusCode == http.StatusOK
	if ok {
		timestamp := time.Now()
		err = json.NewDecoder(resp.Body).Decode(ret)
		if err == nil {
			elapsedTime := time.Since(timestamp).Truncate(time.Millisecond)
			if elapsedTime > time.Millisecond {
				m = fmt.Sprintf("%s json:%s", m, elapsedTime.String())
			}
			s := reflect.ValueOf(ret)
			s = s.Elem()
			RetCode := s.FieldByName("RetCode").Int()
			RetMsg := s.FieldByName("RetMsg").String()
			if RetCode != 0 {
				err = errors.New(fmt.Sprintf("code[%d]: %s", RetCode, RetMsg))
				m = fmt.Sprintf("%s %v", m, err)
			}
		} else {
			m = fmt.Sprintf("%s json unmarshal fail: %v", m, err)
		}
	} else {
		err = errors.New(fmt.Sprintf("http status-code: %d", resp.StatusCode))
		logf("%v", err)
		return
	}
	logf("%s", m)
	return
}

func (this *Client) Get(path string, param any, ret any) error {
	return this.Request(http.MethodGet, path, param, ret)
}

func (this *Client) Post(path string, param any, ret any) error {
	return this.Request(http.MethodPost, path, param, ret)
}

func (this *Client) Delete(path string, param any, ret any) error {
	return this.Request(http.MethodDelete, path, param, ret)
}

func (this *Client) signQuery(src url.Values) url.Values {
	i := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	now := strconv.Itoa(i)
	if src == nil {
		src = url.Values{}
	}
	src.Add("api_key", this.key)
	src.Add("timestamp", now)
	src.Add("sign", makeSignature(src, this.secret))
	return src
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
