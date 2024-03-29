// USDT Perpetual (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-introduction)
package uperpetual

import (
	"fmt"

	"github.com/ginarea/gobybit/bybitv2/iperpetual"
	"github.com/ginarea/gobybit/bybitv2/transport"
)

// USDT Perpetual HTTP client
type Client struct {
	c *transport.Client
}

func NewClient(client *transport.Client) *Client {
	return &Client{c: client}
}

func (this *Client) Transport() *transport.Client {
	return this.c
}

func (this *Client) GetPublic(path string, param any, ret any) error {
	return forwardError(this.c.GetPublic(this.urlPublic(path), param, ret))
}

func (this *Client) Get(path string, param any, ret any) error {
	return forwardError(this.c.Get(this.urlPrivate(path), param, ret))
}

func (this *Client) Post(path string, param any, ret any) error {
	return forwardError(this.c.Post(this.urlPrivate(path), param, ret))
}

func GetPublic[T any](c *Client, path string, param any) (T, error) {
	resp := &Response[T]{}
	err := c.GetPublic(path, param, resp)
	return resp.Result, err
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

func (this *Client) url(access, path string) string {
	return fmt.Sprintf("%s/linear/%s", access, path)
}

func (this *Client) urlPublic(path string) string {
	return this.url("public", path)
}

func (this *Client) urlPrivate(path string) string {
	return this.url("private", path)
}

func (this *Client) iperpetual() *iperpetual.Client {
	return iperpetual.NewClient(this.c)
}
