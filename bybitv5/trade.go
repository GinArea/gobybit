package bybitv5

import "github.com/msw-x/moon/ujson"

type Side string

const (
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

type OrderType string

const (
	Limit  OrderType = "Limit"
	Market OrderType = "Market"
)

// Place Order
// https://bybit-exchange.github.io/docs/v5/order/create-order
type PlaceActiveOrder struct {
	Category         Category
	Symbol           string
	Side             Side
	OrderType        OrderType
	Qty              ujson.Float64 `json:",omitempty"`
	Price            ujson.Float64 `json:",omitempty"`
	IsLeverage       *int          `json:",omitempty"`
	TriggerDirection *int          `json:",omitempty"`
	OrderFilter      string        `json:",omitempty"`
	TriggerPrice     ujson.Float64 `json:",omitempty"`
	TriggerBy        TriggerBy     `json:",omitempty"`
	OrderIv          ujson.Float64 `json:",omitempty"`
	TimeInForce      TimeInForce   `json:",omitempty"`
	PositionIdx      *PositionIdx  `json:",omitempty"`
	OrderLinkId      string        `json:",omitempty"`
	TakeProfit       ujson.Float64 `json:",omitempty"`
	StopLoss         ujson.Float64 `json:",omitempty"`
	TpTriggerBy      TriggerBy     `json:",omitempty"`
	SlTriggerBy      TriggerBy     `json:",omitempty"`
	ReduceOnly       *bool         `json:",omitempty"`
	CloseOnTrigger   *bool         `json:",omitempty"`
	SmpType          SmpType       `json:",omitempty"`
	Mmp              *bool         `json:",omitempty"`
	TpslMode         TpSlMode      `json:",omitempty"`
	TpLimitPrice     ujson.Float64 `json:",omitempty"`
	SlLimitPrice     ujson.Float64 `json:",omitempty"`
	TpOrderType      OrderType     `json:",omitempty"`
	SlOrderType      OrderType     `json:",omitempty"`
}

func (o PlaceActiveOrder) Do(c *Client) Response[OrderId] {
	return Post(c.order(), "create", o, forward[OrderId])
}

type OrderId struct {
	OrderId     string
	OrderLinkId string
}

func (o *Client) PlaceActiveOrder(v PlaceActiveOrder) Response[OrderId] {
	return v.Do(o)
}

// Get Open Orders (real-time)
// https://bybit-exchange.github.io/docs/v5/order/open-order
type GetOpenOrders struct {
	Category    Category
	Symbol      string `json:",omitempty"`
	BaseCoin    string `json:",omitempty"`
	SettleCoin  string `json:",omitempty"`
	OrderId     string `json:",omitempty"`
	OrderLinkId string `json:",omitempty"`
	OpenOnly    *int   `json:",omitempty"`
	OrderFilter string `json:",omitempty"`
	Limit       int    `json:",omitempty"`
	Cursor      string `json:",omitempty"`
}
