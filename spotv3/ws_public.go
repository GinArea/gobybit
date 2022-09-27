package spotv3

import (
	"github.com/tranquiil/bybit/iperpetual"
	"github.com/tranquiil/bybit/transport"
)

type WsPublic struct {
	ws *WsClient
}

func NewWsPublic() *WsPublic {
	return &WsPublic{
		ws: NewWsClient("public", "wss://stream.bybit.com/spot/public/v3"),
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

func (this *WsPublic) SubscribeDepth(symbol iperpetual.Symbol) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicDepth, Interval: "40", Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeDepth(symbol iperpetual.Symbol) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicDepth, Interval: "40", Symbol: &symbol})
}

func (this *WsPublic) SubscribeTrade(symbol iperpetual.Symbol) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicTrade, Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeTrade(symbol iperpetual.Symbol) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicTrade, Symbol: &symbol})
}

func (this *WsPublic) SubscribeKline(symbol iperpetual.Symbol, interval KlineInterval) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicKline, Interval: string(interval), Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeKline(symbol iperpetual.Symbol, interval KlineInterval) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicKline, Interval: string(interval), Symbol: &symbol})
}

func (this *WsPublic) SubscribeTickers(symbol iperpetual.Symbol) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicTickers, Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeTickers(symbol iperpetual.Symbol) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicTickers, Symbol: &symbol})
}

func (this *WsPublic) SubscribeBookTicker(symbol iperpetual.Symbol) bool {
	return this.ws.Subscribe(Subscription{Topic: TopicBookTicker, Symbol: &symbol})
}
func (this *WsPublic) UnsubscribeBookTicker(symbol iperpetual.Symbol) bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicBookTicker, Symbol: &symbol})
}
