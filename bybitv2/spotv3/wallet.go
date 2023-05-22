// Wallet Data Endpoints (https://bybit-exchange.github.io/docs/spot/v3/#t-wallet)
package spotv3

import "github.com/ginarea/gobybit/bybitv2/transport"

// Get Wallet Balance (https://bybit-exchange.github.io/docs/spot/v3/#t-balance)
type Balance struct {
	Coin     string            `json:"coin"`
	CoinID   string            `json:"coinId"`
	CoinName string            `json:"coinName"`
	Total    transport.Float64 `json:"total"`
	Free     transport.Float64 `json:"free"`
	Locked   string            `json:"locked"`
}

func (this *Client) WalletBalance() ([]Balance, error) {
	type result struct {
		Balances []Balance `json:"balances"`
	}
	r, err := Get[result](this, "account", nil)
	return r.Balances, err
}
