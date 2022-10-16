// WebSocket Data (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-websocket)
package iperpetual

import (
	"fmt"

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
	ws := transport.NewWsClient("wss://stream.bybit.com/realtime")
	c := &WsClient{
		log: ulog.New(fmt.Sprintf("ws-iperpetual[%s]", ws.ID())),
		ws:  ws,
	}
	c.public = NewWsPublic(c)
	return c
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

func (this *WsClient) WithAuth(key string, secret string) *WsClient {
	this.private = NewWsPrivate(this, key, secret)
	return this
}

func (this *WsClient) SetOnConnected(onConnected func()) {
	this.onConnected = onConnected
}

func (this *WsClient) SetOnDisconnected(onDisconnected func()) {
	this.onDisconnected = onDisconnected
}

func (this *WsClient) SetOnAuth(onAuth func(bool)) {
	this.onAuth = onAuth
}

func (this *WsClient) Public() *WsPublic {
	return this.public
}

func (this *WsClient) Private() *WsPrivate {
	if this.private == nil {
		moon.Panic("private methods are forbidden")
	}
	return this.private
}

func (this *WsClient) Run() {
	this.log.Debug("run")
	this.ws.SetOnConnected(func() {
		if this.onConnected != nil {
			this.onConnected()
		}
		if this.private == nil {
			this.ready = true
		} else {
			this.log.Info("auth")
			this.private.auth()
		}
		this.public.subscribeAll()
	})
	this.ws.SetOnDisconnected(func() {
		this.ready = false
		if this.onDisconnected != nil {
			this.onDisconnected()
		}
	})
	this.ws.SetOnMessage(this.processMessage)
	this.ws.Run()
}

func (this *WsClient) Connected() bool {
	return this.ws.Connected()
}

func (this *WsClient) Ready() bool {
	return this.ready
}

func (this *WsClient) send(cmd any) bool {
	return this.ws.Send(cmd)
}

func (this *WsClient) subscribe(topic string) bool {
	this.log.Infof("subscribe: topic[%s]", topic)
	return this.send(Request{
		Name: "subscribe",
		Args: []string{topic},
	})
}

func (this *WsClient) unsubscribe(topic string) bool {
	this.log.Infof("unsubscribe: topic[%s]", topic)
	return this.send(Request{
		Name: "unsubscribe",
		Args: []string{topic},
	})
}

func (this *WsClient) processMessage(name string, msg []byte) {
	switch name {
	case "success":
		v := transport.JsonUnmarshal[Responce](msg)
		this.processResponce(v)
	case "topic":
		v := transport.JsonUnmarshal[struct {
			Name string `json:"topic"`
			Type string `json:"type"`
		}](msg)
		this.processTopic(TopicMessage{
			Topic: v.Name,
			Delta: v.Type == "delta",
			Bin:   msg,
		})
	default:
		this.log.Error("unknown message:", name)
	}
}

func (this *WsClient) processResponce(r Responce) {
	if !r.Success {
		this.log.Error(r.RetMsg)
		return
	}
	name := r.RetMsg
	if name == "" {
		name = r.Request.Name
	}
	switch name {
	case "pong":
	case "auth":
		this.log.Info("auth:", ufmt.SuccessFailure(r.Success))
		if this.onAuth != nil {
			this.onAuth(r.Success)
		}
		if this.private != nil {
			this.ready = true
			this.private.subscribeAll()
		}
	case "subscribe":
		this.log.Infof("topic%s subscribe: %s", r.Request.Args, ufmt.SuccessFailure(r.Success))
	case "unsubscribe":
		this.log.Infof("topic%s unsubscribe: %s", r.Request.Args, ufmt.SuccessFailure(r.Success))
	default:
		this.log.Error("unknown response:", name)
	}
}

func (this *WsClient) processTopic(m TopicMessage) {
	ok, err := this.public.processTopic(m)
	if err == nil && this.private != nil && !ok {
		_, err = this.private.processTopic(m)
	}
	if err != nil {
		this.log.Errorf("process topic[%s]: %v", m.Topic, err)
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
