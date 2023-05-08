package spotv3

import (
	"errors"

	"github.com/ginarea/gobybit/transport"
	"golang.org/x/exp/slices"
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

func (o *Error) Timeout() bool {
	return o.Code == 10016
}

func (o *Error) InsufficientBalance() bool {
	codes := []int{
		12131, // Insufficient balance
		12403, // Insufficient available balance. Please make a deposit and try again.
		12406, // Insufficient available balance. Please make a deposit and try again.
		12615, // Insufficient available balance
		12228, // The purchase amount of each order exceeds the estimated maximum purchase amount
		12229, // The sell quantity per order exceeds the estimated maximum sell quantity
		12629, // Amount to borrow has exceeded the user's estimated max. amount to borrow
	}
	return slices.Contains(codes, o.Code)
}
