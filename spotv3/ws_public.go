package spotv3

import (
	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon/ulog"
)

type WsPublic struct {
	ws *WsClient
}

func NewWsPublic() *WsPublic {
	return &WsPublic{
		ws: NewWsClient("wss://stream.bybit.com/spot/public/v3"),
	}
}

func (o *WsPublic) Shutdown() {
	o.ws.Shutdown()
}

func (o *WsPublic) Conf() *transport.WsConf {
	return o.ws.Conf()
}

func (o *WsPublic) WithLog(log *ulog.Log) *WsPublic {
	o.ws.WithLog(log)
	return o
}

func (o *WsPublic) WithProxy(proxy string) *WsPublic {
	o.Conf().SetProxy(proxy)
	return o
}

func (o *WsPublic) Connected() bool {
	return o.ws.Connected()
}

func (o *WsPublic) Run() {
	o.ws.Run()
}

func (o *WsPublic) SetOnConnected(onConnected func()) {
	o.ws.SetOnConnected(onConnected)
}

func (o *WsPublic) SubscribeDepth(symbol string) bool {
	return o.ws.Subscribe(Subscription{Topic: TopicDepth, Interval: "40", Symbol: symbol})
}
func (o *WsPublic) UnsubscribeDepth(symbol string) bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicDepth, Interval: "40", Symbol: symbol})
}

func (o *WsPublic) SubscribeTrade(symbol string) bool {
	return o.ws.Subscribe(Subscription{Topic: TopicTrade, Symbol: symbol})
}
func (o *WsPublic) UnsubscribeTrade(symbol string) bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicTrade, Symbol: symbol})
}

func (o *WsPublic) SubscribeKline(symbol string, interval KlineInterval) bool {
	return o.ws.Subscribe(Subscription{Topic: TopicKline, Interval: string(interval), Symbol: symbol})
}
func (o *WsPublic) UnsubscribeKline(symbol string, interval KlineInterval) bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicKline, Interval: string(interval), Symbol: symbol})
}

func (o *WsPublic) SubscribeTickers(symbol string) bool {
	return o.ws.Subscribe(Subscription{Topic: TopicTickers, Symbol: symbol})
}
func (o *WsPublic) UnsubscribeTickers(symbol string) bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicTickers, Symbol: symbol})
}

func (o *WsPublic) SubscribeBookTicker(symbol string) bool {
	return o.ws.Subscribe(Subscription{Topic: TopicBookTicker, Symbol: symbol})
}
func (o *WsPublic) UnsubscribeBookTicker(symbol string) bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicBookTicker, Symbol: symbol})
}
