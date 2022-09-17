// [Enums Definitions] https://bybit-exchange.github.io/docs/account_asset/#t-enums
package account

// [Account type (from_account_type/to_account_type)] https://bybit-exchange.github.io/docs/account_asset/#account-type-from_account_type-to_account_type
type AccountType string

const (
	AccountContract   AccountType = "CONTRACT"
	AccountSpot       AccountType = "SPOT"
	AccountInvestment AccountType = "INVESTMENT"
	AccountOption     AccountType = "OPTION"
	AccountUnified    AccountType = "UNIFIED"
)

// [Withdraw status (status)] https://bybit-exchange.github.io/docs/account_asset/#withdraw-status-status
type Withdraw string

const (
	WithdrawSecurityCheck       Withdraw = "SecurityCheck"
	WithdrawPending             Withdraw = "Pending"
	WithdrawSuccess             Withdraw = "success"
	WithdrawCancelByUser        Withdraw = "CancelByUser"
	WithdrawReject              Withdraw = "Reject"
	WithdrawFail                Withdraw = "Fail"
	WithdrawBlockchainConfirmed Withdraw = "BlockchainConfirmed"
)

// [Currency (currency/coin)] https://bybit-exchange.github.io/docs/account_asset/#currency-currency-coin
type Currency string

const (
	BTC  Currency = "BTC"
	ETH  Currency = "ETH"
	EOS  Currency = "EOS"
	XRP  Currency = "XRP"
	USDT Currency = "USDT"
	DOT  Currency = "DOT"
	DOGE Currency = "DOGE"
	LTC  Currency = "LTC"
	XLM  Currency = "XLM"
	USD  Currency = "USD"
)

// [Operator type] https://bybit-exchange.github.io/docs/account_asset/#operator-type
type OperatorType string

const (
	OperatorSystem        OperatorType = "SYSTEM"
	OperatorUser          OperatorType = "USER"
	OperatorAdmin         OperatorType = "ADMIN"
	OperatorAffiliateUser OperatorType = "AFFILIATE_USER"
)

// [Transfer type (type)] https://bybit-exchange.github.io/docs/account_asset/#transfer-type-type
type TransferType string

const (
	TransferIn  TransferType = "IN"
	TransferOut TransferType = "OUT"
)

// [Transfer status (status)] https://bybit-exchange.github.io/docs/account_asset/#transfer-status-status
type TransferStatus string

const (
	TransferSuccess TransferStatus = "SUCCESS"
	TransferPending TransferStatus = "PENDING"
	TransferFailed  TransferStatus = "FAILED"
)

// [Page direction (direction)] https://bybit-exchange.github.io/docs/account_asset/#page-direction-direction
type PageDirection string

const (
	PagePrev PageDirection = "Prev"
	PageNext PageDirection = "Next"
)
