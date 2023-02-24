package spotv3

import (
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
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if o.ws.Ready() {
		o.ws.subscribe(topic)
	}
	o.subscriptions[topic] = f
}

func (o *WsSection) unsubscribe(topic string) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if o.ws.Ready() {
		o.ws.unsubscribe(topic)
	}
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
	f, _ := o.subscriptions[m.Topic]
	ok = f != nil
	if ok {
		err = f(m.Bin, m.Delta)
	}
	return
}

type SubscriptionFunc func(m []byte, delta bool) error

type Subscriptions map[string]SubscriptionFunc
