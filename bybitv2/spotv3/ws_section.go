package spotv3

import (
	"strings"
	"sync"
)

type WsSectionClient interface {
	Ready() bool
	subscribe(string) bool
	unsubscribe(string) bool
}

type WsSection struct {
	ws            WsSectionClient
	mutex         sync.Mutex
	subscriptions Subscriptions
}

func NewWsSection(ws WsSectionClient) *WsSection {
	return &WsSection{
		ws:            ws,
		subscriptions: make(Subscriptions),
	}
}

func (o *WsSection) subscribe(topic string, f SubscriptionFunc) {
	if o.ws.Ready() {
		o.ws.subscribe(topic)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.subscriptions[topic] = f
}

func (o *WsSection) unsubscribe(topic string) {
	if o.ws.Ready() {
		o.ws.unsubscribe(topic)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	delete(o.subscriptions, topic)
}

func (o *WsSection) subscribeAll() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for topic, _ := range o.subscriptions {
		o.ws.subscribe(topic)
	}
}

func (o *WsSection) processTopic(m TopicMessage) (ok bool, err error) {
	f := o.getFunc(m.Topic)
	ok = f != nil
	if ok {
		err = f(m.Bin, m.Delta)
	}
	return
}

func (o *WsSection) getFunc(name string) (f SubscriptionFunc) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for topic, fn := range o.subscriptions {
		if strings.HasPrefix(name, topic) {
			f = fn
			break
		}
	}
	return
}

type SubscriptionFunc func(m []byte, delta bool) error

type Subscriptions map[string]SubscriptionFunc
