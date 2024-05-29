package bybitv5

import (
	"encoding/json"

	"github.com/msw-x/moon/ujson"
)

type PublicHeader struct {
	Ts   int64
	Type string
}

func (o PublicHeader) IsDelta() bool {
	return o.Type == "delta"
}

type PrivateHeader struct {
	CreationTime int64
	Id           string
}

type Topic[T any] struct {
	Topic string
	PublicHeader
	PrivateHeader
	Data T
}

type RawTopic Topic[json.RawMessage]

func UnmarshalRawTopic[T any](raw RawTopic) (ret Topic[T], err error) {
	ret.Topic = raw.Topic
	ret.PublicHeader = raw.PublicHeader
	ret.PrivateHeader = raw.PrivateHeader
	err = json.Unmarshal(raw.Data, &ret.Data)
	return
}

type TradeShot struct {
	Timestamp     int64         `json:"ts"`
	Symbol        string        `json:"s"`
	Side          Side          `json:"S"`
	Size          ujson.Float64 `json:"v"`
	Price         ujson.Float64 `json:"p"`
	TickDirection TickDirection `json:"L"` // Unique field for future
	Id            string        `json:"i"`
	Block         bool          `json:"BT"`
}

type TickerShot struct {
	Symbol            string
	TickDirection     TickDirection
	Price24HPcnt      ujson.StringFloat64
	LastPrice         ujson.StringFloat64
	PrevPrice24H      ujson.StringFloat64
	HighPrice24H      ujson.StringFloat64
	LowPrice24H       ujson.StringFloat64
	PrevPrice1H       ujson.StringFloat64
	MarkPrice         ujson.StringFloat64
	IndexPrice        ujson.StringFloat64
	OpenInterest      ujson.StringFloat64
	OpenInterestValue ujson.StringFloat64
	Turnover24H       ujson.StringFloat64
	Volume24H         ujson.StringFloat64
	NextFundingTime   ujson.StringFloat64
	FundingRate       ujson.StringFloat64
	Bid1Price         ujson.StringFloat64
	Bid1Size          ujson.StringFloat64
	Ask1Price         ujson.StringFloat64
	Ask1Size          ujson.StringFloat64
}

type TickerOptionShot struct {
	Symbol                 string
	BidPrice               ujson.StringFloat64
	BidSize                ujson.StringFloat64
	BidIv                  ujson.StringFloat64
	AskPrice               ujson.StringFloat64
	AskSize                ujson.StringFloat64
	AskIv                  ujson.StringFloat64
	LastPrice              ujson.StringFloat64
	HighPrice24H           ujson.StringFloat64
	LowPrice24H            ujson.StringFloat64
	MarkPrice              ujson.StringFloat64
	IndexPrice             ujson.StringFloat64
	MarkPriceIv            ujson.StringFloat64
	UnderlyingPrice        ujson.StringFloat64
	OpenInterest           ujson.StringFloat64
	Turnover24H            ujson.StringFloat64
	Volume24H              ujson.StringFloat64
	TotalVolume            ujson.StringFloat64
	TotalTurnover          ujson.StringFloat64
	Delta                  ujson.StringFloat64
	Gamma                  ujson.StringFloat64
	Vega                   ujson.StringFloat64
	Theta                  ujson.StringFloat64
	PredictedDeliveryPrice ujson.StringFloat64
	Change24H              ujson.StringFloat64
}

type TickerSpotShot struct {
	Symbol        string
	LastPrice     ujson.StringFloat64
	HighPrice24H  ujson.StringFloat64
	LowPrice24H   ujson.StringFloat64
	PrevPrice24H  ujson.StringFloat64
	Volume24H     ujson.StringFloat64
	Turnover24H   ujson.StringFloat64
	Price24HPcnt  ujson.StringFloat64
	UsdIndexPrice ujson.StringFloat64
}

type KlineShot struct {
	Start     int64
	End       int64
	Interval  Interval
	Open      ujson.Float64
	Close     ujson.Float64
	High      ujson.Float64
	Low       ujson.Float64
	Volume    ujson.Float64
	Turnover  ujson.Float64
	Confirm   bool
	Timestamp int64
}

type LiquidationShot struct {
	Price       ujson.Float64
	Side        Side
	Size        ujson.Float64
	Symbol      string
	UpdatedTime int64
}

type LtTickerShot struct {
	Symbol       string
	LastPrice    ujson.Float64
	HighPrice24h ujson.Float64
	LowPrice24h  ujson.Float64
	PrevPrice24h ujson.Float64
	Price24hPcnt ujson.Float64
}

type LtNavShot struct {
	Symbol         string
	Time           int64
	Nav            ujson.Float64
	BasketPosition ujson.Float64
	Leverage       ujson.Float64
	BasketLoan     ujson.Float64
	Circulation    ujson.Float64
	Basket         ujson.Float64
}

type PositionShot struct {
	Position
	Category   Category
	EntryPrice ujson.Float64
}

type ExecutionShot struct {
	Category        Category
	Symbol          ujson.StringFloat64
	ExecFee         ujson.StringFloat64
	ExecId          ujson.StringFloat64
	ExecPrice       ujson.StringFloat64
	ExecQty         ujson.StringFloat64
	ExecType        ExecType
	ExecValue       string
	IsMaker         bool
	FeeRate         ujson.StringFloat64
	TradeIv         ujson.StringFloat64
	MarkIv          string
	BlockTradeId    string
	MarkPrice       ujson.StringFloat64
	IndexPrice      ujson.StringFloat64
	UnderlyingPrice ujson.StringFloat64
	LeavesQty       ujson.StringFloat64
	OrderId         string
	OrderLinkId     string
	OrderPrice      ujson.StringFloat64
	OrderQty        ujson.StringFloat64
	OrderType       OrderType
	StopOrderType   StopOrderType
	Side            Side
	ExecTime        string
	IsLeverage      string
	ClosedSize      string
}

type OrderShot struct {
	Category Category
	Order
}

type WalletShot struct {
	AccountType            AccountType
	AccountLTV             ujson.StringFloat64
	AccountIMRate          ujson.StringFloat64
	AccountMMRate          ujson.StringFloat64
	TotalEquity            ujson.StringFloat64
	TotalWalletBalance     ujson.StringFloat64
	TotalMarginBalance     ujson.StringFloat64
	TotalAvailableBalance  ujson.StringFloat64
	TotalPerpUPL           ujson.StringFloat64
	TotalInitialMargin     ujson.StringFloat64
	TotalMaintenanceMargin ujson.StringFloat64
	Coin                   []WalletCoin
}

type WalletCoin struct {
	Coin                string
	Equity              ujson.StringFloat64
	UsdValue            ujson.StringFloat64
	WalletBalance       ujson.StringFloat64
	Free                ujson.StringFloat64
	Locked              ujson.StringFloat64
	AvailableToWithdraw ujson.StringFloat64
	AvailableToBorrow   ujson.StringFloat64
	BorrowAmount        ujson.StringFloat64
	AccruedInterest     ujson.StringFloat64
	TotalOrderIm        ujson.StringFloat64
	TotalPositionIm     ujson.StringFloat64
	TotalPositionMm     ujson.StringFloat64
	UnrealisedPnl       ujson.StringFloat64
	CumRealisedPnl      ujson.StringFloat64
	Bonus               ujson.StringFloat64
}

type GreekShot struct {
	BaseCoin   string
	TotalDelta ujson.StringFloat64
	TotalGamma ujson.StringFloat64
	TotalVega  ujson.StringFloat64
	TotalTheta ujson.StringFloat64
}
