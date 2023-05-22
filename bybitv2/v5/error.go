package v5

import (
	"errors"

	"github.com/ginarea/gobybit/bybitv2/transport"
	"golang.org/x/exp/slices"
)

// https://bybit-exchange.github.io/docs/v5/error
type Error struct {
	transport.Err
}

func (o *Error) ApiKeyInvalid() bool {
	codes := []int{
		10003, // API key is invalid
		10004, // Error sign, please check your signature generation algorithm.
		10005, // Permission denied, please check your API key permissions
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) ApiKeyExpired() bool {
	return o.Code == 33004
}

func (o *Error) TooManyVisits() bool {
	return o.Code == 10006
}

func (o *Error) UnmatchedIp() bool {
	return o.Code == 10010
}

func (o *Error) KycNeeded() bool {
	return o.Code == 131004
}

func (o *Error) InsufficientBalance() bool {
	codes := []int{
		110004, // Wallet balance is insufficient
		110007, // Available balance is insufficient
		110012, // Insufficient available balance
		110045, // Wallet balance is insufficient
		110051, // The user's available balance cannot cover the lowest price of the current market
		110052, // Your available balance is insufficient to set the price
		110053, // The user's available balance cannot cover the current market price and upper limit price
		170033, // margin Insufficient account balance
		170131, // Balance insufficient
		175003, // Insufficient available balance. Please make a deposit and try again.
		175006, // Insufficient available balance. Please make a deposit and try again.
		176015, // Insufficient available balance
	}
	return slices.Contains(codes, o.Code)
}

func forwardError(err error) error {
	var terr *transport.Error
	if errors.As(err, &terr) {
		return &Error{Err: terr.Err}
	}
	return err
}
