package iperpetual

type WsPublic struct {
	WsSection
}

func NewWsPublic(client *WsClient) *WsPublic {
	c := &WsPublic{}
	c.init(client)
	return c
}

func (this WsPublic) OrderBook25(symbol string) *WsDeltaExecutor[[]OrderBookShot] {
	return NewWsDeltaExecutor[[]OrderBookShot](&this.WsSection, Subscription{Topic: TopicOrderBook25, Symbol: symbol})
}

func (this WsPublic) OrderBook200(symbol string) *WsDeltaExecutor[[]OrderBookShot] {
	return NewWsDeltaExecutor[[]OrderBookShot](&this.WsSection, Subscription{Topic: TopicOrderBook200, Interval: "100ms", Symbol: symbol})
}

func (this WsPublic) Trade() *WsExecutor[[]TradeShot] {
	return NewWsExecutor[[]TradeShot](&this.WsSection, Subscription{Topic: TopicTrade})
}

func (this WsPublic) Insurance() *WsExecutor[[]InsuranceShot] {
	return NewWsExecutor[[]InsuranceShot](&this.WsSection, Subscription{Topic: TopicInsurance})
}

func (this WsPublic) Instrument(symbol string) *WsDeltaExecutor[InstrumentShot] {
	return NewWsDeltaExecutor[InstrumentShot](&this.WsSection, Subscription{Topic: TopicInstrument, Interval: "100ms", Symbol: symbol})
}

func (this WsPublic) Kline(symbol string, interval KlineInterval) *WsExecutor[[]KlineShot] {
	return NewWsExecutor[[]KlineShot](&this.WsSection, Subscription{Topic: TopicKline, Interval: string(interval), Symbol: symbol})
}

func (this WsPublic) Liquidation() *WsExecutor[LiquidationShot] {
	return NewWsExecutor[LiquidationShot](&this.WsSection, Subscription{Topic: TopicLiquidation})
}
