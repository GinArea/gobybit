// [Inverse Futures] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures
package ifutures

import (
	"fmt"

	"github.com/tranquiil/bybit"
)

type Client struct {
	client *bybit.Client
}

func NewClient(Clientclient *bybit.Client) *Client {
	return &Client{
		client: client,
	}
}

// Place Active Order
func (this *Client) PlaceActiveOrder(v PlaceActiveOrder) {
	v.Do(this)
}

func (this *Client) url(path string) string {
	return fmt.Sprintf("/futures/private/%s", path)
}
