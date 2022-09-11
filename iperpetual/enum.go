// [Enums Definitions] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-enums
package iperpetual

// [Side (side)]: https://bybit-exchange.github.io/docs/futuresV2/inverse/#side-side
type Side string

const (
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

// [Symbol (symbol)] https://bybit-exchange.github.io/docs/futuresV2/inverse/#symbol-symbol
type Symbol string

// [Currency (currency/coin)] https://bybit-exchange.github.io/docs/futuresV2/inverse/#currency-currency-coin
type Currency string

const (
	BTC  Currency = "BTC"
	ETH  Currency = "ETH"
	EOS  Currency = "EOS"
	XRP  Currency = "XRP"
	DOT  Currency = "DOT"
	USDT Currency = "USDT"
)

// [Contract Status (status)] https://bybit-exchange.github.io/docs/futuresV2/inverse/#contract-status-status
type ContractStatus string

const (
	Trading  ContractStatus = "Trading"
	Settling ContractStatus = "Settling"
	Closed   ContractStatus = "Closed"
)
