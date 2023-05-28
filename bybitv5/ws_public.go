package bybitv5

import (
	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPublic struct {
	c             *WsClient
	category      Category
	onConnected   func()
	subscriptions *Subscriptions[PublicTopic]
}

func NewWsPublic() *WsPublic {
	o := new(WsPublic)
	o.c = NewWsClient()
	o.subscriptions = NewSubscriptions[PublicTopic](o)
	return o
}

func (o *WsPublic) Close() {
	o.c.Close()
}

func (o *WsPublic) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsPublic) WithLog(log *ulog.Log) *WsPublic {
	o.c.WithLog(log)
	return o
}

func (o *WsPublic) WithBaseUrl(url string) *WsPublic {
	o.c.WithBaseUrl(url)
	return o
}

func (o *WsPublic) WithByTickUrl() *WsPublic {
	return o.WithBaseUrl(MainBaseByTickWsUrl)
}

func (o *WsPublic) WithProxy(proxy string) *WsPublic {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsPublic) WithLogRequest(enable bool) *WsPublic {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsPublic) WithLogResponse(enable bool) *WsPublic {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsPublic) WithOnConnected(f func()) {
	o.onConnected = f
}

func (o *WsPublic) WithOnDisconnected(f func()) {
	o.c.WithOnDisconnected(f)
}

func (o *WsPublic) WithCategory(category Category) *WsPublic {
	o.category = category
	return o
}

func (o *WsPublic) Linear() *WsPublic {
	return o.WithCategory(Linear)
}

func (o *WsPublic) Inverse() *WsPublic {
	return o.WithCategory(Inverse)
}

func (o *WsPublic) Spot() *WsPublic {
	return o.WithCategory(Spot)
}

func (o *WsPublic) Option() *WsPublic {
	return o.WithCategory(Option)
}

func (o *WsPublic) Run() {
	o.c.WithPath(uhttp.UrlJoin("v5", "public", o.category))
	o.c.WithOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
		o.subscriptions.subscribeAll()
	})
	o.c.WithOnResponse(o.onResponse)
	o.c.WithOnTopic(o.onTopic)
	o.c.Run()
}

func (o *WsPublic) Connected() bool {
	return o.c.Connected()
}

func (o *WsPublic) Ready() bool {
	return o.Connected()
}

func (o *WsPublic) OrderBook(symbol string, depth int) {
	o.subscribe(topicNameExt("orderbook", symbol, depth))
	//NewWsDeltaExecutor[OrderbookShot]
}

func (o *WsPublic) Trade(symbol string) *Executor[PublicTopic, []TradeShot] {
	return NewExecutor[PublicTopic, []TradeShot](topicName("publicTrade", symbol), o.subscriptions)
}

func (o *WsPublic) Ticker(symbol string) *Executor[PublicTopic, Ticker] {
	return NewExecutor[PublicTopic, Ticker](topicName("tickers", symbol), o.subscriptions)
	//o.subscribe(topicName("tickers", symbol))
	//o.subscribe("tickers", symbol)
}

func (o *WsPublic) Kline(symbol string, interval Interval) {
	//o.subscribeExt("kline", symbol, interval)
}

func (o *WsPublic) Liquidation(symbol string) {
	//o.subscribe("liquidation", symbol)
}

func (o *WsPublic) LtKline(symbol string, interval Interval) {
	//o.subscribeExt("kline_lt", symbol, interval)
}

func (o *WsPublic) LtTicker(symbol string) {
	//o.subscribe("tickers_lt", symbol)
}

func (o *WsPublic) LtNav(symbol string) {
	//o.subscribe("lt", symbol)
}

func (o *WsPublic) subscribe(topic string) {
	o.c.Subscribe(topic)
}

func (o *WsPublic) unsubscribe(topic string) {
	o.c.Unsubscribe(topic)
}

func (o *WsPublic) onResponse(r WsResponse) error {
	r.Log(o.c.Log())
	return nil
}

func (o *WsPublic) onTopic(data []byte) error {
	return o.subscriptions.processTopic(data)
}

func topicName(name string, symbol string) string {
	return ufmt.JoinWith(".", name, symbol)
}

func topicNameExt(name string, symbol string, param any) string {
	return ufmt.JoinWith(".", name, param, symbol)
}
