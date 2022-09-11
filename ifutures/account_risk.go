// [Risk Limit] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-risklimit
package ifutures

import "github.com/tranquiil/bybit/iperpetual"

// [Set Risk Limit] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-setrisklimit
// symbol       Required string  Symbol
// risk_id      Required integer Risk ID
// position_idx          integer Position idx, used to identify positions in different position modes:
type SetRiskLimit struct {
	Symbol      iperpetual.Symbol `param:"symbol"`
	RiskID      int               `param:"risk_id"`
	PositionIdx *PositionIdx      `param:"position_idx"`
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
