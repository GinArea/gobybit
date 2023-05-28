package bybitv5

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type Subscriptions[T SubscriptionTopic] struct {
	c     SubscriptionClient
	mutex sync.Mutex
	funcs SubscriptionFuncs[T]
}

func NewSubscriptions[T SubscriptionTopic](c SubscriptionClient) *Subscriptions[T] {
	o := new(Subscriptions[T])
	o.c = c
	o.funcs = make(SubscriptionFuncs[T])
	return o
}

func (o *Subscriptions[T]) subscribe(topic string, f SubscriptionFunc[T]) {
	if o.c.Ready() {
		o.c.subscribe(topic)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.funcs[topic] = f
}

func (o *Subscriptions[T]) unsubscribe(topic string) {
	if o.c.Ready() {
		o.c.unsubscribe(topic)
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	delete(o.funcs, topic)
}

func (o *Subscriptions[T]) subscribeAll() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	for topic, _ := range o.funcs {
		o.c.subscribe(topic)
	}
}

func (o *Subscriptions[T]) processTopic(data []byte) (err error) {
	var topic T
	err = json.Unmarshal(data, &topic)
	if err == nil {
		f := o.getFunc(topic.Name())
		if f == nil {
			err = fmt.Errorf("subscription of topic[%s] not found", topic.Name())
		} else {
			err = f(topic)
		}
	}
	return
}

func (o *Subscriptions[T]) getFunc(name string) (f SubscriptionFunc[T]) {
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

type SubscriptionTopic interface {
	Name() string
	RawData() []byte
}

type SubscriptionClient interface {
	Ready() bool
	subscribe(string)
	unsubscribe(string)
}

type SubscriptionFunc[T SubscriptionTopic] func(T) error

type SubscriptionFuncs[T SubscriptionTopic] map[string]SubscriptionFunc[T]
