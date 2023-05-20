package gobybit

import "github.com/ginarea/gobybit/transport"

// Get All Coins Balance (https://bybit-exchange.github.io/docs/v5/asset/all-balance)
// accountType Required string Account type
// memberId             string User Id. It is required when you use master api key to check sub account coin balance
// coin                 string Coin name
// withBonus            string Whether query bonus or not. 0(default)：false; 1：true
type GetCoinsBalance struct {
	AccountType AccountType
	MemberId    *string
	Coin        *string
	WithBonus   *string
}

func (o GetCoinsBalance) Do(c *Client) Response[CoinsBalance] {
	return Get(c.asset(), "transfer/query-account-coins-balance", o, forward[CoinsBalance])
}

type CoinsBalance struct {
	AccountType AccountType
	MemberId    string
	Balance     []CoinBalance
}

type CoinBalance struct {
	Coin            string
	TransferBalance transport.Float64
	WalletBalance   transport.Float64
	Bonus           string
}

func (o *Client) GetCoinsBalance(v GetCoinsBalance) Response[CoinsBalance] {
	return v.Do(o)
}

func (o *Client) GetAccountCoinsBalance(accountType AccountType) Response[CoinsBalance] {
	return o.GetCoinsBalance(GetCoinsBalance{
		AccountType: accountType,
	})
}
