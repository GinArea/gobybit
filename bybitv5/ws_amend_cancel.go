package bybitv5

// WsAmendOrder is the arg of the trade-WS order.amend operation
// (https://bybit-exchange.github.io/docs/v5/websocket/trade/amend-order).
// Price/Qty are strings: Bybit requires string args on the trade WS.
type WsAmendOrder struct {
	Category    Category
	Symbol      string
	OrderId     string `json:",omitempty"`
	OrderLinkId string `json:",omitempty"`
	Price       string `json:",omitempty"`
	Qty         string `json:",omitempty"`
}

// WsCancelOrder is the arg of the trade-WS order.cancel operation
// (https://bybit-exchange.github.io/docs/v5/websocket/trade/cancel-order).
type WsCancelOrder struct {
	Category    Category
	Symbol      string
	OrderId     string `json:",omitempty"`
	OrderLinkId string `json:",omitempty"`
}
