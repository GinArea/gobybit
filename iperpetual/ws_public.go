package iperpetual

type WsPublic struct {
	WsSection
}

func NewWsPublic(client *WsClient) *WsPublic {
	c := &WsPublic{}
	c.init(client)
	return c
}

func (o WsPublic) OrderBook25(symbol string) *WsDeltaExecutor[[]OrderBookShot] {
	return NewWsDeltaExecutor[[]OrderBookShot](&o.WsSection, Subscription{Topic: TopicOrderBook25, Symbol: symbol})
}

func (o WsPublic) OrderBook200(symbol string) *WsDeltaExecutor[[]OrderBookShot] {
	return NewWsDeltaExecutor[[]OrderBookShot](&o.WsSection, Subscription{Topic: TopicOrderBook200, Interval: "100ms", Symbol: symbol})
}

func (o WsPublic) Trade() *WsExecutor[[]TradeShot] {
	return NewWsExecutor[[]TradeShot](&o.WsSection, Subscription{Topic: TopicTrade})
}

func (o WsPublic) Insurance() *WsExecutor[[]InsuranceShot] {
	return NewWsExecutor[[]InsuranceShot](&o.WsSection, Subscription{Topic: TopicInsurance})
}

func (o WsPublic) Instrument(symbol string) *WsDeltaExecutor[InstrumentShot] {
	return NewWsDeltaExecutor[InstrumentShot](&o.WsSection, Subscription{Topic: TopicInstrument, Interval: "100ms", Symbol: symbol})
}

func (o WsPublic) Kline(symbol string, interval KlineInterval) *WsExecutor[[]KlineShot] {
	return NewWsExecutor[[]KlineShot](&o.WsSection, Subscription{Topic: TopicKline, Interval: string(interval), Symbol: symbol})
}

func (o WsPublic) Liquidation() *WsExecutor[LiquidationShot] {
	return NewWsExecutor[LiquidationShot](&o.WsSection, Subscription{Topic: TopicLiquidation})
}
