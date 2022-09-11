package bybit

import (
	"github.com/tranquiil/bybit/spot"
	"github.com/tranquiil/bybit/transport"
)

type Client struct {
	c *transport.Client
}

func NewClient(url, key, secret string) *Client {
	return &Client{
		c: transport.NewClient(url, key, secret),
	}
}

func (this *Client) WithProxy(proxy string) *Client {
	this.c.WithProxy(proxy)
	return this
}

func (this *Client) Spot() *spot.Client {
	return spot.NewClient(this.c)
}
