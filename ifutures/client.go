// [Inverse Futures] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures
package ifutures

import "github.com/tranquiil/bybit"

type Client struct {
	client *bybit.Client
}

// New creates a new client
func New(client *bybit.Client) *Client {
	return &Client{
		client:  client,
		version: 2,
	}
}

/*
// Place Active Order
func (this *Client) PlaceActiveOrder() (any, bool) {
	// coin string Currency alias. Returns all wallet balances if not passed
	resp := Response[map[Coin]Balance]{}
	err := this.client.Get(this.urlPrivate("wallet/balance"), bybit.UrlParam{
		//"coin": coin,
	}, &resp)
	return resp.Result, err == nil
}

// Place Active Order Request
type PlaceActiveOrder struct {
	Side        Side          `json:"side"`
	Symbol      SymbolInverse `json:"symbol"`
	OrderType   OrderType     `json:"order_type"`
	Qty         int           `json:"qty"`
	TimeInForce TimeInForce   `json:"time_in_force"`

	Price          *float64 `json:"price,omitempty"`
	TakeProfit     *float64 `json:"take_profit,omitempty"`
	StopLoss       *float64 `json:"stop_loss,omitempty"`
	ReduceOnly     *bool    `json:"reduce_only,omitempty"`
	CloseOnTrigger *bool    `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string  `json:"order_link_id,omitempty"`
}

func (this *PlaceActiveOrder) Do(client *Client) {

}


//ifutures.PlaceActiveOrder{}.Do(client)
//client.PlaceActiveOrder(PlaceActiveOrder{})


// CreateOrderParam :
type CreateOrderParam struct {

}*/
