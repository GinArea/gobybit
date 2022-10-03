// Account Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-accountdata)
package uperpetual

type PositionIdx int

const (
	OneWay   PositionIdx = 0
	BuySide  PositionIdx = 1
	SellSide PositionIdx = 1
)
