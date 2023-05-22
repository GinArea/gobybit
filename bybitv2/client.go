package gobybit

import (
	"time"

	"github.com/ginarea/gobybit/account"
	"github.com/ginarea/gobybit/ifutures"
	"github.com/ginarea/gobybit/iperpetual"
	"github.com/ginarea/gobybit/spot"
	"github.com/ginarea/gobybit/spotv3"
	"github.com/ginarea/gobybit/transport"
	"github.com/ginarea/gobybit/uperpetual"
	v5 "github.com/ginarea/gobybit/v5"
	"github.com/msw-x/moon/ulog"
)

type Client struct {
	c *transport.Client
}

func NewClient() *Client {
	return &Client{
		c: transport.NewClient(),
	}
}

func (o *Client) WithUrl(url string) *Client {
	o.c.WithUrl(url)
	return o
}

func (o *Client) WithByTickUrl() *Client {
	o.c.WithByTickUrl()
	return o
}

func (o *Client) WithAuth(key, secret string) *Client {
	o.c.WithAuth(key, secret)
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

func (o *Client) WithLog(log *ulog.Log) *Client {
	o.c.WithLog(log)
	return o
}

func (o *Client) WithLogUri(logUri bool) *Client {
	o.c.WithLogUri(logUri)
	return o
}

func (o *Client) WithLogRequest(logRequest bool) *Client {
	o.c.WithLogRequest(logRequest)
	return o
}

func (o *Client) WithLogResponse(logResponse bool) *Client {
	o.c.WithLogResponse(logResponse)
	return o
}

func (o *Client) WithOnHttpError(onHttpError func(err error, attempt int) bool) *Client {
	o.c.WithOnHttpError(onHttpError)
	return o
}

func (o *Client) HasProxy() bool {
	return o.c.HasProxy()
}

func (o *Client) Key() string {
	return o.c.Key()
}

func (o *Client) Secret() string {
	return o.c.Secret()
}

func (o *Client) InversePerpetual() *iperpetual.Client {
	return iperpetual.NewClient(o.c)
}

func (o *Client) UsdtPerpetual() *uperpetual.Client {
	return uperpetual.NewClient(o.c)
}

func (o *Client) InverseFutures() *ifutures.Client {
	return ifutures.NewClient(o.c)
}

func (o *Client) Spot() *spot.Client {
	return spot.NewClient(o.c)
}

func (o *Client) Spotv3() *spotv3.Client {
	return spotv3.NewClient(o.c)
}

func (o *Client) AccountAsset() *account.Client {
	return account.NewClient(o.c)
}

func (o *Client) V5() *v5.Client {
	return v5.NewClient(o.c)
}
