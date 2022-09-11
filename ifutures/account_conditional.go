// [Conditional Orders] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-conditionalorders
package ifutures

import "github.com/tranquiil/bybit/iperpetual"

type ConditionalOrderBase struct {
	UserID      int               `json:"user_id"`
	Symbol      iperpetual.Symbol `json:"symbol"`
	Side        Side              `json:"side"`
	OrderType   OrderType         `json:"order_type"`
	Price       float64           `json:"price"`
	Qty         int               `json:"qty"`
	TimeInForce TimeInForce       `json:"time_in_force"`
	TriggerBy   TriggerPrice      `json:"trigger_by"`
	StopPx      string            `json:"stop_px"`
	BasePrice   string            `json:"base_price"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
}

type ConditionalOrderProfitLoss struct {
	TakeProfit float64      `json:"take_profit"`
	StopLoss   float64      `json:"stop_loss"`
	TpTrigger  TriggerPrice `json:"tp_trigger_by"`
	SlTrigger  TriggerPrice `json:"sl_trigger_by"`
}

// [Place Conditional Order] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-placecond
type PlaceConditionalOrder struct {
	Side           Side              `param:"side"`
	Symbol         iperpetual.Symbol `param:"symbol"`
	OrderType      OrderType         `param:"order_type"`
	Qty            int               `param:"qty"`
	TimeInForce    TimeInForce       `param:"time_in_force"`
	BasePrice      string            `param:"base_price"`
	StopPx         string            `param:"stop_px"`
	PositionIdx    *PositionIdx      `param:"position_idx"`
	Price          *float64          `param:"price"`
	CloseOnTrigger *bool             `param:"close_on_trigger"`
	OrderLinkID    *string           `param:"order_link_id"`
	TakeProfit     *float64          `param:"take_profit"`
	StopLoss       *float64          `param:"stop_loss"`
	TpTrigger      *TriggerPrice     `param:"tp_trigger_by"`
	SlTrigger      *TriggerPrice     `param:"sl_trigger_by"`
	TriggerBy      *TriggerPrice     `param:"trigger_by"`
}

type ConditionalOrderCreated struct {
	ConditionalOrderProfitLoss
	Remark       string  `json:"remark"`
	RejectReason string  `json:"reject_reason"`
	LeavesQty    float64 `json:"leaves_qty"`
	LeavesValue  string  `json:"leaves_value"`
	StopOrderID  string  `json:"stop_order_id"`
	OrderLinkID  string  `json:"order_link_id"`
}

func (this *PlaceConditionalOrder) Do(client *Client) (ConditionalOrderCreated, bool) {
	return Post[ConditionalOrderCreated](client, "stop-order/create", this)
}

func (this *Client) PlaceConditionalOrder(v PlaceConditionalOrder) (ConditionalOrderCreated, bool) {
	return v.Do(this)
}

// [Get Conditional Order] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-getcond
func (this OrderList) DoConditional(client *Client) (ConditionalOrderListResult, bool) {
	return Get[ConditionalOrderListResult](client, "stop-order/list", this)
}

type ConditionalOrderListResult struct {
	Items  []ConditionalOrderItem `json:"data"`
	Cursor string                 `json:"cursor"`
}

type ConditionalOrderItem struct {
	ConditionalOrderProfitLoss
	StopOrderStatus OrderStatus `json:"stop_order_status"`
	StopOrderID     string      `json:"stop_order_id"`
	OrderLinkID     string      `json:"order_link_id"`
	StopOrderType   StopOrder   `json:"stop_order_type"`
	PositionIdx     PositionIdx `json:"position_idx"`
}

func (this *Client) ConditionalOrderList(v OrderList) (ConditionalOrderListResult, bool) {
	return v.DoConditional(this)
}

// [Cancel Conditional Order] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-cancelcond
func (this CancelOrder) DoConditional(client *Client) (string, bool) {
	type result struct {
		StopOrderID string `json:"stop_order_id"`
	}
	r, ok := Post[result](client, "stop-order/cancel", this)
	return r.StopOrderID, ok
}

func (this *Client) CancelConditionalOrder(v CancelOrder) (string, bool) {
	return v.DoConditional(this)
}

// [Cancel All Conditional Orders] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-cancelallcond
func (this CancelAllOrders) DoConditional(client *Client) ([]ConditionalCancelOrderItem, bool) {
	return Post[[]ConditionalCancelOrderItem](client, "stop-order/cancelAll", this)
}

type ConditionalCancelOrderItem struct {
	ConditionalOrderProfitLoss
	OrderID           string      `json:"clOrdID"`
	CrossStatus       string      `json:"cross_status"`
	CrossSeq          int         `json:"cross_seq"`
	ExpectedDirection string      `json:"expected_direction"`
	CreateType        CreateType  `json:"create_type"`
	CancelType        CancelType  `json:"cancel_type"`
	OrderStatus       OrderStatus `json:"order_status"`
	LeavesQty         float64     `json:"leaves_qty"`
	LeavesValue       string      `json:"leaves_value"`
	StopOrderType     StopOrder   `json:"stop_order_type"`
}

func (this *Client) CancelAllConditionalOrders(symbol iperpetual.Symbol) ([]ConditionalCancelOrderItem, bool) {
	return CancelAllOrders{Symbol: symbol}.DoConditional(this)
}

// [Replace Conditional Order] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-replacecond
type ReplaceConditionalOrder struct {
	Symbol       iperpetual.Symbol `param:"symbol"`
	OrderID      *string           `param:"stop_order_id"`
	OrderLinkId  *string           `param:"order_link_id"`
	Qty          *int              `param:"p_r_qty"`
	Price        *string           `param:"p_r_price"`
	TriggerPrice *string           `param:"p_r_trigger_price"`
	TakeProfit   *float64          `param:"take_profit"`
	StopLoss     *float64          `param:"stop_loss"`
	TpTrigger    *TriggerPrice     `param:"tp_trigger_by"`
	SlTrigger    *TriggerPrice     `param:"sl_trigger_by"`
}

func (this ReplaceConditionalOrder) Do(client *Client) (string, bool) {
	type result struct {
		StopOrderID string `json:"stop_order_id"`
	}
	r, ok := Post[result](client, "stop-order/replace", this)
	return r.StopOrderID, ok
}

func (this *Client) ReplaceConditionalOrder(v ReplaceConditionalOrder) (string, bool) {
	return v.Do(this)
}

// [Query Conditional Order (real-time)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-querycond
// When only symbol is passed, the response uses a different structure.
func (this QueryOrder) DoConditional(client *Client) ([]ConditionalOrder, bool) {
	if this.OnlySymbol() {
		return Get[[]ConditionalOrder](client, "stop-order", this)
	}
	r, ok := Get[ConditionalOrder](client, "stop-order", this)
	return []ConditionalOrder{r}, ok
}

type ConditionalOrder struct {
	ConditionalOrderProfitLoss
	CumExecQty   float64                   `json:"cum_exec_qty"`
	CumExecValue float64                   `json:"cum_exec_value"`
	CumExecFee   float64                   `json:"cum_exec_fee"`
	OrderID      string                    `json:"order_id"`
	RejectReason string                    `json:"reject_reason"`
	OrderStatus  OrderStatus               `json:"order_status"`
	LeavesQty    float64                   `json:"leaves_qty"`
	LeavesValue  string                    `json:"leaves_value"`
	CancelType   CancelType                `json:"cancel_type"`
	OrderLinkID  string                    `json:"order_link_id"`
	PositionIdx  PositionIdx               `json:"position_idx"`
	ExtFields    ConditionalOrderExtFields `json:"ext_fields"`
}

type ConditionalOrderExtFields struct {
	oreqNum int64 `json:"o_req_num"`
}

func (this *Client) QueryConditionalOrder(v QueryOrder) ([]ConditionalOrder, bool) {
	return v.DoConditional(this)
}
