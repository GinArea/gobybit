package iperpetual

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type WsPrivate struct {
	WsSection
	key    string
	secret string
}

func NewWsPrivate(client *WsClient, key string, secret string) *WsPrivate {
	c := &WsPrivate{
		key:    key,
		secret: secret,
	}
	c.init(client)
	return c
}

func (this WsPrivate) Position() *WsExecutor[[]PositionShot] {
	return NewWsExecutor[[]PositionShot](&this.WsSection, Subscription{Topic: TopicPosition})
}

func (this WsPrivate) Execution() *WsExecutor[[]ExecutionShot] {
	return NewWsExecutor[[]ExecutionShot](&this.WsSection, Subscription{Topic: TopicExecution})
}

func (this WsPrivate) Order() *WsExecutor[[]OrderShot] {
	return NewWsExecutor[[]OrderShot](&this.WsSection, Subscription{Topic: TopicOrder})
}

func (this WsPrivate) StopOrder() *WsExecutor[[]StopOrderShot] {
	return NewWsExecutor[[]StopOrderShot](&this.WsSection, Subscription{Topic: TopicStopOrder})
}

func (this WsPrivate) Wallet() *WsExecutor[[]WalletShot] {
	return NewWsExecutor[[]WalletShot](&this.WsSection, Subscription{Topic: TopicWallet})
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
	this.ws.send(cmd)
}
