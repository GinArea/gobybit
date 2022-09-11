package transport

import (
	"bytes"
	"time"
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

func (this *WsClient) Shutdown() {
	this.ws.Shutdown()
}

func (this *WsClient) Conf() *WsConf {
	return this.ws.Conf()
}

func (this *WsClient) WithProxy(proxy string) *WsClient {
	this.Conf().SetProxy(proxy)
	return this
}

func (this *WsClient) Connected() bool {
	return this.ws.Connected()
}

func (this *WsClient) Run() {
	this.ws.SetOnMessage(this.processMessage)
	this.ws.Run()
	go func() {
		for this.ws.Do() {
			this.ws.Sleep(this.heartbeatTimeout)
			if !this.ws.Do() {
				break
			}
			if !this.ws.Connected() {
				continue
			}
			this.ping()
		}
	}()
}

func (this *WsClient) ID() string {
	return this.ws.ID()
}

func (this *WsClient) SetOnMessage(onMessage func(string, []byte)) {
	this.onMessage = onMessage
}

func (this *WsClient) SetOnConnected(onConnected func()) {
	this.ws.SetOnConnected(onConnected)
}

func (this *WsClient) SetOnDisconnected(onDisconnected func()) {
	this.ws.SetOnDisconnected(onDisconnected)
}

func (this *WsClient) Send(cmd any) bool {
	return this.ws.Send(cmd)
}

func (this *WsClient) ping() bool {
	return this.ws.Send(struct {
		Cmd string `json:"op"`
	}{
		Cmd: "ping",
	})
}

func (this *WsClient) processMessage(msg []byte) {
	prefix := []byte(`{"`)
	if bytes.HasPrefix(msg, prefix) {
		m := bytes.TrimPrefix(msg, prefix)
		i := bytes.IndexByte(m, '"')
		if i > -1 {
			name := string(m[:i])
			if this.onMessage != nil {
				this.onMessage(name, msg)
			}
			return
		}
	}
	panic("message type not detected")
}
