package uperpetual

import (
	"github.com/ginarea/gobybit/transport"
)

type WsPublic struct {
	ws *WsClient
}

func NewWsPublic() *WsPublic {
	return &WsPublic{
		ws: NewWsClient("public", "wss://stream.bybit.com/realtime_public"),
	}
}

func (this *WsPublic) Shutdown() {
	this.ws.Shutdown()
}

func (this *WsPublic) Conf() *transport.WsConf {
	return this.ws.Conf()
}

func (this *WsPublic) WithProxy(proxy string) *WsPublic {
	this.Conf().SetProxy(proxy)
	return this
}

func (this *WsPublic) Connected() bool {
	return this.ws.Connected()
}

func (this *WsPublic) Run() {
	this.ws.Run()
}

func (this *WsPublic) SetOnConnected(onConnected func()) {
	this.ws.SetOnConnected(onConnected)
}

func (this *WsPublic) SubscribeOrderBook25(symbol string) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicOrderBook25, Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeOrderBook25(symbol string) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicOrderBook25, Symbol: &symbol})
}

func (this *WsPublic) SubscribeOrderBook200(symbol string) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicOrderBook200, Interval: "100ms", Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeOrderBook200(symbol string) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicOrderBook200, Interval: "100ms", Symbol: &symbol})
}

func (this *WsPublic) SubscribeTrade(symbol string) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicTrade, Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeTrade(symbol string) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicTrade, Symbol: &symbol})
}

func (this *WsPublic) SubscribeInstrument(symbol string) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicInstrument, Interval: "100ms", Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeInstrument(symbol string) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicInstrument, Interval: "100ms", Symbol: &symbol})
}

func (this *WsPublic) SubscribeKline(symbol string, interval KlineInterval) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicKline, Interval: string(interval), Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeKline(symbol string, interval KlineInterval) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicKline, Interval: string(interval), Symbol: &symbol})
}

func (this *WsPublic) SubscribeLiquidation(symbol string) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicLiquidation, Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeLiquidation(symbol string) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicLiquidation, Symbol: &symbol})
}
