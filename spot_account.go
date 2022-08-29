package bybit

type SpotOrderBase struct {
	AccountId   string      `json:"accountId"`
	OrderID     string      `json:"orderId"`
	OrderLinkID string      `json:"orderLinkId"`
	Symbol      SymbolSpot  `json:"symbol"`
	SymbolName  string      `json:"symbolName"`
	Price       string      `json:"price"`
	OrigQty     string      `json:"origQty"`
	ExecutedQty string      `json:"executedQty"`
	OrderType   OrderType   `json:"type"`
	OrderStatus OrderStatus `json:"status"`
	TimeInForce TimeInForce `json:"timeInForce"`
}

type SpotOrderCreated struct {
	SpotOrderBase
	TransactTime string `json:"transactTime"`
}

type SpotOrderCancelled struct {
	SpotOrderCreated
	Side string `json:"side"`
}

type SpotOrderHistory struct {
	SpotOrderBase
	ExchangeId          string `json:"exchangeId"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
}

type SpotOrder struct {
	SpotOrderHistory
	Locked string `json:"locked"`
}

type CancelOrderID struct {
	OrderID string `json:"orderId"`
	Code    int    `json:"code"`
}

type Trade struct {
	ID              string     `json:"id"`
	Symbol          SymbolSpot `json:"symbol"`
	SymbolName      string     `json:"symbolName"`
	OrderID         string     `json:"orderId"`
	TicketID        string     `json:"ticketId"`
	MatchOrderID    string     `json:"matchOrderId"`
	Price           string     `json:"price"`
	Qty             string     `json:"qty"`
	Commission      string     `json:"commission"`
	CommissionAsset string     `json:"commissionAsset"`
	Time            string     `json:"time"`
	IsBuyer         bool       `json:"isBuyer"`
	IsMaker         bool       `json:"isMaker"`
	Fee             TradeFee   `json:"fee"`
	FeeTokenId      string     `json:"feeTokenId"`
	FeeAmount       string     `json:"feeAmount"`
	MakerRebate     string     `json:"makerRebate"`
	ExecutionTime   string     `json:"executionTime"`
}

type TradeFee struct {
	TokenId   string `json:"feeTokenId"`
	TokenName string `json:"feeTokenName"`
	Fee       string `json:"fee"`
}

func (this *Spot) PlaceOrder(symbol SymbolSpot, side Side, typ OrderType, qty int, price int) (SpotOrderCreated, bool) {
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
	param := UrlParam{
		"symbol": symbol,
		"side":   side,
		"type":   typ,
		//"recv_window":   5000000,
		"time_in_force": "GTC",
	}
	if qty > 0 {
		param["qty"] = qty
	}
	if price > 0 {
		param["price"] = price
	}
	resp := &Response[SpotOrderCreated]{}
	err := this.client.Post(this.url("order"), param, resp)
	return resp.Result, err == nil
}

func (this *Spot) GetOrder(orderID string) (SpotOrder, bool) {
	//orderId     string Order ID. Required if not passing orderLinkId
	//orderLinkId string Unique user-set order ID. Required if not passing orderId
	resp := Response[SpotOrder]{}
	err := this.client.Get(this.url("order"), UrlParam{
		"orderId": orderID,
		//"symbol": symbol,
	}, &resp)
	return resp.Result, err == nil
}

func (this *Spot) CancelOrder(orderID string) (SpotOrderCancelled, bool) {
	//orderId     string Order ID. Required if not passing orderLinkId
	//orderLinkId string Unique user-set order ID. Required if not passing orderId
	resp := &Response[SpotOrderCancelled]{}
	err := this.client.Delete(this.url("order"), UrlParam{"orderId": orderID}, resp)
	return resp.Result, err == nil
}

func (this *Spot) FastCancelOrder(symbol SymbolSpot) (bool, bool) {
	//symbolId    Required string Name of the trading pair
	//orderId              string Order ID. Required if not passing orderLinkId
	//orderLinkId          string Unique user-set order ID. Required if not passing orderId
	resp := &Response[struct {
		IsCancelled bool `json:"isCancelled"`
	}]{}
	err := this.client.Delete(this.url("order/fast"), UrlParam{
		"symbolId": symbol,
	}, resp)
	return resp.Result.IsCancelled, err == nil
}

func (this *Spot) BatchCancelOrder(symbol string) (bool, bool) {
	// symbol     Required string Name of the trading pair
	// side                string Order direction
	// orderTypes          string Order type. Use commas to indicate multiple order types, eg LIMIT,LIMIT_MAKER. Default: LIMIT
	resp := &Response[struct {
		Success bool `json:"success"`
	}]{}
	err := this.client.Delete(this.url("order/batch-cancel"), UrlParam{"symbol": symbol}, resp)
	return resp.Result.Success, err == nil
}

func (this *Spot) BatchFastCancelOrder(symbol string) (bool, bool) {
	// symbol     Required string Name of the trading pair
	// side                string Order direction
	// orderTypes          string Order type. Use commas to indicate multiple order types, eg LIMIT,LIMIT_MAKER. Default: LIMIT
	resp := &Response[struct {
		Success bool `json:"success"`
	}]{}
	err := this.client.Delete(this.url("order/batch-fast-cancel"), UrlParam{"symbol": symbol}, resp)
	return resp.Result.Success, err == nil
}

func (this *Spot) BatchCancelOrderByID(ID []string) ([]CancelOrderID, bool) {
	// orderIds Required string Order ID, use commas to indicate multiple orderIds. Maximum of 100 ids.
	resp := &Response[[]CancelOrderID]{}
	err := this.client.Delete(this.url("order/batch-cancel-by-ids"), UrlParam{"orderIds": ID}, resp)
	return resp.Result, err == nil
}

func (this *Spot) OpenOrders() ([]SpotOrderBase, bool) {
	// symbol  string  Name of the trading pair
	// orderId string  Specify orderId to return all the orders that orderId of which are smaller than this particular one for pagination purpose
	// limit   integer Default value is 500, max 500
	resp := Response[[]SpotOrderBase]{}
	err := this.client.Get(this.url("open-orders"), UrlParam{}, &resp)
	return resp.Result, err == nil
}

func (this *Spot) OrderHistory(orderID string) ([]SpotOrderHistory, bool) {
	// symbol    string  Name of the trading pair
	// orderId   string  Specify orderId to return all the orders that orderId of which are smaller than this particular one for pagination purpose
	// limit     integer Default value is 500, max 500
	// startTime long    Start time, unit in millisecond
	// endTime   long    End time, unit in millisecond
	resp := Response[[]SpotOrderHistory]{}
	err := this.client.Get(this.url("history-orders"), UrlParam{"orderId": orderID}, &resp)
	return resp.Result, err == nil
}

func (this *Spot) TradeHistory(symbol string) ([]Trade, bool) {
	// symbol       string  Name of the trading pair
	// limit        integer Default value is 50, max 50
	// fromTicketId integer Query greater than the trade ID. (fromTicketId < trade ID)
	// toTicketId   integer Query smaller than the trade ID. (trade ID < toTicketId)
	// orderId      integer Order ID
	// startTime    long    Start time, unit in millisecond
	// endTime      long    End time, unit in millisecond
	resp := Response[[]Trade]{}
	err := this.client.Get(this.url("myTrades"), UrlParam{"symbol": symbol}, &resp)
	return resp.Result, err == nil
}
