package iperpetual

type WsInstant[T any] struct {
	executor WsExecutorInterface[T]
	onUpdate func(T)
	v        *T
}

func NewWsInstant[T any](executor WsExecutorInterface[T]) *WsInstant[T] {
	i := &WsInstant[T]{
		executor: executor,
	}
	executor.Subscribe(func(v T) {
		i.v = &v
		if i.onUpdate != nil {
			i.onUpdate(v)
		}
	})
	return i
}

func (o *WsInstant[T]) Empty() bool {
	return o.v == nil
}

func (o *WsInstant[T]) Has() bool {
	return !o.Empty()
}

func (o *WsInstant[T]) Value() T {
	return *o.v
}

func (o *WsInstant[T]) OnUpdate(onUpdate func(T)) {
	o.onUpdate = onUpdate
}

func (o *WsInstant[T]) Unsubscribe() {
	o.executor.Unsubscribe()
}

type WsExecutorInterface[T any] interface {
	Subscribe(func(T))
	Unsubscribe()
}
