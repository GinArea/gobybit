package bybitv5

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type Subscriptions struct {
	c     SubscriptionClient
	mutex sync.Mutex
	funcs SubscriptionFuncs
}

func NewSubscriptions(c SubscriptionClient) *Subscriptions {
	o := new(Subscriptions)
	o.c = c
	o.funcs = make(SubscriptionFuncs)
	return o
}

func (o *Subscriptions) subscribe(topic string, f SubscriptionFunc) {
	if o.c.Ready() {
		o.c.subscribe(topic)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.funcs[topic] = f
}

func (o *Subscriptions) unsubscribe(topic string) {
	if o.c.Ready() {
		o.c.unsubscribe(topic)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	delete(o.funcs, topic)
}

func (o *Subscriptions) subscribeAll() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for topic, _ := range o.funcs {
		o.c.subscribe(topic)
	}
}

func (o *Subscriptions) processTopic(data []byte) (err error) {
	var topic RawTopic
	err = json.Unmarshal(data, &topic)
	if err == nil {
		f := o.getFunc(topic.Topic)
		if f == nil {
			err = fmt.Errorf("subscription of topic[%s] not found", topic.Topic)
		} else {
			err = f(topic)
		}
	}
	return
}

func (o *Subscriptions) getFunc(name string) (f SubscriptionFunc) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for topic, fn := range o.funcs {
		if strings.HasPrefix(name, topic) {
			f = fn
			break
		}
	}
	return
}

type SubscriptionClient interface {
	Ready() bool
	subscribe(string)
	unsubscribe(string)
}

type SubscriptionFunc func(RawTopic) error

type SubscriptionFuncs map[string]SubscriptionFunc
