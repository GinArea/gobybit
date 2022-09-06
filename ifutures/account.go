package ifutures

// Place Active Order
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
