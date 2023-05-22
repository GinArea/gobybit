// Market Data Endpoints (https://bybit-exchange.github.io/docs/spot/v3/#t-marketdata)
package spotv3

import (
	"errors"

	"github.com/ginarea/gobybit/transport"
)

// Query Symbol (https://bybit-exchange.github.io/docs/spot/v3/#t-spot_querysymbol)
type SymbolInfo struct {
	Name              string            `json:"name"`
	Alias             string            `json:"alias"`
	BaseCoin          string            `json:"baseCoin"`
	QuoteCoin         string            `json:"quoteCoin"`
	BasePrecision     transport.Float64 `json:"basePrecision"`
	QuotePrecision    transport.Float64 `json:"quotePrecision"`
	MinTradeQty       transport.Float64 `json:"minTradeQty"`
	MinTradeAmt       transport.Float64 `json:"minTradeAmt"`
	MaxTradeQty       transport.Float64 `json:"maxTradeQty"`
	MaxTradeAmt       transport.Float64 `json:"maxTradeAmt"`
	MinPricePrecision transport.Float64 `json:"minPricePrecision"`
	Category          string            `json:"category"`
	ShowStatus        string            `json:"showStatus"`
	Innovation        string            `json:"innovation"`
}

func (this *Client) QuerySymbol() ([]SymbolInfo, error) {
	type result struct {
		List []SymbolInfo `json:"list"`
	}
	r, err := GetPublic[result](this, "symbols", nil)
	return r.List, err
}

func (this *Client) QuerySymbolNames() ([]string, error) {
	result, err := this.QuerySymbol()
	names := make([]string, len(result))
	for n, s := range result {
		names[n] = s.Name
	}
	return names, err
}

// Order Book (https://bybit-exchange.github.io/docs/spot/v3/#t-orderbook)
//
//	symbol Required string  Name of the trading pair
//	limit           integer Default value is 100
type OrderBook struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (this OrderBook) Do(client *Client) (OrderBookResult, error) {
	return GetPublic[OrderBookResult](client, "quote/depth", this)
}

type OrderBookResult struct {
	Time uint64 `json:"time"`
	Bids [][]string
	Asks [][]string
}

func (this *Client) OrderBook(v OrderBook) (OrderBookResult, error) {
	return v.Do(this)
}

// Merged Order Book (https://bybit-exchange.github.io/docs/spot/v3/#t-mergedorderbook)
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
	return GetPublic[OrderBookResult](client, "quote/depth/merged", this)
}

func (this *Client) MergedOrderBook(v MergedOrderBook) (OrderBookResult, error) {
	return v.Do(this)
}

// Public Trading Records (https://bybit-exchange.github.io/docs/spot/v3/#t-publictradingrecords)
//
//	symbol Required string  Name of the trading pair
//	limit           integer Default value is 60, max 60
type PublicTradingRecords struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (this PublicTradingRecords) Do(client *Client) ([]PublicTradingRecord, error) {
	type result struct {
		List []PublicTradingRecord `json:"list"`
	}
	r, err := GetPublic[result](client, "quote/trades", this)
	return r.List, err
}

type PublicTradingRecord struct {
	Price        string `json:"price"`
	Time         uint64 `json:"time"`
	Qty          string `json:"qty"`
	IsBuyerMaker int    `json:"isBuyerMaker"`
}

func (this *Client) PublicTradingRecords(v PublicTradingRecords) ([]PublicTradingRecord, error) {
	return v.Do(this)
}

// Query Kline (https://bybit-exchange.github.io/docs/spot/v3/#t-querykline)
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

func (this QueryKline) Do(client *Client) ([]KlineData, error) {
	type result struct {
		List []KlineData `json:"list"`
	}
	r, err := GetPublic[result](client, "quote/kline", this)
	return r.List, err
}

type KlineData struct {
	Timestamp     uint64 `json:"t"`
	Symbol        string `json:"s"`
	Alias         string `json:"sn"`
	ClosePrice    string `json:"c"`
	HighPrice     string `json:"h"`
	LowPrice      string `json:"l"`
	OpenPrice     string `json:"o"`
	TradingVolume string `json:"v"`
}

func (this *Client) QueryKline(v QueryKline) ([]KlineData, error) {
	return v.Do(this)
}

// Latest Information for Symbol (https://bybit-exchange.github.io/docs/spot/v3/#t-spot_latestsymbolinfo)
type SymbolLatestInformation struct {
	Symbol *string `param:"symbol"`
}

func (this SymbolLatestInformation) Do(client *Client) ([]LatestInformation, error) {
	path := "quote/ticker/24hr"
	if this.Symbol == nil {
		type result struct {
			List []LatestInformation `json:"list"`
		}
		r, err := GetPublic[result](client, path, this)
		return r.List, err
	}
	r, err := GetPublic[LatestInformation](client, path, this)
	return []LatestInformation{r}, err
}

type LatestInformation struct {
	Time               uint64            `json:"t"`
	Symbol             string            `json:"s"`
	LastTradedPrice    transport.Float64 `json:"lp"`
	HighPrice          string            `json:"h"`
	LowPrice           string            `json:"l"`
	OpenPrice          string            `json:"o"`
	BestBidPrice       transport.Float64 `json:"bp"`
	BestAskPrice       transport.Float64 `json:"ap"`
	TradingVolume      string            `json:"v"`
	TradingQuoteVolume string            `json:"qv"`
}

func (this *Client) SymbolLatestInformation(symbol *string) ([]LatestInformation, error) {
	return SymbolLatestInformation{Symbol: symbol}.Do(this)
}

func (this *Client) OneSymbolLatestInformation(symbol string) (i LatestInformation, err error) {
	ret, err := this.SymbolLatestInformation(&symbol)
	if err == nil {
		if len(ret) == 1 {
			i = ret[0]
		} else {
			err = errors.New("symbol latest len != 1")
		}
	}
	return
}

// Last Traded Price (https://bybit-exchange.github.io/docs/spot/v3/#t-lasttradedprice)
type LastTradedPrice struct {
	Symbol string `param:"symbol"`
}

func (this LastTradedPrice) Do(client *Client) (SymbolPrice, error) {
	return GetPublic[SymbolPrice](client, "quote/ticker/price", this)
}

type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (this *Client) LastTradedPrice(symbol string) (SymbolPrice, error) {
	return LastTradedPrice{Symbol: symbol}.Do(this)
}

// Best Bid/Ask Price (https://bybit-exchange.github.io/docs/spot/v3/#t-bestbidask)
type BestBidAskPrice struct {
	Symbol string `param:"symbol"`
}

func (this BestBidAskPrice) Do(client *Client) (BestBidAskPriceResult, error) {
	return GetPublic[BestBidAskPriceResult](client, "quote/ticker/bookTicker", this)
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
