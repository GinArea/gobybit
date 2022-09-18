// [USDT Perpetual] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-introduction
package uperpetual

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

func (this *Client) GetPublic(path string, param any, ret any) error {
	return this.c.Get(this.urlPublic(path), param, ret)
}

func (this *Client) Get(path string, param any, ret any) error {
	return this.c.Get(this.urlPrivate(path), param, ret)
}

func (this *Client) Post(path string, param any, ret any) error {
	return this.c.Post(this.urlPrivate(path), param, ret)
}

func GetPublic[T any](c *Client, path string, param any) (T, bool) {
	resp := &Response[T]{}
	err := c.GetPublic(path, param, resp)
	return resp.Result, err == nil
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
