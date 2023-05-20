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

func (o *Client) Transport() *transport.Client {
	return o.c
}

func (o *Client) GetPub(path string, param any, ret any) error {
	return forwardError(o.c.GetPublic(o.url(path), param, ret))
}

func (o *Client) Get(path string, param any, ret any) error {
	return forwardError(o.c.Get(o.url(path), param, ret))
}

func (o *Client) Post(path string, param any, ret any) error {
	return forwardError(o.c.Post(o.url(path), param, ret))
}

func GetPub[T any](c *Client, path string, param any) (T, error) {
	resp := &Response[T]{}
	err := c.GetPub(path, param, resp)
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

func (o *Client) url(path string) string {
	return fmt.Sprintf("v5/%s", path)
}
