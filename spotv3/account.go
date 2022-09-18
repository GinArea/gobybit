// [Account Data Endpoints] https://bybit-exchange.github.io/docs/spot/v3/#t-accountdata
package spotv3

type OrderBase struct {
	AccountId   string      `json:"accountId"`
	OrderID     string      `json:"orderId"`
	OrderLinkID string      `json:"orderLinkId"`
	Symbol      Symbol      `json:"symbol"`
	CreatedTime uint64      `json:"createTime"`
	Price       string      `json:"orderPrice"`
	OrigQty     string      `json:"orderQty"`
	OrderType   OrderType   `json:"orderType"`
	Side        Side        `json:"side"`
	OrderStatus OrderStatus `json:"status"`
	TimeInForce TimeInForce `json:"timeInForce"`
}

// [Place Active Order] https://bybit-exchange.github.io/docs/spot/v1/#t-placeactive
// symbol      Required string Name of the trading pair
// qty         Required number Order quantity (for market orders: when side is Buy, this is in the quote currency.
//                        Otherwise, qty is in the base currency. For example, on BTCUSDT a Buy order is in USDT, otherwise it's in BTC.
//                        For limit orders, the qty is always in the base currency.)
// side        Required string Order direction
// type        Required string Order type
// timeInForce          string Time in force
// price                number Order price. When the type field is MARKET, the price field is optional. When the type field is LIMIT or LIMIT_MAKER,
//                        the price field is required
// orderLinkId          string User-generated order ID
type PlaceOrder struct {
	Symbol        Symbol       `json:"symbol"`
	Qty           int          `json:"orderQty"`
	Side          Side         `json:"side"`
	Type          OrderType    `json:"orderType"`
	TimeInForce   *TimeInForce `json:"timeInForce"`
	Price         *Price       `json:"orderPrice"`
	OrderLinkId   *string      `json:"orderLinkId"`
	OrderCategory *int         `json:"orderCategory"`
	TriggerPrice  *string      `json:"triggerPrice"`
}

func (this PlaceOrder) Do(client *Client) (OrderCreated, bool) {
	return Post[OrderCreated](client, "order", this)
}

type OrderCreated struct {
	OrderBase
	OrderCategory int    `json:"orderCategory"`
	TriggerPrice  string `json:"triggerPrice"`
}

func (this *Client) PlaceOrder(v PlaceOrder) (OrderCreated, bool) {
	return v.Do(this)
}

// [Get Active Order] https://bybit-exchange.github.io/docs/spot/v3/#t-getactive
// orderId     string Order ID. Required if not passing orderLinkId
// orderLinkId string Unique user-set order ID. Required if not passing orderId
type GetOrder struct {
	OrderId     *string `param:"orderId"`
	OrderLinkId *string `param:"orderLinkId"`
}

func (this GetOrder) Do(client *Client) (Order, bool) {
	return Get[Order](client, "order", this)
}

type Order struct {
	OpenedOrder
	Locked string `json:"locked"`
}

func (this *Client) GetOrder(v GetOrder) (Order, bool) {
	return v.Do(this)
}

// [Cancel Active Order] https://bybit-exchange.github.io/docs/spot/v3/#t-cancelactive
// orderId     string Order ID. Required if not passing orderLinkId
// orderLinkId string Unique user-set order ID. Required if not passing orderId
type CancelOrder struct {
	OrderId     *string `param:"orderId"`
	OrderLinkId *string `param:"orderLinkId"`
}

func (this CancelOrder) Do(client *Client) (OrderCancelled, bool) {
	return Post[OrderCancelled](client, "cancel-order", this)
}

type OrderCancelled struct {
	OrderBase
	ExecQty string `json:"execQty"`
}

func (this *Client) CancelOrder(v CancelOrder) (OrderCancelled, bool) {
	return v.Do(this)
}

// [Batch Cancel Active Order] https://bybit-exchange.github.io/docs/spot/v3/#t-batchcancelactiveorder
// symbol     Required string Name of the trading pair
// side                string Order direction
// orderTypes          string Order type. Use commas to indicate multiple order types, eg LIMIT,LIMIT_MAKER. Default: LIMIT
type BatchCancelOrder struct {
	Symbol Symbol     `param:"symbol"`
	Side   *Side      `param:"side"`
	Type   *OrderType `param:"orderTypes"`
}

func (this BatchCancelOrder) Do(client *Client) (bool, bool) {
	type result struct {
		Success string `json:"success"`
	}
	r, ok := Post[result](client, "cancel-orders", this)
	return r.Success == "1", ok
}

func (this *Client) BatchCancelOrder(v BatchCancelOrder) (bool, bool) {
	return v.Do(this)
}

// [Batch Cancel Active Order By IDs] https://bybit-exchange.github.io/docs/spot/v3/#t-batchcancelactiveorderbyids
// orderIds Required string Order ID, use commas to indicate multiple orderIds. Maximum of 100 ids.
type BatchCancelOrderByID struct {
	ID []string `param:"orderIds"`
}

func (this BatchCancelOrderByID) Do(client *Client) ([]CancelOrderID, bool) {
	type result struct {
		List []CancelOrderID `json:"list"`
	}
	r, ok := Post[result](client, "cancel-orders-by-ids", this)
	return r.List, ok
}

type CancelOrderID struct {
	OrderID string `json:"orderId"`
	Code    string `json:"code"`
}

func (this *Client) BatchCancelOrderByID(ID []string) ([]CancelOrderID, bool) {
	return BatchCancelOrderByID{ID: ID}.Do(this)
}

// [Open Orders] https://bybit-exchange.github.io/docs/spot/v3/#t-openorders
// symbol  string  Name of the trading pair
// orderId string  Specify orderId to return all the orders that orderId of which are smaller than this particular one for pagination purpose
// limit   integer Default value is 500, max 500
type OpenOrders struct {
	Symbol  *Symbol `param:"symbol"`
	OrderID *string `param:"orderId"`
	Limit   *int    `param:"limit"`
}

func (this OpenOrders) Do(client *Client) ([]any, bool) {
	type result struct {
		List []any `json:"list"`
	}
	r, ok := Get[result](client, "open-orders", this)
	return r.List, ok
}

func (this *Client) OpenOrders(v OpenOrders) ([]any, bool) {
	return v.Do(this)
}

// [Order History] https://bybit-exchange.github.io/docs/spot/v3/#t-orderhistory
// symbol    string  Name of the trading pair
// orderId   string  Specify orderId to return all the orders that orderId of which are smaller than this particular one for pagination purpose
// limit     integer Default value is 500, max 500
// startTime long    Start time, unit in millisecond
// endTime   long    End time, unit in millisecond
type OrderHistory struct {
	Symbol    *Symbol `param:"symbol"`
	OrderID   *string `param:"orderId"`
	Limit     *int    `param:"limit"`
	StartTime *uint64 `param:"startTime"`
	EndTime   *uint64 `param:"endTime"`
}

func (this OrderHistory) Do(client *Client) ([]OpenedOrder, bool) {
	type result struct {
		List []OpenedOrder `json:"list"`
	}
	r, ok := Get[result](client, "history-orders", this)
	return r.List, ok
}

type OpenedOrder struct {
	OrderBase
	ExecQty             string `json:"execQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	UpdateTime          uint64 `json:"updateTime"`
	IsWorking           string `json:"isWorking"`
}

func (this *Client) OrderHistory(v OrderHistory) ([]OpenedOrder, bool) {
	return v.Do(this)
}

// [Trade History] https://bybit-exchange.github.io/docs/spot/v1/#t-tradehistory
// symbol       string  Name of the trading pair
// limit        integer Default value is 50, max 50
// fromTicketId integer Query greater than the trade ID. (fromTicketId < trade ID)
// toTicketId   integer Query smaller than the trade ID. (trade ID < toTicketId)
// orderId      integer Order ID
// startTime    long    Start time, unit in millisecond
// endTime      long    End time, unit in millisecond
type TradeHistory struct {
	Symbol       *Symbol `param:"symbol"`
	Limit        *int    `param:"limit"`
	FromTicketId *int    `param:"fromTicketId"`
	ToTicketId   *int    `param:"toTicketId"`
	OrderID      *string `param:"orderId"`
	StartTime    *uint64 `param:"startTime"`
	EndTime      *uint64 `param:"endTime"`
}

func (this TradeHistory) Do(client *Client) ([]Trade, bool) {
	type result struct {
		List []Trade `json:"list"`
	}
	r, ok := Get[result](client, "my-trades", this)
	return r.List, ok
}

type Trade struct {
	ID            string `json:"id"`
	Symbol        Symbol `json:"symbol"`
	OrderID       string `json:"orderId"`
	TradeID       string `json:"tradeId"`
	Price         string `json:"orderPrice"`
	Qty           string `json:"orderQty"`
	ExecFee       string `json:"execFee"`
	FeeTokenId    string `json:"feeTokenId"`
	CreatedTime   string `json:"createdTime"`
	IsBuyer       string `json:"isBuyer"`
	IsMaker       string `json:"isMaker"`
	MatchOrderID  string `json:"matchOrderId"`
	MakerRebate   string `json:"makerRebate"`
	ExecutionTime string `json:"executionTime"`
}

func (this *Client) TradeHistory(v TradeHistory) ([]Trade, bool) {
	return v.Do(this)
}
