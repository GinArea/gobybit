package transport

import (
	"bytes"
	"time"

	"github.com/msw-x/moon/app"
	"github.com/msw-x/moon/ulog"
)

type WsClient struct {
	ws               *WsConn
	heartbeatTimeout time.Duration
	onMessage        func(string, []byte)
}

func NewWsClient(url string) *WsClient {
	return &WsClient{
		ws:               NewWsConn(url),
		heartbeatTimeout: time.Second * 20,
	}
}

func (o *WsClient) Shutdown() {
	o.ws.Shutdown()
}

func (o *WsClient) WithLog(log *ulog.Log) *WsClient {
	o.ws.WithLog(log)
	return o
}

func (o *WsClient) Conf() *WsConf {
	return o.ws.Conf()
}

func (o *WsClient) WithProxy(proxy string) *WsClient {
	o.Conf().SetProxy(proxy)
	return o
}

func (o *WsClient) Connected() bool {
	return o.ws.Connected()
}

func (o *WsClient) Run() {
	o.ws.SetOnMessage(o.processMessage)
	o.ws.Run()
	app.Go(func() {
		for o.ws.Do() {
			o.ws.Sleep(o.heartbeatTimeout)
			if !o.ws.Do() {
				break
			}
			if !o.ws.Connected() {
				continue
			}
			o.ping()
		}
	})
}

func (o *WsClient) SetOnMessage(onMessage func(string, []byte)) {
	o.onMessage = onMessage
}

func (o *WsClient) SetOnConnected(onConnected func()) {
	o.ws.SetOnConnected(onConnected)
}

func (o *WsClient) SetOnDisconnected(onDisconnected func()) {
	o.ws.SetOnDisconnected(onDisconnected)
}

func (o *WsClient) Send(cmd any) bool {
	return o.ws.Send(cmd)
}

func (o *WsClient) ping() bool {
	return o.ws.Send(struct {
		Cmd string `json:"op"`
	}{
		Cmd: "ping",
	})
}

func (o *WsClient) processMessage(msg []byte) {
	prefix := []byte(`{"`)
	if bytes.HasPrefix(msg, prefix) {
		m := bytes.TrimPrefix(msg, prefix)
		i := bytes.IndexByte(m, '"')
		if i > -1 {
			name := string(m[:i])
			if o.onMessage != nil {
				o.onMessage(name, msg)
			}
			return
		}
	}
	panic("message type not detected")
}
