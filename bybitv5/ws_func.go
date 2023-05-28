package bybitv5

import (
	"encoding/json"
)

func WsFunc[T any](data []byte, f func(T)) error {
	var v T
	err := json.Unmarshal(data, &v)
	if err == nil {
		if f != nil {
			f(v)
		}
	}
	return err
}

func WsFuncDelta[T any, D any](data []byte, f func(T), delta bool, fd func(D)) error {
	if delta {
		return WsFunc(data, fd)
	}
	return WsFunc(data, f)
}
