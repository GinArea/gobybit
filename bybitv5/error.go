package bybitv5

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Error (https://bybit-exchange.github.io/docs/v5/error)
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

func (o *Error) RequestParameterError() bool {
	return o.Code == 10001
}

func (o *Error) ServiceTemporarilyUnavailable() bool {
	return o.Code == 50001 // with http 503
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

func (o *Error) SymbolIsNotWhitelisted() bool {
	return o.Code == 10029
}

func (o *Error) KycNeeded() bool {
	codes := []int{
		10024,  // Compliance rules triggered
		131004, // KYC needed
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) Timeout() bool {
	codes := []int{
		10000,  // Server Timeout
		10016,  // Request timeout
		170007, // Timeout waiting for response from backend server.
		170146, // Order creation timeout
		170147, // Order cancellation timeout
		177002, // Timeout
	}
	return slices.Contains(codes, o.Code)
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

func (o *Error) OrderLinkedIdIsDuplicate() bool {
	return o.Code == 110072
}

func (o *Error) ReduceOnlyRuleNotSatisfied() bool {
	return o.Code == 110017
}
