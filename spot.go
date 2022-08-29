package bybit

import "fmt"

type Spot struct {
	client  *Client
	name    string
	version int
}

func NewSpot(client *Client) *Spot {
	return &Spot{
		client:  client,
		name:    "spot",
		version: 1,
	}
}

type Balance struct {
	Coin     string `json:"coin"`
	CoinID   string `json:"coinId"`
	CoinName string `json:"coinName"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Locked   string `json:"locked"`
}

func (this *Spot) WalletBalance() ([]Balance, bool) {
	resp := Response[struct {
		Balances []Balance `json:"balances"`
	}]{}
	err := this.client.Get(this.url("account"), UrlParam{}, &resp)
	return resp.Result.Balances, err == nil
}

func (this *Spot) ServerTime() (uint64, bool) {
	resp := Response[struct {
		Time uint64 `json:"serverTime"`
	}]{}
	err := this.client.Get(this.url("time"), UrlParam{}, &resp)
	return resp.Result.Time, err == nil
}

func (this *Spot) url(method string) string {
	return fmt.Sprintf("%s/v%d/%s", this.name, this.version, method)
}

func (this *Spot) urlQuote(method string) string {
	return fmt.Sprintf("%s/quote/v%d/%s", this.name, this.version, method)
}
