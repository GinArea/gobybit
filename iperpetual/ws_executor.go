package iperpetual

type WsExecutor struct {
	section *WsSection
	topic   string
}

func (this *WsExecutor) Init(section *WsSection, subscription Subscription) {
	this.section = section
	this.topic = subscription.String()
}

func (this *WsExecutor) Unsubscribe() {
	this.section.unsubscribe(this.topic)
}

type WsMonoExecutor[T any] struct {
	WsExecutor
}

func NewWsMonoExecutor[T any](section *WsSection, subscription Subscription) *WsMonoExecutor[T] {
	e := &WsMonoExecutor[T]{}
	e.Init(section, subscription)
	return e
}

func (this *WsMonoExecutor[T]) Subscribe(onShot func(T)) {
	this.section.subscribe(this.topic, func(m []byte, delta bool) error {
		return WsFunc(m, onShot)
	})
}

type WsDualExecutor[T any, TD any] struct {
	WsExecutor
}

func NewWsDualExecutor[T any, TD any](section *WsSection, subscription Subscription) *WsDualExecutor[T, TD] {
	e := &WsDualExecutor[T, TD]{}
	e.Init(section, subscription)
	return e
}

func (this *WsDualExecutor[T, TD]) Subscribe(onShot func(T), onDelta func(TD)) {
	this.section.subscribe(this.topic, func(m []byte, delta bool) error {
		return WsFuncDelta(m, onShot, delta, onDelta)
	})
}
