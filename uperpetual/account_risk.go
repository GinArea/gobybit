// Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-risklimit)
package uperpetual

// Get Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-getrisklimit)
type GetRiskLimit struct {
	Symbol *string `param:"symbol"`
}

func (this GetRiskLimit) Do(client *Client) ([]RiskLimitItem, bool) {
	return GetPublic[[]RiskLimitItem](client, "risk-limit", this)
}

type RiskLimitItem struct {
	ID             int      `json:"id"`
	Symbol         string   `json:"symbol"`
	Limit          int      `json:"limit"`
	MaintainMargin float64  `json:"maintain_margin"`
	StartingMargin float64  `json:"starting_margin"`
	Section        []string `json:"section"`
	IsLowestRisk   int      `json:"is_lowest_risk"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	MaxLeverage    float64  `json:"max_leverage"`
}

func (this *Client) GetRiskLimit(symbol *string) ([]RiskLimitItem, bool) {
	return GetRiskLimit{Symbol: symbol}.Do(this)
}

// Set Risk Limit (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-setrisklimit)
//
//	symbol  Required string  Symbol
//	risk_id Required integer Risk ID
type SetRiskLimit struct {
	Symbol      string       `param:"symbol"`
	Side        Side         `param:"side"`
	RiskID      int          `param:"risk_id"`
	PositionIdx *PositionIdx `param:"position_idx"`
}

func (this SetRiskLimit) Do(client *Client) (int, bool) {
	type result struct {
		RiskID int `json:"risk_id"`
	}
	r, ok := Post[result](client, "position/set-risk", this)
	return r.RiskID, ok
}

func (this *Client) SetRiskLimit(v SetRiskLimit) (int, bool) {
	return v.Do(this)
}
