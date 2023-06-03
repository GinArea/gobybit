package bybitv5

import "github.com/msw-x/moon/ujson"

// Get Wallet Balance
// https://bybit-exchange.github.io/docs/v5/account/wallet-balance
//
//	accountType Required string Account type
//	coin                 string Coin name
type GetWallatBalance struct {
	AccountType AccountType
	Coin        string `url:",omitempty"`
}

type WalletBalance struct {
	AccountType            AccountType
	TotalEquity            ujson.Float64
	AccountImRate          ujson.Float64
	TotalMarginBalance     ujson.Float64
	TotalInitialMargin     ujson.Float64
	TotalAvailableBalance  ujson.Float64
	AccountMmRate          ujson.Float64
	TotalPerpUpl           ujson.Float64
	TotalWalletBalance     ujson.Float64
	AccountLtv             ujson.Float64
	TotalMaintenanceMargin ujson.Float64
	Coin                   []WalletCoinBalance
}

type WalletCoinBalance struct {
	Coin                string
	AvailableToBorrow   ujson.Float64
	Bonus               ujson.Float64
	AccruedInterest     ujson.Float64
	AvailableToWithdraw ujson.Float64
	TotalOrderIm        ujson.Float64
	Equity              ujson.Float64
	TotalPositionMm     ujson.Float64
	UsdValue            ujson.Float64
	UnrealisedPnl       ujson.Float64
	BorrowAmount        ujson.Float64
	TotalPositionIm     ujson.Float64
	WalletBalance       ujson.Float64
	CumRealisedPnl      ujson.Float64
	Free                ujson.Float64
	Locked              ujson.Float64
}

func (o GetWallatBalance) Do(c *Client) Response[[]WalletBalance] {
	type result struct {
		List []WalletBalance
	}
	return Get(c.account(), "wallet-balance", o, func(r result) ([]WalletBalance, error) {
		return r.List, nil
	})
}

func (o *Client) GetWallatBalance(v GetWallatBalance) Response[[]WalletBalance] {
	return v.Do(o)
}

func (o *Client) GetAccountWallatBalance(accountType AccountType) Response[[]WalletBalance] {
	return o.GetWallatBalance(GetWallatBalance{AccountType: accountType})
}
