package bybitv5

type Executor[T any] struct {
	topic         string
	subscriptions *Subscriptions
}

func NewExecutor[T any](topic string, subscriptions *Subscriptions) *Executor[T] {
	o := new(Executor[T])
	o.topic = topic
	o.subscriptions = subscriptions
	return o
}

func (o *Executor[T]) Subscribe(onShot func(Topic[T])) {
	o.subscriptions.subscribe(o.topic, func(raw RawTopic) error {
		topic, err := UnmarshalRawTopic[T](raw)
		if err == nil {
			onShot(topic)
		}
		return err
	})
}

func (o *Executor[T]) Unsubscribe() {
	o.subscriptions.unsubscribe(o.topic)
}
