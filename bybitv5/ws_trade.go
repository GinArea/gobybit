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

// TradeRequest sends an arbitrary trade-WS operation
// (order.create / order.amend / order.cancel) with an explicit reqId:
// the API echoes it in the ack (<= 36 chars, must be unique — a
// duplicate is rejected with code 20006).
func (o *WsTrade) TradeRequest(op string, reqId string, arg any) {
	ts, window := o.ts.Get()
	o.c.SendJson(WsRequest[any]{
		RequestId: reqId,
		Header: &WsHeader{
			Timestamp:  ts,
			RecvWindow: window,
			Referer:    o.referer,
		},
		Operation: op,
		Args:      []any{arg},
	})
}

func (o *WsTrade) onOperation(r WsTradeResponse) (ok bool, err error) {
	if r.OperationIs("order.create") || r.OperationIs("order.amend") || r.OperationIs("order.cancel") {
		ok = true
		if o.onResponse != nil {
			o.onResponse(r)
		}
	}
	return
}
