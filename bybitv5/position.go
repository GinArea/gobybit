package bybitv5

import "github.com/msw-x/moon/ujson"

// Get Position Info
// https://bybit-exchange.github.io/docs/v5/position
type GetPositions struct {
	Category   Category
	Symbol     string `url:",omitempty"`
	BaseCoin   string `url:",omitempty"`
	SettleCoin string `url:",omitempty"`
	Limit      int    `url:",omitempty"`
	Cursor     string `url:",omitempty"`
}

type Position struct {
	PositionIdx      PositionIdx
	RiskId           int
	RiskLimitValue   ujson.Int64
	Symbol           string
	Side             Side
	Size             ujson.Float64
	AvgPrice         ujson.Float64
	PositionValue    ujson.Float64
	TradeMode        TradeMode
	PositionStatus   PositionStatus
	AutoAddMargin    ujson.Bool
	AdlRankIndicator int
	Leverage         ujson.Float64
	PositionBalance  ujson.Float64
	MarkPrice        ujson.Float64
	LiqPrice         ujson.Float64
	BustPrice        ujson.Float64
	PositionMm       ujson.Float64
	PositionIm       ujson.Float64
	TpslMode         TpSlMode
	TakeProfit       ujson.Float64
	StopLoss         ujson.Float64
	TrailingStop     ujson.Float64
	UnrealisedPnl    ujson.Float64
	CumRealisedPnl   ujson.Float64
	CreatedTime      ujson.TimeMs
	UpdatedTime      ujson.TimeMs
}

func (o GetPositions) Do(c *Client) Response[[]Position] {
	type result struct {
		Category       Category
		NextPageCursor string
		List           []Position
	}
	return Get(c.position(), "list", o, func(r result) ([]Position, error) {
		return r.List, nil
	})
}

func (o *Client) GetPositions(v GetPositions) Response[[]Position] {
	return v.Do(o)
}
