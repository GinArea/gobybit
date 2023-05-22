package spotv3

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ginarea/gobybit/bybitv2/transport"
	"github.com/msw-x/moon/ulog"
)

type WsPrivate struct {
	ws     *WsClient
	s      *WsSection
	key    string
	secret string
	ready  bool
}

func NewWsPrivate(key string, secret string) *WsPrivate {
	o := &WsPrivate{
		ws:     NewWsClient("spot/private/v3"),
		key:    key,
		secret: secret,
	}
	o.s = NewWsSection(o)
	return o
}

func (o *WsPrivate) Shutdown() {
	o.ws.Shutdown()
}

func (o *WsPrivate) Conf() *transport.WsConf {
	return o.ws.Conf()
}

func (o *WsPrivate) WithLog(log *ulog.Log) *WsPrivate {
	o.ws.WithLog(log)
	return o
}

func (o *WsPrivate) WithByTickUrl() *WsPrivate {
	o.ws.WithByTickUrl()
	return o
}

func (o *WsPrivate) WithProxy(proxy string) *WsPrivate {
	o.Conf().SetProxy(proxy)
	return o
}

func (o *WsPrivate) SetOnDialError(onDialError func(error) bool) {
	o.ws.SetOnDialError(onDialError)
}

func (o *WsPrivate) Connected() bool {
	return o.ws.Connected()
}

func (o *WsPrivate) Ready() bool {
	return o.ready
}

func (o *WsPrivate) Run() {
	o.ws.SetOnConnected(func() {
		o.auth()
	})
	o.ws.SetOnDisconnected(func() {
		o.ready = false
	})
	o.ws.SetOnAuth(func(ok bool) {
		o.ready = ok
		o.s.subscribeAll()
	})
	o.ws.SetOnTopicMessage(func(topic TopicMessage) error {
		ok, err := o.s.processTopic(topic)
		if !ok && err == nil {
			err = fmt.Errorf("unknown topic: %s", topic.Topic)
		}
		return err
	})
	o.ws.Run()
}

func (o *WsPrivate) Order() *WsExecutor[[]OrderShot] {
	return NewWsExecutor[[]OrderShot](o.s, Subscription{Topic: TopicOrder})
}

func (o *WsPrivate) StopOrder() *WsExecutor[[]StopOrderShot] {
	return NewWsExecutor[[]StopOrderShot](o.s, Subscription{Topic: TopicStopOrder})
}

func (o *WsPrivate) Ticket() *WsExecutor[[]TicketShot] {
	return NewWsExecutor[[]TicketShot](o.s, Subscription{Topic: TopicTicket})
}

func (o *WsPrivate) Outbound() *WsExecutor[[]OutboundShot] {
	return NewWsExecutor[[]OutboundShot](o.s, Subscription{Topic: TopicOutbound})
}

func (o *WsPrivate) auth() {
	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	sig := hmac.New(sha256.New, []byte(o.secret))
	sig.Write([]byte(req))
	signature := hex.EncodeToString(sig.Sum(nil))
	cmd := struct {
		Name string `json:"op"`
		Args []any  `json:"args"`
	}{
		Name: "auth",
		Args: []any{
			o.key,
			expires,
			signature,
		},
	}
	o.ws.Send(cmd)
}

func (o *WsPrivate) subscribe(topic string) bool {
	return o.ws.subscribe(topic)
}

func (o *WsPrivate) unsubscribe(topic string) bool {
	return o.ws.unsubscribe(topic)
}
