// Market Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketdata)
package iperpetual

// Query Symbol (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querysymbol)
type SymbolInfo struct {
	Name            string         `json:"name"`
	Alias           string         `json:"alias"`
	Status          ContractStatus `json:"status"`
	BaseCurrency    string         `json:"base_currency"`
	QuoteCurrency   string         `json:"quote_currency"`
	PriceScale      int            `json:"price_scale"`
	TakerFee        string         `json:"taker_fee"`
	MakerFee        string         `json:"maker_fee"`
	FundingInterval int            `json:"funding_interval"`
	LeverageFilter  LeverageFilter `json:"leverage_filter"`
	PriceFilter     PriceFilter    `json:"price_filter"`
	LotSizeFilter   LotSizeFilter  `json:"lot_size_filter"`
}

type LeverageFilter struct {
	Min  int    `json:"min_leverage"`
	Max  int    `json:"max_leverage"`
	Step string `json:"leverage_step"`
}

type PriceFilter struct {
	Min      string `json:"min_price"`
	Max      string `json:"max_price"`
	TickSize string `json:"tick_size"`
}

type LotSizeFilter struct {
	MaxTradingQty         float64 `json:"max_trading_qty"`
	MinTradingQty         float64 `json:"min_trading_qty"`
	QtyStep               float64 `json:"qty_step"`
	PostOnlyMaxTradingQty string  `json:"post_only_max_trading_qty"`
}

func (this *Client) QuerySymbol() ([]SymbolInfo, bool) {
	return GetPublic[[]SymbolInfo](this, "symbols", nil)
}

func (this *Client) QuerySymbolNames() ([]string, bool) {
	result, ok := this.QuerySymbol()
	names := make([]string, len(result))
	for n, s := range result {
		names[n] = s.Name
	}
	return names, ok
}

// Order Book (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-orderbook)
type OrderBook struct {
	Symbol string `param:"symbol"`
}

func (this OrderBook) Do(client *Client) ([]OrderBookItem, bool) {
	return GetPublic[[]OrderBookItem](client, "orderBook/L2", this)
}

type OrderBookItem struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Size   int    `json:"size"`
	Side   Side   `json:"side"`
}

func (this *Client) OrderBook(symbol string) ([]OrderBookItem, bool) {
	return OrderBook{Symbol: symbol}.Do(this)
}

// Query Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querykline)
//   symbol    Required string  Symbol
//   interval  Required string  Data refresh interval. Enum : 1 3 5 15 30 60 120 240 360 720 "D" "M" "W"
//   from      Required integer From timestamp in seconds
//   limit              integer Limit for data size per page, max size is 200. Default as showing 200 pieces of data per page
type QueryKline struct {
	Symbol   string        `param:"symbol"`
	Interval KlineInterval `param:"interval"`
	From     int64         `param:"from"`
	Limit    *int          `param:"limit"`
}

func (this QueryKline) Do(client *Client) ([]KlineItem, bool) {
	return GetPublic[[]KlineItem](client, "kline/list", this)
}

type KlineItem struct {
	Symbol   string        `json:"symbol"`
	Interval KlineInterval `json:"interval"`
	OpenTime uint64        `json:"open_time"`
	Open     string        `json:"open"`
	High     string        `json:"high"`
	Low      string        `json:"low"`
	Close    string        `json:"close"`
	Volume   string        `json:"volume"`
	Turnover string        `json:"turnover"`
}

func (this *Client) QueryKline(v QueryKline) ([]KlineItem, bool) {
	return v.Do(this)
}

// Latest Information for Symbol (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-latestsymbolinfo)
type SymbolLatestInformation struct {
	Symbol *string `param:"symbol"`
}

func (this SymbolLatestInformation) Do(client *Client) ([]LatestInformation, bool) {
	return GetPublic[[]LatestInformation](client, "tickers", this)
}

type LatestInformation struct {
	Symbol                 string        `json:"symbol"`
	BidPrice               string        `json:"bid_price"`
	AskPrice               string        `json:"ask_price"`
	LastPrice              string        `json:"last_price"`
	LastTickDirection      TickDirection `json:"last_tick_direction"`
	PrevPrice24h           string        `json:"prev_price_24h"`
	Price24hPcnt           string        `json:"price_24h_pcnt"`
	HighPrice24h           string        `json:"high_price_24h"`
	LowPrice24h            string        `json:"low_price_24h"`
	PrevPrice1h            string        `json:"prev_price_1h"`
	Price1hPcnt            string        `json:"price_1h_pcnt"`
	MarkPrice              string        `json:"mark_price"`
	IndexPrice             string        `json:"index_price"`
	OpenInterest           float64       `json:"open_interest"`
	OpenValue              string        `json:"open_value"`
	TotalTurnover          string        `json:"total_turnover"`
	Turnover24h            string        `json:"turnover_24h"`
	TotalVolume            float64       `json:"total_volume"`
	Volume24h              float64       `json:"volume_24h"`
	FundingRate            string        `json:"funding_rate"`
	PredictedFundingRate   string        `json:"predicted_funding_rate"`
	NextFundingTime        string        `json:"next_funding_time"`
	CountdownHour          int           `json:"countdown_hour"`
	DeliveryFeeRate        string        `json:"delivery_fee_rate"`
	PredictedDeliveryPrice string        `json:"predicted_delivery_price"`
	DeliveryTime           string        `json:"delivery_time"`
}

func (this *Client) SymbolLatestInformation(symbol *string) ([]LatestInformation, bool) {
	return SymbolLatestInformation{Symbol: symbol}.Do(this)
}

// Public Trading Records (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-publictradingrecords)
//   symbol Required string  Symbol
//   limit           integer Limit for data size, max size is 1000. Default size is 500
type PublicTradingRecords struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (this PublicTradingRecords) Do(client *Client) ([]PublicTradingRecord, bool) {
	return GetPublic[[]PublicTradingRecord](client, "trading-records", this)
}

type PublicTradingRecord struct {
	ID     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Qty    int     `json:"qty"`
	Side   Side    `json:"side"`
	Time   string  `json:"time"`
}

func (this *Client) PublicTradingRecords(v PublicTradingRecords) ([]PublicTradingRecord, bool) {
	return v.Do(this)
}

// Query Mark Price Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-markpricekline)
//
// Query mark price kline (like Query Kline but for mark price)
func (this QueryKline) DoMark(client *Client) ([]MarkKlineItem, bool) {
	return GetPublic[[]MarkKlineItem](client, "mark-price-kline", this)
}

type MarkKlineItem struct {
	Symbol   string        `json:"symbol"`
	Interval KlineInterval `json:"period"`
	OpenTime uint64        `json:"start_at"`
	Open     int           `json:"open"`
	High     int           `json:"high"`
	Low      int           `json:"low"`
	Close    int           `json:"close"`
}

func (this *Client) QueryMarkKline(v QueryKline) ([]MarkKlineItem, bool) {
	return v.DoMark(this)
}

// Query Index Price Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-queryindexpricekline)
//
// Index price kline. Tracks BTC spot prices, with a frequency of every second
func (this QueryKline) DoIndex(client *Client) ([]IndexKlineItem, bool) {
	return GetPublic[[]IndexKlineItem](client, "index-price-kline", this)
}

type IndexKlineItem struct {
	Symbol   string        `json:"symbol"`
	Interval KlineInterval `json:"period"`
	OpenTime uint64        `json:"open_time"`
	Open     string        `json:"open"`
	High     string        `json:"high"`
	Low      string        `json:"low"`
	Close    string        `json:"close"`
}

func (this *Client) QueryIndexKline(v QueryKline) ([]IndexKlineItem, bool) {
	return v.DoIndex(this)
}

// Query Premium Index Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querypremiumindexkline)
//
// Premium index kline. Tracks the premium / discount of BTC perpetual contracts relative to the mark price per minute
func (this QueryKline) DoPremium(client *Client) ([]IndexKlineItem, bool) {
	return GetPublic[[]IndexKlineItem](client, "premium-index-kline", this)
}

func (this *Client) QueryPremiumKline(v QueryKline) ([]IndexKlineItem, bool) {
	return v.DoPremium(this)
}
