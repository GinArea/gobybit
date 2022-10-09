package transport

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func JsonUnmarshal[V any](jsonBlob []byte) V {
	var v V
	err := json.Unmarshal(jsonBlob, &v)
	if err != nil {
		panic(fmt.Sprint("json unmarshal: ", err))
	}
	return v
}

type Float64 float64

func (this *Float64) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, `"`)
	f, err := strconv.ParseFloat(s, 64)
	*this = Float64(f)
	return err
}

func (this *Float64) Value() float64 {
	return float64(*this)
}
