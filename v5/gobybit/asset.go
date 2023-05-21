package gobybit

import "github.com/ginarea/gobybit/transport"

// Get Coin Exchange Records
// https://bybit-exchange.github.io/docs/v5/asset/exchange
//
//	fromCoin string  The currency to convert from. e.g,BTC
//	toCoin   string  The currency to convert to. e.g,USDT
//	limit    integer Limit for data size per page. [1, 50]. Default: 10
//	cursor   string  Cursor. Used for pagination
type GetCoinRecords struct {
	FromCoin *string
	ToCoin   *string
	Limit    *int
	Cursor   *string
}

func (o GetCoinRecords) Do(c *Client) Response[[]CoinRecord] {
	type result struct {
		OrderBody      []CoinRecord
		NextPageCursor string
	}
	return Get(c.asset(), "exchange/order-record", o, func(r result) ([]CoinRecord, error) {
		return r.OrderBody, nil
	})
}

type CoinRecord struct {
	Fromcoin     string
	Fromamount   string
	Tocoin       string
	Toamount     string
	Exchangerate string
	Createdtime  string
	Exchangetxid string
}

func (o *Client) GetCoinRecords(v GetCoinRecords) Response[[]CoinRecord] {
	return v.Do(o)
}

// Get Delivery Record
// https://bybit-exchange.github.io/docs/v5/asset/delivery
//
// category Required string  Product type. option, linear
// symbol            string  Symbol name
// expDate           string  Expiry date. 25MAR22. Default: return all
// limit             integer Limit for data size per page. [1, 50]. Default: 20
// cursor            string  Cursor. Used for pagination
type GetDeliveryRecords struct {
	Category Category
	Symbol   *string
	Expdate  *string
	Limit    *int
	Cursor   *string
}

func (o GetDeliveryRecords) Do(c *Client) Response[[]DeliveryRecord] {
	type result struct {
		Category       Category
		NextPageCursor string
		List           []DeliveryRecord
	}
	return Get(c.asset(), "delivery-record", o, func(r result) ([]DeliveryRecord, error) {
		return r.List, nil
	})
}

type DeliveryRecord struct {
	Symbol        string
	Side          Side
	DeliveryTime  int64
	Strike        string
	Fee           string
	Position      string
	DeliveryPrice string
	DeliveryRpl   string
}

func (o *Client) GetDeliveryRecords(v GetDeliveryRecords) Response[[]DeliveryRecord] {
	return v.Do(o)
}

// Get USDC Session Settlement
// https://bybit-exchange.github.io/docs/v5/asset/settlement
// category Required string  Product type. linear
// symbol            string  Symbol name
// limit             integer Limit for data size per page. [1, 50]. Default: 20
// cursor            string  Cursor. Used for pagination
type GetSettlement struct {
	Category Category
	Symbol   *string
	Limit    *int
	Cursor   *string
}

func (o GetSettlement) Do(c *Client) Response[[]Settlement] {
	type result struct {
		Category       Category
		NextPageCursor string
		List           []Settlement
	}
	return Get(c.asset(), "settlement-record", o, func(r result) ([]Settlement, error) {
		return r.List, nil
	})
}

type Settlement struct {
	RealisedPnl     string
	Symbol          string
	Side            string
	MarkPrice       string
	Size            string
	CreatedTime     string
	SessionAvgPrice string
}

func (o *Client) GetSettlement(v GetSettlement) Response[[]Settlement] {
	return v.Do(o)
}

// Get All Coins Balance
// https://bybit-exchange.github.io/docs/v5/asset/all-balance
//
//	accountType Required string Account type
//	memberId             string User Id. It is required when you use master api key to check sub account coin balance
//	coin                 string Coin name
//	withBonus            string Whether query bonus or not. 0(default)：false; 1：true
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
