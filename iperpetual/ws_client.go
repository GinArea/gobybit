// WebSocket Data (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-websocket)
package iperpetual

import (
	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon"
	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
)

type WsClient struct {
	log            *ulog.Log
	ws             *transport.WsClient
	public         *WsPublic
	private        *WsPrivate
	ready          bool
	onConnected    func()
	onDisconnected func()
	onAuth         func(bool)
}

func NewWsClient() *WsClient {
	ws := transport.NewWsClient("realtime")
	c := &WsClient{
		log: ulog.Empty(),
		ws:  ws,
	}
	c.public = NewWsPublic(c)
	return c
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
	o.log = log
	o.ws.WithLog(log)
	return o
}

func (o *WsClient) WithProxy(proxy string) *WsClient {
	o.Conf().SetProxy(proxy)
	return o
}

func (o *WsClient) WithAuth(key string, secret string) *WsClient {
	o.private = NewWsPrivate(o, key, secret)
	return o
}

func (o *WsClient) SetOnConnected(onConnected func()) {
	o.onConnected = onConnected
}

func (o *WsClient) SetOnDisconnected(onDisconnected func()) {
	o.onDisconnected = onDisconnected
}

func (o *WsClient) SetOnDialError(onDialError func(error) bool) {
	o.ws.SetOnDialError(onDialError)
}

func (o *WsClient) SetOnAuth(onAuth func(bool)) {
	o.onAuth = onAuth
}

func (o *WsClient) Public() *WsPublic {
	return o.public
}

func (o *WsClient) Private() *WsPrivate {
	if o.private == nil {
		moon.Panic("private methods are forbidden")
	}
	return o.private
}

func (o *WsClient) Run() {
	o.log.Debug("run")
	o.ws.SetOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
		if o.private == nil {
			o.ready = true
		} else {
			o.log.Info("auth")
			o.private.auth()
		}
		o.public.subscribeAll()
	})
	o.ws.SetOnDisconnected(func() {
		o.ready = false
		if o.onDisconnected != nil {
			o.onDisconnected()
		}
	})
	o.ws.SetOnMessage(o.processMessage)
	o.ws.Run()
}

func (o *WsClient) Connected() bool {
	return o.ws.Connected()
}

func (o *WsClient) Ready() bool {
	return o.ready
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

func (o *WsClient) processMessage(name string, msg []byte) {
	switch name {
	case "success":
		v := transport.JsonUnmarshal[Responce](msg)
		o.processResponce(v)
	case "topic":
		v := transport.JsonUnmarshal[struct {
			Name string `json:"topic"`
			Type string `json:"type"`
		}](msg)
		o.processTopic(TopicMessage{
			Topic: v.Name,
			Delta: v.Type == "delta",
			Bin:   msg,
		})
	default:
		o.log.Error("unknown message:", name)
	}
}

func (o *WsClient) processResponce(r Responce) {
	if !r.Success {
		o.log.Error(r.RetMsg)
		return
	}
	name := r.RetMsg
	if name == "" {
		name = r.Request.Name
	}
	switch name {
	case "pong":
	case "auth":
		o.log.Info("auth:", ufmt.SuccessFailure(r.Success))
		if o.onAuth != nil {
			o.onAuth(r.Success)
		}
		if o.private != nil {
			o.ready = true
			o.private.subscribeAll()
		}
	case "subscribe":
		o.log.Infof("topic%s subscribe: %s", r.Request.Args, ufmt.SuccessFailure(r.Success))
	case "unsubscribe":
		o.log.Infof("topic%s unsubscribe: %s", r.Request.Args, ufmt.SuccessFailure(r.Success))
	default:
		o.log.Error("unknown response:", name)
	}
}

func (o *WsClient) processTopic(m TopicMessage) {
	ok, err := o.public.processTopic(m)
	if err == nil && o.private != nil && !ok {
		ok, err = o.private.processTopic(m)
	}
	if err == nil {
		if !ok {
			o.log.Warningf("process topic[%s]: not found", m.Topic)
		}
	} else {
		o.log.Errorf("process topic[%s]: %v", m.Topic, err)
	}
	return
}

type Request struct {
	Name string   `json:"op"`
	Args []string `json:"args"`
}

type Responce struct {
	Success bool    `json:"success"`
	RetMsg  string  `json:"ret_msg"`
	ConnID  string  `json:"conn_id"`
	Request Request `json:"request"`
}

type TopicMessage struct {
	Topic string
	Delta bool
	Bin   []byte
}
