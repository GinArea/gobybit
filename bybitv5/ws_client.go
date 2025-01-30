package bybitv5

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsClient[T WsResponse] struct {
	c          *uws.Client
	onResponce func(T) error
	onTopic    func([]byte) error
}

func NewWsClient[T WsResponse]() *WsClient[T] {
	o := new(WsClient[T])
	o.c = uws.NewClient(MainBaseWsUrl)
	return o
}

func (o *WsClient[T]) Close() {
	o.c.Close()
}

func (o *WsClient[T]) Log() *ulog.Log {
	return o.c.Log()
}

func (o *WsClient[T]) Transport() *uws.Options {
	return &o.c.Options
}

func (o *WsClient[T]) WithLog(log *ulog.Log) *WsClient[T] {
	o.c.WithLog(log)
	return o
}

func (o *WsClient[T]) WithBaseUrl(url string) *WsClient[T] {
	o.c.WithBase(url)
	return o
}

func (o *WsClient[T]) WithByTickUrl() *WsClient[T] {
	return o.WithBaseUrl(MainBaseByTickWsUrl)
}

func (o *WsClient[T]) WithPath(path string) *WsClient[T] {
	o.c.WithPath(path)
	return o
}

func (o *WsClient[T]) WithProxy(proxy string) *WsClient[T] {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsClient[T]) WithLogRequest(enable bool) *WsClient[T] {
	o.Transport().LogSent.Size = enable
	o.Transport().LogSent.Data = enable
	return o
}

func (o *WsClient[T]) WithLogResponse(enable bool) *WsClient[T] {
	o.Transport().LogRecv.Size = enable
	o.Transport().LogRecv.Data = enable
	return o
}

func (o *WsClient[T]) WithOnDialDelay(f func() time.Duration) *WsClient[T] {
	o.c.WithOnDialDelay(f)
	return o
}

func (o *WsClient[T]) WithOnDialError(f func(error) bool) *WsClient[T] {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsClient[T]) WithOnConnected(f func()) *WsClient[T] {
	o.c.WithOnConnected(f)
	return o
}

func (o *WsClient[T]) WithOnDisconnected(f func()) *WsClient[T] {
	o.c.WithOnDisconnected(f)
	return o
}

func (o *WsClient[T]) WithOnResponse(f func(T) error) *WsClient[T] {
	o.onResponce = f
	return o
}

func (o *WsClient[T]) WithOnTopic(f func([]byte) error) *WsClient[T] {
	o.onTopic = f
	return o
}

func (o *WsClient[T]) Run() {
	o.c.WithOnPing(o.ping)
	o.c.WithOnMessage(o.onMessage)
	o.c.Run()
}

func (o *WsClient[T]) Connected() bool {
	return o.c.Connected()
}

func (o *WsClient[T]) Send(r WsRequest[string]) {
	o.c.SendJson(r)
}

func (o *WsClient[T]) SendOrder(r WsRequest[WsPlaceOrder]) {
	o.c.SendJson(r)
}

func (o *WsClient[T]) Subscribe(s string) {
	o.SubscribeBatch([]string{s})
}

func (o *WsClient[T]) SubscribeBatch(l []string) {
	o.Send(WsRequest[string]{
		Operation: "subscribe",
		Args:      l,
	})
}

func (o *WsClient[T]) Unsubscribe(s string) {
	o.UnsubscribeBatch([]string{s})
}

func (o *WsClient[T]) UnsubscribeBatch(l []string) {
	o.Send(WsRequest[string]{
		Operation: "unsubscribe",
		Args:      l,
	})
}
func (o *WsClient[T]) ping() {
	o.Send(WsRequest[string]{
		Operation: "ping",
	})
}

func (o *WsClient[T]) onMessage(messageType int, data []byte) {
	log := o.c.Log()
	if messageType != websocket.TextMessage {
		log.Warning("invalid message type:", uws.MessageTypeString(messageType))
		return
	}
	var r T
	err := json.Unmarshal(data, &r)
	if err == nil {
		if r.IsOperateion() {
			if o.onResponce != nil {
				err = o.onResponce(r)
			}
		} else {
			if o.onTopic != nil {
				err = o.onTopic(data)
			}
		}
	}
	if err != nil {
		log.Error(err)
	}
}
