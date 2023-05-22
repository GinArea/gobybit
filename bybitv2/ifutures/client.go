// Inverse Futures (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures)
package ifutures

import (
	"fmt"

	"github.com/ginarea/gobybit/iperpetual"
	"github.com/ginarea/gobybit/transport"
)

// Inverse Futures HTTP client
type Client struct {
	c *transport.Client
}

func NewClient(client *transport.Client) *Client {
	return &Client{c: client}
}

func (this *Client) Transport() *transport.Client {
	return this.c
}

func (this *Client) Get(path string, param any, ret any) error {
	return forwardError(this.c.Get(this.url(path), param, ret))
}

func (this *Client) Post(path string, param any, ret any) error {
	return forwardError(this.c.Post(this.url(path), param, ret))
}

func Get[T any](c *Client, path string, param any) (T, error) {
	resp := &Response[T]{}
	err := c.Get(path, param, resp)
	return resp.Result, err
}

func Post[T any](c *Client, path string, param any) (T, error) {
	resp := &Response[T]{}
	err := c.Post(path, param, resp)
	return resp.Result, err
}

func (this *Client) url(path string) string {
	return fmt.Sprintf("futures/private/%s", path)
}

func (this *Client) iperpetual() *iperpetual.Client {
	return iperpetual.NewClient(this.c)
}
