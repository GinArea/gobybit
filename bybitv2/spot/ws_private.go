package spot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ginarea/gobybit/bybitv2/transport"
	"github.com/msw-x/moon/ulog"
)

type WsPrivate struct {
	log    *ulog.Log
	ws     *transport.WsClient
	key    string
	secret string
	userID string
	onAuth func(bool)
}

func NewWsPrivate(url string, key string, secret string) *WsPrivate {
	ws := transport.NewWsClient(url)
	return &WsPrivate{
		log:    ulog.Empty(),
		ws:     ws,
		key:    key,
		secret: secret,
	}
}

func (this *WsPrivate) Shutdown() {
	this.log.Debug("shutdown")
	this.ws.Shutdown()
}

func (this *WsPrivate) Conf() *transport.WsConf {
	return this.ws.Conf()
}

func (this *WsPrivate) WithLog(log *ulog.Log) *WsPrivate {
	this.ws.WithLog(log)
	return this
}

func (this *WsPrivate) WithProxy(proxy string) *WsPrivate {
	this.ws.WithProxy(proxy)
	return this
}

func (this *WsPrivate) Run() {
	this.log.Debug("run")
	this.ws.SetOnConnected(func() {
		this.auth()
	})
	this.ws.SetOnDisconnected(func() {
		this.userID = ""
	})
	this.ws.SetOnMessage(this.processMessage)
	this.ws.Run()
}

func (this *WsPrivate) SetOnAuth(onAuth func(bool)) {
	this.onAuth = onAuth
}

func (this *WsPrivate) auth() {
	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	sig := hmac.New(sha256.New, []byte(this.secret))
	sig.Write([]byte(req))
	signature := hex.EncodeToString(sig.Sum(nil))
	cmd := struct {
		Name string `json:"op"`
		Args []any  `json:"args"`
	}{
		Name: "auth",
		Args: []any{
			this.key,
			expires,
			signature,
		},
	}
	this.ws.Send(cmd)
}

func (this *WsPrivate) processMessage(name string, msg []byte) {
	switch name {
	case "pong":
		v := transport.JsonUnmarshal[struct {
			Pong string `json:"pong"`
		}](msg)
		this.log.Debug("pong:", v.Pong)
	case "auth":
		v := transport.JsonUnmarshal[struct {
			Auth   string `json:"auth"`
			UserID string `json:"userId"`
		}](msg)
		this.log.Infof("auth[%s] user[%s]", v.Auth, v.UserID)
		success := v.Auth == "success"
		this.userID = v.UserID
		if this.onAuth != nil {
			this.onAuth(success)
		}
	case "code":
		v := transport.JsonUnmarshal[struct {
			Code        string `json:"code"`
			Description string `json:"desc"`
		}](msg)
		this.log.Warningf("code[%s]: %s", v.Code, v.Description)
	default:
		panic("unknown message type")
	}
}
