package bybitv5

import (
	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPrivate struct {
	c              *WsClient
	s              *Sign
	ready          bool
	onReady        func()
	onConnected    func()
	onDisconnected func()
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
	return o.WithBaseUrl(MainBaseByTickWsUrl)
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

func (o *WsPrivate) WithOnReady(f func()) {
	o.onReady = f
}

func (o *WsPrivate) WithOnConnected(f func()) {
	o.onConnected = f
}

func (o *WsPrivate) WithOnDisconnected(f func()) {
	o.onDisconnected = f
}

func (o *WsPrivate) WithAuth(key, secret string) *WsPrivate {
	o.s = NewSign(key, secret)
	return o
}

func (o *WsPrivate) Run() {
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

func (o *WsPrivate) Connected() bool {
	return o.c.Connected()
}

func (o *WsPrivate) Ready() bool {
	return o.ready
}

func (o *WsPrivate) Position() {

}

func (o *WsPrivate) Execution() {

}

func (o *WsPrivate) Order() {

}

func (o *WsPrivate) Wallet() {

}

func (o *WsPrivate) Greek() {

}

func (o *WsPrivate) auth() {
	o.c.Send(WsRequest{
		Operation: "auth",
		Args:      o.s.WebSocket(),
	})
}

func (o *WsPrivate) setReady() {
	o.ready = true
	if o.onReady != nil {
		o.onReady()
	}
}

func (o *WsPrivate) onResponse(r WsResponse) error {
	log := o.c.Log()
	if r.Operation == "auth" {
		log.Info("auth:", ufmt.SuccessFailure(r.Success))
		if r.Success {
			o.setReady()
			//o.private.subscribeAll()
		}
	}
	r.Log(log)
	return nil
}

func (o *WsPrivate) onTopic(data []byte) error {
	return nil
}
