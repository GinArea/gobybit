package bybit

import (
	"fmt"
	"time"
)

type Symbol struct {
	Name              string `json:"name": "BTCUSDT",`
	Alias             string `json:"alias": "BTCUSDT",`
	BaseCurrency      string `json:"baseCurrency": "BTC",`
	QuoteCurrency     string `json:"quoteCurrency": "USDT",`
	BasePrecision     string `json:"basePrecision": "0.000001",`
	QuotePrecision    string `json:"quotePrecision": "0.00000001",`
	MinTradeQuantity  string `json:"minTradeQuantity": "0.00004",`
	MinTradeAmount    string `json:"minTradeAmount": "1",`
	MaxTradeQuantity  string `json:"maxTradeQuantity": "46.13",`
	MaxTradeAmount    string `json:"maxTradeAmount": "820000",`
	MinPricePrecision string `json:"minPricePrecision": "0.01",`
	Category          int    `json:"category"`
	ShowStatus        bool   `json:"showStatus"`
	Innovation        bool   `json:"innovation"`
}

type OrderBook struct {
	Time uint64 `json:"time"`
	Bids [][]string
	Asks [][]string
}

type PublicTradingRecord struct {
	Price        string `json:"price"`
	Time         uint64 `json:"time"`
	Qty          string `json:"qty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

type LatestInformation struct {
	Time         uint64     `json:"time"`
	Symbol       SymbolSpot `json:"symbol"`
	BestBidPrice string     `json:"bestBidPrice"`
	BestAskPrice string     `json:"bestAskPrice"`
	Volume       string     `json:"volume"`
	QuoteVolume  string     `json:"quoteVolume"`
	LastPrice    string     `json:"lastPrice"`
	HighPrice    string     `json:"highPrice"`
	LowPrice     string     `json:"lowPrice"`
	OpenPrice    string     `json:"openPrice"`
}

type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type BestBidAskPrice struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     uint64 `json:"time"`
}

func (this *Spot) Symbols() ([]Symbol, bool) {
	resp := &Response[[]Symbol]{}
	err := this.client.Get(this.url("symbols"), UrlParam{}, resp)
	return resp.Result, err == nil
}

func (this *Spot) OrderBook(symbol SymbolSpot) (OrderBook, bool) {
	// symbol Required string  Name of the trading pair
	// limit           integer Default value is 100
	resp := &Response[OrderBook]{}
	err := this.client.Get(this.urlQuote("depth"), UrlParam{
		"symbol": symbol,
		//"limit":  100,
	}, resp)
	return resp.Result, err == nil
}

func (this *Spot) MergedOrderBook(symbol SymbolSpot) (OrderBook, bool) {
	// symbol Required string  Name of the trading pair
	// scale           int     Precision of the merged orderbook, 1 means 1 digit
	// limit           integer Default value is 100
	resp := &Response[OrderBook]{}
	err := this.client.Get(this.urlQuote("depth/merged"), UrlParam{
		"symbol": symbol,
		//"scale":  1,
		//"limit":  100,
	}, resp)
	return resp.Result, err == nil
}

func (this *Spot) PublicTradingRecords(symbol SymbolSpot) ([]PublicTradingRecord, bool) {
	// symbol Required string  Name of the trading pair
	// limit           integer Default value is 60, max 60
	resp := &Response[[]PublicTradingRecord]{}
	err := this.client.Get(this.urlQuote("trades"), UrlParam{"symbol": symbol, "limit": 60}, resp)
	return resp.Result, err == nil
}

func (this *Spot) Kline(symbol SymbolSpot, interval time.Duration) ([][]any, bool) {
	// symbol    Required string  Name of the trading pair
	// interval  Required string  Chart interval
	// limit              integer Default value is 1000, max 1000
	// startTime          number  Start time, unit in millisecond
	// endTime            number  End time, unit in millisecond
	resp := &Response[[][]any]{}
	err := this.client.Get(this.urlQuote("kline"), UrlParam{
		"symbol":   symbol,
		"interval": fmt.Sprintf("%dm", int(interval.Minutes())),
		//"limit":    1000,
		//"startTime"
		//"endTime"
	}, resp)
	return resp.Result, err == nil
}

func (this *Spot) SymbolLatestInformation(symbol SymbolSpot) (LatestInformation, bool) {
	// symbol string Name of the trading pair
	resp := &Response[LatestInformation]{}
	err := this.client.Get(this.urlQuote("ticker/24hr"), UrlParam{"symbol": symbol}, resp)
	return resp.Result, err == nil
}

func (this *Spot) LastTradedPrice(symbol SymbolSpot) (SymbolPrice, bool) {
	// symbol string Name of the trading pair
	resp := &Response[SymbolPrice]{}
	err := this.client.Get(this.urlQuote("ticker/price"), UrlParam{"symbol": symbol}, resp)
	return resp.Result, err == nil
}

func (this *Spot) BestBidAskPrice(symbol SymbolSpot) (BestBidAskPrice, bool) {
	// symbol string Name of the trading pair
	resp := &Response[BestBidAskPrice]{}
	err := this.client.Get(this.urlQuote("ticker/book_ticker"), UrlParam{"symbol": symbol}, resp)
	return resp.Result, err == nil
}
