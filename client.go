package gobybit

import (
	"github.com/ginarea/gobybit/account"
	"github.com/ginarea/gobybit/ifutures"
	"github.com/ginarea/gobybit/iperpetual"
	"github.com/ginarea/gobybit/spot"
	"github.com/ginarea/gobybit/spotv3"
	"github.com/ginarea/gobybit/transport"
	"github.com/ginarea/gobybit/uperpetual"
)

type Client struct {
	c *transport.Client
}

func NewClient() *Client {
	return &Client{
		c: transport.NewClient(),
	}
}

func (this *Client) WithUrl(url string) *Client {
	this.c.WithUrl(url)
	return this
}

func (this *Client) WithByTickUrl() *Client {
	this.c.WithByTickUrl()
	return this
}

func (this *Client) WithAuth(key, secret string) *Client {
	this.c.WithAuth(key, secret)
	return this
}

func (this *Client) WithProxy(proxy string) *Client {
	this.c.WithProxy(proxy)
	return this
}

func (this *Client) WithLogUri(logUri bool) *Client {
	this.c.WithLogUri(logUri)
	return this
}

func (this *Client) WithLogResponse(logResponse bool) *Client {
	this.c.WithLogResponse(logResponse)
	return this
}

func (this *Client) InversePerpetual() *iperpetual.Client {
	return iperpetual.NewClient(this.c)
}

func (this *Client) UsdtPerpetual() *uperpetual.Client {
	return uperpetual.NewClient(this.c)
}

func (this *Client) InverseFutures() *ifutures.Client {
	return ifutures.NewClient(this.c)
}

func (this *Client) Spot() *spot.Client {
	return spot.NewClient(this.c)
}

func (this *Client) Spotv3() *spotv3.Client {
	return spotv3.NewClient(this.c)
}

func (this *Client) AccountAsset() *account.Client {
	return account.NewClient(this.c)
}
