package transport

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
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

func (this Float64) Value() float64 {
	return float64(this)
}

func (this Float64) Ptr() *float64 {
	v := this.Value()
	return &v
}

func (this Float64) IsZero() bool {
	return this.Value() == 0
}

func (this Float64) IsNotZero() bool {
	return this.Value() != 0
}

type Time time.Time

func (this *Time) UnmarshalJSON(b []byte) error {
	// convert uint timestamp to time.Time
	return nil
}
