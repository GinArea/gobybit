package bybitv5

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

type Sign struct {
	key    string
	secret string
	ts     *Timestamp
}

func NewSign(key, secret string) *Sign {
	o := new(Sign)
	o.key = key
	o.secret = secret
	o.ts = NewTimestamp()
	return o
}

func (o *Sign) Params(v url.Values) {
	ts, window := o.ts.Get()
	if v == nil {
		v = make(url.Values)
	}
	v.Add("api_key", o.key)
	v.Add("timestamp", ts)
	v.Add("recv_window", window)
	v.Add("sign", signParams(v, o.secret))
}

func (o *Sign) HeaderGet(h http.Header, v url.Values) {
	o.header(h, encodeSortParams(v))
}

func (o *Sign) HeaderPost(h http.Header, body []byte) {
	o.header(h, string(body[:]))
}

func (o *Sign) WebSocket() []string {
	expires := o.ts.Expires()
	sign := signHmac(fmt.Sprintf("GET/realtime%d", expires), o.secret)
	return []string{
		o.key,
		strconv.Itoa(expires),
		sign,
	}
}

func (o *Sign) header(h http.Header, s string) {
	ts, window := o.ts.Get()
	h.Set("X-BAPI-API-KEY", o.key)
	h.Set("X-BAPI-TIMESTAMP", ts)
	h.Set("X-BAPI-RECV-WINDOW", window)
	h.Set("X-BAPI-SIGN", signHmac(ts+o.key+window+s, o.secret))
}

func signHmac(s, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

func encodeParam(name, value string) string {
	params := url.Values{}
	params.Add(name, value)
	return params.Encode()
}

func encodeSortParams(src url.Values) (s string) {
	if len(src) == 0 {
		return
	}
	keys := make([]string, len(src))
	i := 0
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		s += encodeParam(k, src.Get(k)) + "&"
	}
	s = s[0 : len(s)-1]
	return
}

func signParams(v url.Values, secret string) string {
	return signHmac(encodeSortParams(v), secret)
}
