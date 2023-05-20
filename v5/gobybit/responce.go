package gobybit

type Response[T any] struct {
	Time  uint64
	Data  T
	Limit RateLimit
	Error error
}

func (o *Response[T]) Ok() bool {
	return o.Error == nil
}

func (o *Response[T]) SetErrorIfNil(err error) {
	if o.Error == nil {
		o.Error = err
	}
}

type response[T any] struct {
	RetCode int
	RetMsg  string
	Time    uint64
	Result  T
}

func (o *response[T]) Error() error {
	e := Error{
		Code: o.RetCode,
		Text: o.RetMsg,
	}
	return e.Std()
}
