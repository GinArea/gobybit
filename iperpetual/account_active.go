// [Active Orders] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-activeorders
package iperpetual

type OrderMain struct {
	UserID      int         `json:"user_id"`
	Symbol      Symbol      `json:"symbol"`
	Side        Side        `json:"side"`
	OrderType   OrderType   `json:"order_type"`
	Price       float64     `json:"price"`
	Qty         int         `json:"qty"`
	TimeInForce TimeInForce `json:"time_in_force"`
	OrderStatus OrderStatus `json:"order_status"`
	LeavesQty   float64     `json:"leaves_qty"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

type OrderBase struct {
	OrderMain
	OrderID      string  `json:"order_id"`
	OrderLinkID  string  `json:"order_link_id"`
	CumExecQty   float64 `json:"cum_exec_qty"`
	CumExecValue float64 `json:"cum_exec_value"`
	CumExecFee   float64 `json:"cum_exec_fee"`
	RejectReason string  `json:"reject_reason"`
}
type OrderProfitLoss struct {
	TakeProfit float64      `json:"take_profit"`
	StopLoss   float64      `json:"stop_loss"`
	TpTrigger  TriggerPrice `json:"tp_trigger_by"`
	SlTrigger  TriggerPrice `json:"sl_trigger_by"`
}

// [Place Active Order] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-placeactive
type PlaceActiveOrder struct {
	Side           Side          `param:"side"`
	Symbol         Symbol        `param:"symbol"`
	OrderType      OrderType     `param:"order_type"`
	Qty            int           `param:"qty"`
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
	LastExecTime  string `json:"last_exec_time"`
	LastExecPrice string `json:"last_exec_price"`
}

func (this *PlaceActiveOrder) Do(client *Client) (OrderCreated, bool) {
	return Post[OrderCreated](client, "order/create", this)
}

func (this *Client) PlaceActiveOrder(v PlaceActiveOrder) (OrderCreated, bool) {
	return v.Do(this)
}

// [Get Active Order] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getactive
type OrderList struct {
	Symbol      Symbol       `param:"symbol"`
	OrderStatus *OrderStatus `param:"order_status"`
	Direction   *Direction   `param:"direction"`
	Limit       *int         `param:"limit"`
	Cursor      *string      `param:"cursor"`
}

func (this OrderList) Do(client *Client) (OrderListResult, bool) {
	return GetPrivate[OrderListResult](client, "order/list", this)
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

func (this *Client) OrderList(v OrderList) (OrderListResult, bool) {
	return v.Do(this)
}

// [Cancel Active Order] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-cancelactive
type CancelOrder struct {
	Symbol      Symbol  `param:"symbol"`
	OrderId     *string `param:"order_id"`
	OrderLinkId *string `param:"order_link_id"`
}

func (this CancelOrder) Do(client *Client) (OrderCancelled, bool) {
	return Post[OrderCancelled](client, "order/cancel", this)
}

type OrderCancelled struct {
	OrderCreated
}

func (this *Client) CancelOrder(v CancelOrder) (OrderCancelled, bool) {
	return v.Do(this)
}

// [Cancel All Active Orders] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelallactive
type CancelAllOrders struct {
	Symbol Symbol `param:"symbol"`
}

func (this CancelAllOrders) Do(client *Client) ([]CancelOrderItem, bool) {
	return Post[[]CancelOrderItem](client, "order/cancelAll", this)
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

func (this *Client) CancelAllOrders(symbol Symbol) ([]CancelOrderItem, bool) {
	return CancelAllOrders{Symbol: symbol}.Do(this)
}

// [Replace Active Order] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-replaceactive
type ReplaceOrder struct {
	Symbol      Symbol        `param:"symbol"`
	OrderId     *string       `param:"order_id"`
	OrderLinkId *string       `param:"order_link_id"`
	Qty         *int          `param:"p_r_qty"`
	Price       *string       `param:"p_r_price"`
	TakeProfit  *float64      `param:"take_profit"`
	StopLoss    *float64      `param:"stop_loss"`
	TpTrigger   *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger   *TriggerPrice `param:"sl_trigger_by"`
}

func (this ReplaceOrder) Do(client *Client) (string, bool) {
	type result struct {
		OrderID string `json:"order_id"`
	}
	r, ok := Post[result](client, "order/replace", this)
	return r.OrderID, ok
}

func (this *Client) ReplaceOrder(v ReplaceOrder) (string, bool) {
	return v.Do(this)
}

// [Query Active Order (real-time)] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-queryactive
// Query real-time active order information. If only order_id or order_link_id are passed,
// a single order will be returned; otherwise, returns up to 500 unfilled orders.
type QueryOrder struct {
	Symbol      Symbol  `param:"symbol"`
	OrderId     *string `param:"order_id"`
	OrderLinkId *string `param:"order_link_id"`
}

func (this QueryOrder) OnlySymbol() bool {
	return this.OrderId == nil && this.OrderLinkId == nil
}

// When only symbol is passed, the response uses a different structure.
func (this QueryOrder) Do(client *Client) ([]Order, bool) {
	if this.OnlySymbol() {
		return GetPrivate[[]Order](client, "order", this)
	}
	r, ok := GetPrivate[Order](client, "order", this)
	return []Order{r}, ok
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

func (this *Client) QueryOrder(v QueryOrder) ([]Order, bool) {
	return v.Do(this)
}
