package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/msw-x/moon/ulog"
)

type WsPrivate struct {
	log    *ulog.Log
	ws     *WsClient
	key    string
	secret string
	userID string
	onAuth func(bool)
}

func NewWsPrivate(url string, key string, secret string) *WsPrivate {
	ws := NewWsClient(url)
	return &WsPrivate{
		log:    ulog.NewLog(fmt.Sprintf("ws-private[%s]", ws.ID())),
		ws:     ws,
		key:    key,
		secret: secret,
	}
}

func (this *WsPrivate) Shutdown() {
	this.log.Info("shutdown")
	this.ws.Shutdown()
}

func (this *WsPrivate) Conf() *WsConf {
	return this.ws.Conf()
}

func (this *WsPrivate) WithProxy(proxy string) *WsPrivate {
	this.ws.WithProxy(proxy)
	return this
}

func (this *WsPrivate) Run() {
	this.log.Info("run")
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
		v := jsonUnmarshal[struct {
			Pong string `json:"pong"`
		}](msg)
		this.log.Debug("pong:", v.Pong)
	case "auth":
		v := jsonUnmarshal[struct {
			Auth   string `json:"auth"`
			UserId string `json:"userId"`
		}](msg)
		this.log.Infof("auth[%s] user[%s]", v.Auth, v.UserId)
		success := v.Auth == "success"
		this.userID = v.UserId
		if this.onAuth != nil {
			this.onAuth(success)
		}
	case "code":
		v := jsonUnmarshal[struct {
			Code        string `json:"code"`
			Description string `json:"desc"`
		}](msg)
		this.log.Infof("code[%s]: %s", v.Code, v.Description)
	default:
		panic("unknown message type")
	}
}
