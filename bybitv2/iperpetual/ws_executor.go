package iperpetual

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

func (o *WsExecutor[T]) Instant() *WsInstant[T] {
	return NewWsInstant[T](o)
}

type WsDeltaExecutor[T any] struct {
	WsExecutor[T]
}

func NewWsDeltaExecutor[T any](section *WsSection, subscription Subscription) *WsDeltaExecutor[T] {
	e := &WsDeltaExecutor[T]{}
	e.Init(section, subscription)
	return e
}

func (o *WsDeltaExecutor[T]) SubscribeWithDelta(onShot func(T), onDelta func(Delta)) {
	o.section.subscribe(o.topic, func(m []byte, delta bool) error {
		return WsFuncDelta(m, onShot, delta, onDelta)
	})
}

func (o *WsDeltaExecutor[T]) Subscribe(onShot func(T)) {
	var current T
	o.SubscribeWithDelta(func(shot T) {
		current = shot
		onShot(current)
	}, func(delta Delta) {
		if delta.HasData() {
			WsDeltaApply(&current, delta)
			onShot(current)
		}
	})
}

func (o *WsDeltaExecutor[T]) Instant() *WsInstant[T] {
	return NewWsInstant[T](o)
}
