package iperpetual

import (
	"sync"
)

type WsSection struct {
	ws            *WsClient
	mutex         sync.Mutex
	subscriptions Subscriptions
}

func (this *WsSection) init(client *WsClient) {
	this.ws = client
	this.subscriptions = make(Subscriptions)
}

func (this *WsSection) subscribe(topic string, f SubscriptionFunc) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.ws.Ready() {
		this.ws.subscribe(topic)
	}
	this.subscriptions[topic] = f
}

func (this *WsSection) unsubscribe(topic string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.ws.Ready() {
		this.ws.unsubscribe(topic)
	}
	delete(this.subscriptions, topic)
}

func (this *WsSection) subscribeAll() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	for topic, _ := range this.subscriptions {
		this.ws.subscribe(topic)
	}
}

func (this *WsSection) processTopic(m TopicMessage) (ok bool, err error) {
	f, _ := this.subscriptions[m.Topic]
	ok = f != nil
	if ok {
		err = f(m.Bin, m.Delta)
	}
	return
}

type SubscriptionFunc func(m []byte, delta bool) error

type Subscriptions map[string]SubscriptionFunc
