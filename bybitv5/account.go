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

// Get Account Info
// https://bybit-exchange.github.io/docs/v5/account/account-info
type GetAccountInfo struct {
}

type AccountInfo struct {
	MarginMode          string
	DcpStatus           string
	TimeWindow          int
	SmpGroup            int
	IsMasterTrader      bool
	UnifiedMarginStatus int
	SpotHedgingStatus   string
	UpdatedTime         string
}

func (o GetAccountInfo) Do(c *Client) Response[AccountInfo] {
	return Get(c.account(), "info", o, forward[AccountInfo])
}

func (o *Client) GetAccountInfo() Response[AccountInfo] {
	return GetAccountInfo{}.Do(o)
}

// Get Transaction Log
// https://bybit-exchange.github.io/docs/v5/account/transaction-log
//
//	accountType string  Account Type. UNIFIED
//	category    string  Product type. spot,linear,option
//	currency    string  Currency
//	baseCoin    string  BaseCoin. e.g., BTC of BTCPERP
//	type        string  Types of transaction logs
//	startTime   integer The start timestamp (ms)
//	endTime     integer The end timestamp (ms)
//	limit       integer Limit for data size per page. [1, 50]. Default: 20
//	cursor      string  Cursor. Use the nextPageCursor token from the response to retrieve the next page of the result set
type GetTransactionLog struct {
	AccountType AccountType `url:",omitempty"`
	Category    Category    `url:",omitempty"`
	Currency    string      `url:",omitempty"`
	BaseCoin    string      `url:",omitempty"`
	Type        Type        `url:",omitempty"`
	StartTime   int         `url:",omitempty"`
	EndTime     int         `url:",omitempty"`
	Limit       int         `url:",omitempty"`
	Cursor      string      `url:",omitempty"`
}

type TransactionLog struct {
	Symbol          string
	Side            Side
	Funding         ujson.Float64
	OrderLinkId     string
	OrderId         string
	Fee             ujson.Float64
	Change          ujson.Float64
	CashFlow        ujson.Float64
	TransactionTime string
	Type            Type
	FeeRate         ujson.Float64
	BonusChange     ujson.Float64
	Size            ujson.Float64
	Qty             ujson.Float64
	CashBalance     ujson.Float64
	Currency        string
	Category        Category
	TradePrice      ujson.Float64
	TradeId         string
}

func (o GetTransactionLog) Do(c *Client) Response[[]TransactionLog] {
	type result struct {
		NextPageCursor string
		List           []TransactionLog
	}
	return Get(c.account(), "transaction-log", o, func(r result) ([]TransactionLog, error) {
		return r.List, nil
	})
}

func (o *Client) GetTransactionLog(v GetTransactionLog) Response[[]TransactionLog] {
	return v.Do(o)
}
