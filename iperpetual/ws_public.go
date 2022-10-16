package iperpetual

type WsPublic struct {
	WsSection
}

func NewWsPublic(client *WsClient) *WsPublic {
	c := &WsPublic{}
	c.init(client)
	return c
}

func (this WsPublic) OrderBook25(symbol string) *WsDualExecutor[[]OrderBookShot] {
	return NewWsDualExecutor[[]OrderBookShot](&this.WsSection, Subscription{Topic: TopicOrderBook25, Symbol: symbol})
}

func (this WsPublic) OrderBook200(symbol string) *WsDualExecutor[[]OrderBookShot] {
	return NewWsDualExecutor[[]OrderBookShot](&this.WsSection, Subscription{Topic: TopicOrderBook200, Interval: "100ms", Symbol: symbol})
}

func (this WsPublic) Trade() *WsMonoExecutor[[]TradeShot] {
	return NewWsMonoExecutor[[]TradeShot](&this.WsSection, Subscription{Topic: TopicTrade})
}

func (this WsPublic) Insurance() *WsMonoExecutor[[]InsuranceShot] {
	return NewWsMonoExecutor[[]InsuranceShot](&this.WsSection, Subscription{Topic: TopicInsurance})
}

func (this WsPublic) Instrument(symbol string) *WsDualExecutor[InstrumentShot] {
	return NewWsDualExecutor[InstrumentShot](&this.WsSection, Subscription{Topic: TopicInstrument, Interval: "100ms", Symbol: symbol})
}

func (this WsPublic) Kline(symbol string, interval KlineInterval) *WsMonoExecutor[[]KlineShot] {
	return NewWsMonoExecutor[[]KlineShot](&this.WsSection, Subscription{Topic: TopicKline, Interval: string(interval), Symbol: symbol})
}

func (this WsPublic) Liquidation() *WsMonoExecutor[LiquidationShot] {
	return NewWsMonoExecutor[LiquidationShot](&this.WsSection, Subscription{Topic: TopicLiquidation})
}
