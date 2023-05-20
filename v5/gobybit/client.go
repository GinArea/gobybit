package gobybit

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ulog"
)

type Client struct {
	c *uhttp.Client
	s *Sign
}

func NewClient() *Client {
	o := new(Client)
	o.c = uhttp.NewClient()
	o.WithBaseUrl(transport.MainBaseUrl)
	o.WithPath("v5")
	return o
}

func (o *Client) Clone() *Client {
	r := new(Client)
	r.c = o.c.Clone()
	r.s = o.s
	return r
}

func (o *Client) WithBaseUrl(url string) *Client {
	o.c.WithBase(url)
	return o
}

func (o *Client) WithByTickUrl() *Client {
	return o.WithBaseUrl(transport.MainBaseByTickUrl)
}

func (o *Client) WithPath(path string) *Client {
	o.c.WithPath(path)
	return o
}

func (o *Client) WithAppendPath(path string) *Client {
	o.c.WithAppendPath(path)
	return o
}

func (o *Client) WithAuth(key, secret string) *Client {
	o.s = NewSign(key, secret)
	return o
}

func (o *Client) WithProxy(proxy string) *Client {
	o.c.WithProxy(proxy)
	return o
}

func (o *Client) WithTimeout(timeout time.Duration) *Client {
	o.c.WithTimeout(timeout)
	return o
}

func (o *Client) WithTraceFormat(log *ulog.Log, f uhttp.Format) *Client {
	o.c.WithTraceFormat(log, f)
	return o
}

func (o *Client) market() *Client {
	return o.Clone().WithAppendPath("market")
}

func (o *Client) asset() *Client {
	return o.Clone().WithAppendPath("asset")
}

func request[R, T any](c *Client, method string, path string, request any, transform func(R) (T, error), sign bool) (r Response[T]) {
	var perf *uhttp.Performer
	switch method {
	case http.MethodGet:
		perf = c.c.Get(path).Params(request)
	case http.MethodPost:
		perf = c.c.Post(path).Json(request)
	default:
		r.Error = fmt.Errorf("forbidden method: %s", method)
		return
	}
	if sign && c.s != nil {
		perf.Request.Header = make(http.Header)
		switch method {
		case http.MethodGet:
			c.s.HeaderGet(perf.Request.Header, perf.Request.Params)
		case http.MethodPost:
			c.s.HeaderPost(perf.Request.Header, perf.Request.Body)
		}
	}
	h := perf.Do()
	if h.Error == nil {
		if h.BodyExists() {
			raw := new(response[R])
			h.Json(raw)
			r.Time = raw.Time
			r.Error = raw.Error()
			if r.Ok() {
				r.Data, r.Error = transform(raw.Result)
			}
		}
		if sign {
			r.SetErrorIfNil(h.HeaderTo(&r.Limit))
		}
	} else {
		r.Error = h.Error
	}
	return
}

func GetPub[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, false)
}

func Get[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, true)
}
