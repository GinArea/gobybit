package bybitv5

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPrivate struct {
	c             *WsAuthClient[WsBaseResponse]
	onReady       func()
	subscriptions *Subscriptions
}

func NewWsPrivate(key, secret string) *WsPrivate {
	o := new(WsPrivate)
	o.c = NewWsAuthClient[WsBaseResponse]("private", key, secret)
	o.subscriptions = NewSubscriptions(o)
	o.c.WithOnReady(func() {
		o.subscriptions.subscribeAll()
		if o.onReady != nil {
			o.onReady()
		}
	})
	o.c.SetOnTopic(o.subscriptions.processTopic)
	return o
}

func (o *WsPrivate) Close() {
	o.c.Close()
}

func (o *WsPrivate) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsPrivate) WithLog(log *ulog.Log) *WsPrivate {
	o.c.WithLog(log)
	return o
}

func (o *WsPrivate) WithBaseUrl(url string) *WsPrivate {
	o.c.WithBaseUrl(url)
	return o
}

func (o *WsPrivate) WithByTickUrl() *WsPrivate {
	o.c.WithByTickUrl()
	return o
}

func (o *WsPrivate) WithProxy(proxy string) *WsPrivate {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsPrivate) WithLogRequest(enable bool) *WsPrivate {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsPrivate) WithLogResponse(enable bool) *WsPrivate {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsPrivate) WithOnDialError(f func(error) bool) *WsPrivate {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsPrivate) WithOnReady(f func()) *WsPrivate {
	o.onReady = f
	return o
}

func (o *WsPrivate) WithOnConnected(f func()) *WsPrivate {
	o.c.WithOnConnected(f)
	return o
}

func (o *WsPrivate) WithOnDisconnected(f func()) *WsPrivate {
	o.c.WithOnDisconnected(f)
	return o
}

func (o *WsPrivate) Run() {
	o.c.Run()
}

func (o *WsPrivate) Connected() bool {
	return o.c.Connected()
}

func (o *WsPrivate) Ready() bool {
	return o.c.Ready()
}

func (o *WsPrivate) Position() *Executor[[]PositionShot] {
	return NewExecutor[[]PositionShot]("position", o.subscriptions)
}

func (o *WsPrivate) Execution() *Executor[[]ExecutionShot] {
	return NewExecutor[[]ExecutionShot]("execution", o.subscriptions)
}

func (o *WsPrivate) Order() *Executor[[]OrderShot] {
	return NewExecutor[[]OrderShot]("order", o.subscriptions)
}

func (o *WsPrivate) Wallet() *Executor[[]WalletShot] {
	return NewExecutor[[]WalletShot]("wallet", o.subscriptions)
}

func (o *WsPrivate) Greek() *Executor[[]GreekShot] {
	return NewExecutor[[]GreekShot]("greeks", o.subscriptions)
}

func (o *WsPrivate) subscribe(topic string) {
	o.c.Subscribe(topic)
}

func (o *WsPrivate) unsubscribe(topic string) {
	o.c.Unsubscribe(topic)
}
