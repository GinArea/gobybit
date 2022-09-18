// [Inverse Futures] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures
package ifutures

import (
	"fmt"

	"github.com/tranquiil/bybit/iperpetual"
	"github.com/tranquiil/bybit/transport"
)

type Client struct {
	c *transport.Client
}

func NewClient(client *transport.Client) *Client {
	return &Client{c: client}
}

func (this *Client) Get(path string, param any, ret any) error {
	return this.c.Get(this.url(path), param, ret)
}

func (this *Client) Post(path string, param any, ret any) error {
	return this.c.Post(this.url(path), param, ret)
}

func Get[T any](c *Client, path string, param any) (T, bool) {
	resp := &Response[T]{}
	err := c.Get(path, param, resp)
	return resp.Result, err == nil
}

func Post[T any](c *Client, path string, param any) (T, bool) {
	resp := &Response[T]{}
	err := c.Post(path, param, resp)
	return resp.Result, err == nil
}

func (this *Client) url(path string) string {
	return fmt.Sprintf("futures/private/%s", path)
}

func (this *Client) iperpetual() *iperpetual.Client {
	return iperpetual.NewClient(this.c)
}
