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

func (this *WsInstant[T]) Empty() bool {
	return this.v == nil
}

func (this *WsInstant[T]) Value() T {
	return *this.v
}

func (this *WsInstant[T]) OnUpdate(onUpdate func(T)) {
	this.onUpdate = onUpdate
}

func (this *WsInstant[T]) Unsubscribe() {
	this.executor.Unsubscribe()
}

type WsExecutorInterface[T any] interface {
	Subscribe(func(T))
	Unsubscribe()
}
