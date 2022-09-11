package ifutures

import "github.com/tranquiil/bybit"

// [Place Active Order] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-placeactive
type PlaceActiveOrder struct {
	Side        Side        `json:"side"`
	Symbol      Symbol      `json:"symbol"`
	OrderType   OrderType   `json:"order_type"`
	Qty         int         `json:"qty"`
	TimeInForce TimeInForce `json:"time_in_force"`

	Price          *float64 `json:"price,omitempty"`
	TakeProfit     *float64 `json:"take_profit,omitempty"`
	StopLoss       *float64 `json:"stop_loss,omitempty"`
	ReduceOnly     *bool    `json:"reduce_only,omitempty"`
	CloseOnTrigger *bool    `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string  `json:"order_link_id,omitempty"`
}

type PlaceActiveOrderResult struct {
	//"user_id": 533285,
	//"order_id": "fe333932-d76a-4246-9dd2-856bd3c9273e",
	//"symbol": "BTCUSDM22",
	//"side": "Buy",
	//"order_type": "Limit",
	//"price": 20000,
	//"qty": 200,
	//"time_in_force": "GoodTillCancel",
	//"order_status": "Created",
	//"last_exec_time": 0,
	//"last_exec_price": 0,
	//"leaves_qty": 200,
	//"cum_exec_qty": 0,
	//"cum_exec_value": 0,
	//"cum_exec_fee": 0,
	//"reject_reason": "EC_NoError",
	//"order_link_id": "",
	//"created_at": "2022-06-23T06:27:42.324Z",
	//"updated_at": "2022-06-23T06:27:42.324Z",
	//"take_profit": "23000.00",
	//"stop_loss": "17000.00",
	//"tp_trigger_by": "MarkPrice",
	//"sl_trigger_by": "MarkPrice"
}

func (this *PlaceActiveOrder) Do(client *Client) {
	resp := Response[PlaceActiveOrderResult]{}
	param := bybit.NewUrlParam()
	err := client.Get(this.url("order/create"), param, &resp)
}
