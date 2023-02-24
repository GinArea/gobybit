package spotv3

import (
	"strings"

	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon"
	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
)

type WsClient struct {
	log            *ulog.Log
	ws             *transport.WsClient
	onAuth         func(bool)
	onTopicMessage func(TopicMessage) error
}

func NewWsClient(url string) *WsClient {
	ws := transport.NewWsClient(url)
	return &WsClient{
		log: ulog.Empty(),
		ws:  ws,
	}
}

func (o *WsClient) Shutdown() {
	o.log.Debug("shutdown")
	o.ws.Shutdown()
}

func (o *WsClient) Conf() *transport.WsConf {
	return o.ws.Conf()
}

func (o *WsClient) WithUrl(url string) *WsClient {
	o.ws.WithUrl(url)
	return o
}

func (o *WsClient) WithByTickUrl() *WsClient {
	o.ws.WithByTickUrl()
	return o
}

func (o *WsClient) WithLog(log *ulog.Log) *WsClient {
	o.ws.WithLog(log)
	return o
}

func (o *WsClient) WithProxy(proxy string) *WsClient {
	o.Conf().SetProxy(proxy)
	return o
}

func (o *WsClient) SetOnConnected(onConnected func()) {
	o.ws.SetOnConnected(onConnected)
}

func (o *WsClient) SetOnDisconnected(onConnected func()) {
	o.ws.SetOnDisconnected(onConnected)
}

func (o *WsClient) SetOnDialError(onDialError func(error) bool) {
	o.ws.SetOnDialError(onDialError)
}

func (o *WsClient) SetOnAuth(onAuth func(bool)) {
	o.onAuth = onAuth
}

func (o *WsClient) SetOnTopicMessage(onTopicMessage func(TopicMessage) error) {
	o.onTopicMessage = onTopicMessage
}

func (o *WsClient) Run() {
	o.log.Debug("run")
	o.ws.SetOnMessage(o.processMessage)
	o.ws.Run()
}

func (o *WsClient) Connected() bool {
	return o.ws.Connected()
}

func (o *WsClient) send(cmd any) bool {
	return o.ws.Send(cmd)
}

func (o *WsClient) subscribe(topic string) bool {
	o.log.Infof("subscribe: topic[%s]", topic)
	return o.send(Request{
		Name: "subscribe",
		Args: []string{topic},
	})
}

func (o *WsClient) unsubscribe(topic string) bool {
	o.log.Infof("unsubscribe: topic[%s]", topic)
	return o.send(Request{
		Name: "unsubscribe",
		Args: []string{topic},
	})
}

// /////
func (o *WsClient) Send(cmd any) bool {
	return o.ws.Send(cmd)
}

func (o *Subscription) Request(operation string) Request {
	return Request{
		Name: operation,
		Args: []string{o.String()},
	}
}

func (o *WsClient) Subscribe(s Subscription) bool {
	o.log.Infof("subscribe: topic[%s]", s.Topic)
	return o.ws.Send(s.Request("subscribe"))
}

func (o *WsClient) Unsubscribe(s Subscription) bool {
	o.log.Infof("unsubscribe: topic[%s]", s.Topic)
	return o.ws.Send(s.Request("unsubscribe"))
}

/////////

func (o *WsClient) processMessage(name string, msg []byte) {
	v := transport.JsonUnmarshal[Responce](msg)
	if v.IsTopic() {
		o.processTopic(TopicMessage{
			Topic: v.Topic,
			Delta: v.Type == "delta",
			Bin:   msg,
		})
	} else {
		o.processResponce(v)
	}
}

func (o *WsClient) processResponce(r Responce) {
	name := r.RetMsg
	if name == "" {
		name = r.Name
	}
	if !r.Success && name != "pong" {
		o.log.Error(r.RetMsg)
		return
	}
	o.log.Debug("response:", name)
	switch name {
	case "pong":
	case "auth":
		o.log.Info("auth:", ufmt.SuccessFailure(r.Success))
		if o.onAuth != nil {
			o.onAuth(r.Success)
		}
	case "subscribe":
		o.log.Infof("topic%s subscribe: %s", r.Args, ufmt.SuccessFailure(r.Success))
	case "unsubscribe":
		o.log.Infof("topic%s unsubscribe: %s", r.Args, ufmt.SuccessFailure(r.Success))
	default:
		o.log.Error("unknown response:", name)
	}
}

func (o *WsClient) processTopic(topic TopicMessage) {
	if o.onTopicMessage != nil {
		err := o.onTopicMessage(topic)
		if err != nil {
			o.log.Error("process topic:", err)
		}
		return
	}
	s := strings.Split(topic.Topic, ".")
	name := s[0]
	switch TopicName(name) {
	// public
	case TopicDepth:
		transport.JsonUnmarshal[Topic[DepthShot]](topic.Bin)
	case TopicTrade:
		transport.JsonUnmarshal[Topic[TradeShot]](topic.Bin)
	case TopicKline:
		transport.JsonUnmarshal[Topic[KlineShot]](topic.Bin)
	case TopicTickers:
		transport.JsonUnmarshal[Topic[TickersShot]](topic.Bin)
	case TopicBookTicker:
		transport.JsonUnmarshal[Topic[BookTickerShot]](topic.Bin)
	// private
	case TopicOutbound:
		transport.JsonUnmarshal[Topic[[]OutboundSnapshot]](topic.Bin)
	case TopicOrder:
		transport.JsonUnmarshal[Topic[[]OrderSnapshot]](topic.Bin)
	case TopicStopOrder:
		transport.JsonUnmarshal[Topic[[]StopOrderSnapshot]](topic.Bin)
	case TopicTicket:
		transport.JsonUnmarshal[Topic[[]TicketSnapshot]](topic.Bin)
	default:
		moon.Panic("unknown topic:", topic.Topic)
	}
}

type Request struct {
	Name  string   `json:"op"`
	Args  []string `json:"args,omitempty"`
	ReqID string   `json:"req_id,omitempty"`
}

type Responce struct {
	Name    string `json:"op"`
	Args    []any  `json:"args"`
	ReqID   string `json:"req_id"`
	ConnID  string `json:"conn_id"`
	Success bool   `json:"success"`
	RetMsg  string `json:"ret_msg"`
	Topic   string `json:"topic"`
	Type    string `json:"type"`
}

func (o *Responce) IsTopic() bool {
	return o.Topic != ""
}

type TopicMessage struct {
	Topic string
	Delta bool
	Bin   []byte
}
