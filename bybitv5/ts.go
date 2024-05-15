package bybitv5

import (
	"strconv"
	"time"
)

type Timestamp struct {
	timeShift  int
	recvWindow int
}

func NewTimestamp() *Timestamp {
	o := new(Timestamp)
	o.timeShift = -10000
	o.recvWindow = 20000
	return o
}

func (o *Timestamp) Get() (ts, window string) {
	i := nowUtcMs() + o.timeShift
	ts = strconv.Itoa(i)
	window = strconv.Itoa(o.recvWindow)
	return
}

func (o *Timestamp) Expires() int {
	return nowUtcMs() + o.timeShift + o.recvWindow
}

func nowUtcMs() int {
	return int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
}
