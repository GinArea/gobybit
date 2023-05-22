package spotv3

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
