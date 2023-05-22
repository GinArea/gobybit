// Conditional Orders (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-conditionalorders)
package uperpetual

// Place Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-placecond)
type PlaceConditionalOrder struct {
	Side           Side          `param:"side"`
	Symbol         string        `param:"symbol"`
	OrderType      OrderType     `param:"order_type"`
	Qty            int           `param:"qty"`
	BasePrice      string        `param:"base_price"`
	StopPx         string        `param:"stop_px"`
	TimeInForce    TimeInForce   `param:"time_in_force"`
	TriggerBy      TriggerPrice  `param:"trigger_by"`
	ReduceOnly     bool          `param:"reduce_only"`
	CloseOnTrigger bool          `param:"close_on_trigger"`
	Price          *float64      `param:"price"`
	OrderLinkID    *string       `param:"order_link_id"`
	TakeProfit     *float64      `param:"take_profit"`
	StopLoss       *float64      `param:"stop_loss"`
	TpTrigger      *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger      *TriggerPrice `param:"sl_trigger_by"`
	PositionIdx    *PositionIdx  `param:"position_idx"`
}

type ConditionalOrderCreated struct {
	ConditionalOrder
	PositionIdx PositionIdx `json:"position_idx"`
}

func (this *PlaceConditionalOrder) Do(client *Client) (ConditionalOrderCreated, error) {
	return Post[ConditionalOrderCreated](client, "stop-order/create", this)
}

func (this *Client) PlaceConditionalOrder(v PlaceConditionalOrder) (ConditionalOrderCreated, error) {
	return v.Do(this)
}

// Get Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-getcond)
func (this OrderList) DoConditional(client *Client) (ConditionalOrderListResult, error) {
	return Get[ConditionalOrderListResult](client, "stop-order/list", this)
}

type ConditionalOrderListResult struct {
	Items       []ConditionalOrderItem `json:"data"`
	CurrentPage int                    `json:"current_page"`
}

type ConditionalOrderItem struct {
	OrderID      string       `json:"stop_order_id"`
	UserID       int          `json:"user_id"`
	Symbol       string       `json:"symbol"`
	Side         Side         `json:"side"`
	OrderType    OrderType    `json:"order_type"`
	Price        float64      `json:"price"`
	Qty          int          `json:"qty"`
	TimeInForce  TimeInForce  `json:"time_in_force"`
	OrderStatus  OrderStatus  `json:"order_status"`
	TriggerPrice float64      `json:"trigger_price"`
	OrderLinkID  string       `json:"order_link_id"`
	CreatedTime  string       `json:"created_time"`
	UpdatedTime  string       `json:"updated_time"`
	BasePrice    float64      `json:"base_price"`
	TriggerBy    TriggerPrice `json:"trigger_by"`
	TpTrigger    TriggerPrice `json:"tp_trigger_by"`
	SlTrigger    TriggerPrice `json:"sl_trigger_by"`
	TakeProfit   float64      `json:"take_profit"`
	StopLoss     float64      `json:"stop_loss"`
}

func (this *Client) ConditionalOrderList(v OrderList) (ConditionalOrderListResult, error) {
	return v.DoConditional(this)
}

// Cancel Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelcond)
func (this CancelOrder) DoConditional(client *Client) (string, error) {
	type result struct {
		StopOrderID string `json:"stop_order_id"`
	}
	r, err := Post[result](client, "stop-order/cancel", this)
	return r.StopOrderID, err
}

func (this *Client) CancelConditionalOrder(v CancelOrder) (string, error) {
	return v.DoConditional(this)
}

// Cancel All Conditional Orders (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-cancelallcond)
func (this CancelAllOrders) DoConditional(client *Client) ([]string, error) {
	return Post[[]string](client, "stop-order/cancel-all", this)
}

func (this *Client) CancelAllConditionalOrders(symbol string) ([]string, error) {
	return CancelAllOrders{Symbol: symbol}.DoConditional(this)
}

// Replace Conditional Order (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-replacecond)
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

func (this ReplaceConditionalOrder) Do(client *Client) (string, error) {
	type result struct {
		StopOrderID string `json:"stop_order_id"`
	}
	r, err := Post[result](client, "stop-order/replace", this)
	return r.StopOrderID, err
}

func (this *Client) ReplaceConditionalOrder(v ReplaceConditionalOrder) (string, error) {
	return v.Do(this)
}

// Query Conditional Order (real-time) (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-querycond)
//
// When only symbol is passed, the response uses a different structure.
func (this QueryOrder) DoConditional(client *Client) ([]ConditionalOrder, error) {
	if this.OnlySymbol() {
		return Get[[]ConditionalOrder](client, "stop-order/search", this)
	}
	r, err := Get[ConditionalOrder](client, "stop-order/search", this)
	return []ConditionalOrder{r}, err
}

type ConditionalOrder struct {
	ConditionalOrderItem
	ReduceOnly     bool `json:"reduce_only"`
	CloseOnTrigger bool `json:"close_on_trigger"`
}

func (this *Client) QueryConditionalOrder(v QueryOrder) ([]ConditionalOrder, error) {
	return v.DoConditional(this)
}
