package spotv3

type WsExecutor[T any] struct {
	section *WsSection
	topic   string
}

func NewWsExecutor[T any](section *WsSection, subscription Subscription) *WsExecutor[T] {
	e := &WsExecutor[T]{}
	e.Init(section, subscription)
	return e
}

func (o *WsExecutor[T]) Init(section *WsSection, subscription Subscription) {
	o.section = section
	o.topic = subscription.String()
}

func (o *WsExecutor[T]) Subscribe(onShot func(T)) {
	o.section.subscribe(o.topic, func(m []byte, delta bool) error {
		return WsFunc(m, onShot)
	})
}

func (o *WsExecutor[T]) Unsubscribe() {
	o.section.unsubscribe(o.topic)
}
