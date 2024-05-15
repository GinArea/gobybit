package bybitv5

type WsPlaceOrder struct {
	Category    Category
	Symbol      string
	Side        Side
	OrderType   OrderType
	Qty         string
	Price       string      `json:",omitempty"`
	TimeInForce TimeInForce `json:",omitempty"`
}
