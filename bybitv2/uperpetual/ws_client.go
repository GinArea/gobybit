package uperpetual

import (
	"strings"

	"github.com/ginarea/gobybit/bybitv2/transport"
	"github.com/msw-x/moon/uerr"
	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
)

type WsClient struct {
	log         *ulog.Log
	ws          *transport.WsClient
	onConnected func()
	onAuth      func(bool)
}

func NewWsClient(name string, url string) *WsClient {
	ws := transport.NewWsClient(url)
	return &WsClient{
		log: ulog.Empty(),
		ws:  ws,
	}
}

func (this *WsClient) Shutdown() {
	this.log.Debug("shutdown")
	this.ws.Shutdown()
}

func (this *WsClient) Conf() *transport.WsConf {
	return this.ws.Conf()
}

func (o *WsClient) WithUrl(url string) *WsClient {
	o.ws.WithUrl(url)
	return o
}

func (o *WsClient) WithByTickUrl() *WsClient {
	o.ws.WithByTickUrl()
	return o
}

func (this *WsClient) WithLog(log *ulog.Log) *WsClient {
	this.ws.WithLog(log)
	return this
}

func (this *WsClient) WithProxy(proxy string) *WsClient {
	this.Conf().SetProxy(proxy)
	return this
}

func (this *WsClient) Connected() bool {
	return this.ws.Connected()
}

func (this *WsClient) Run() {
	this.log.Debug("run")
	this.ws.SetOnMessage(this.processMessage)
	this.ws.Run()
}

func (this *WsClient) SetOnConnected(onConnected func()) {
	this.ws.SetOnConnected(onConnected)
}

func (this *WsClient) SetOnDisconnected(onConnected func()) {
	this.ws.SetOnDisconnected(onConnected)
}

func (this *WsClient) SetOnAuth(onAuth func(bool)) {
	this.onAuth = onAuth
}

func (this *WsClient) Send(cmd any) bool {
	return this.ws.Send(cmd)
}

type Subscription struct {
	Topic    TopicName
	Interval string
	Symbol   *string
}

func (this *Subscription) String() string {
	s := []string{string(this.Topic)}
	if this.Interval != "" {
		s = append(s, this.Interval)
	}
	if this.Symbol != nil {
		s = append(s, string(*this.Symbol))
	}
	return ufmt.JoinSliceWith(".", s)
}

func (this *Subscription) Request(operation string) Request {
	return Request{
		Operation: operation,
		Args:      []string{this.String()},
	}
}

func (this *WsClient) Subscribe(s Subscription) bool {
	this.log.Infof("subscribe: topic[%s]", s.Topic)
	return this.ws.Send(s.Request("subscribe"))
}

func (this *WsClient) Unsubscribe(s Subscription) bool {
	this.log.Infof("unsubscribe: topic[%s]", s.Topic)
	return this.ws.Send(s.Request("unsubscribe"))
}

type Request struct {
	Operation string   `json:"op"`
	Args      []string `json:"args"`
}

type Responce struct {
	Success bool    `json:"success"`
	RetMsg  string  `json:"ret_msg"`
	ConnID  string  `json:"conn_id"`
	Request Request `json:"request"`
}

func (this *WsClient) processMessage(name string, msg []byte) {
	this.log.Debug("name:", name)
	switch name {
	case "success":
		v := transport.JsonUnmarshal[Responce](msg)
		this.processResponce(v)
	case "topic":
		v := transport.JsonUnmarshal[struct {
			Name string `json:"topic"`
			Type string `json:"type"`
		}](msg)
		s := strings.Split(v.Name, ".")
		name := s[0]
		this.processTopic(TopicName(name), v.Type == "delta", msg)
	default:
		uerr.Panic("unknown message:", name)
	}
}

func (this *WsClient) processResponce(r Responce) {
	if !r.Success {
		this.log.Error(r.RetMsg)
		return
	}
	name := r.RetMsg
	if name == "" {
		name = r.Request.Operation
	}
	this.log.Debug("response:", name, "success:", r.Success)
	switch name {
	case "pong":
	case "auth":
		if this.onAuth != nil {
			this.onAuth(r.Success)
		}
	case "subscribe":
	case "unsubscribe":
	default:
		uerr.Panic("unknown response:", name)
	}
}

func (this *WsClient) processTopic(topic TopicName, delta bool, msg []byte) {
	type OrderBookResult struct {
		OrderBook []OrderBookSnapshot `json:"order_book"`
	}
	switch topic {
	// public
	case TopicOrderBook25:
		if delta {
			transport.JsonUnmarshal[Topic[OrderBookDelta]](msg)
		} else {
			transport.JsonUnmarshal[Topic[OrderBookResult]](msg)
		}
	case TopicOrderBook200:
		if delta {
			transport.JsonUnmarshal[Topic[OrderBookDelta]](msg)
		} else {
			transport.JsonUnmarshal[Topic[OrderBookResult]](msg)
		}
	case TopicTrade:
		transport.JsonUnmarshal[Topic[[]TradeSnapshot]](msg)
	case TopicInstrument:
		if delta {
			transport.JsonUnmarshal[Topic[InstrumentDelta]](msg)
		} else {
			transport.JsonUnmarshal[Topic[InstrumentSnapshot]](msg)
		}
	case TopicKline:
		transport.JsonUnmarshal[Topic[[]KlineSnapshot]](msg)
	case TopicLiquidation:
		transport.JsonUnmarshal[Topic[LiquidationSnapshot]](msg)
	//private
	case TopicPosition:
		transport.JsonUnmarshal[Topic[[]PositionSnapshot]](msg)
	case TopicExecution:
		transport.JsonUnmarshal[Topic[[]ExecutionSnapshot]](msg)
	case TopicOrder:
		transport.JsonUnmarshal[Topic[[]OrderSnapshot]](msg)
	case TopicStopOrder:
		transport.JsonUnmarshal[Topic[[]StopOrderSnapshot]](msg)
	case TopicWallet:
		transport.JsonUnmarshal[Topic[[]WalletSnapshot]](msg)
	default:
		uerr.Panic("unknown topic:", topic)
	}
}
