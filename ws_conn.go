package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/msw-x/moon"
	"github.com/msw-x/moon/syn"
	"github.com/msw-x/moon/ulog"
)

type WsConn struct {
	log            *ulog.Log
	do             *syn.Do
	url            string
	ws             *websocket.Conn
	mutex          sync.Mutex
	msg            []byte
	conf           *WsConf
	onMessage      func([]byte)
	onConnected    func()
	onDisconnected func()
}

func NewWsConn(url string) *WsConn {
	conn := &WsConn{
		do:   syn.NewDo(),
		url:  url,
		conf: NewWsConf(),
	}
	conn.log = ulog.NewLog(fmt.Sprintf("ws-conn[%s]", conn.ID()))
	return conn
}

func (this *WsConn) Shutdown() {
	this.log.Debug("shutdown...")
	this.do.Cancel()
	if this.Connected() {
		this.ws.Close()
	}
	this.log.Debug("stop")
	this.do.Stop()
	this.log.Debug("shutdown completed")
}

func (this *WsConn) Connected() bool {
	return this.ws != nil
}

func (this *WsConn) Conf() *WsConf {
	return this.conf
}

func (this *WsConn) Run() {
	go this.run()
}

func (this *WsConn) ID() string {
	id := fmt.Sprintf("%p", this)
	return id[5:]
}

func (this *WsConn) Do() bool {
	return this.do.Do()
}

func (this *WsConn) Sleep(duration time.Duration) {
	this.do.Sleep(duration)
}

func (this *WsConn) SetOnMessage(onMessage func([]byte)) {
	this.onMessage = onMessage
}

func (this *WsConn) SetOnConnected(onConnected func()) {
	this.onConnected = onConnected
}

func (this *WsConn) SetOnDisconnected(onDisconnected func()) {
	this.onDisconnected = onDisconnected
}

func (this *WsConn) Send(v any) bool {
	if !this.do.Do() {
		return false
	}
	bin, err := json.Marshal(v)
	if err != nil {
		this.log.Errorf("json: %v", err)
		return false
	}
	err = this.write(bin)
	if err != nil {
		this.log.Errorf("send: %v", err)
		this.setConnected(nil)
		return false
	}
	return true
}

func (this *WsConn) run() {
	if this.url == "" {
		this.log.Warning("disabled")
		this.do.Cancel()
		return
	}
	defer moon.Recover(func(err string) {
		this.log.Errorf(err)
	})
	defer this.log.Debug("completed")
	defer this.do.Notify()
	this.log.Debug("run")
	for this.do.Do() {
		this.connectAndRun()
	}
}

func (this *WsConn) connectAndRun() {
	this.log.Info("dial:", this.url)
	dialer := websocket.Dialer{
		HandshakeTimeout: this.conf.HandshakeTimeout,
	}
	if this.conf.Proxy != nil {
		this.log.Debug("proxy:", this.conf.Proxy)
		dialer.Proxy = http.ProxyURL(this.conf.Proxy)
	}
	c, _, err := dialer.Dial(this.url, nil)
	if err != nil {
		this.log.Error("dial:", err)
		this.do.Sleep(time.Second * 10)
		return
	}
	defer c.Close()
	defer this.setConnected(nil)
	this.setConnected(c)
	this.ws.SetPongHandler(func(text string) error {
		this.log.Debug("pong")
		return nil
	})
	for this.do.Do() {
		if !this.Connected() {
			break
		}
		var err error
		this.msg, err = this.read()
		if err != nil {
			if this.do.Do() {
				this.log.Error("read:", err)
			}
			break
		}
		if len(this.msg) > 0 {
			err = this.processMessage()
			if err != nil {
				this.log.Error("process message:", err)
			}
		}
	}
}

func (this *WsConn) setConnected(ws *websocket.Conn) {
	if this.ws != ws {
		this.ws = ws
		if ws == nil {
			this.log.Info("disconnected")
			if this.onDisconnected != nil {
				this.onDisconnected()
			}
		} else {
			this.log.Info("connected")
			if this.onConnected != nil {
				this.onConnected()
			}
		}
	}
}

func (this *WsConn) ping() error {
	ws := this.ws
	if ws == nil {
		return errors.New("empty socket")
	}
	return this.ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(this.conf.WriteTimeout))
}

func (this *WsConn) read() ([]byte, error) {
	ws := this.ws
	if ws == nil {
		return []byte{}, errors.New("empty socket")
	}
	ws.SetReadDeadline(time.Now().Add(this.conf.ReadTimeout))
	_, msg, err := ws.ReadMessage()
	if this.conf.LogIO && this.do.Do() {
		this.log.Debugf("recv: %d B: %s\n", len(msg), string(msg))
	}
	return msg, err
}

func (this *WsConn) write(buf []byte) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	ws := this.ws
	if ws == nil {
		return errors.New("empty socket")
	}
	ws.SetWriteDeadline(time.Now().Add(this.conf.WriteTimeout))
	err := ws.WriteMessage(websocket.TextMessage, buf)
	if this.conf.LogIO {
		this.log.Debugf("sent: %d B: %s\n", len(buf), string(buf))
	}
	return err
}

func (this *WsConn) processMessage() (err error) {
	defer moon.Recover(func(err string) {
		this.log.Error("process message:", err)
	})
	if this.onMessage != nil {
		this.onMessage(this.msg)
	}
	return nil
}
