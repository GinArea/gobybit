// Market Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketdata)
package iperpetual

import (
	"errors"

	"github.com/ginarea/gobybit/transport"
)

// Query Symbol (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querysymbol)
type SymbolInfo struct {
	Name            string            `json:"name"`
	Alias           string            `json:"alias"`
	Status          ContractStatus    `json:"status"`
	BaseCurrency    string            `json:"base_currency"`
	QuoteCurrency   string            `json:"quote_currency"`
	PriceScale      float64           `json:"price_scale"`
	TakerFee        transport.Float64 `json:"taker_fee"`
	MakerFee        transport.Float64 `json:"maker_fee"`
	FundingInterval float64           `json:"funding_interval"`
	LeverageFilter  LeverageFilter    `json:"leverage_filter"`
	PriceFilter     PriceFilter       `json:"price_filter"`
	LotSizeFilter   LotSizeFilter     `json:"lot_size_filter"`
}

type LeverageFilter struct {
	Min  int               `json:"min_leverage"`
	Max  int               `json:"max_leverage"`
	Step transport.Float64 `json:"leverage_step"`
}

type PriceFilter struct {
	Min      transport.Float64 `json:"min_price"`
	Max      transport.Float64 `json:"max_price"`
	TickSize transport.Float64 `json:"tick_size"`
}

type LotSizeFilter struct {
	MaxTradingQty         float64           `json:"max_trading_qty"`
	MinTradingQty         float64           `json:"min_trading_qty"`
	QtyStep               float64           `json:"qty_step"`
	PostOnlyMaxTradingQty transport.Float64 `json:"post_only_max_trading_qty"`
}

func (o *Client) QuerySymbol() ([]SymbolInfo, error) {
	return GetPublic[[]SymbolInfo](o, "symbols", nil)
}

func (o *Client) QuerySymbolNames() ([]string, error) {
	result, err := o.QuerySymbol()
	names := make([]string, len(result))
	for n, s := range result {
		names[n] = s.Name
	}
	return names, err
}

// Order Book (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-orderbook)
type OrderBook struct {
	Symbol string `param:"symbol"`
}

func (o OrderBook) Do(client *Client) ([]OrderBookItem, error) {
	return GetPublic[[]OrderBookItem](client, "orderBook/L2", o)
}

type OrderBookItem struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Size   int    `json:"size"`
	Side   Side   `json:"side"`
}

func (o *Client) OrderBook(symbol string) ([]OrderBookItem, error) {
	return OrderBook{Symbol: symbol}.Do(o)
}

// Query Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querykline)
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

func (o QueryKline) Do(client *Client) ([]KlineItem, error) {
	return GetPublic[[]KlineItem](client, "kline/list", o)
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

func (o *Client) QueryKline(v QueryKline) ([]KlineItem, error) {
	return v.Do(o)
}

// Latest Information for Symbol (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-latestsymbolinfo)
type SymbolLatestInformation struct {
	Symbol *string `param:"symbol"`
}

func (o SymbolLatestInformation) Do(client *Client) ([]LatestInformation, error) {
	return GetPublic[[]LatestInformation](client, "tickers", o)
}

type LatestInformation struct {
	Symbol                 string            `json:"symbol"`
	BidPrice               transport.Float64 `json:"bid_price"`
	AskPrice               transport.Float64 `json:"ask_price"`
	LastPrice              transport.Float64 `json:"last_price"`
	LastTickDirection      TickDirection     `json:"last_tick_direction"`
	PrevPrice24h           transport.Float64 `json:"prev_price_24h"`
	Price24hPcnt           transport.Float64 `json:"price_24h_pcnt"`
	HighPrice24h           transport.Float64 `json:"high_price_24h"`
	LowPrice24h            transport.Float64 `json:"low_price_24h"`
	PrevPrice1h            transport.Float64 `json:"prev_price_1h"`
	Price1hPcnt            transport.Float64 `json:"price_1h_pcnt"`
	MarkPrice              transport.Float64 `json:"mark_price"`
	IndexPrice             transport.Float64 `json:"index_price"`
	OpenInterest           float64           `json:"open_interest"`
	OpenValue              transport.Float64 `json:"open_value"`
	TotalTurnover          transport.Float64 `json:"total_turnover"`
	Turnover24h            transport.Float64 `json:"turnover_24h"`
	TotalVolume            float64           `json:"total_volume"`
	Volume24h              float64           `json:"volume_24h"`
	FundingRate            transport.Float64 `json:"funding_rate"`
	PredictedFundingRate   transport.Float64 `json:"predicted_funding_rate"`
	NextFundingTime        transport.Time    `json:"next_funding_time"`
	CountdownHour          int               `json:"countdown_hour"`
	DeliveryFeeRate        string            `json:"delivery_fee_rate"`
	PredictedDeliveryPrice string            `json:"predicted_delivery_price"`
	DeliveryTime           string            `json:"delivery_time"`
}

func (o *Client) SymbolLatestInformation(symbol *string) ([]LatestInformation, error) {
	return SymbolLatestInformation{Symbol: symbol}.Do(o)
}

func (o *Client) OneSymbolLatestInformation(symbol string) (i LatestInformation, err error) {
	ret, err := o.SymbolLatestInformation(&symbol)
	if err == nil {
		if len(ret) == 1 {
			i = ret[0]
		} else {
			err = errors.New("symbol latest len != 1")
		}
	}
	return
}

// Public Trading Records (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-publictradingrecords)
//
//	symbol Required string  Symbol
//	limit           integer Limit for data size, max size is 1000. Default size is 500
type PublicTradingRecords struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (o PublicTradingRecords) Do(client *Client) ([]PublicTradingRecord, error) {
	return GetPublic[[]PublicTradingRecord](client, "trading-records", o)
}

type PublicTradingRecord struct {
	ID     int     `json:"id"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Qty    int     `json:"qty"`
	Side   Side    `json:"side"`
	Time   string  `json:"time"`
}

func (o *Client) PublicTradingRecords(v PublicTradingRecords) ([]PublicTradingRecord, error) {
	return v.Do(o)
}

// Query Mark Price Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-markpricekline)
//
// Query mark price kline (like Query Kline but for mark price)
func (o QueryKline) DoMark(client *Client) ([]MarkKlineItem, error) {
	return GetPublic[[]MarkKlineItem](client, "mark-price-kline", o)
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

func (o *Client) QueryMarkKline(v QueryKline) ([]MarkKlineItem, error) {
	return v.DoMark(o)
}

// Query Index Price Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-queryindexpricekline)
//
// Index price kline. Tracks BTC spot prices, with a frequency of every second
func (o QueryKline) DoIndex(client *Client) ([]IndexKlineItem, error) {
	return GetPublic[[]IndexKlineItem](client, "index-price-kline", o)
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

func (o *Client) QueryIndexKline(v QueryKline) ([]IndexKlineItem, error) {
	return v.DoIndex(o)
}

// Query Premium Index Kline (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querypremiumindexkline)
//
// Premium index kline. Tracks the premium / discount of BTC perpetual contracts relative to the mark price per minute
func (o QueryKline) DoPremium(client *Client) ([]IndexKlineItem, error) {
	return GetPublic[[]IndexKlineItem](client, "premium-index-kline", o)
}

func (o *Client) QueryPremiumKline(v QueryKline) ([]IndexKlineItem, error) {
	return v.DoPremium(o)
}
