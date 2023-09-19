package bybitv5

import (
	"time"

	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ulog"
)

type Client struct {
	c                *uhttp.Client
	s                *Sign
	referer          string
	onTransportError OnTransportError
}

func NewClient() *Client {
	o := new(Client)
	o.c = uhttp.NewClient()
	o.WithBaseUrl(MainBaseUrl)
	o.WithPath("v5")
	o.WithReferer("GinArea")
	return o
}

func (o *Client) Clone() *Client {
	r := new(Client)
	r.c = o.c.Clone()
	r.s = o.s
	r.referer = o.referer
	r.onTransportError = o.onTransportError
	return r
}

func (o *Client) WithBaseUrl(url string) *Client {
	o.c.WithBase(url)
	return o
}

func (o *Client) WithByTickUrl() *Client {
	return o.WithBaseUrl(MainBaseByTickUrl)
}

func (o *Client) WithPath(path string) *Client {
	o.c.WithPath(path)
	return o
}

func (o *Client) WithAppendPath(path string) *Client {
	o.c.WithAppendPath(path)
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

func (o *Client) WithAuth(key, secret string) *Client {
	o.s = NewSign(key, secret)
	return o
}

func (o *Client) WithReferer(referer string) *Client {
	o.referer = referer
	return o
}

func (o *Client) WithOnTransportError(f OnTransportError) *Client {
	o.onTransportError = f
	return o
}

func (o *Client) market() *Client {
	return o.Clone().WithAppendPath("market")
}

func (o *Client) order() *Client {
	return o.Clone().WithAppendPath("order")
}

func (o *Client) position() *Client {
	return o.Clone().WithAppendPath("position")
}

func (o *Client) account() *Client {
	return o.Clone().WithAppendPath("account")
}

func (o *Client) asset() *Client {
	return o.Clone().WithAppendPath("asset")
}

func (o *Client) user() *Client {
	return o.Clone().WithAppendPath("user")
}

type OnTransportError func(err error, statusCode int, attempt int) bool
