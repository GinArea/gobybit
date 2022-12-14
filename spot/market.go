// Market Data Endpoints (https://bybit-exchange.github.io/docs/spot/v1/#t-marketdata)
package spot

// Query Symbol (https://bybit-exchange.github.io/docs/spot/v1/#t-spot_querysymbol)
type SymbolInfo struct {
	Name              string `json:"name"`
	Alias             string `json:"alias"`
	BaseCurrency      string `json:"baseCurrency"`
	QuoteCurrency     string `json:"quoteCurrency"`
	BasePrecision     string `json:"basePrecision"`
	QuotePrecision    string `json:"quotePrecision"`
	MinTradeQuantity  string `json:"minTradeQuantity"`
	MinTradeAmount    string `json:"minTradeAmount"`
	MaxTradeQuantity  string `json:"maxTradeQuantity"`
	MaxTradeAmount    string `json:"maxTradeAmount"`
	MinPricePrecision string `json:"minPricePrecision"`
	Category          int    `json:"category"`
	ShowStatus        bool   `json:"showStatus"`
	Innovation        bool   `json:"innovation"`
}

func (this *Client) QuerySymbol() ([]SymbolInfo, error) {
	return GetPublic[[]SymbolInfo](this, "symbols", nil)
}

func (this *Client) QuerySymbolNames() ([]string, error) {
	result, err := this.QuerySymbol()
	names := make([]string, len(result))
	for n, s := range result {
		names[n] = s.Name
	}
	return names, err
}

// Order Book (https://bybit-exchange.github.io/docs/spot/v1/#t-orderbook)
//
//	symbol Required string  Name of the trading pair
//	limit           integer Default value is 100
type OrderBook struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (this OrderBook) Do(client *Client) (OrderBookResult, error) {
	return GetQuote[OrderBookResult](client, "depth", this)
}

type OrderBookResult struct {
	Time uint64 `json:"time"`
	Bids [][]string
	Asks [][]string
}

func (this *Client) OrderBook(v OrderBook) (OrderBookResult, error) {
	return v.Do(this)
}

// Merged Order Book (https://bybit-exchange.github.io/docs/spot/v1/#t-mergedorderbook)
//
//	symbol Required string  Name of the trading pair
//	scale           int     Precision of the merged orderbook, 1 means 1 digit
//	limit           integer Default value is 100
type MergedOrderBook struct {
	Symbol string `param:"symbol"`
	Scale  *int   `param:"scale"`
	Limit  *int   `param:"limit"`
}

func (this MergedOrderBook) Do(client *Client) (OrderBookResult, error) {
	return GetQuote[OrderBookResult](client, "depth/merged", this)
}

func (this *Client) MergedOrderBook(v MergedOrderBook) (OrderBookResult, error) {
	return v.Do(this)
}

// Public Trading Records (https://bybit-exchange.github.io/docs/spot/v1/#t-publictradingrecords)
//
//	symbol Required string  Name of the trading pair
//	limit           integer Default value is 60, max 60
type PublicTradingRecords struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (this PublicTradingRecords) Do(client *Client) ([]PublicTradingRecord, error) {
	return GetQuote[[]PublicTradingRecord](client, "trades", this)
}

type PublicTradingRecord struct {
	Price        string `json:"price"`
	Time         uint64 `json:"time"`
	Qty          string `json:"qty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

func (this *Client) PublicTradingRecords(v PublicTradingRecords) ([]PublicTradingRecord, error) {
	return v.Do(this)
}

// Query Kline (https://bybit-exchange.github.io/docs/spot/v1/#t-querykline)
//
//	symbol    Required string  Name of the trading pair
//	interval  Required string  Chart interval
//	limit              integer Default value is 1000, max 1000
//	startTime          number  Start time, unit in millisecond
//	endTime            number  End time, unit in millisecond
type QueryKline struct {
	Symbol    string        `param:"symbol"`
	Interval  KlineInterval `param:"interval"`
	Limit     *int          `param:"limit"`
	StartTime *int          `param:"startTime"`
	EndTime   *int          `param:"endTime"`
}

func (this QueryKline) Do(client *Client) ([][]any, error) {
	return GetQuote[[][]any](client, "kline", this)
}

func (this *Client) QueryKline(v QueryKline) ([][]any, error) {
	return v.Do(this)
}

// Latest Information for Symbol (https://bybit-exchange.github.io/docs/spot/v1/#t-spot_latestsymbolinfo)
type SymbolLatestInformation struct {
	Symbol *string `param:"symbol"`
}

func (this SymbolLatestInformation) Do(client *Client) ([]LatestInformation, error) {
	path := "ticker/24hr"
	if this.Symbol == nil {
		return GetQuote[[]LatestInformation](client, path, this)
	}
	r, err := GetQuote[LatestInformation](client, path, this)
	return []LatestInformation{r}, err
}

type LatestInformation struct {
	Time         uint64 `json:"time"`
	Symbol       string `json:"symbol"`
	BestBidPrice string `json:"bestBidPrice"`
	BestAskPrice string `json:"bestAskPrice"`
	Volume       string `json:"volume"`
	QuoteVolume  string `json:"quoteVolume"`
	LastPrice    string `json:"lastPrice"`
	HighPrice    string `json:"highPrice"`
	LowPrice     string `json:"lowPrice"`
	OpenPrice    string `json:"openPrice"`
}

func (this *Client) SymbolLatestInformation(symbol *string) ([]LatestInformation, error) {
	return SymbolLatestInformation{Symbol: symbol}.Do(this)
}

// Last Traded Price (https://bybit-exchange.github.io/docs/spot/v1/#t-lasttradedprice)
type LastTradedPrice struct {
	Symbol string `param:"symbol"`
}

func (this LastTradedPrice) Do(client *Client) (SymbolPrice, error) {
	return GetQuote[SymbolPrice](client, "ticker/price", this)
}

type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (this *Client) LastTradedPrice(symbol string) (SymbolPrice, error) {
	return LastTradedPrice{Symbol: symbol}.Do(this)
}

// Best Bid/Ask Price (https://bybit-exchange.github.io/docs/spot/v1/#t-bestbidask)
type BestBidAskPrice struct {
	Symbol string `param:"symbol"`
}

func (this BestBidAskPrice) Do(client *Client) (BestBidAskPriceResult, error) {
	return GetQuote[BestBidAskPriceResult](client, "ticker/book_ticker", this)
}

type BestBidAskPriceResult struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     uint64 `json:"time"`
}

func (this *Client) BestBidAskPrice(symbol string) (BestBidAskPriceResult, error) {
	return BestBidAskPrice{Symbol: symbol}.Do(this)
}
