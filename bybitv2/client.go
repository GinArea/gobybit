package bybitv2

import (
	"time"

	"github.com/ginarea/gobybit/bybitv2/account"
	"github.com/ginarea/gobybit/bybitv2/ifutures"
	"github.com/ginarea/gobybit/bybitv2/iperpetual"
	"github.com/ginarea/gobybit/bybitv2/spot"
	"github.com/ginarea/gobybit/bybitv2/spotv3"
	"github.com/ginarea/gobybit/bybitv2/transport"
	"github.com/ginarea/gobybit/bybitv2/uperpetual"
	v5 "github.com/ginarea/gobybit/bybitv2/v5"
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

func (o *Client) WithOnTransportError(f transport.OnTransportError) *Client {
	o.c.WithOnTransportError(f)
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
