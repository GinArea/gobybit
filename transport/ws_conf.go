package transport

import (
	"fmt"
	"net/url"
	"time"
)

type WsConf struct {
	Proxy            *url.URL
	HandshakeTimeout time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	LogRecv          bool
	LogSent          bool
}

func NewWsConf() *WsConf {
	return &WsConf{
		HandshakeTimeout: time.Second * 10,
		ReadTimeout:      time.Second * 30,
		WriteTimeout:     time.Second * 5,
	}
}

func (o *WsConf) SetProxy(proxy string) {
	var err error
	o.Proxy, err = ParseProxy(proxy)
	if err != nil {
		panic(fmt.Sprintf("set proxy fail: %v", err))
	}
}

func (o *WsConf) ResetProxy() {
	o.Proxy = nil
}
