package spotv3

import (
	"fmt"
	"strings"

	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon"
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
		log: ulog.New(fmt.Sprintf("ws-%s[%s]", name, ws.ID())),
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
	Args      []string `json:"args,omitempty"`
	ReqID     string   `json:"req_id,omitempty"`
}

type Responce struct {
	Operation string `json:"op"`
	Args      []any  `json:"args"`
	ReqID     string `json:"req_id"`
	ConnID    string `json:"conn_id"`
	Success   bool   `json:"success"`
	RetMsg    string `json:"ret_msg"`
	Topic     string `json:"topic"`
	Type      string `json:"type"`
}

func (this *Responce) IsTopic() bool {
	return this.Topic != ""
}

func (this *WsClient) processMessage(name string, msg []byte) {
	v := transport.JsonUnmarshal[Responce](msg)
	if v.IsTopic() {
		s := strings.Split(v.Topic, ".")
		name := s[0]
		this.processTopic(TopicName(name), v.Type == "delta", msg)
	} else {
		this.processResponce(v)
	}
}

func (this *WsClient) processResponce(r Responce) {
	name := r.RetMsg
	if name == "" {
		name = r.Operation
	}
	if !r.Success && name != "pong" {
		this.log.Error(r.RetMsg)
		return
	}
	this.log.Debug("response:", name)
	switch name {
	case "pong":
	case "auth":
		if this.onAuth != nil {
			this.onAuth(r.Success)
		}
	case "subscribe":
	case "unsubscribe":
	default:
		moon.Panic("unknown response:", name)
	}
}

func (this *WsClient) processTopic(topic TopicName, delta bool, msg []byte) {
	switch topic {
	// public
	case TopicDepth:
		transport.JsonUnmarshal[Topic[DepthDelta]](msg)
	case TopicTrade:
		transport.JsonUnmarshal[Topic[TradeDelta]](msg)
	case TopicKline:
		transport.JsonUnmarshal[Topic[KlineDelta]](msg)
	case TopicTickers:
		transport.JsonUnmarshal[Topic[TickersDelta]](msg)
	case TopicBookTicker:
		transport.JsonUnmarshal[Topic[BookTickerDelta]](msg)
	// private
	case TopicOutbound:
		transport.JsonUnmarshal[Topic[[]OutboundSnapshot]](msg)
	case TopicOrder:
		transport.JsonUnmarshal[Topic[[]OrderSnapshot]](msg)
	case TopicStopOrder:
		transport.JsonUnmarshal[Topic[[]StopOrderSnapshot]](msg)
	case TopicTicket:
		transport.JsonUnmarshal[Topic[[]TicketSnapshot]](msg)
	default:
		moon.Panic("unknown topic:", topic)
	}
}
