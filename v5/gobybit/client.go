package gobybit

import (
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

func GetPub[R, T any](c *Client, path string, request any, transform func(R) (T, error)) (r Response[T]) {
	h := c.c.Get(path).Params(request).Do()
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
		r.SetErrorIfNil(h.HeaderTo(&r.Limit))
	} else {
		r.Error = h.Error
	}
	return
}
