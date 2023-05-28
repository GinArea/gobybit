package bybitv5

type Executor[T SubscriptionTopic, V any] struct {
	topic         string
	subscriptions *Subscriptions[T]
}

func NewExecutor[T SubscriptionTopic, V any](topic string, subscriptions *Subscriptions[T]) *Executor[T, V] {
	o := new(Executor[T, V])
	o.topic = topic
	o.subscriptions = subscriptions
	return o
}

func (o *Executor[T, V]) Subscribe(onShot func(V)) {
	o.subscriptions.subscribe(o.topic, func(topic T) error {
		return WsFunc(topic.RawData(), onShot)
	})
}

func (o *Executor[T, V]) Unsubscribe() {
	o.subscriptions.unsubscribe(o.topic)
}
