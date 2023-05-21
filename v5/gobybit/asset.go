package gobybit

import (
	"github.com/google/uuid"
	"github.com/msw-x/moon/ujson"
)

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

// Get Asset Info
// https://bybit-exchange.github.io/docs/v5/asset/asset-info
//
//	accountType Required string Account type. SPOT
//	coin                 string Coin name
type GetAssetInfo struct {
	AccountType AccountType
	Coin        *string
}

func (o GetAssetInfo) Do(c *Client) Response[SpotAssetInfo] {
	type result struct {
		Spot SpotAssetInfo
	}
	return Get(c.asset(), "transfer/query-asset-info", o, func(r result) (SpotAssetInfo, error) {
		return r.Spot, nil
	})
}

type SpotAssetInfo struct {
	Status string
	Assets []AssetInfo
}

type AssetInfo struct {
	Coin     string
	Frozen   string
	Free     string
	Withdraw string
}

func (o *Client) GetAssetInfo(v GetAssetInfo) Response[SpotAssetInfo] {
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
	TransferBalance ujson.Float64
	WalletBalance   ujson.Float64
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

// Get Single Coin Balance
// https://bybit-exchange.github.io/docs/v5/asset/account-coin-balance
//
//	accountType            Required string  Account type
//	coin                   Required string  Coin name
//	memberId                        string  User Id. It is required when you use master api key to check sub account coin balance
//	withBonus                       integer Whether query bonus or not. 0(default)：false; 1：true
//	withTransferSafeAmount          integer Whether query delay withdraw/transfer safe amount
type GetCoinBalance struct {
	AccountType            AccountType
	Coin                   string
	MemberId               *string
	WithBonus              *int
	WithTransferSafeAmount *int
}

func (o GetCoinBalance) Do(c *Client) Response[SingleCoinBalance] {
	return Get(c.asset(), "transfer/query-account-coin-balance", o, forward[SingleCoinBalance])
}

type SingleCoinBalance struct {
	AccountType AccountType
	MemberId    string
	BizType     int
	AccountId   string
	Balance     CoinBalanceExt
}

type CoinBalanceExt struct {
	CoinBalance
	TransferSafeAmount string
}

func (o *Client) GetCoinBalance(v GetCoinBalance) Response[SingleCoinBalance] {
	return v.Do(o)
}

// Get Transferable Coin
// https://bybit-exchange.github.io/docs/v5/asset/transferable-coin
//
//	fromAccountType Required string From account type
//	toAccountType   Required string To account type
type GetTransferableCoin struct {
	FromAccountType AccountType
	ToAccountType   AccountType
}

func (o GetTransferableCoin) Do(c *Client) Response[[]string] {
	type result struct {
		List []string
	}
	return Get(c.asset(), "transfer/query-transfer-coin-list", o, func(r result) ([]string, error) {
		return r.List, nil
	})
}

func (o *Client) GetTransferableCoin(v GetTransferableCoin) Response[[]string] {
	return v.Do(o)
}

// Create Internal Transfer
// https://bybit-exchange.github.io/docs/v5/asset/create-inter-transfer
//
//	transferId      Required string UUID. Please manually generate a UUID
//	coin            Required string Coin
//	amount          Required string Amount
//	fromAccountType Required string From account type
//	toAccountType   Required string To account type
type CreateInternalTransfer struct {
	TransferId      uuid.UUID
	Coin            string
	Amount          ujson.Float64
	FromAccountType AccountType
	ToAccountType   AccountType
}

func (o CreateInternalTransfer) Do(c *Client) Response[uuid.UUID] {
	if o.TransferId == uuid.Nil {
		o.TransferId = uuid.New()
	}
	type result struct {
		TransferId uuid.UUID
	}
	return Post(c.asset(), "transfer/inter-transfer", o, func(r result) (uuid.UUID, error) {
		return r.TransferId, nil
	})
}

func (o *Client) CreateInternalTransfer(v CreateInternalTransfer) Response[uuid.UUID] {
	return v.Do(o)
}
