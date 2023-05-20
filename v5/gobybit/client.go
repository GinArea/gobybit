package gobybit

import (
	"time"

	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ulog"
)

type Client struct {
	c          *uhttp.Client
	key        string
	secret     string
	timeShift  int
	recvWindow int
}

func NewClient() *Client {
	o := new(Client)
	o.timeShift = -10000
	o.recvWindow = 20000
	return o
}

func (o *Client) WithUrl(url string) *Client {
	o.c.WithUrl(url)
	return o
}

func (o *Client) WithByTickUrl() *Client {
	return o.WithUrl(transport.MainBaseByTickUrl)
}

func (o *Client) WithPath(path string) *Client {
	o.c = o.c.WithPath(path)
	return o
}

func (o *Client) WithAuth(key, secret string) *Client {
	o.key = key
	o.secret = secret
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

func (o *Client) GetPub(path string, request any, response any) error {
	r := o.c.Get(path).Params(request).Do()
	if r.BodyExists() {
		r.Json(response)
	}
	return r.Error
}
