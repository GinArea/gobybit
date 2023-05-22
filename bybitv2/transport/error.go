package transport

import (
	"fmt"
)

type Error struct {
	Err
}

type Err struct {
	Code int
	Text string
}

func (o *Err) Empty() bool {
	return o.Code == 0
}

func (o *Err) Error() string {
	return fmt.Sprintf("code[%d]: %s", o.Code, o.Text)
}
