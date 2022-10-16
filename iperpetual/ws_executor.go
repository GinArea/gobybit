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

type WsDeltaExecutor[T any] struct {
	WsExecutor[T]
}

func NewWsDeltaExecutor[T any](section *WsSection, subscription Subscription) *WsDeltaExecutor[T] {
	e := &WsDeltaExecutor[T]{}
	e.Init(section, subscription)
	return e
}

func (this *WsDeltaExecutor[T]) Subscribe(onShot func(T), onDelta func(Delta)) {
	this.section.subscribe(this.topic, func(m []byte, delta bool) error {
		return WsFuncDelta(m, onShot, delta, onDelta)
	})
}

/*
func (this *WsDeltaExecutor[T, TD]) Subscribe(onShot func(T)) {
	var current T
	this.SubscribeDual(func(shot T) {
		current = shot
		onShot(shot)
	}, func(delta TD) {
		WsApplyDelta(&current, delta)
	})
}
*/
