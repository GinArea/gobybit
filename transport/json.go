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

type Timestamp time.Time

func (o *Timestamp) UnmarshalJSON(b []byte) error {
	// convert uint timestamp to time.Time
	return nil
}

func (o Timestamp) Std() time.Time {
	return time.Time(o)
}

type Time time.Time

func (o *Time) UnmarshalJSON(b []byte) error {
	// convert tim to time.Time
	return nil
}

func (o Time) Std() time.Time {
	return time.Time(o)
}
