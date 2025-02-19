package bybitv5

import "github.com/msw-x/moon/ujson"

// Place Order
// https://bybit-exchange.github.io/docs/v5/order/create-order
type PlaceOrder struct {
	Category         Category
	Symbol           string
	Side             Side
	OrderType        OrderType
	Qty              ujson.StringFloat64 `json:",omitempty"`
	Price            ujson.Float64       `json:",omitempty"`
	IsLeverage       *int                `json:",omitempty"`
	TriggerDirection *int                `json:",omitempty"`
	OrderFilter      string              `json:",omitempty"`
	TriggerPrice     ujson.Float64       `json:",omitempty"`
	TriggerBy        TriggerBy           `json:",omitempty"`
	OrderIv          ujson.Float64       `json:",omitempty"`
	TimeInForce      TimeInForce         `json:",omitempty"`
	PositionIdx      *PositionIdx        `json:",omitempty"`
	OrderLinkId      string              `json:",omitempty"`
	TakeProfit       ujson.Float64       `json:",omitempty"`
	StopLoss         ujson.Float64       `json:",omitempty"`
	TpTriggerBy      TriggerBy           `json:",omitempty"`
	SlTriggerBy      TriggerBy           `json:",omitempty"`
	ReduceOnly       *bool               `json:",omitempty"`
	CloseOnTrigger   *bool               `json:",omitempty"`
	SmpType          SmpType             `json:",omitempty"`
	Mmp              *bool               `json:",omitempty"`
	TpslMode         TpSlMode            `json:",omitempty"`
	TpLimitPrice     ujson.Float64       `json:",omitempty"`
	SlLimitPrice     ujson.Float64       `json:",omitempty"`
	TpOrderType      OrderType           `json:",omitempty"`
	SlOrderType      OrderType           `json:",omitempty"`
}

type OrderId struct {
	OrderId     string
	OrderLinkId string
}

func (o PlaceOrder) Do(c *Client) Response[OrderId] {
	return Post(c.order(), "create", o, forward[OrderId])
}

func (o *Client) PlaceOrder(v PlaceOrder) Response[OrderId] {
	return v.Do(o)
}

// todo: change to PlaceOrder
type PlaceOrder2 struct {
	Category         Category
	Symbol           string
	Side             Side
	OrderType        OrderType
	Qty              ujson.StringFloat64 `json:",omitempty"`
	Price            ujson.StringFloat64 `json:",omitempty"`
	IsLeverage       *int                `json:",omitempty"`
	TriggerDirection *int                `json:",omitempty"`
	OrderFilter      string              `json:",omitempty"`
	TriggerPrice     ujson.Float64       `json:",omitempty"`
	TriggerBy        TriggerBy           `json:",omitempty"`
	OrderIv          ujson.Float64       `json:",omitempty"`
	TimeInForce      TimeInForce         `json:",omitempty"`
	PositionIdx      *PositionIdx        `json:",omitempty"`
	OrderLinkId      string              `json:",omitempty"`
	TakeProfit       ujson.Float64       `json:",omitempty"`
	StopLoss         ujson.Float64       `json:",omitempty"`
	TpTriggerBy      TriggerBy           `json:",omitempty"`
	SlTriggerBy      TriggerBy           `json:",omitempty"`
	ReduceOnly       *bool               `json:",omitempty"`
	CloseOnTrigger   *bool               `json:",omitempty"`
	SmpType          SmpType             `json:",omitempty"`
	Mmp              *bool               `json:",omitempty"`
	TpslMode         TpSlMode            `json:",omitempty"`
	TpLimitPrice     ujson.Float64       `json:",omitempty"`
	SlLimitPrice     ujson.Float64       `json:",omitempty"`
	TpOrderType      OrderType           `json:",omitempty"`
	SlOrderType      OrderType           `json:",omitempty"`
}

func (o PlaceOrder2) Do(c *Client) Response[OrderId] {
	return Post(c.order(), "create", o, forward[OrderId])
}

func (o *Client) PlaceOrder2(v PlaceOrder2) Response[OrderId] {
	return v.Do(o)
}

// Amend Order
// https://bybit-exchange.github.io/docs/v5/order/amend-order
type AmendOrder struct {
	Category     Category
	Symbol       string
	OrderId      string        `json:",omitempty"`
	OrderLinkId  string        `json:",omitempty"`
	OrderIv      ujson.Float64 `json:",omitempty"`
	TriggerPrice ujson.Float64 `json:",omitempty"`
	Qty          ujson.Float64 `json:",omitempty"`
	Price        ujson.Float64 `json:",omitempty"`
	TakeProfit   ujson.Float64 `json:",omitempty"`
	StopLoss     ujson.Float64 `json:",omitempty"`
	TpTriggerBy  TriggerBy     `json:",omitempty"`
	SlTriggerBy  TriggerBy     `json:",omitempty"`
	TriggerBy    TriggerBy     `json:",omitempty"`
	TpLimitPrice ujson.Float64 `json:",omitempty"`
	SlLimitPrice ujson.Float64 `json:",omitempty"`
}

func (o AmendOrder) Do(c *Client) Response[OrderId] {
	return Post(c.order(), "amend", o, forward[OrderId])
}

func (o *Client) AmendOrder(v AmendOrder) Response[OrderId] {
	return v.Do(o)
}

// Cancel Order
// https://bybit-exchange.github.io/docs/v5/order/cancel-order
type CancelOrder struct {
	Category    Category
	Symbol      string
	OrderId     string `json:",omitempty"`
	OrderLinkId string `json:",omitempty"`
	OrderFilter string `json:",omitempty"`
}

func (o CancelOrder) Do(c *Client) Response[OrderId] {
	return Post(c.order(), "cancel", o, forward[OrderId])
}

func (o *Client) CancelOrder(v CancelOrder) Response[OrderId] {
	return v.Do(o)
}

// Get Open Orders (real-time)
// https://bybit-exchange.github.io/docs/v5/order/open-order
type GetOpenOrders struct {
	Category    Category
	Symbol      string `url:",omitempty"`
	BaseCoin    string `url:",omitempty"`
	SettleCoin  string `url:",omitempty"`
	OrderId     string `url:",omitempty"`
	OrderLinkId string `url:",omitempty"`
	OpenOnly    *int   `url:",omitempty"`
	OrderFilter string `url:",omitempty"`
	Limit       int    `url:",omitempty"`
	Cursor      string `url:",omitempty"`
}

type Order struct {
	OrderId            string
	OrderLinkId        string
	BlockTradeId       string
	Symbol             string
	Price              ujson.Float64
	Qty                ujson.Float64
	Side               Side
	IsLeverage         string
	PositionIdx        PositionIdx
	OrderStatus        OrderStatus
	CancelType         CancelType
	RejectReason       RejectReason
	AvgPrice           ujson.Float64
	LeavesQty          ujson.Float64
	LeavesValue        string
	CumExecQty         ujson.Float64
	CumExecValue       ujson.Float64
	CumExecFee         ujson.Float64
	FeeCurrency        string
	TimeInForce        TimeInForce
	OrderType          OrderType
	StopOrderType      StopOrderType
	OrderIv            string
	TriggerPrice       ujson.StringFloat64
	TakeProfit         ujson.StringFloat64
	StopLoss           ujson.StringFloat64
	TpTriggerBy        TriggerBy
	SlTriggerBy        TriggerBy
	TriggerDirection   int
	TriggerBy          TriggerBy
	LastPriceOnCreated ujson.StringFloat64
	ReduceOnly         bool
	CloseOnTrigger     bool
	SmpType            SmpType
	SmpGroup           int
	SmpOrderId         string
	TpslMode           TpSlMode
	TpLimitPrice       ujson.StringFloat64
	SlLimitPrice       ujson.StringFloat64
	PlaceType          string
	CreatedTime        ujson.TimeMs
	UpdatedTime        ujson.TimeMs
}

func (o Order) GetCumExecValue() float64 {
	return o.CumExecQty.Value() * o.AvgPrice.Value()
}

func (o GetOpenOrders) Do(c *Client) Response[[]Order] {
	type result struct {
		Category       Category
		NextPageCursor string
		List           []Order
	}
	return Get(c.order(), "realtime", o, func(r result) ([]Order, error) {
		return r.List, nil
	})
}

func (o *Client) GetOpenOrders(v GetOpenOrders) Response[[]Order] {
	return v.Do(o)
}

// Cancel All Orders
// https://bybit-exchange.github.io/docs/v5/order/cancel-all
type CancelAllOrders struct {
	Category    Category
	Symbol      string `json:",omitempty"`
	BaseCoin    string `json:",omitempty"`
	SettleCoin  string `json:",omitempty"`
	OrderFilter string `json:",omitempty"`
}

func (o CancelAllOrders) Do(c *Client) Response[OrderId] {
	return Post(c.order(), "cancel-all", o, forward[OrderId])
}

func (o CancelAllOrders) DoSpot(c *Client) Response[bool] {
	type result struct {
		Success ujson.Bool
	}
	return Post(c.order(), "cancel-all", o, func(r result) (bool, error) {
		return r.Success.Value(), nil
	})
}

func (o *Client) CancelAllOrders(v CancelAllOrders) Response[OrderId] {
	return v.Do(o)
}

func (o *Client) CancelAllOrdersSpot(v CancelAllOrders) Response[bool] {
	return v.DoSpot(o)
}

// Get Order History
// https://bybit-exchange.github.io/docs/v5/order/order-list
type GetOrderHistory struct {
	Category    Category
	Symbol      string      `url:",omitempty"`
	BaseCoin    string      `url:",omitempty"`
	OrderId     string      `url:",omitempty"`
	OrderLinkId string      `url:",omitempty"`
	OrderFilter string      `url:",omitempty"`
	OrderStatus OrderStatus `url:",omitempty"`
	StartTime   int         `url:",omitempty"`
	EndTime     int         `url:",omitempty"`
	Limit       int         `url:",omitempty"`
	Cursor      string      `url:",omitempty"`
}

func (o GetOrderHistory) Do(c *Client) Response[[]Order] {
	type result struct {
		Category       Category
		NextPageCursor string
		List           []Order
	}
	return Get(c.order(), "history", o, func(r result) ([]Order, error) {
		return r.List, nil
	})
}

func (o *Client) GetOrderHistory(v GetOrderHistory) Response[[]Order] {
	return v.Do(o)
}
