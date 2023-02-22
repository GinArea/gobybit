// Enums Definitions (https://bybit-exchange.github.io/docs/v5/enum)
package v5

// (https://bybit-exchange.github.io/docs/v5/enum#accounttype)
type AccountType string

const (
	AccountContract   AccountType = "CONTRACT"
	AccountSpot       AccountType = "SPOT"
	AccountInvestment AccountType = "INVESTMENT"
	AccountOption     AccountType = "OPTION"
	AccountUnified    AccountType = "UNIFIED"
	AccountFund       AccountType = "FUND"
)
