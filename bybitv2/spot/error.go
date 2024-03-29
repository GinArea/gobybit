package spot

import (
	"errors"

	"github.com/ginarea/gobybit/bybitv2/transport"
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
