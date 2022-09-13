// [Active Orders] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-activeorders
package iperpetual

// [Place Active Order] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-placeactive
type PlaceActiveOrder struct {
	Side           Side          `param:"side"`
	Symbol         Symbol        `param:"symbol"`
	OrderType      OrderType     `param:"order_type"`
	Qty            int           `param:"qty"`
	TimeInForce    TimeInForce   `param:"time_in_force"`
	Price          *float64      `param:"price"`
	CloseOnTrigger *bool         `param:"close_on_trigger"`
	OrderLinkID    *string       `param:"order_link_id"`
	TakeProfit     *float64      `param:"take_profit"`
	StopLoss       *float64      `param:"stop_loss"`
	TpTrigger      *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger      *TriggerPrice `param:"sl_trigger_by"`
	ReduceOnly     *bool         `param:"reduce_only"`
}

type OrderCreated struct {
}

func (this *PlaceActiveOrder) Do(client *Client) (OrderCreated, bool) {
	return Post[OrderCreated](client, "order/create", this)
}

func (this *Client) PlaceActiveOrder(v PlaceActiveOrder) (OrderCreated, bool) {
	return v.Do(this)
}
