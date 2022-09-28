// [Spot] https://bybit-exchange.github.io/docs/spot/v1
package spot

import (
	"fmt"

	"github.com/ginarea/gobybit/transport"
)

type Client struct {
	c *transport.Client
}

func NewClient(client *transport.Client) *Client {
	return &Client{c: client}
}

func (this *Client) GetPublic(path string, param any, ret any) error {
	return this.c.GetPublic(this.url(path), param, ret)
}

func (this *Client) GetQuote(path string, param any, ret any) error {
	return this.c.GetPublic(this.urlQuote(path), param, ret)
}

func (this *Client) Get(path string, param any, ret any) error {
	return this.c.Get(this.url(path), param, ret)
}

func (this *Client) Post(path string, param any, ret any) error {
	return this.c.Post(this.url(path), param, ret)
}

func (this *Client) Delete(path string, param any, ret any) error {
	return this.c.Delete(this.url(path), param, ret)
}

func GetPublic[T any](c *Client, path string, param any) (T, bool) {
	resp := &Response[T]{}
	err := c.GetPublic(path, param, resp)
	return resp.Result, err == nil
}

func GetQuote[T any](c *Client, path string, param any) (T, bool) {
	resp := &Response[T]{}
	err := c.GetQuote(path, param, resp)
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

func Delete[T any](c *Client, path string, param any) (T, bool) {
	resp := &Response[T]{}
	err := c.Delete(path, param, resp)
	return resp.Result, err == nil
}

func (this *Client) url(uri string) string {
	return fmt.Sprintf("spot/v1/%s", uri)
}

func (this *Client) urlQuote(uri string) string {
	return fmt.Sprintf("spot/quote/v1/%s", uri)
}
