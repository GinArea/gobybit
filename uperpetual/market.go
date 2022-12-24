// Market Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-marketdata)
package uperpetual

// Query Symbol (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-querysymbol)
// using iperpetual

// Order Book (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-orderbook)
// using iperpetual

// Query Kline (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-querykline)
//
//	symbol    Required string  Symbol
//	interval  Required string  Data refresh interval. Enum : 1 3 5 15 30 60 120 240 360 720 "D" "M" "W"
//	from      Required integer From timestamp in seconds
//	limit              integer Limit for data size per page, max size is 200. Default as showing 200 pieces of data per page
type QueryKline struct {
	Symbol   string        `param:"symbol"`
	Interval KlineInterval `param:"interval"`
	From     int64         `param:"from"`
	Limit    *int          `param:"limit"`
}

func (this QueryKline) Do(client *Client) ([]KlineItem, bool) {
	return GetPublic[[]KlineItem](client, "kline", this)
}

type KlineItem struct {
	ID       int           `json:"id"`
	Symbol   string        `json:"symbol"`
	Period   KlineInterval `json:"period"`
	StartAt  uint64        `json:"start_at"`
	Volume   float64       `json:"volume"`
	Open     string        `json:"open"`
	High     string        `json:"high"`
	Low      string        `json:"low"`
	Close    string        `json:"close"`
	Interval KlineInterval `json:"interval"`
	OpenTime uint64        `json:"open_time"`
	Turnover string        `json:"turnover"`
}

func (this *Client) QueryKline(v QueryKline) ([]KlineItem, bool) {
	return v.Do(this)
}

// Latest Information for Symbol (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-latestsymbolinfo)
// using iperpetual

// Public Trading Records (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-publictradingrecords)
//
//	symbol Required string  Symbol
//	limit           integer Limit for data size, max size is 1000. Default size is 500
type PublicTradingRecords struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (this PublicTradingRecords) Do(client *Client) ([]PublicTradingRecord, bool) {
	return GetPublic[[]PublicTradingRecord](client, "recent-trading-records", this)
}

type PublicTradingRecord struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Price        float64 `json:"price"`
	Qty          float64 `json:"qty"`
	Side         Side    `json:"side"`
	Time         string  `json:"time"`
	TradeTime    uint64  `json:"trade_time_ms"`
	IsBlockTrade bool    `json:"is_block_trade"`
}

func (this *Client) PublicTradingRecords(v PublicTradingRecords) ([]PublicTradingRecord, bool) {
	return v.Do(this)
}

// Liquidated Orders (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-query_liqrecords)

// Get the Last Funding Rate (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-fundingrate)
//
// The funding rate is generated every 8 hours at 00:00 UTC, 08:00 UTC and 16:00 UTC.
// For example, if a request is sent at 12:00 UTC, the funding rate generated earlier that day at 08:00 UTC will be sent
type GetLastFundingRate struct {
	Symbol string `param:"symbol"`
}

func (this GetLastFundingRate) Do(client *Client) (LastFundingRate, bool) {
	return GetPublic[LastFundingRate](client, "funding/prev-funding-rate", this)
}

type LastFundingRate struct {
	Symbol      string  `json:"symbol"`
	FundingRate float64 `json:"funding_rate"`
	Timestamp   string  `json:"funding_rate_timestamp"`
}

func (this *Client) GetLastFundingRate(symbol string) (LastFundingRate, bool) {
	return GetLastFundingRate{Symbol: symbol}.Do(this)
}

// Query Mark Price Kline (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-markpricekline)
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

// Query Index Price Kline (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-queryindexpricekline)
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

// Query Premium Index Kline (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-querypremiumindexkline)
//
// Premium index kline. Tracks the premium / discount of BTC perpetual contracts relative to the mark price per minute
func (this QueryKline) DoPremium(client *Client) ([]IndexKlineItem, bool) {
	return GetPublic[[]IndexKlineItem](client, "premium-index-kline", this)
}

func (this *Client) QueryPremiumKline(v QueryKline) ([]IndexKlineItem, bool) {
	return v.DoPremium(this)
}
