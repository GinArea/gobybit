package iperpetual

// Get Wallet Balance (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-balance)
//
// coin string Currency alias. Returns all wallet balances if not passed
type WalletBalance struct {
	Currency *string `param:"coin"`
}

func (this WalletBalance) Do(client *Client) (map[string]Balance, bool) {
	return Get[map[string]Balance](client, "wallet/balance", this)
}

type Balance struct {
	Equity           float32 `json:"equity"`
	AvailableBalance float32 `json:"available_balance"`
	UsedMargin       float32 `json:"used_margin"`
	OrderMargin      float32 `json:"order_margin"`
	PositionMargin   float32 `json:"position_margin"`
	OccClosingFee    float32 `json:"occ_closing_fee"`
	OccFundingFee    float32 `json:"occ_funding_fee"`
	WalletBalance    float32 `json:"wallet_balance"`
	RealisedPnl      float32 `json:"realised_pnl"`
	UnrealisedPnl    float32 `json:"unrealised_pnl"`
	CumRealisedPnl   float32 `json:"cum_realised_pnl"`
	GivenCash        float32 `json:"given_cash"`
	ServiceCash      float32 `json:"service_cash"`
}

func (this *Client) WalletBalance(currency *string) (map[string]Balance, bool) {
	return WalletBalance{Currency: currency}.Do(this)
}
