package iperpetual

import (
	"encoding/json"
)

func WsFunc[T any, F func(T)](m []byte, f F) error {
	var v Topic[T]
	err := json.Unmarshal(m, &v)
	if err != nil {
		return err
	}
	if f != nil {
		f(v.Data)
	}
	return nil
}

func WsFuncDelta[T any, F func(T), TD any, FD func(TD)](m []byte, f F, delta bool, fd FD) error {
	if delta {
		return WsFunc(m, fd)
	}
	return WsFunc(m, f)
}
