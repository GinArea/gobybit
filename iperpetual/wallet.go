package iperpetual

// Get Wallet Balance (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-balance)
//
// coin string Currency alias. Returns all wallet balances if not passed
type WalletBalance struct {
	Currency *string `param:"coin"`
}

func (o WalletBalance) Do(client *Client) (map[string]Balance, error) {
	return Get[map[string]Balance](client, "wallet/balance", o)
}

type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

func (o *Client) WalletBalance(currency *string) (map[string]Balance, error) {
	return WalletBalance{Currency: currency}.Do(o)
}
