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

func (o WsPrivate) Position() *WsExecutor[[]PositionShot] {
	return NewWsExecutor[[]PositionShot](&o.WsSection, Subscription{Topic: TopicPosition})
}

func (o WsPrivate) Execution() *WsExecutor[[]ExecutionShot] {
	return NewWsExecutor[[]ExecutionShot](&o.WsSection, Subscription{Topic: TopicExecution})
}

func (o WsPrivate) Order() *WsExecutor[[]OrderShot] {
	return NewWsExecutor[[]OrderShot](&o.WsSection, Subscription{Topic: TopicOrder})
}

func (o WsPrivate) StopOrder() *WsExecutor[[]StopOrderShot] {
	return NewWsExecutor[[]StopOrderShot](&o.WsSection, Subscription{Topic: TopicStopOrder})
}

func (o WsPrivate) Wallet() *WsExecutor[[]WalletShot] {
	return NewWsExecutor[[]WalletShot](&o.WsSection, Subscription{Topic: TopicWallet})
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
	o.ws.send(cmd)
}
