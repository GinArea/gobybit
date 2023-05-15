// Inverse Perpetual (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-introduction)
package iperpetual

import (
	"fmt"

	"github.com/ginarea/gobybit/transport"
)

// Inverse Perpetual HTTP client
type Client struct {
	c         *transport.Client
	rateLimit int
}

func NewClient(client *transport.Client) *Client {
	return &Client{c: client}
}

func (o *Client) WithRateLimit(rateLimit int) *Client {
	o.rateLimit = rateLimit
	return o
}

func (o *Client) Transport() *transport.Client {
	return o.c
}

func (o *Client) GetPublic(path string, param any, ret any) error {
	return forwardError(o.c.GetPublic(o.urlPublic(path), param, ret))
}

func (o *Client) Get(path string, param any, ret any) error {
	return forwardError(o.c.Get(o.urlPrivate(path), param, ret))
}

func (o *Client) Post(path string, param any, ret any) error {
	return forwardError(o.c.Post(o.urlPrivate(path), param, ret))
}

func GetPublic[T any](c *Client, path string, param any) (T, error) {
	resp := &Response[T]{}
	err := c.GetPublic(path, param, resp)
	return resp.Result, err
}

func Get[T any](c *Client, path string, param any) (T, error) {
	resp := &Response[T]{}
	err := c.Get(path, param, resp)
	if err == nil {
		err = c.checkRateLimit(resp.RateLimit, resp.RateLimitStatus)
	}
	return resp.Result, err
}

func Post[T any](c *Client, path string, param any) (T, error) {
	resp := &Response[T]{}
	err := c.Post(path, param, resp)
	if err == nil {
		err = c.checkRateLimit(resp.RateLimit, resp.RateLimitStatus)
	}
	return resp.Result, err
}

func (o *Client) url(access, path string) string {
	return fmt.Sprintf("v2/%s/%s", access, path)
}

func (o *Client) urlPublic(path string) string {
	return o.url("public", path)
}

func (o *Client) urlPrivate(path string) string {
	return o.url("private", path)
}

func (o *Client) checkRateLimit(limit, status int) error {
	if limit > 0 && status < o.rateLimit {
		return &transport.RateLimitError{
			Limit:  limit,
			Status: status,
		}
	}
	return nil
}
