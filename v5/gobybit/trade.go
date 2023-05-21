package gobybit

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
	Qty              *float64     `json:",omitempty"` // string
	Price            *float64     `json:",omitempty"` // string
	IsLeverage       *int         `json:",omitempty"`
	TriggerDirection *int         `json:",omitempty"`
	OrderFilter      *string      `json:",omitempty"`
	TriggerPrice     *string      `json:",omitempty"`
	TriggerBy        *TriggerBy   `json:",omitempty"`
	OrderIv          *string      `json:",omitempty"`
	TimeInForce      *TimeInForce `json:",omitempty"`
	PositionIdx      *PositionIdx `json:",omitempty"`
	OrderLinkId      *string      `json:",omitempty"`
	TakeProfit       *string      `json:",omitempty"`
	StopLoss         *string      `json:",omitempty"`
	TpTriggerBy      *TriggerBy   `json:",omitempty"`
	SlTriggerBy      *TriggerBy   `json:",omitempty"`
	ReduceOnly       *bool        `json:",omitempty"`
	CloseOnTrigger   *bool        `json:",omitempty"`
	SmpType          *SmpType     `json:",omitempty"`
	Mmp              *bool        `json:",omitempty"`
	TpslMode         *string      `json:",omitempty"`
	TpLimitPrice     *string      `json:",omitempty"`
	SlLimitPrice     *string      `json:",omitempty"`
	TpOrderType      *string      `json:",omitempty"`
	SlOrderType      *string      `json:",omitempty"`
}
