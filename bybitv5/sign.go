package bybitv5

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

func (o *Sign) Params(v url.Values) {
	ts, window := o.timestamp()
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

func (o *Sign) header(h http.Header, s string) {
	ts, window := o.timestamp()
	h.Set("X-BAPI-API-KEY", o.key)
	h.Set("X-BAPI-TIMESTAMP", ts)
	h.Set("X-BAPI-RECV-WINDOW", window)
	h.Set("X-BAPI-SIGN", signHmac(ts+o.key+window+s, o.secret))
}

func (o *Sign) timestamp() (ts, window string) {
	i := int(time.Now().UTC().UnixNano()/int64(time.Millisecond)) + o.timeShift
	ts = strconv.Itoa(i)
	window = strconv.Itoa(o.recvWindow)
	return
}

func signHmac(s, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
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
		s += k + "=" + src.Get(k) + "&"
	}
	s = s[0 : len(s)-1]
	return
}

func signParams(v url.Values, secret string) string {
	return signHmac(encodeSortParams(v), secret)
}
