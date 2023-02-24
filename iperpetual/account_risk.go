// Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-risklimit)
package iperpetual

// Get Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getrisklimit)
type GetRiskLimit struct {
	Symbol *string `param:"symbol"`
}

func (o GetRiskLimit) Do(client *Client) ([]RiskLimitItem, error) {
	return GetPublic[[]RiskLimitItem](client, "risk-limit/list", o)
}

type RiskLimitItem struct {
	ID             int      `json:"id"`
	Symbol         string   `json:"symbol"`
	Limit          int      `json:"limit"`
	MaintainMargin string   `json:"maintain_margin"`
	StartingMargin string   `json:"starting_margin"`
	Section        []string `json:"section"`
	IsLowestRisk   int      `json:"is_lowest_risk"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	MaxLeverage    string   `json:"max_leverage"`
}

func (o *Client) GetRiskLimit(symbol *string) ([]RiskLimitItem, error) {
	return GetRiskLimit{Symbol: symbol}.Do(o)
}

// Set Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-setrisklimit)
//
//	symbol  Required string  Symbol
//	risk_id Required integer Risk ID
type SetRiskLimit struct {
	Symbol string `param:"symbol"`
	RiskID int    `param:"risk_id"`
}

func (o SetRiskLimit) Do(client *Client) (int, error) {
	type result struct {
		RiskID int `json:"risk_id"`
	}
	r, err := Post[result](client, "position/risk-limit", o)
	return r.RiskID, err
}

func (o *Client) SetRiskLimit(v SetRiskLimit) (int, error) {
	return v.Do(o)
}
