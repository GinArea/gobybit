// [Inverse Perpetual] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-introduction
package iperpetual

import (
	"fmt"

	"github.com/tranquiil/bybit"
)

type Client struct {
	client  *bybit.Client
	version int
}

func NewClient(client *bybit.Client) *Client {
	return &Client{
		client:  client,
		version: 2,
	}
}

type Balance struct {
	Equity           float32 `json:"equity"`
	AvailableBalance float32 `json:"available_balance"`
	UsedMargin       float32 `json:"used_margin"`
	OrderMargin      float32 `json:"order_margin"`
	PositionMargin   float32 `json:"position_margin"`
	OccClosingFee    float32 `json:"occ_closing_fee"`
	OccFundingFee    float32 `json:"occ_funding_fee"`
	WalletBalance    float32 `json:"wallet_balance"`
	RealisedPnl      float32 `json:"realised_pnl"`
	UnrealisedPnl    float32 `json:"unrealised_pnl"`
	CumRealisedPnl   float32 `json:"cum_realised_pnl"`
	GivenCash        float32 `json:"given_cash"`
	ServiceCash      float32 `json:"service_cash"`
}

/*
func (this *Client) WalletBalance() (any, bool) {
	// coin string Currency alias. Returns all wallet balances if not passed
	resp := Response[map[Coin]Balance]{}
	err := this.client.Get(this.urlPrivate("wallet/balance"), bybit.UrlParam{
		//"coin": coin,
	}, &resp)
	return resp.Result, err == nil
}

func (this *Client) ServerTime() (string, bool) {
	resp := Response[any]{}
	err := this.client.Get(this.urlPublic("time"), bybit.UrlParam{}, &resp)
	return resp.TimeNow, err == nil
}
*/
func (this *Client) url(access, path string) string {
	return fmt.Sprintf("v%d/%s/%s", this.version, access, path)
}

func (this *Client) urlPublic(path string) string {
	return this.url("public", path)
}

func (this *Client) urlPrivate(path string) string {
	return this.url("private", path)
}
