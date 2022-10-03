// Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-risklimit)
package iperpetual

// Get Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getrisklimit)
type GetRiskLimit struct {
	Symbol *Symbol `param:"symbol"`
}

func (this GetRiskLimit) Do(client *Client) ([]RiskLimitItem, bool) {
	return GetPublic[[]RiskLimitItem](client, "risk-limit/list", this)
}

type RiskLimitItem struct {
	ID             int      `json:"id"`
	Symbol         Symbol   `json:"symbol"`
	Limit          int      `json:"limit"`
	MaintainMargin string   `json:"maintain_margin"`
	StartingMargin string   `json:"starting_margin"`
	Section        []string `json:"section"`
	IsLowestRisk   int      `json:"is_lowest_risk"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	MaxLeverage    string   `json:"max_leverage"`
}

func (this *Client) GetRiskLimit(symbol *Symbol) ([]RiskLimitItem, bool) {
	return GetRiskLimit{Symbol: symbol}.Do(this)
}

// Set Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-setrisklimit)
//   symbol  Required string  Symbol
//   risk_id Required integer Risk ID
type SetRiskLimit struct {
	Symbol Symbol `param:"symbol"`
	RiskID int    `param:"risk_id"`
}

func (this SetRiskLimit) Do(client *Client) (int, bool) {
	type result struct {
		RiskID int `json:"risk_id"`
	}
	r, ok := Post[result](client, "position/risk-limit", this)
	return r.RiskID, ok
}

func (this *Client) SetRiskLimit(v SetRiskLimit) (int, bool) {
	return v.Do(this)
}
