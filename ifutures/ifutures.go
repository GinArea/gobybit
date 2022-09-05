package ifutures

import (
	"fmt"

	"github.com/tranquiil/bybit"
)

type InverseFutures struct {
	client  *Client
	name    string
	version int
}

func New(client *bybit.Client) *InverseFutures {
	return &InverseFutures{
		client:  client,
		name:    "private",
		version: 2,
	}
}

type Balance struct {
}

func (this *InverseFutures) WalletBalance() (any, bool) {
	// symbol Required string  Name of the trading pair
	// scale           int     Precision of the merged orderbook, 1 means 1 digit
	// limit           integer Default value is 100
	resp := Response[any]{}
	err := this.client.Get(this.url("wallet/balance"), UrlParam{}, &resp)
	//return resp.Result.Balances, err == nil
	return resp.Result, err == nil
}

/*
func (this *Spot) ServerTime() (uint64, bool) {
	resp := Response[struct {
		Time uint64 `json:"serverTime"`
	}]{}
	err := this.client.Get(this.url("time"), UrlParam{}, &resp)
	return resp.Result.Time, err == nil
}
*/
func (this *Spot) url(method string) string {
	return fmt.Sprintf("%s/v%d/%s", this.name, this.version, method)
}
