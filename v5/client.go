// v5 (https://bybit-exchange.github.io/docs/v5/intro)
package v5

import (
	"fmt"

	"github.com/ginarea/gobybit/transport"
)

// v5 HTTP client
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
	return forwardError(this.c.GetPublic(this.url(path), param, ret))
}

func (this *Client) Get(path string, param any, ret any) error {
	return forwardError(this.c.Get(this.url(path), param, ret))
}

func (this *Client) Post(path string, param any, ret any) error {
	return forwardError(this.c.Post(this.url(path), param, ret))
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

func (this *Client) url(path string) string {
	return fmt.Sprintf("v5/%s", path)
}
