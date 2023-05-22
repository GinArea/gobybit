package spotv3

import (
	"fmt"

	"github.com/ginarea/gobybit/bybitv2/transport"
	"github.com/msw-x/moon/ulog"
)

type WsPublic struct {
	ws *WsClient
	s  *WsSection
}

func NewWsPublic() *WsPublic {
	o := &WsPublic{
		ws: NewWsClient("spot/public/v3"),
	}
	o.s = NewWsSection(o)
	return o
}

func (o *WsPublic) Shutdown() {
	o.ws.Shutdown()
}

func (o *WsPublic) Conf() *transport.WsConf {
	return o.ws.Conf()
}

func (o *WsPublic) WithByTickUrl() *WsPublic {
	o.ws.WithByTickUrl()
	return o
}

func (o *WsPublic) WithLog(log *ulog.Log) *WsPublic {
	o.ws.WithLog(log)
	return o
}

func (o *WsPublic) WithProxy(proxy string) *WsPublic {
	o.Conf().SetProxy(proxy)
	return o
}

func (o *WsPublic) SetOnDialError(onDialError func(error) bool) {
	o.ws.SetOnDialError(onDialError)
}

func (o *WsPublic) Connected() bool {
	return o.ws.Connected()
}

func (o *WsPublic) Ready() bool {
	return o.Connected()
}

func (o *WsPublic) Run() {
	o.ws.SetOnConnected(o.s.subscribeAll)
	o.ws.SetOnTopicMessage(func(topic TopicMessage) error {
		ok, err := o.s.processTopic(topic)
		if !ok && err == nil {
			err = fmt.Errorf("unknown topic: %s", topic.Topic)
		}
		return err
	})
	o.ws.Run()
}

func (o *WsPublic) Depth(symbol string) *WsExecutor[DepthShot] {
	return NewWsExecutor[DepthShot](o.s, Subscription{Topic: TopicDepth, Interval: "40", Symbol: symbol})
}

func (o *WsPublic) Trade(symbol string) *WsExecutor[TradeShot] {
	return NewWsExecutor[TradeShot](o.s, Subscription{Topic: TopicTrade, Symbol: symbol})
}

func (o *WsPublic) Kline(symbol string, interval KlineInterval) *WsExecutor[KlineShot] {
	return NewWsExecutor[KlineShot](o.s, Subscription{Topic: TopicKline, Interval: string(interval), Symbol: symbol})
}

func (o *WsPublic) Tickers(symbol string) *WsExecutor[TickersShot] {
	return NewWsExecutor[TickersShot](o.s, Subscription{Topic: TopicTickers, Symbol: symbol})
}

func (o *WsPublic) BookTicker(symbol string) *WsExecutor[BookTickerShot] {
	return NewWsExecutor[BookTickerShot](o.s, Subscription{Topic: TopicBookTicker, Symbol: symbol})
}

func (o *WsPublic) subscribe(topic string) bool {
	return o.ws.subscribe(topic)
}

func (o *WsPublic) unsubscribe(topic string) bool {
	return o.ws.unsubscribe(topic)
}
