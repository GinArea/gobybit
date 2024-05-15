package bybitv5

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsTrade struct {
	c          *WsAuthClient[WsTradeResponse]
	ts         *Timestamp
	referer    string
	onResponse func(WsTradeResponse)
}

func NewWsTrade(key, secret string) *WsTrade {
	o := new(WsTrade)
	o.c = NewWsAuthClient[WsTradeResponse]("trade", key, secret)
	o.ts = NewTimestamp()
	o.c.SetOnOperation(o.onOperation)
	return o
}

func (o *WsTrade) Close() {
	o.c.Close()
}

func (o *WsTrade) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsTrade) WithLog(log *ulog.Log) *WsTrade {
	o.c.WithLog(log)
	return o
}

func (o *WsTrade) WithBaseUrl(url string) *WsTrade {
	o.c.WithBaseUrl(url)
	return o
}

func (o *WsTrade) WithByTickUrl() *WsTrade {
	o.c.WithByTickUrl()
	return o
}

func (o *WsTrade) WithProxy(proxy string) *WsTrade {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsTrade) WithLogRequest(enable bool) *WsTrade {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsTrade) WithLogResponse(enable bool) *WsTrade {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsTrade) WithOnDialError(f func(error) bool) *WsTrade {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsTrade) WithOnReady(f func()) *WsTrade {
	o.c.WithOnReady(f)
	return o
}

func (o *WsTrade) WithOnConnected(f func()) *WsTrade {
	o.c.WithOnConnected(f)
	return o
}

func (o *WsTrade) WithOnDisconnected(f func()) *WsTrade {
	o.c.WithOnDisconnected(f)
	return o
}

func (o *WsTrade) WithOnResponse(f func(WsTradeResponse)) *WsTrade {
	o.onResponse = f
	return o
}

func (o *WsTrade) Run() {
	o.c.Run()
}

func (o *WsTrade) Connected() bool {
	return o.c.Connected()
}

func (o *WsTrade) Ready() bool {
	return o.c.Ready()
}

func (o *WsTrade) PlaceOrder(v WsPlaceOrder) {
	ts, window := o.ts.Get()
	o.c.SendOrder(WsRequest[WsPlaceOrder]{
		Operation: "order.create",
		Header: &WsHeader{
			Timestamp:  ts,
			RecvWindow: window,
			Referer:    o.referer,
		},
		Args: []WsPlaceOrder{v},
	})
}

func (o *WsTrade) onOperation(r WsTradeResponse) (ok bool, err error) {
	if r.OperationIs("order.create") {
		ok = true
		if o.onResponse != nil {
			o.onResponse(r)
		}
	}
	return
}
