// Account Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-accountdata)
package iperpetual

type PositionIdx int

const (
	OneWay   PositionIdx = 0
	BuySide  PositionIdx = 1
	SellSide PositionIdx = 1
)

type Direction string

const (
	Prev Direction = "prev"
	Next Direction = "next"
)
