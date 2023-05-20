package gobybit

type request[T any] interface {
	Do(*Client) Response[T]
}

func Request[T any](c *Client, r request[T]) Response[T] {
	return r.Do(c)
}

func Request2[T any, R request[T]](c *Client, r R) Response[T] {
	return r.Do(c)
}

type request3[T any] interface {
	Do(*Client) T
}

func Request3[T any](c *Client, r request3[T]) T {
	return r.Do(c)
}

//func (o *Client) Do Request[T any]
