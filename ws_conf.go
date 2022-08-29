package bybit

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
	LogIO            bool
}

func NewWsConf() *WsConf {
	return &WsConf{
		HandshakeTimeout: time.Second * 10,
		ReadTimeout:      time.Second * 30,
		WriteTimeout:     time.Second * 5,
		LogIO:            false,
	}
}

func (this *WsConf) SetProxy(proxy string) {
	var err error
	this.Proxy, err = ParseProxy(proxy)
	if err != nil {
		panic(fmt.Sprintf("set proxy fail: %v", err))
	}
}
