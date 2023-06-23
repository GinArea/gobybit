// Active Orders (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-activeorders)
package iperpetual

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ginarea/gobybit/bybitv2/transport"
)

type OrderMain struct {
	UserID      int               `json:"user_id"`
	Symbol      string            `json:"symbol"`
	Side        Side              `json:"side"`
	OrderType   OrderType         `json:"order_type"`
	Price       transport.Float64 `json:"price"`
	Qty         int               `json:"qty"`
	TimeInForce TimeInForce       `json:"time_in_force"`
	OrderStatus OrderStatus       `json:"order_status"`
	LeavesQty   float64           `json:"leaves_qty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type OrderBase struct {
	OrderMain
	OrderID      string            `json:"order_id"`
	OrderLinkID  string            `json:"order_link_id"`
	CumExecQty   transport.Float64 `json:"cum_exec_qty"`
	CumExecValue transport.Float64 `json:"cum_exec_value"`
	CumExecFee   transport.Float64 `json:"cum_exec_fee"`
	RejectReason string            `json:"reject_reason"`
}
type OrderProfitLoss struct {
	TakeProfit string       `json:"take_profit"`
	StopLoss   string       `json:"stop_loss"`
	TpTrigger  TriggerPrice `json:"tp_trigger_by"`
	SlTrigger  TriggerPrice `json:"sl_trigger_by"`
}

// Place Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-placeactive)
type PlaceActiveOrder struct {
	Side           Side          `param:"side"`
	Symbol         string        `param:"symbol"`
	OrderType      OrderType     `param:"order_type"`
	Qty            float64       `param:"qty"`
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
	OrderBase
	OrderProfitLoss
	LastExecTime  transport.Float64 `json:"last_exec_time"`
	LastExecPrice transport.Float64 `json:"last_exec_price"`
}

func (o PlaceActiveOrder) Do(client *Client) (OrderCreated, error) {
	return Post[OrderCreated](client, "order/create", o)
}

func (o *Client) PlaceActiveOrder(v PlaceActiveOrder) (OrderCreated, error) {
	return v.Do(o)
}

// Get Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getactive)
type OrderList struct {
	Symbol      string       `param:"symbol"`
	OrderStatus *OrderStatus `param:"order_status"`
	Direction   *Direction   `param:"direction"`
	Limit       *int         `param:"limit"`
	Cursor      *string      `param:"cursor"`
}

func (o OrderList) Do(client *Client) (OrderListResult, error) {
	return Get[OrderListResult](client, "order/list", o)
}

type OrderListResult struct {
	Items  []OrderItem `json:"data"`
	Cursor string      `json:"cursor"`
}

type OrderItem struct {
	OrderBase
	OrderProfitLoss
	LeavesValue string      `json:"leaves_value"`
	PositionIdx PositionIdx `json:"position_idx"`
}

func (o *Client) OrderList(v OrderList) (OrderListResult, error) {
	return v.Do(o)
}

// Cancel Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-cancelactive)
type CancelOrder struct {
	Symbol      string  `param:"symbol"`
	OrderId     *string `param:"order_id"`
	OrderLinkId *string `param:"order_link_id"`
}

func (o CancelOrder) Do(client *Client) (OrderCancelled, error) {
	return Post[OrderCancelled](client, "order/cancel", o)
}

type OrderCancelled struct {
	OrderCreated
}

func (o *Client) CancelOrder(v CancelOrder) (OrderCancelled, error) {
	return v.Do(o)
}

// Cancel All Active Orders (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelallactive)
type CancelAllOrders struct {
	Symbol string `param:"symbol"`
}

func (o CancelAllOrders) Do(client *Client) ([]CancelOrderItem, error) {
	return Post[[]CancelOrderItem](client, "order/cancelAll", o)
}

type CancelOrderItem struct {
	OrderMain
	OrderID     string      `json:"clOrdID"`
	LeavesValue string      `json:"leaves_value"`
	CreateType  CreateType  `json:"create_type"`
	CancelType  CancelType  `json:"cancel_type"`
	CrossStatus OrderStatus `json:"cross_status"`
	CrossSeq    int         `json:"cross_seq"`
	OrderLinkID string      `оыщт:"order_link_id"`
}

func (o *Client) CancelAllOrders(symbol string) ([]CancelOrderItem, error) {
	return CancelAllOrders{Symbol: symbol}.Do(o)
}

// Replace Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-replaceactive)
type ReplaceOrder struct {
	Symbol      string        `param:"symbol"`
	OrderID     *string       `param:"order_id"`
	OrderLinkID *string       `param:"order_link_id"`
	Qty         *int          `param:"p_r_qty"`
	Price       *string       `param:"p_r_price"`
	TakeProfit  *float64      `param:"take_profit"`
	StopLoss    *float64      `param:"stop_loss"`
	TpTrigger   *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger   *TriggerPrice `param:"sl_trigger_by"`
}

func (o ReplaceOrder) Do(client *Client) (string, error) {
	type result struct {
		OrderID string `json:"order_id"`
	}
	r, err := Post[result](client, "order/replace", o)
	return r.OrderID, err
}

func (o *Client) ReplaceOrder(v ReplaceOrder) (string, error) {
	return v.Do(o)
}

// Query Active Order (real-time) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-queryactive)
//
// Query real-time active order information. If only order_id or order_link_id are passed,
// a single order will be returned; otherwise, returns up to 500 unfilled orders.
type QueryOrder struct {
	Symbol      string  `param:"symbol"`
	OrderID     *string `param:"order_id"`
	OrderLinkID *string `param:"order_link_id"`
}

func (o QueryOrder) OnlySymbol() bool {
	return o.OrderID == nil && o.OrderLinkID == nil
}

// When only symbol is passed, the response uses a different structure.
func (o QueryOrder) Do(client *Client) (l []Order, err error) {
	var r json.RawMessage
	r, err = Get[json.RawMessage](client, "order", o)
	if len(r) > 0 && r[0] == '[' {
		err = json.Unmarshal(r, &l)
	} else {
		l = make([]Order, 1)
		err = json.Unmarshal(r, &l[0])
	}
	return
}

type Order struct {
	OrderCancelled
	LeavesValue string         `json:"leaves_value"`
	PositionIdx PositionIdx    `json:"position_idx"`
	CancelType  CancelType     `json:"cancel_type"`
	ExtFields   OrderExtFields `json:"ext_fields"`
}

type OrderExtFields struct {
	oreqNum  int64  `json:"o_req_num"`
	XreqType string `json:"xreq_type"`
}

func (o *Client) QueryOrder(v QueryOrder) ([]Order, error) {
	return v.Do(o)
}

func (o *Client) QueryOrderByID(symbol string, orderID string) (i Order, err error) {
	ret, err := o.QueryOrder(QueryOrder{
		Symbol:  symbol,
		OrderID: &orderID,
	})
	if err == nil {
		if len(ret) == 1 {
			i = ret[0]
		} else {
			err = fmt.Errorf("query order result len != 1 (%d)", len(ret))
		}
	}
	return
}

func (o *Client) QueryOrderByLinkID(symbol string, orderLinkID string) (i Order, err error) {
	ret, err := o.QueryOrder(QueryOrder{
		Symbol:      symbol,
		OrderLinkID: &orderLinkID,
	})
	if err == nil {
		if len(ret) == 1 {
			i = ret[0]
		} else {
			err = fmt.Errorf("query order result len != 1 (%d)", len(ret))
		}
	}
	return
}
