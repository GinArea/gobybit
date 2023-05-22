package iperpetual

import (
	"errors"

	"github.com/ginarea/gobybit/bybitv2/transport"
	"golang.org/x/exp/slices"
)

type Error struct {
	transport.Err
}

func (o *Error) ApiKeyInvalid() bool {
	codes := []int{
		10002, // Request not authorized
		10003, // Invalid API key
		10004, // Invalid sign
		10005, // Permission denied for current API key
		20014, // Invalid API key format
		20015, // Invalid API key or IP
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

func (o *Error) Timeout() bool {
	return o.Code == 10016
}

func (o *Error) InsufficientBalance() bool {
	codes := []int{
		30010, // Insufficient wallet balance
		30022, // Estimated buy liq_price cannot be higher than current mark_price
		30023, // Estimated sell liq_price cannot be lower than current mark_price
		30031, // Insufficient available balance for order cost
		30042, // Insufficient wallet balance
		30049, // Insufficient available balance
		30067, // Insufficient available balance
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
