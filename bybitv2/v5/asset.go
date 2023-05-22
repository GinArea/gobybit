package v5

import "github.com/ginarea/gobybit/transport"

type GetCoinsBalance struct {
	AccountType AccountType `param:"accountType"`
	MemberId    *string     `param:"memberId"`
	Coin        *string     `param:"coin"`
	WithBonus   *bool       `param:"withBonus"`
}

func (o GetCoinsBalance) Do(client *Client) (CoinsBalance, error) {
	return Get[CoinsBalance](client, "asset/transfer/query-account-coins-balance", o)
}

// (Get All Coins Balance) https://bybit-exchange.github.io/docs/v5/asset/all-balance
type CoinsBalance struct {
	AccountType AccountType
	MemberId    string
	Balances    []CoinBalance
}

type CoinBalance struct {
	Coin            string
	TransferBalance transport.Float64
	WalletBalance   transport.Float64
	Bonus           string
}

func (o *Client) GetCoinsBalance(v GetCoinsBalance) (CoinsBalance, error) {
	return v.Do(o)
}

func (o *Client) GetAccountCoinsBalance(accountType AccountType) (CoinsBalance, error) {
	return o.GetCoinsBalance(GetCoinsBalance{
		AccountType: accountType,
	})
}
