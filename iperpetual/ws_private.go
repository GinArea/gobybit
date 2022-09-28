package iperpetual

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ginarea/gobybit/transport"
)

type WsPrivate struct {
	ws     *WsClient
	key    string
	secret string
}

func NewWsPrivate(key string, secret string) *WsPrivate {
	return &WsPrivate{
		ws:     NewWsClient("private"),
		key:    key,
		secret: secret,
	}
}

func (this *WsPrivate) Shutdown() {
	this.ws.Shutdown()
}

func (this *WsPrivate) Conf() *transport.WsConf {
	return this.ws.Conf()
}

func (this *WsPrivate) WithProxy(proxy string) *WsPrivate {
	this.Conf().SetProxy(proxy)
	return this
}

func (this *WsPrivate) Connected() bool {
	return this.ws.Connected()
}

func (this *WsPrivate) Run() {
	this.ws.SetOnConnected(func() {
		this.auth()
	})
	this.ws.Run()
}

func (this *WsPrivate) SetOnAuth(onAuth func(bool)) {
	this.ws.SetOnAuth(onAuth)
}

func (this *WsPrivate) auth() {
	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	sig := hmac.New(sha256.New, []byte(this.secret))
	sig.Write([]byte(req))
	signature := hex.EncodeToString(sig.Sum(nil))
	cmd := struct {
		Name string `json:"op"`
		Args []any  `json:"args"`
	}{
		Name: "auth",
		Args: []any{
			this.key,
			expires,
			signature,
		},
	}
	this.ws.Send(cmd)
}

func (this *WsPrivate) SubscribePosition() bool {
	return this.ws.Subscribe(Subscription{Topic: TopicPosition})
}
func (this *WsPrivate) UnsubscribePosition() bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicPosition})
}

func (this *WsPrivate) SubscribeExecution() bool {
	return this.ws.Subscribe(Subscription{Topic: TopicExecution})
}
func (this *WsPrivate) UnsubcribeExecution() bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicExecution})
}

func (this *WsPrivate) SubscribeOrder() bool {
	return this.ws.Subscribe(Subscription{Topic: TopicOrder})
}
func (this *WsPrivate) UnsubscribeOrder() bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicOrder})
}

func (this *WsPrivate) SubscribeStopOrder() bool {
	return this.ws.Subscribe(Subscription{Topic: TopicStopOrder})
}
func (this *WsPrivate) UnsubscribeStopOrder() bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicStopOrder})
}

func (this *WsPrivate) SubscribeWallet() bool {
	return this.ws.Subscribe(Subscription{Topic: TopicWallet})
}
func (this *WsPrivate) UnsubscribeWallet() bool {
	return this.ws.Unsubscribe(Subscription{Topic: TopicWallet})
}
