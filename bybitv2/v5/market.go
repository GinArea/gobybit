package v5

import "fmt"

// Get Kline (https://bybit-exchange.github.io/docs/v5/market/kline)
//
// category Required string  Product type. spot,linear,inverse
// symbol   Required string  Symbol name
// interval Required string  Kline interval. 1,3,5,15,30,60,120,240,360,720,D,M,W
// start             integer The start timestamp (ms)
// end               integer The end timestamp (ms)
// limit             integer Limit for data size per page. [1, 200]. Default: 200
type GetKline struct {
	Category Category `param:"category"`
	Symbol   string   `param:"symbol"`
	Interval Interval `param:"interval"`
	Start    *int     `param:"start"`
	End      *int     `param:"end"`
	Limit    *int     `param:"limit"`
}

func (o GetKline) Do(c *Client) ([]KlineExt, error) {
	return getKline(c, "market/kline", o, UnmarshalKlineExt)
}

func getKlineRaw(c *Client, path string, q GetKline) (list [][]string, err error) {
	type result struct {
		Category Category
		Symbol   string
		List     [][]string
	}
	var r result
	r, err = Get[result](c, path, q)
	list = r.List
	return
}

func getKline[T any](c *Client, path string, q GetKline, unmarshal func([]string) (T, error)) (v []T, err error) {
	var l [][]string
	l, err = getKlineRaw(c, path, q)
	if err == nil {
		v = make([]T, len(l))
		for n, s := range l {
			v[n], err = unmarshal(s)
			if err != nil {
				break
			}
		}
	}
	return
}

type Kline struct {
	StartTime  string
	OpenPrice  string
	HighPrice  string
	LowPrice   string
	ClosePrice string
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

type KlineExt struct {
	Kline
	Volume   string
	Turnover string
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

func (o *Client) GetKline(v GetKline) ([]KlineExt, error) {
	return v.Do(o)
}

// Get Mark Price Kline (https://bybit-exchange.github.io/docs/v5/market/mark-kline)
func (o GetKline) DoMarkPrice(c *Client) ([]Kline, error) {
	return getKline(c, "market/mark-price-kline", o, UnmarshalKline)
}

func (o *Client) GetMarkPriceKline(v GetKline) ([]Kline, error) {
	return v.DoMarkPrice(o)
}

// Get Index Price Kline (https://bybit-exchange.github.io/docs/v5/market/index-kline)
func (o GetKline) DoIndexPrice(c *Client) ([]Kline, error) {
	return getKline(c, "market/index-price-kline", o, UnmarshalKline)
}

func (o *Client) GetIndexPriceKline(v GetKline) ([]Kline, error) {
	return v.DoIndexPrice(o)
}

// Get Premium Index Price Kline (https://bybit-exchange.github.io/docs/v5/market/preimum-index-kline)
func (o GetKline) DoPremiumIndexPrice(c *Client) ([]Kline, error) {
	return getKline(c, "market/premium-index-price-kline", o, UnmarshalKline)
}

func (o *Client) GetPremiumIndexPriceKline(v GetKline) ([]Kline, error) {
	return v.DoPremiumIndexPrice(o)
}

// Get Instruments Info (https://bybit-exchange.github.io/docs/v5/market/instrument)
//
// category Required string  Product type. spot,linear,inverse
// symbol            string  Symbol name
// status            string  Symbol status filter, spot/linear/inverse has Trading only
// baseCoin          string  Base coin. linear,inverse,option only
// limit             integer Limit for data size per page. [1, 1000]. Default: 500
// cursor            string  Cursor. Used for pagination
type GetInstruments struct {
	Category Category `param:"category"`
	Symbol   *string  `param:"symbol"`
	Status   *Status  `param:"status"`
	BaseCoin *string  `param:"baseCoin"`
	Limit    *int     `param:"limit"`
	Cursor   *string  `param:"cursor"`
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

func (o GetInstruments) Do(c *Client) ([]Instrument, error) {
	type result struct {
		Category       Category
		List           []Instrument
		NextPageCursor string
	}
	r, err := Get[result](c, "market/instruments-info", o)
	return r.List, err
}

func (o *Client) GetInstruments(v GetInstruments) ([]Instrument, error) {
	return v.Do(o)
}

// Get Orderbook (https://bybit-exchange.github.io/docs/v5/market/orderbook)
//
// category Required string  Product type. spot,linear,inverse,option
// symbol   Required string  Symbol name
// limit             integer Limit size for each bid and ask
//
//	spot: [1, 50]. Default: 1.
//	linear&inverse: [1, 200]. Default: 25.
//	option: [1, 25]. Default: 1.
type GetOrderbook struct {
	Category Category `param:"category"`
	Symbol   string   `param:"symbol"`
	Limit    *int     `param:"limit"`
}

type Orderbook struct {
	Symbol    string     `json:"s"`
	Bid       [][]string `json:"b"`
	Ask       [][]string `json:"a"`
	Timestamp int        `json:"ts"`
	UpdateID  int        `json:"u"`
}

func (o GetOrderbook) Do(c *Client) (Orderbook, error) {
	return Get[Orderbook](c, "market/orderbook", o)
}

func (o *Client) GetOrderbook(v GetOrderbook) (Orderbook, error) {
	return v.Do(o)
}

// Get Tickers (https://bybit-exchange.github.io/docs/v5/market/tickers)
//
// category Required string Product type. spot,linear,inverse,option
// symbol            string Symbol name
// baseCoin          string Base coin. For option only
// expDate           string Expiry date. e.g., 25DEC22. For option only
type GetTickers struct {
	Category Category `param:"category"`
	Symbol   *string  `param:"symbol"`
	BaseCoin *string  `param:"baseCoin"`
	ExpDate  *string  `param:"expDate"`
}

type Ticker struct {
	Symol                  string
	LastPrice              string
	IndexPrice             string
	MarkPrice              string
	PrevPrice24h           string
	Price24hPcnt           string
	HighPrice24h           string
	LowPrice24h            string
	PrevPrice1h            string
	OpenInterest           string
	OpenInterestValue      string
	Turnover24h            string
	Volume24h              string
	FundingRate            string
	NextFundingTime        string
	PredictedDeliveryPrice string
	BasisRate              string
	DeliveryFeeRate        string
	DeliveryTime           string
	Ask1Size               string
	Bid1Price              string
	Ask1Price              string
	Bid1Size               string
	Basis                  string
}

func (o GetTickers) Do(c *Client) ([]Ticker, error) {
	type result struct {
		Category Category
		List     []Ticker
	}
	r, err := Get[result](c, "market/tickers", o)
	return r.List, err
}

func (o *Client) GetTickers(v GetTickers) ([]Ticker, error) {
	return v.Do(o)
}

// Get Funding Rate History (https://bybit-exchange.github.io/docs/v5/market/history-fund-rate)
//
// category  Required string  Product type. linear,inverse
// symbol    Required string  Symbol name
// startTime          integer The start timestamp (ms)
// endTime            integer The end timestamp (ms)
// limit              integer Limit for data size per page. [1, 200]. Default: 200
type GetFundingRateHistory struct {
	Category  Category `param:"category"`
	Symbol    string   `param:"symbol"`
	StartTime *int     `param:"startTime"`
	EndTime   *int     `param:"endTime"`
	Limit     *int     `param:"limit"`
}

type FundingRateHistory struct {
	Symbol               string
	FundingRate          string
	FundingRateTimestamp string
}

func (o GetFundingRateHistory) Do(c *Client) ([]FundingRateHistory, error) {
	type result struct {
		Category Category
		List     []FundingRateHistory
	}
	r, err := Get[result](c, "market/funding/history", o)
	return r.List, err
}

func (o *Client) GetFundingRateHistory(v GetFundingRateHistory) ([]FundingRateHistory, error) {
	return v.Do(o)
}

// Get Public Trading History (https://bybit-exchange.github.io/docs/v5/market/recent-trade)
//
// category   Required string  Product type. spot,linear,inverse,option
// symbol              string  Symbol name
// baseCoin            string  Base coin. For option only. If not passed, return BTC data by default
// optionType          string  Option type. Call or Put. For option only
// limit               integer
type GetPublicTradingHistory struct {
	Category   Category `param:"category"`
	Symbol     *string  `param:"symbol"`
	BaseCoin   *string  `param:"baseCoin"`
	OptionType *string  `param:"optionType"`
	Limit      *int     `param:"limit"`
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

func (o GetPublicTradingHistory) Do(c *Client) ([]PublicTradingHistory, error) {
	type result struct {
		Category Category
		List     []PublicTradingHistory
	}
	r, err := Get[result](c, "market/recent-trade", o)
	return r.List, err
}

func (o *Client) GetPublicTradingHistory(v GetPublicTradingHistory) ([]PublicTradingHistory, error) {
	return v.Do(o)
}

// Get Open Interest (https://bybit-exchange.github.io/docs/v5/market/open-interest)
//
// category     Required string  Product type. linear,inverse
// symbol       Required string  Symbol name
// intervalTime Required string  Interval. 5min,15min,30min,1h,4h,1d
// startTime             integer The start timestamp (ms)
// endTime               integer The end timestamp (ms)
// limit                 integer Limit for data size per page. [1, 200]. Default: 50
// cursor                string  Cursor. Used for pagination
type GetOpenInterest struct {
	Category     Category     `param:"category"`
	Symbol       string       `param:"symbol"`
	IntervalTime IntervalTime `param:"intervalTime"`
	StartTime    *int         `param:"startTime"`
	EndTime      *int         `param:"endTime"`
	Limit        *int         `param:"limit"`
	Cursor       *string      `param:"cursor"`
}

type OpenInterest struct {
	OpenInterest string
	Timestamp    string
}

func (o GetOpenInterest) Do(c *Client) ([]OpenInterest, error) {
	type result struct {
		Category Category
		Symbol   string
		List     []OpenInterest
	}
	r, err := Get[result](c, "market/open-interest", o)
	return r.List, err
}

func (o *Client) GetOpenInterest(v GetOpenInterest) ([]OpenInterest, error) {
	return v.Do(o)
}

// Get Historical Volatility (https://bybit-exchange.github.io/docs/v5/market/iv)
//
// category  Required string  Product type. option
// baseCoin           string  Base coin. Default: return BTC data
// period             integer Period
// startTime          integer The start timestamp (ms)
// endTime            integer The end timestamp (ms)
type GetHistoricalVolatility struct {
	Category  Category `param:"category"`
	BaseCoin  *string  `param:"baseCoin"`
	Period    *Period  `param:"period"`
	StartTime *int     `param:"startTime"`
	EndTime   *int     `param:"endTime"`
}

type HistoricalVolatility struct {
	Period int
	Value  string
	Time   string
}

func (o GetHistoricalVolatility) Do(c *Client) ([]HistoricalVolatility, error) {
	return Get[[]HistoricalVolatility](c, "market/historical-volatility", o)
}

func (o *Client) GetHistoricalVolatility(v GetHistoricalVolatility) ([]HistoricalVolatility, error) {
	return v.Do(o)
}

// Get Insurance (https://bybit-exchange.github.io/docs/v5/market/insurance)
//
// coin string coin. Default: return all insurance coins
type GetInsurance struct {
	Coin *string `param:"coin"`
}

type Insurance struct {
	Coin    string
	Balance string
	Value   string
}

func (o GetInsurance) Do(c *Client) ([]Insurance, error) {
	type result struct {
		UpdatedTime string
		List        []Insurance
	}
	r, err := Get[result](c, "market/insurance", o)
	return r.List, err
}

func (o *Client) GetInsurance(v GetInsurance) ([]Insurance, error) {
	return v.Do(o)
}

// Get Risk Limit (https://bybit-exchange.github.io/docs/v5/market/risk-limit)
//
// category Required string  Product type. linear,inverse
// symbol            string  Symbol name
type GetRiskLimit struct {
	Category Category `param:"category"`
	Symbol   *string  `param:"symbol"`
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

func (o GetRiskLimit) Do(c *Client) ([]RiskLimit, error) {
	type result struct {
		Category Category
		List     []RiskLimit
	}
	r, err := Get[result](c, "market/risk-limit", o)
	return r.List, err
}

func (o *Client) GetRiskLimit(v GetRiskLimit) ([]RiskLimit, error) {
	return v.Do(o)
}

// Get Delivery Price (https://bybit-exchange.github.io/docs/v5/market/delivery-price)
//
// category Required string  Product type. spot,linear,inverse
// symbol            string  Symbol name
// baseCoin          string  Base coin. Default: BTC. valid for option only
// limit             integer Limit for data size per page. [1, 200]. Default: 50
// cursor            string  Cursor. Used for pagination
type GetDeliveryPrice struct {
	Category Category `param:"category"`
	Symbol   *string  `param:"symbol"`
	BaseCoin *string  `param:"baseCoin"`
	Limit    *int     `param:"limit"`
	Cursor   *string  `param:"cursor"`
}

type DeliveryPrice struct {
	Symbol        string
	DeliveryPrice string
	DeliveryTime  string
}

func (o GetDeliveryPrice) Do(c *Client) ([]DeliveryPrice, error) {
	type result struct {
		Category       Category
		NextPageCursor string
		List           []DeliveryPrice
	}
	r, err := Get[result](c, "market/delivery-price", o)
	return r.List, err
}

func (o *Client) GetDeliveryPrice(v GetDeliveryPrice) ([]DeliveryPrice, error) {
	return v.Do(o)
}
