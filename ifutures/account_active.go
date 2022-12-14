// Active Orders (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-activeorders)
package ifutures

type OrderMain struct {
	UserID      int         `json:"user_id"`
	Symbol      string      `json:"symbol"`
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

// Place Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-placeactive)
type PlaceActiveOrder struct {
	Side           Side          `param:"side"`
	Symbol         string        `param:"symbol"`
	OrderType      OrderType     `param:"order_type"`
	Qty            int           `param:"qty"`
	TimeInForce    TimeInForce   `param:"time_in_force"`
	PositionIdx    *PositionIdx  `param:"position_idx"`
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
	OrderCancelled
	OrderProfitLoss
}

func (this *PlaceActiveOrder) Do(client *Client) (OrderCreated, error) {
	return Post[OrderCreated](client, "order/create", this)
}

func (this *Client) PlaceActiveOrder(v PlaceActiveOrder) (OrderCreated, error) {
	return v.Do(this)
}

// Get Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-getactive)
//
//	symbol       Required string  Name of the trading pair
//	order_status          string  Queries orders of all statuses if order_status not provided.
//	                              If you want to query orders with specific statuses, you can pass the
//	                              order_status split by ',' (eg Filled,New).
//	direction             string  Search direction. prev: prev page, next: next page. Defaults to next
//	limit                 integer Limit for data size per page, max size is 50. Default as showing 20 pieces of data per page
//	cursor                string  Page turning mark. Use return cursor. Sign using origin data, in request please use urlencode
type OrderList struct {
	Symbol      string       `param:"symbol"`
	OrderStatus *OrderStatus `param:"order_status"`
	Direction   *Direction   `param:"direction"`
	Limit       *int         `param:"limit"`
	Cursor      *string      `param:"cursor"`
}

func (this OrderList) Do(client *Client) (OrderListResult, error) {
	return Get[OrderListResult](client, "order/list", this)
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

func (this *Client) OrderList(v OrderList) (OrderListResult, error) {
	return v.Do(this)
}

// Cancel Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-cancelactive)
type CancelOrder struct {
	Symbol      string  `param:"symbol"`
	OrderId     *string `param:"order_id"`
	OrderLinkId *string `param:"order_link_id"`
}

func (this CancelOrder) Do(client *Client) (OrderCancelled, error) {
	return Post[OrderCancelled](client, "order/cancel", this)
}

type OrderCancelled struct {
	OrderBase
	LastExecTime  string `json:"last_exec_time"`
	LastExecPrice string `json:"last_exec_price"`
}

func (this *Client) CancelOrder(v CancelOrder) (OrderCancelled, error) {
	return v.Do(this)
}

// Cancel All Active Orders (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-cancelallactive)
type CancelAllOrders struct {
	Symbol string `param:"symbol"`
}

func (this CancelAllOrders) Do(client *Client) ([]CancelOrderItem, error) {
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
}

func (this *Client) CancelAllOrders(symbol string) ([]CancelOrderItem, error) {
	return CancelAllOrders{Symbol: symbol}.Do(this)
}

// Replace Active Order (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-replaceactive)
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

func (this ReplaceOrder) Do(client *Client) (string, error) {
	type result struct {
		OrderID string `json:"order_id"`
	}
	r, err := Post[result](client, "order/replace", this)
	return r.OrderID, err
}

func (this *Client) ReplaceOrder(v ReplaceOrder) (string, error) {
	return v.Do(this)
}

// Query Active Order (real-time) (https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-queryactive)
//
// Query real-time active order information. If only order_id or order_link_id are passed,
// a single order will be returned; otherwise, returns up to 500 unfilled orders.
type QueryOrder struct {
	Symbol      string  `param:"symbol"`
	OrderID     *string `param:"order_id"`
	OrderLinkID *string `param:"order_link_id"`
}

func (this QueryOrder) OnlySymbol() bool {
	return this.OrderID == nil && this.OrderLinkID == nil
}

// When only symbol is passed, the response uses a different structure.
func (this QueryOrder) Do(client *Client) ([]Order, error) {
	if this.OnlySymbol() {
		return Get[[]Order](client, "order", this)
	}
	r, err := Get[Order](client, "order", this)
	return []Order{r}, err
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

func (this *Client) QueryOrder(v QueryOrder) ([]Order, error) {
	return v.Do(this)
}
