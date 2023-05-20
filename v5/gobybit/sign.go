package gobybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Sign struct {
	key        string
	secret     string
	timeShift  int
	recvWindow int
}

func NewSign(key, secret string) *Sign {
	o := new(Sign)
	o.key = key
	o.secret = secret
	o.timeShift = -10000
	o.recvWindow = 20000
	return o
}

func (o *Sign) Query(v url.Values) url.Values {
	ts, window := o.timestamp()
	if v == nil {
		v = make(url.Values)
	}
	v.Add("api_key", o.key)
	v.Add("timestamp", ts)
	v.Add("recv_window", window)
	v.Add("sign", sign(v, o.secret))
	return v
}

func (o *Sign) Header(v url.Values) func(http.Header) {
	ts, window := o.timestamp()
	return func(h http.Header) {
		h.Set("X-BAPI-API-KEY", o.key)
		h.Set("X-BAPI-TIMESTAMP", ts)
		h.Set("X-BAPI-RECV-WINDOW", window)
		h.Set("X-BAPI-SIGN", sign(v, o.secret))
	}
}

func (o *Sign) timestamp() (ts, window string) {
	i := int(time.Now().UTC().UnixNano()/int64(time.Millisecond)) + o.timeShift
	ts = strconv.Itoa(i)
	window = strconv.Itoa(o.recvWindow)
	return
}

func sign(src url.Values, key string) string {
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
