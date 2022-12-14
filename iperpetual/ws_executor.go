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

func (this *WsExecutor[T]) Init(section *WsSection, subscription Subscription) {
	this.section = section
	this.topic = subscription.String()
}

func (this *WsExecutor[T]) Subscribe(onShot func(T)) {
	this.section.subscribe(this.topic, func(m []byte, delta bool) error {
		return WsFunc(m, onShot)
	})
}

func (this *WsExecutor[T]) Unsubscribe() {
	this.section.unsubscribe(this.topic)
}

func (this *WsExecutor[T]) Instant() *WsInstant[T] {
	return NewWsInstant[T](this)
}

type WsDeltaExecutor[T any] struct {
	WsExecutor[T]
}

func NewWsDeltaExecutor[T any](section *WsSection, subscription Subscription) *WsDeltaExecutor[T] {
	e := &WsDeltaExecutor[T]{}
	e.Init(section, subscription)
	return e
}

func (this *WsDeltaExecutor[T]) SubscribeWithDelta(onShot func(T), onDelta func(Delta)) {
	this.section.subscribe(this.topic, func(m []byte, delta bool) error {
		return WsFuncDelta(m, onShot, delta, onDelta)
	})
}

func (this *WsDeltaExecutor[T]) Subscribe(onShot func(T)) {
	var current T
	this.SubscribeWithDelta(func(shot T) {
		current = shot
		onShot(current)
	}, func(delta Delta) {
		if delta.HasData() {
			WsDeltaApply(&current, delta)
			onShot(current)
		}
	})
}

func (this *WsDeltaExecutor[T]) Instant() *WsInstant[T] {
	return NewWsInstant[T](this)
}
