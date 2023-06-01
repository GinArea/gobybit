package bybitv5

import (
	"fmt"

	"github.com/msw-x/moon/ujson"
)

// Get Kline
// https://bybit-exchange.github.io/docs/v5/market/kline
//
//	category Required string  Product type. spot,linear,inverse
//	symbol   Required string  Symbol name
//	interval Required string  Kline interval. 1,3,5,15,30,60,120,240,360,720,D,M,W
//	start             integer The start timestamp (ms)
//	end               integer The end timestamp (ms)
//	limit             integer Limit for data size per page. [1, 200]. Default: 200
type GetKline struct {
	Category Category
	Symbol   string
	Interval Interval
	Start    int `url:",omitempty"`
	End      int `url:",omitempty"`
	Limit    int `url:",omitempty"`
}

type Kline struct {
	StartTime  string
	OpenPrice  string
	HighPrice  string
	LowPrice   string
	ClosePrice string
}

type KlineExt struct {
	Kline
	Volume   string
	Turnover string
}

type klineResult struct {
	Category Category
	Symbol   string
	List     [][]string
}

func getKline[T any](c *Client, path string, q GetKline, unmarshal func([]string) (T, error)) Response[[]T] {
	return GetPub(c.market(), path, q, func(r klineResult) ([]T, error) {
		return transformList(r.List, unmarshal)
	})
}

func UnmarshalKline(s []string) (r Kline, err error) {
	requiredLen := 5
	currentLen := len(s)
	if currentLen == requiredLen {
		r.StartTime = s[0]
		r.OpenPrice = s[1]
		r.HighPrice = s[2]
		r.LowPrice = s[3]
		r.ClosePrice = s[4]
	} else {
		err = fmt.Errorf("kline list len is %d, but required %d", currentLen, requiredLen)
	}
	return
}

func UnmarshalKlineExt(s []string) (r KlineExt, err error) {
	requiredLen := 7
	currentLen := len(s)
	if currentLen == requiredLen {
		r.Kline, err = UnmarshalKline(s[:5])
		if err == nil {
			r.Volume = s[5]
			r.Turnover = s[6]
		}
	} else {
		err = fmt.Errorf("kline list len is %d, but required %d", currentLen, requiredLen)
	}
	return
}

func (o GetKline) Do(c *Client) Response[[]KlineExt] {
	return getKline(c, "kline", o, UnmarshalKlineExt)
}

func (o *Client) GetKline(v GetKline) Response[[]KlineExt] {
	return v.Do(o)
}

// Get Mark Price Kline
// https://bybit-exchange.github.io/docs/v5/market/mark-kline
func (o GetKline) DoMarkPrice(c *Client) Response[[]Kline] {
	return getKline(c, "mark-price-kline", o, UnmarshalKline)
}

func (o *Client) GetMarkPriceKline(v GetKline) Response[[]Kline] {
	return v.DoMarkPrice(o)
}

// Get Index Price Kline
// https://bybit-exchange.github.io/docs/v5/market/index-kline
func (o GetKline) DoIndexPrice(c *Client) Response[[]Kline] {
	return getKline(c, "index-price-kline", o, UnmarshalKline)
}

func (o *Client) GetIndexPriceKline(v GetKline) Response[[]Kline] {
	return v.DoIndexPrice(o)
}

// Get Premium Index Price Kline
// https://bybit-exchange.github.io/docs/v5/market/preimum-index-kline
func (o GetKline) DoPremiumIndexPrice(c *Client) Response[[]Kline] {
	return getKline(c, "premium-index-price-kline", o, UnmarshalKline)
}

func (o *Client) GetPremiumIndexPriceKline(v GetKline) Response[[]Kline] {
	return v.DoPremiumIndexPrice(o)
}

// Get Instruments Info
// https://bybit-exchange.github.io/docs/v5/market/instrument
//
//	category Required string  Product type. spot,linear,inverse
//	symbol            string  Symbol name
//	status            string  Symbol status filter, spot/linear/inverse has Trading only
//	baseCoin          string  Base coin. linear,inverse,option only
//	limit             integer Limit for data size per page. [1, 1000]. Default: 500
//	cursor            string  Cursor. Used for pagination
type GetInstruments struct {
	Category Category
	Symbol   string `url:",omitempty"`
	Status   Status `url:",omitempty"`
	BaseCoin string `url:",omitempty"`
	Limit    int    `url:",omitempty"`
	Cursor   string `url:",omitempty"`
}

type Instrument struct {
	Symbol             string
	ContractType       ContractType
	Status             Status
	BaseCoin           string
	QuoteCoin          string
	LaunchTime         string
	DeliveryTime       string
	DeliveryFeeRate    string
	PriceScale         string
	LeverageFilter     LeverageFilter
	PriceFilter        PriceFilter
	LotSizeFilter      LotSizeFilter
	UnifiedMarginTrade bool
	FundingInterval    int
	SettleCoin         string
}

type LeverageFilter struct {
	MinLeverage  string
	MaxLeverage  string
	LeverageStep string
}

type PriceFilter struct {
	MaxPrice string
	MinPrice string
	TickSize string
}

type LotSizeFilter struct {
	MaxOrderQty         string
	MinOrderQty         string
	QtyStep             string
	PostOnlyMaxOrderQty string
}

func (o GetInstruments) Do(c *Client) Response[[]Instrument] {
	type result struct {
		Category       Category
		List           []Instrument
		NextPageCursor string
	}
	return GetPub(c.market(), "instruments-info", o, func(r result) ([]Instrument, error) {
		return r.List, nil
	})
}

func (o *Client) GetInstruments(v GetInstruments) Response[[]Instrument] {
	return v.Do(o)
}

// Get Orderbook
// https://bybit-exchange.github.io/docs/v5/market/orderbook
//
//	category Required string  Product type. spot,linear,inverse,option
//	symbol   Required string  Symbol name
//	limit             integer Limit size for each bid and ask
//	                  spot: [1, 50]. Default: 1.
//	                  linear&inverse: [1, 200]. Default: 25.
//	                  option: [1, 25]. Default: 1.
type GetOrderbook struct {
	Category Category
	Symbol   string
	Limit    int `url:",omitempty"`
}

type Orderbook struct {
	Symbol    string            `json:"s"`
	Bid       [][]ujson.Float64 `json:"b"`
	Ask       [][]ujson.Float64 `json:"a"`
	Timestamp int               `json:"ts"`
	UpdateId  int               `json:"u"`
}

func (o GetOrderbook) Do(c *Client) Response[Orderbook] {
	return GetPub(c.market(), "orderbook", o, forward[Orderbook])
}

func (o *Client) GetOrderbook(v GetOrderbook) Response[Orderbook] {
	return v.Do(o)
}

// Get Tickers
// https://bybit-exchange.github.io/docs/v5/market/tickers
//
//	category Required string Product type. spot,linear,inverse,option
//	symbol            string Symbol name
//	baseCoin          string Base coin. For option only
//	expDate           string Expiry date. e.g., 25DEC22. For option only
type GetTickers struct {
	Category Category
	Symbol   string `url:",omitempty"`
	BaseCoin string `url:",omitempty"`
	ExpDate  string `url:",omitempty"`
}

type Ticker struct {
	Symbol                 string
	LastPrice              ujson.StringFloat64
	IndexPrice             ujson.StringFloat64
	MarkPrice              ujson.StringFloat64
	PrevPrice24h           ujson.StringFloat64
	Price24hPcnt           ujson.StringFloat64
	HighPrice24h           ujson.StringFloat64
	LowPrice24h            ujson.StringFloat64
	PrevPrice1h            ujson.StringFloat64
	OpenInterest           ujson.StringFloat64
	OpenInterestValue      ujson.StringFloat64
	Turnover24h            ujson.StringFloat64
	Volume24h              ujson.StringFloat64
	FundingRate            ujson.StringFloat64
	NextFundingTime        string
	PredictedDeliveryPrice ujson.StringFloat64
	BasisRate              ujson.StringFloat64
	DeliveryFeeRate        ujson.StringFloat64
	DeliveryTime           string
	Ask1Size               ujson.StringFloat64
	Bid1Price              ujson.StringFloat64
	Ask1Price              ujson.StringFloat64
	Bid1Size               ujson.StringFloat64
	Basis                  string
}

func (o GetTickers) Do(c *Client) Response[[]Ticker] {
	type result struct {
		Category Category
		List     []Ticker
	}
	return GetPub(c.market(), "tickers", o, func(r result) ([]Ticker, error) {
		return r.List, nil
	})
}

func (o *Client) GetTickers(v GetTickers) Response[[]Ticker] {
	return v.Do(o)
}

// Get Funding Rate History
// https://bybit-exchange.github.io/docs/v5/market/history-fund-rate
//
//	category  Required string  Product type. linear,inverse
//	symbol    Required string  Symbol name
//	startTime          integer The start timestamp (ms)
//	endTime            integer The end timestamp (ms)
//	limit              integer Limit for data size per page. [1, 200]. Default: 200
type GetFundingRateHistory struct {
	Category  Category
	Symbol    string
	StartTime int `url:",omitempty"`
	EndTime   int `url:",omitempty"`
	Limit     int `url:",omitempty"`
}

type FundingRateHistory struct {
	Symbol               string
	FundingRate          string
	FundingRateTimestamp string
}

func (o GetFundingRateHistory) Do(c *Client) Response[[]FundingRateHistory] {
	type result struct {
		Category Category
		List     []FundingRateHistory
	}
	return GetPub(c.market(), "funding/history", o, func(r result) ([]FundingRateHistory, error) {
		return r.List, nil
	})
}

func (o *Client) GetFundingRateHistory(v GetFundingRateHistory) Response[[]FundingRateHistory] {
	return v.Do(o)
}

// Get Public Trading History
// https://bybit-exchange.github.io/docs/v5/market/recent-trade
//
//	category   Required string  Product type. spot,linear,inverse,option
//	symbol              string  Symbol name
//	baseCoin            string  Base coin. For option only. If not passed, return BTC data by default
//	optionType          string  Option type. Call or Put. For option only
//	limit               integer
type GetPublicTradingHistory struct {
	Category   Category
	Symbol     string `url:",omitempty"`
	BaseCoin   string `url:",omitempty"`
	OptionType string `url:",omitempty"`
	Limit      int    `url:",omitempty"`
}

type PublicTradingHistory struct {
	ExecId       string
	Symbol       string
	Price        string
	Size         string
	Side         string
	Time         string
	IsBlockTrade bool
}

func (o GetPublicTradingHistory) Do(c *Client) Response[[]PublicTradingHistory] {
	type result struct {
		Category Category
		List     []PublicTradingHistory
	}
	return GetPub(c.market(), "recent-trade", o, func(r result) ([]PublicTradingHistory, error) {
		return r.List, nil
	})
}

func (o *Client) GetPublicTradingHistory(v GetPublicTradingHistory) Response[[]PublicTradingHistory] {
	return v.Do(o)
}

// Get Open Interest
// https://bybit-exchange.github.io/docs/v5/market/open-interest
//
//	category     Required string  Product type. linear,inverse
//	symbol       Required string  Symbol name
//	intervalTime Required string  Interval. 5min,15min,30min,1h,4h,1d
//	startTime             integer The start timestamp (ms)
//	endTime               integer The end timestamp (ms)
//	limit                 integer Limit for data size per page. [1, 200]. Default: 50
//	cursor                string  Cursor. Used for pagination
type GetOpenInterest struct {
	Category     Category
	Symbol       string
	IntervalTime IntervalTime
	StartTime    int    `url:",omitempty"`
	EndTime      int    `url:",omitempty"`
	Limit        int    `url:",omitempty"`
	Cursor       string `url:",omitempty"`
}

type OpenInterest struct {
	OpenInterest string
	Timestamp    string
}

func (o GetOpenInterest) Do(c *Client) Response[[]OpenInterest] {
	type result struct {
		Category Category
		Symbol   string
		List     []OpenInterest
	}
	return GetPub(c.market(), "open-interest", o, func(r result) ([]OpenInterest, error) {
		return r.List, nil
	})
}

func (o *Client) GetOpenInterest(v GetOpenInterest) Response[[]OpenInterest] {
	return v.Do(o)
}

// Get Historical Volatility
// https://bybit-exchange.github.io/docs/v5/market/iv
//
//	category  Required string  Product type. option
//	baseCoin           string  Base coin. Default: return BTC data
//	period             integer Period
//	startTime          integer The start timestamp (ms)
//	endTime            integer The end timestamp (ms)
type GetHistoricalVolatility struct {
	Category  Category
	BaseCoin  string `url:",omitempty"`
	Period    Period `url:",omitempty"`
	StartTime int    `url:",omitempty"`
	EndTime   int    `url:",omitempty"`
}

type HistoricalVolatility struct {
	Period int
	Value  string
	Time   string
}

func (o GetHistoricalVolatility) Do(c *Client) Response[[]HistoricalVolatility] {
	return GetPub(c.market(), "historical-volatility", o, forward[[]HistoricalVolatility])
}

func (o *Client) GetHistoricalVolatility(v GetHistoricalVolatility) Response[[]HistoricalVolatility] {
	return v.Do(o)
}

// Get Insurance
// https://bybit-exchange.github.io/docs/v5/market/insurance
//
//	coin string coin. Default: return all insurance coins
type GetInsurance struct {
	Coin string `url:",omitempty"`
}

type Insurance struct {
	Coin    string
	Balance string
	Value   string
}

func (o GetInsurance) Do(c *Client) Response[[]Insurance] {
	type result struct {
		UpdatedTime string
		List        []Insurance
	}
	return GetPub(c.market(), "insurance", o, func(r result) ([]Insurance, error) {
		return r.List, nil
	})
}

func (o *Client) GetInsurance(v GetInsurance) Response[[]Insurance] {
	return v.Do(o)
}

// Get Risk Limit
// https://bybit-exchange.github.io/docs/v5/market/risk-limit
//
//	category Required string Product type. linear,inverse
//	symbol            string Symbol name
type GetRiskLimit struct {
	Category Category
	Symbol   string `url:",omitempty"`
}

type RiskLimit struct {
	Id                int
	Symbol            string
	RiskLimitValue    string
	MaintenanceMargin string
	InitialMargin     string
	IsLowestRisk      int
	MaxLeverage       string
}

func (o GetRiskLimit) Do(c *Client) Response[[]RiskLimit] {
	type result struct {
		Category Category
		List     []RiskLimit
	}
	return GetPub(c.market(), "risk-limit", o, func(r result) ([]RiskLimit, error) {
		return r.List, nil
	})
}

func (o *Client) GetRiskLimit(v GetRiskLimit) Response[[]RiskLimit] {
	return v.Do(o)
}

// Get Delivery Price
// https://bybit-exchange.github.io/docs/v5/market/delivery-price
//
//	category Required string  Product type. spot,linear,inverse
//	symbol            string  Symbol name
//	baseCoin          string  Base coin. Default: BTC. valid for option only
//	limit             integer Limit for data size per page. [1, 200]. Default: 50
//	cursor            string  Cursor. Used for pagination
type GetDeliveryPrice struct {
	Category Category
	Symbol   string `url:",omitempty"`
	BaseCoin string `url:",omitempty"`
	Limit    int    `url:",omitempty"`
	Cursor   string `url:",omitempty"`
}

type DeliveryPrice struct {
	Symbol        string
	DeliveryPrice string
	DeliveryTime  string
}

func (o GetDeliveryPrice) Do(c *Client) Response[[]DeliveryPrice] {
	type result struct {
		Category       Category
		NextPageCursor string
		List           []DeliveryPrice
	}
	return GetPub(c.market(), "delivery-price", o, func(r result) ([]DeliveryPrice, error) {
		return r.List, nil
	})
}

func (o *Client) GetDeliveryPrice(v GetDeliveryPrice) Response[[]DeliveryPrice] {
	return v.Do(o)
}
