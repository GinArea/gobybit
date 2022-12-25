package spot

import (
	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon/ulog"
)

type WsPublic struct {
	ws            *WsPublicTiny
	subscriptions map[TopicName]topicSymbolsSet
}

func NewWsPublic(url string) *WsPublic {
	return &WsPublic{
		ws:            NewWsPublicTiny(url),
		subscriptions: make(map[TopicName]topicSymbolsSet),
	}
}

func (this *WsPublic) Shutdown() {
	this.ws.Shutdown()
}

func (this *WsPublic) Conf() *transport.WsConf {
	return this.ws.Conf()
}

func (this *WsPublic) WithLog(log *ulog.Log) *WsPublic {
	this.ws.WithLog(log)
	return this
}

func (this *WsPublic) WithProxy(proxy string) *WsPublic {
	this.ws.WithProxy(proxy)
	return this
}

func (this *WsPublic) Run() {
	this.ws.SetOnConnected(func() {
		for topic, symbols := range this.subscriptions {
			for symbol := range symbols {
				this.ws.Subscribe(topic, symbol)
			}
		}
	})
	this.ws.Run()
}

func (this *WsPublic) Subscribe(topic TopicName, symbol string) {
	sub, ok := this.subscriptions[topic]
	if !ok {
		sub = make(topicSymbolsSet)
	}
	sub[symbol] = true
	this.subscriptions[topic] = sub
	if this.ws.Connected() {
		this.ws.Subscribe(topic, symbol)
	}
}

func (this *WsPublic) Unsubscribe(topic TopicName, symbol string) {
	sub, ok := this.subscriptions[topic]
	if ok {
		delete(sub, symbol)
	}
	if this.ws.Connected() {
		this.ws.Unsubscribe(topic, symbol)
	}
}

type topicSymbolsSet map[string]bool
