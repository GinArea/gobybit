// [Active Orders] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-activeorders
package uperpetual

import "github.com/ginarea/gobybit/iperpetual"

// [Place Active Order] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-placeactive
type PlaceActiveOrder struct {
	Side           Side              `param:"side"`
	Symbol         iperpetual.Symbol `param:"symbol"`
	OrderType      OrderType         `param:"order_type"`
	Qty            int               `param:"qty"`
	TimeInForce    TimeInForce       `param:"time_in_force"`
	ReduceOnly     bool              `param:"reduce_only"`
	CloseOnTrigger bool              `param:"close_on_trigger"`
	Price          *float64          `param:"price"`
	OrderLinkID    *string           `param:"order_link_id"`
	TakeProfit     *float64          `param:"take_profit"`
	StopLoss       *float64          `param:"stop_loss"`
	TpTrigger      *TriggerPrice     `param:"tp_trigger_by"`
	SlTrigger      *TriggerPrice     `param:"sl_trigger_by"`
	PositionIdx    *PositionIdx      `param:"position_idx"`
}

type OrderCreated struct {
	Order
	PositionIdx PositionIdx `json:"position_idx"`
}

func (this *PlaceActiveOrder) Do(client *Client) (OrderCreated, bool) {
	return Post[OrderCreated](client, "order/create", this)
}

func (this *Client) PlaceActiveOrder(v PlaceActiveOrder) (OrderCreated, bool) {
	return v.Do(this)
}

// [Get Active Order] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-getactive
type OrderList struct {
	Symbol      iperpetual.Symbol `param:"symbol"`
	OrderID     *string           `param:"order_id"`
	OrderLinkID *string           `param:"order_link_id"`
	Order       *SortOrder        `param:"order"`
	Page        *int              `param:"page"`
	Limit       *int              `param:"limit"`
	OrderStatus *OrderStatus      `param:"order_status"`
}

func (this OrderList) Do(client *Client) (OrderListResult, bool) {
	return Get[OrderListResult](client, "order/list", this)
}

type OrderListResult struct {
	Items       []Order `json:"data"`
	CurrentPage int     `json:"current_page"`
}

func (this *Client) OrderList(v OrderList) (OrderListResult, bool) {
	return v.Do(this)
}

// [Cancel Active Order] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelactive
type CancelOrder struct {
	Symbol      iperpetual.Symbol `param:"symbol"`
	OrderId     *string           `param:"order_id"`
	OrderLinkId *string           `param:"order_link_id"`
}

func (this CancelOrder) Do(client *Client) (string, bool) {
	type result struct {
		OrderID string `json:"order_id"`
	}
	r, ok := Post[result](client, "order/cancel", this)
	return r.OrderID, ok
}

func (this *Client) CancelOrder(v CancelOrder) (string, bool) {
	return v.Do(this)
}

// [Cancel All Active Orders] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelallactive
type CancelAllOrders struct {
	Symbol iperpetual.Symbol `param:"symbol"`
}

func (this CancelAllOrders) Do(client *Client) ([]string, bool) {
	return Post[[]string](client, "order/cancel-all", this)
}

func (this *Client) CancelAllOrders(symbol iperpetual.Symbol) ([]string, bool) {
	return CancelAllOrders{Symbol: symbol}.Do(this)
}

// [Replace Active Order] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-replaceactive
type ReplaceOrder struct {
	Symbol      iperpetual.Symbol `param:"symbol"`
	OrderId     *string           `param:"order_id"`
	OrderLinkId *string           `param:"order_link_id"`
	Qty         *int              `param:"p_r_qty"`
	Price       *string           `param:"p_r_price"`
	TakeProfit  *float64          `param:"take_profit"`
	StopLoss    *float64          `param:"stop_loss"`
	TpTrigger   *TriggerPrice     `param:"tp_trigger_by"`
	SlTrigger   *TriggerPrice     `param:"sl_trigger_by"`
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

// [Query Active Order (real-time)] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-queryactive
type QueryOrder struct {
	Symbol      iperpetual.Symbol `param:"symbol"`
	OrderId     *string           `param:"order_id"`
	OrderLinkId *string           `param:"order_link_id"`
}

func (this QueryOrder) OnlySymbol() bool {
	return this.OrderId == nil && this.OrderLinkId == nil
}

// When only symbol is passed, the response uses a different structure.
func (this QueryOrder) Do(client *Client) ([]Order, bool) {
	if this.OnlySymbol() {
		return Get[[]Order](client, "order/search", this)
	}
	r, ok := Get[Order](client, "order/search", this)
	return []Order{r}, ok
}

type Order struct {
	OrderID        string            `json:"order_id"`
	UserID         int               `json:"user_id"`
	Symbol         iperpetual.Symbol `json:"symbol"`
	Side           Side              `json:"side"`
	OrderType      OrderType         `json:"order_type"`
	Price          float64           `json:"price"`
	Qty            int               `json:"qty"`
	TimeInForce    TimeInForce       `json:"time_in_force"`
	OrderStatus    OrderStatus       `json:"order_status"`
	LastExecPrice  string            `json:"last_exec_price"`
	CumExecQty     float64           `json:"cum_exec_qty"`
	CumExecValue   float64           `json:"cum_exec_value"`
	CumExecFee     float64           `json:"cum_exec_fee"`
	ReduceOnly     bool              `json:"reduce_only"`
	CloseOnTrigger bool              `json:"close_on_trigger"`
	OrderLinkID    string            `json:"order_link_id"`
	CreatedTime    string            `json:"created_time"`
	UpdatedTime    string            `json:"updated_time"`
	TakeProfit     float64           `json:"take_profit"`
	StopLoss       float64           `json:"stop_loss"`
	TpTrigger      TriggerPrice      `json:"tp_trigger_by"`
	SlTrigger      TriggerPrice      `json:"sl_trigger_by"`
}

func (this *Client) QueryOrder(v QueryOrder) ([]Order, bool) {
	return v.Do(this)
}
