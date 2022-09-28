package spot

import (
	"fmt"

	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon/ulog"
)

type WsPublicTiny struct {
	log *ulog.Log
	ws  *transport.WsClient
}

func NewWsPublicTiny(url string) *WsPublicTiny {
	ws := transport.NewWsClient(url)
	return &WsPublicTiny{
		log: ulog.New(fmt.Sprintf("ws-public[%s]", ws.ID())),
		ws:  ws,
	}
}

func (this *WsPublicTiny) Shutdown() {
	this.log.Debug("shutdown")
	this.ws.Shutdown()
}

func (this *WsPublicTiny) Conf() *transport.WsConf {
	return this.ws.Conf()
}

func (this *WsPublicTiny) WithProxy(proxy string) *WsPublicTiny {
	this.ws.WithProxy(proxy)
	return this
}

func (this *WsPublicTiny) Connected() bool {
	return this.ws.Connected()
}

func (this *WsPublicTiny) Run() {
	this.log.Debug("run")
	this.ws.SetOnMessage(this.processMessage)
	this.ws.Run()
}

func (this *WsPublicTiny) SetOnConnected(onConnected func()) {
	this.ws.SetOnConnected(onConnected)
}

func (this *WsPublicTiny) Subscribe(topic TopicName, symbol Symbol) bool {
	this.log.Infof("subscribe: topic[%s] symbol[%s] ", symbol, topic)
	return this.ws.Send(Topic{
		Name:  topic,
		Event: TopicEventSub,
		Params: TopicParams{
			Symbol: symbol,
		},
	})
}

func (this *WsPublicTiny) Unsubscribe(topic TopicName, symbol Symbol) bool {
	this.log.Infof("unsubscribe: topic[%s] symbol[%s]", symbol, topic)
	return this.ws.Send(Topic{
		Name:  topic,
		Event: TopicEventCancel,
		Params: TopicParams{
			Symbol: symbol,
		},
	})
}

func (this *WsPublicTiny) processMessage(name string, msg []byte) {
	switch name {
	case "pong":
		v := transport.JsonUnmarshal[struct {
			Pong int `json:"pong"`
		}](msg)
		this.log.Debug("pong:", v.Pong)
	case "code":
		v := transport.JsonUnmarshal[struct {
			Code        string `json:"code"`
			Description string `json:"desc"`
		}](msg)
		this.log.Warningf("code[%s]: %s", v.Code, v.Description)
	case "topic":
		v := transport.JsonUnmarshal[TopicSubscribtion](msg)
		if v.HasCode() {
			//success := v.Ok()
			this.log.Infof("topic %s (%s) subscribtion: %s (%s)", v.Topic.Name, v.Params.Symbol, v.Message, v.Code)
		} else {
			this.processTopic(msg)
		}
	default:
		panic("unknown message type")
	}
}

func (this *WsPublicTiny) processTopic(msg []byte) {
	v := transport.JsonUnmarshal[TopicNotification[any]](msg)
	var data any
	switch v.Topic {
	case TopicDepth:
		v := transport.JsonUnmarshal[TopicNotification[TopicDataDepth]](msg)
		data = v.Data
	case TopicKline:
		v := transport.JsonUnmarshal[TopicNotification[TopicDataKline]](msg)
		data = v.Data
	case TopicTrade:
		v := transport.JsonUnmarshal[TopicNotification[TopicDataTrade]](msg)
		data = v.Data
	case TopicBookTicker:
		v := transport.JsonUnmarshal[TopicNotification[TopicDataBookTicker]](msg)
		data = v.Data
	case TopicRealtimes:
		v := transport.JsonUnmarshal[TopicNotification[TopicDataRealtimes]](msg)
		data = v.Data
	default:
		panic(fmt.Sprintf("unknown topic name: %s", v.Topic))
	}
	this.log.Infof("topic %s (%s) notify: %+v", v.Topic, v.Params.Symbol, data)
}
