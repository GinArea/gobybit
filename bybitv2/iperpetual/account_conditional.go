// Conditional Orders (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-conditionalorders)
package iperpetual

type ConditionalOrderBase struct {
	UserID      int          `json:"user_id"`
	Symbol      string       `json:"symbol"`
	Side        Side         `json:"side"`
	OrderType   OrderType    `json:"order_type"`
	Price       float64      `json:"price"`
	Qty         int          `json:"qty"`
	TimeInForce TimeInForce  `json:"time_in_force"`
	TriggerBy   TriggerPrice `json:"trigger_by"`
	StopPx      string       `json:"stop_px"`
	BasePrice   string       `json:"base_price"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

type ConditionalOrderProfitLoss struct {
	TakeProfit float64      `json:"take_profit"`
	StopLoss   float64      `json:"stop_loss"`
	TpTrigger  TriggerPrice `json:"tp_trigger_by"`
	SlTrigger  TriggerPrice `json:"sl_trigger_by"`
}

// Place Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-placecond)
type PlaceConditionalOrder struct {
	Side           Side          `param:"side"`
	Symbol         string        `param:"symbol"`
	OrderType      OrderType     `param:"order_type"`
	Qty            float64       `param:"qty"`
	TimeInForce    TimeInForce   `param:"time_in_force"`
	BasePrice      string        `param:"base_price"`
	StopPx         string        `param:"stop_px"`
	Price          *float64      `param:"price"`
	CloseOnTrigger *bool         `param:"close_on_trigger"`
	OrderLinkID    *string       `param:"order_link_id"`
	TakeProfit     *float64      `param:"take_profit"`
	StopLoss       *float64      `param:"stop_loss"`
	TpTrigger      *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger      *TriggerPrice `param:"sl_trigger_by"`
	TriggerBy      *TriggerPrice `param:"trigger_by"`
}

type ConditionalOrderCreated struct {
	ConditionalOrderBase
	ConditionalOrderProfitLoss
	Remark       string  `json:"remark"`
	RejectReason string  `json:"reject_reason"`
	LeavesQty    float64 `json:"leaves_qty"`
	LeavesValue  string  `json:"leaves_value"`
	StopOrderID  string  `json:"stop_order_id"`
	OrderLinkID  string  `json:"order_link_id"`
}

func (o *PlaceConditionalOrder) Do(client *Client) (ConditionalOrderCreated, error) {
	return Post[ConditionalOrderCreated](client, "stop-order/create", o)
}

func (o *Client) PlaceConditionalOrder(v PlaceConditionalOrder) (ConditionalOrderCreated, error) {
	return v.Do(o)
}

// Get Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getcond)
func (o OrderList) DoConditional(client *Client) (ConditionalOrderListResult, error) {
	return Get[ConditionalOrderListResult](client, "stop-order/list", o)
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

func (o *Client) ConditionalOrderList(v OrderList) (ConditionalOrderListResult, error) {
	return v.DoConditional(o)
}

// Cancel Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelcond)
func (o CancelOrder) DoConditional(client *Client) (string, error) {
	type result struct {
		StopOrderID string `json:"stop_order_id"`
	}
	r, err := Post[result](client, "stop-order/cancel", o)
	return r.StopOrderID, err
}

func (o *Client) CancelConditionalOrder(v CancelOrder) (string, error) {
	return v.DoConditional(o)
}

// Cancel All Conditional Orders (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelallcond)
func (o CancelAllOrders) DoConditional(client *Client) ([]ConditionalCancelOrderItem, error) {
	return Post[[]ConditionalCancelOrderItem](client, "stop-order/cancelAll", o)
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

func (o *Client) CancelAllConditionalOrders(symbol string) ([]ConditionalCancelOrderItem, error) {
	return CancelAllOrders{Symbol: symbol}.DoConditional(o)
}

// Replace Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-replacecond)
type ReplaceConditionalOrder struct {
	Symbol       string        `param:"symbol"`
	OrderID      *string       `param:"stop_order_id"`
	OrderLinkID  *string       `param:"order_link_id"`
	Qty          *int          `param:"p_r_qty"`
	Price        *string       `param:"p_r_price"`
	TriggerPrice *string       `param:"p_r_trigger_price"`
	TakeProfit   *float64      `param:"take_profit"`
	StopLoss     *float64      `param:"stop_loss"`
	TpTrigger    *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger    *TriggerPrice `param:"sl_trigger_by"`
}

func (o ReplaceConditionalOrder) Do(client *Client) (string, error) {
	type result struct {
		StopOrderID string `json:"stop_order_id"`
	}
	r, err := Post[result](client, "stop-order/replace", o)
	return r.StopOrderID, err
}

func (o *Client) ReplaceConditionalOrder(v ReplaceConditionalOrder) (string, error) {
	return v.Do(o)
}

// Query Conditional Order (real-time) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querycond)
//
// When only symbol is passed, the response uses a different structure.
func (o QueryOrder) DoConditional(client *Client) ([]ConditionalOrder, error) {
	if o.OnlySymbol() {
		return Get[[]ConditionalOrder](client, "stop-order", o)
	}
	r, err := Get[ConditionalOrder](client, "stop-order", o)
	return []ConditionalOrder{r}, err
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

func (o *Client) QueryConditionalOrder(v QueryOrder) ([]ConditionalOrder, error) {
	return v.DoConditional(o)
}
