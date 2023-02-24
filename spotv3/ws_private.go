package spotv3

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ginarea/gobybit/transport"
	"github.com/msw-x/moon/ulog"
)

type WsPrivate struct {
	ws     *WsClient
	key    string
	secret string
}

func NewWsPrivate(key string, secret string) *WsPrivate {
	return &WsPrivate{
		ws:     NewWsClient("wss://stream.bybit.com/spot/private/v3"),
		key:    key,
		secret: secret,
	}
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

func (o *WsPrivate) WithProxy(proxy string) *WsPrivate {
	o.Conf().SetProxy(proxy)
	return o
}

func (o *WsPrivate) Connected() bool {
	return o.ws.Connected()
}

func (o *WsPrivate) Run() {
	o.ws.SetOnConnected(func() {
		o.auth()
	})
	o.ws.Run()
}

func (o *WsPrivate) SetOnAuth(onAuth func(bool)) {
	o.ws.SetOnAuth(onAuth)
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

func (o *WsPrivate) SubscribeOutbound() bool {
	return o.ws.Subscribe(Subscription{Topic: TopicOutbound})
}
func (o *WsPrivate) UnsubscribeOutbound() bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicOutbound})
}

func (o *WsPrivate) SubscribeOrder() bool {
	return o.ws.Subscribe(Subscription{Topic: TopicOrder})
}
func (o *WsPrivate) UnsubscribeOrder() bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicOrder})
}

func (o *WsPrivate) SubscribeStopOrder() bool {
	return o.ws.Subscribe(Subscription{Topic: TopicStopOrder})
}
func (o *WsPrivate) UnsubscribeStopOrder() bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicStopOrder})
}

func (o *WsPrivate) SubscribeTicket() bool {
	return o.ws.Subscribe(Subscription{Topic: TopicTicket})
}
func (o *WsPrivate) UnsubscribeTicket() bool {
	return o.ws.Unsubscribe(Subscription{Topic: TopicTicket})
}
