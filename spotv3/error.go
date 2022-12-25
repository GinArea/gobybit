package spotv3

import (
	"errors"

	"github.com/ginarea/gobybit/transport"
)

type Error struct {
	transport.Err
}

func forwardError(err error) error {
	var terr *transport.Error
	if errors.As(err, &terr) {
		return &Error{Err: terr.Err}
	}
	return err
}
