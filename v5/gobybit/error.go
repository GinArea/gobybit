package gobybit

import (
	"fmt"
)

type Error struct {
	Code int
	Text string
}

func (o *Error) Std() error {
	if o.Empty() {
		return nil
	} else {
		return o
	}
}

func (o *Error) Empty() bool {
	return o.Code == 0
}

func (o *Error) Error() string {
	return fmt.Sprintf("code[%d]: %s", o.Code, o.Text)
}
