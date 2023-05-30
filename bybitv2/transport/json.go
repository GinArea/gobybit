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

func (o *Float64) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, `"`)
	if s == "null" {
		s = ""
	}
	if s == "" {
		*o = 0
		return nil
	}
	f, err := strconv.ParseFloat(s, 64)
	*o = Float64(f)
	return err
}

func (o Float64) Value() float64 {
	return float64(o)
}

func (o Float64) Ptr() *float64 {
	v := o.Value()
	return &v
}

func (o Float64) IsZero() bool {
	return o.Value() == 0
}

func (o Float64) IsNotZero() bool {
	return o.Value() != 0
}

type TimeMs time.Time

func (o *TimeMs) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, `"`)
	if s == "" {
		*o = TimeMs{}
		return nil
	}
	i, err := strconv.ParseInt(s, 10, 64)
	t := time.Unix(0, i*int64(time.Millisecond))
	*o = TimeMs(t)
	return err
}

func (o TimeMs) Std() time.Time {
	return time.Time(o)
}
