package bybitv5

import (
	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsAuthClient[T WsResponse] struct {
	c              *WsClient[T]
	s              *Sign
	ready          bool
	onReady        func()
	onConnected    func()
	onDisconnected func()
	onTopic        func([]byte) error
	onOperation    func(T) (bool, error)
}

func NewWsAuthClient[T WsResponse](path, key, secret string) *WsAuthClient[T] {
	o := new(WsAuthClient[T])
	o.c = NewWsClient[T]()
	o.s = NewSign(key, secret)
	o.c.WithPath("v5/" + path)
	return o
}

func (o *WsAuthClient[T]) Close() {
	o.c.Close()
}

func (o *WsAuthClient[T]) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsAuthClient[T]) WithLog(log *ulog.Log) *WsAuthClient[T] {
	o.c.WithLog(log)
	return o
}

func (o *WsAuthClient[T]) WithBaseUrl(url string) *WsAuthClient[T] {
	o.c.WithBaseUrl(url)
	return o
}

func (o *WsAuthClient[T]) WithByTickUrl() *WsAuthClient[T] {
	return o.WithBaseUrl(MainBaseByTickWsUrl)
}

func (o *WsAuthClient[T]) WithProxy(proxy string) *WsAuthClient[T] {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsAuthClient[T]) WithLogRequest(enable bool) *WsAuthClient[T] {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsAuthClient[T]) WithLogResponse(enable bool) *WsAuthClient[T] {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsAuthClient[T]) WithOnDialError(f func(error) bool) *WsAuthClient[T] {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsAuthClient[T]) WithOnReady(f func()) *WsAuthClient[T] {
	o.onReady = f
	return o
}

func (o *WsAuthClient[T]) WithOnConnected(f func()) *WsAuthClient[T] {
	o.onConnected = f
	return o
}

func (o *WsAuthClient[T]) WithOnDisconnected(f func()) *WsAuthClient[T] {
	o.onDisconnected = f
	return o
}

func (o *WsAuthClient[T]) Run() {
	o.c.WithOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
		o.auth()
	})
	o.c.WithOnDisconnected(func() {
		o.ready = false
		if o.onDisconnected != nil {
			o.onDisconnected()
		}
	})
	o.c.WithOnResponse(o.onResponse)
	o.c.WithOnTopic(o.onTopic)
	o.c.Run()
}

func (o *WsAuthClient[T]) Connected() bool {
	return o.c.Connected()
}

func (o *WsAuthClient[T]) Ready() bool {
	return o.ready
}

func (o *WsAuthClient[T]) SetOnTopic(f func([]byte) error) {
	o.onTopic = f
}

func (o *WsAuthClient[T]) SetOnOperation(f func(T) (bool, error)) {
	o.onOperation = f
}

func (o *WsAuthClient[T]) SendOrder(r WsRequest[WsPlaceOrder]) {
	o.c.SendOrder(r)
}

func (o *WsAuthClient[T]) Subscribe(topic string) {
	o.c.Subscribe(topic)
}

func (o *WsAuthClient[T]) Unsubscribe(topic string) {
	o.c.Unsubscribe(topic)
}

func (o *WsAuthClient[T]) auth() {
	o.c.Send(WsRequest[string]{
		Operation: "auth",
		Args:      o.s.WebSocket(),
	})
}

func (o *WsAuthClient[T]) onResponse(r T) error {
	log := o.c.Log()
	if r.OperationIs("auth") {
		log.Info("auth:", ufmt.SuccessFailure(r.Ok()))
		if r.Ok() {
			o.ready = true
			if o.onReady != nil {
				o.onReady()
			}
		}
	} else if !r.OperationIs("pong") {
		if r.Ok() && o.onOperation != nil {
			ok, err := o.onOperation(r)
			if ok || err != nil {
				return err
			}
		}
		r.Log(log)
	}
	return nil
}
