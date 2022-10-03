// Account Data Endpoints (https://bybit-exchange.github.io/docs/spot/v1/#t-accountdata)
package spot

type OrderBase struct {
	AccountID   string      `json:"accountId"`
	OrderID     string      `json:"orderId"`
	OrderLinkID string      `json:"orderLinkId"`
	Symbol      string      `json:"symbol"`
	SymbolName  string      `json:"symbolName"`
	Price       string      `json:"price"`
	OrigQty     string      `json:"origQty"`
	ExecutedQty string      `json:"executedQty"`
	OrderType   OrderType   `json:"type"`
	OrderStatus OrderStatus `json:"status"`
	TimeInForce TimeInForce `json:"timeInForce"`
}

// Place Active Order (https://bybit-exchange.github.io/docs/spot/v1/#t-placeactive)
//   symbol      Required string Name of the trading pair
//   qty         Required number Order quantity (for market orders: when side is Buy, this is in the quote currency.
//                        Otherwise, qty is in the base currency. For example, on BTCUSDT a Buy order is in USDT, otherwise it's in BTC.
//                        For limit orders, the qty is always in the base currency.)
//   side        Required string Order direction
//   type        Required string Order type
//   timeInForce          string Time in force
//   price                number Order price. When the type field is MARKET, the price field is optional. When the type field is LIMIT or LIMIT_MAKER,
//                        the price field is required
//   orderLinkId          string User-generated order ID
type PlaceOrder struct {
	Symbol      string       `param:"symbol"`
	Qty         int          `param:"qty"`
	Side        Side         `param:"side"`
	Type        OrderType    `param:"type"`
	TimeInForce *TimeInForce `param:"timeInForce"`
	Price       *Price       `param:"price"`
	OrderLinkID *string      `param:"orderLinkId"`
}

func (this PlaceOrder) Do(client *Client) (OrderCreated, bool) {
	return Post[OrderCreated](client, "order", this)
}

type OrderCreated struct {
	OrderBase
	TransactTime string `json:"transactTime"`
}

func (this *Client) PlaceOrder(v PlaceOrder) (OrderCreated, bool) {
	return v.Do(this)
}

// Get Active Order (https://bybit-exchange.github.io/docs/spot/v1/#t-getactive)
//   orderId     string Order ID. Required if not passing orderLinkId
//   orderLinkId string Unique user-set order ID. Required if not passing orderId
type GetOrder struct {
	OrderID     *string `param:"orderId"`
	OrderLinkID *string `param:"orderLinkId"`
}

func (this GetOrder) Do(client *Client) (Order, bool) {
	return Get[Order](client, "order", this)
}

type Order struct {
	OrderHistoryResult
	Locked string `json:"locked"`
}

func (this *Client) GetOrder(v GetOrder) (Order, bool) {
	return v.Do(this)
}

// Cancel Active Order (https://bybit-exchange.github.io/docs/spot/v1/#t-cancelactive)
//   orderId     string Order ID. Required if not passing orderLinkId
//   orderLinkId string Unique user-set order ID. Required if not passing orderId
type CancelOrder struct {
	OrderID     *string `param:"orderId"`
	OrderLinkID *string `param:"orderLinkId"`
}

func (this CancelOrder) Do(client *Client) (OrderCancelled, bool) {
	return Delete[OrderCancelled](client, "order", this)
}

type OrderCancelled struct {
	OrderCreated
	Side string `json:"side"`
}

func (this *Client) CancelOrder(v CancelOrder) (OrderCancelled, bool) {
	return v.Do(this)
}

// Fast Cancel Active Order (https://bybit-exchange.github.io/docs/spot/v1/#t-fastcancelactiveorder)
//   symbolId    Required string Name of the trading pair
//   orderId              string Order ID. Required if not passing orderLinkId
//   orderLinkId          string Unique user-set order ID. Required if not passing orderId
type FastCancelOrder struct {
	Symbol      string  `param:"symbolId"`
	OrderID     *string `param:"orderId"`
	OrderLinkID *string `param:"orderLinkId"`
}

func (this FastCancelOrder) Do(client *Client) (bool, bool) {
	type result struct {
		IsCancelled bool `json:"isCancelled"`
	}
	r, ok := Delete[result](client, "order/fast", this)
	return r.IsCancelled, ok
}

func (this *Client) FastCancelOrder(v FastCancelOrder) (bool, bool) {
	return v.Do(this)
}

// Batch Cancel Active Order (https://bybit-exchange.github.io/docs/spot/v1/#t-batchcancelactiveorder)
//   symbol     Required string Name of the trading pair
//   side                string Order direction
//   orderTypes          string Order type. Use commas to indicate multiple order types, eg LIMIT,LIMIT_MAKER. Default: LIMIT
type BatchCancelOrder struct {
	Symbol string     `param:"symbol"`
	Side   *Side      `param:"side"`
	Type   *OrderType `param:"orderTypes"`
}

func (this BatchCancelOrder) Do(client *Client) (bool, bool) {
	type result struct {
		Success bool `json:"success"`
	}
	r, ok := Delete[result](client, "order/batch-cancel", this)
	return r.Success, ok
}

func (this *Client) BatchCancelOrder(v BatchCancelOrder) (bool, bool) {
	return v.Do(this)
}

// Batch Fast Cancel Active Order (https://bybit-exchange.github.io/docs/spot/v1/#t-batchfastcancelactiveorder)
type BatchFastCancelOrder struct {
	Symbol string     `param:"symbol"`
	Side   *Side      `param:"side"`
	Type   *OrderType `param:"orderTypes"`
}

func (this BatchFastCancelOrder) Do(client *Client) (bool, bool) {
	type result struct {
		Success bool `json:"success"`
	}
	r, ok := Delete[result](client, "order/batch-fast-cancel", this)
	return r.Success, ok
}

func (this *Client) BatchFastCancelOrder(v BatchFastCancelOrder) (bool, bool) {
	return v.Do(this)
}

// Batch Cancel Active Order By IDs (https://bybit-exchange.github.io/docs/spot/v1/#t-batchcancelactiveorderbyids)
//   orderIds Required string Order ID, use commas to indicate multiple orderIds. Maximum of 100 ids.
type BatchCancelOrderByID struct {
	ID []string `param:"orderIds"`
}

func (this BatchCancelOrderByID) Do(client *Client) ([]CancelOrderID, bool) {
	return Delete[[]CancelOrderID](client, "order/batch-cancel-by-ids", this)
}

type CancelOrderID struct {
	OrderID string `json:"orderId"`
	Code    int    `json:"code"`
}

func (this *Client) BatchCancelOrderByID(ID []string) ([]CancelOrderID, bool) {
	return BatchCancelOrderByID{ID: ID}.Do(this)
}

// Open Orders (https://bybit-exchange.github.io/docs/spot/v1/#t-openorders)
//   symbol  string  Name of the trading pair
//   orderId string  Specify orderId to return all the orders that orderId of which are smaller than this particular one for pagination purpose
//   limit   integer Default value is 500, max 500
type OpenOrders struct {
	Symbol  *string `param:"symbol"`
	OrderID *string `param:"orderId"`
	Limit   *int    `param:"limit"`
}

func (this OpenOrders) Do(client *Client) ([]OrderBase, bool) {
	return Get[[]OrderBase](client, "open-orders", this)
}

func (this *Client) OpenOrders(v OpenOrders) ([]OrderBase, bool) {
	return v.Do(this)
}

// Order History (https://bybit-exchange.github.io/docs/spot/v1/#t-orderhistory)
//   symbol    string  Name of the trading pair
//   orderId   string  Specify orderId to return all the orders that orderId of which are smaller than this particular one for pagination purpose
//   limit     integer Default value is 500, max 500
//   startTime long    Start time, unit in millisecond
//   endTime   long    End time, unit in millisecond
type OrderHistory struct {
	Symbol    *string `param:"symbol"`
	OrderID   *string `param:"orderId"`
	Limit     *int    `param:"limit"`
	StartTime *uint64 `param:"startTime"`
	EndTime   *uint64 `param:"endTime"`
}

func (this OrderHistory) Do(client *Client) ([]OrderHistoryResult, bool) {
	return Get[[]OrderHistoryResult](client, "history-orders", this)
}

type OrderHistoryResult struct {
	OrderBase
	ExchangeId          string `json:"exchangeId"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
}

func (this *Client) OrderHistory(v OrderHistory) ([]OrderHistoryResult, bool) {
	return v.Do(this)
}

// Trade History (https://bybit-exchange.github.io/docs/spot/v1/#t-tradehistory)
//   symbol       string  Name of the trading pair
//   limit        integer Default value is 50, max 50
//   fromTicketId integer Query greater than the trade ID. (fromTicketId < trade ID)
//   toTicketId   integer Query smaller than the trade ID. (trade ID < toTicketId)
//   orderId      integer Order ID
//   startTime    long    Start time, unit in millisecond
//   endTime      long    End time, unit in millisecond
type TradeHistory struct {
	Symbol       *string `param:"symbol"`
	Limit        *int    `param:"limit"`
	FromTicketID *int    `param:"fromTicketId"`
	ToTicketID   *int    `param:"toTicketId"`
	OrderID      *string `param:"orderId"`
	StartTime    *uint64 `param:"startTime"`
	EndTime      *uint64 `param:"endTime"`
}

func (this TradeHistory) Do(client *Client) ([]Trade, bool) {
	return Get[[]Trade](client, "myTrades", this)
}

type Trade struct {
	ID              string   `json:"id"`
	Symbol          string   `json:"symbol"`
	SymbolName      string   `json:"symbolName"`
	OrderID         string   `json:"orderId"`
	TicketID        string   `json:"ticketId"`
	MatchOrderID    string   `json:"matchOrderId"`
	Price           string   `json:"price"`
	Qty             string   `json:"qty"`
	Commission      string   `json:"commission"`
	CommissionAsset string   `json:"commissionAsset"`
	Time            string   `json:"time"`
	IsBuyer         bool     `json:"isBuyer"`
	IsMaker         bool     `json:"isMaker"`
	Fee             TradeFee `json:"fee"`
	FeeTokenID      string   `json:"feeTokenId"`
	FeeAmount       string   `json:"feeAmount"`
	MakerRebate     string   `json:"makerRebate"`
	ExecutionTime   string   `json:"executionTime"`
}

type TradeFee struct {
	TokenID   string `json:"feeTokenId"`
	TokenName string `json:"feeTokenName"`
	Fee       string `json:"fee"`
}

func (this *Client) TradeHistory(v TradeHistory) ([]Trade, bool) {
	return v.Do(this)
}
