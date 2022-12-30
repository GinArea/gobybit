package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/msw-x/moon"
	"github.com/msw-x/moon/app"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/usync"
)

type WsConn struct {
	log            *ulog.Log
	do             *usync.Do
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
	return &WsConn{
		log:  ulog.Empty(),
		do:   usync.NewDo(),
		url:  url,
		conf: NewWsConf(),
	}
}

func (o *WsConn) Shutdown() {
	o.log.Debug("shutdown...")
	o.do.Cancel()
	if o.Connected() {
		o.ws.Close()
	}
	o.log.Debug("stop")
	o.do.Stop()
	o.log.Debug("shutdown completed")
}

func (o *WsConn) WithLog(log *ulog.Log) *WsConn {
	o.log = log
	return o
}

func (o *WsConn) Connected() bool {
	return o.ws != nil
}

func (o *WsConn) Conf() *WsConf {
	return o.conf
}

func (o *WsConn) Run() {
	app.Go(o.run)
}

func (o *WsConn) ID() string {
	id := fmt.Sprintf("%p", o)
	return id[5:]
}

func (o *WsConn) Do() bool {
	return o.do.Do()
}

func (o *WsConn) Sleep(duration time.Duration) {
	o.do.Sleep(duration)
}

func (o *WsConn) SetOnMessage(onMessage func([]byte)) {
	o.onMessage = onMessage
}

func (o *WsConn) SetOnConnected(onConnected func()) {
	o.onConnected = onConnected
}

func (o *WsConn) SetOnDisconnected(onDisconnected func()) {
	o.onDisconnected = onDisconnected
}

func (o *WsConn) Send(v any) bool {
	if !o.do.Do() {
		return false
	}
	bin, err := json.Marshal(v)
	if err != nil {
		o.log.Errorf("json: %v", err)
		return false
	}
	err = o.write(bin)
	if err != nil {
		o.log.Errorf("send: %v", err)
		o.setConnected(nil)
		return false
	}
	return true
}

func (o *WsConn) run() {
	if o.url == "" {
		o.log.Warning("disabled")
		o.do.Cancel()
		return
	}
	defer moon.Recover(func(err string) {
		o.log.Errorf(err)
	})
	defer o.log.Debug("completed")
	defer o.do.Notify()
	o.log.Debug("run")
	for o.do.Do() {
		o.connectAndRun()
	}
}

func (o *WsConn) connectAndRun() {
	o.log.Info("dial:", o.url)
	dialer := websocket.Dialer{
		HandshakeTimeout: o.conf.HandshakeTimeout,
	}
	if o.conf.Proxy != nil {
		o.log.Debug("proxy:", o.conf.Proxy)
		dialer.Proxy = http.ProxyURL(o.conf.Proxy)
	}
	c, _, err := dialer.Dial(o.url, nil)
	if err != nil {
		o.log.Error("dial:", err)
		o.do.Sleep(time.Second * 10)
		return
	}
	defer c.Close()
	defer o.setConnected(nil)
	o.setConnected(c)
	o.ws.SetPongHandler(func(text string) error {
		o.log.Debug("pong")
		return nil
	})
	for o.do.Do() {
		if !o.Connected() {
			break
		}
		var err error
		o.msg, err = o.read()
		if err != nil {
			if o.do.Do() {
				o.log.Error("read:", err)
			}
			break
		}
		if len(o.msg) > 0 {
			err = o.processMessage()
			if err != nil {
				o.log.Error("process message:", err)
			}
		}
	}
}

func (o *WsConn) setConnected(ws *websocket.Conn) {
	if o.ws != ws {
		o.ws = ws
		if ws == nil {
			o.log.Info("disconnected")
			if o.onDisconnected != nil {
				o.onDisconnected()
			}
		} else {
			o.log.Info("connected")
			if o.onConnected != nil {
				o.onConnected()
			}
		}
	}
}

func (o *WsConn) ping() error {
	ws := o.ws
	if ws == nil {
		return errors.New("empty socket")
	}
	return o.ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(o.conf.WriteTimeout))
}

func (o *WsConn) read() ([]byte, error) {
	ws := o.ws
	if ws == nil {
		return []byte{}, errors.New("empty socket")
	}
	ws.SetReadDeadline(time.Now().Add(o.conf.ReadTimeout))
	_, msg, err := ws.ReadMessage()
	if o.conf.LogRecv && o.do.Do() {
		o.log.Debugf("recv: %d B: %s", len(msg), string(msg))
	}
	return msg, err
}

func (o *WsConn) write(buf []byte) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	ws := o.ws
	if ws == nil {
		return errors.New("empty socket")
	}
	ws.SetWriteDeadline(time.Now().Add(o.conf.WriteTimeout))
	err := ws.WriteMessage(websocket.TextMessage, buf)
	if o.conf.LogSent {
		o.log.Debugf("sent: %d B: %s", len(buf), string(buf))
	}
	return err
}

func (o *WsConn) processMessage() (err error) {
	defer moon.Recover(func(err string) {
		o.log.Error("process message:", err)
	})
	if o.onMessage != nil {
		o.onMessage(o.msg)
	}
	return nil
}
