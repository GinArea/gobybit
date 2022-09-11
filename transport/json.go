package transport

import (
	"encoding/json"
	"fmt"
)

func JsonUnmarshal[V any](jsonBlob []byte) V {
	var v V
	err := json.Unmarshal(jsonBlob, &v)
	if err != nil {
		panic(fmt.Sprint("json unmarshal: ", err))
	}
	return v
}
