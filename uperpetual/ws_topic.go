package uperpetual

type TopicName string

const (
	// public:
	TopicOrderBook25  TopicName = "orderBookL2_25"
	TopicOrderBook200 TopicName = "orderBook_200"
	TopicTrade        TopicName = "trade"
	TopicInstrument   TopicName = "instrument_info"
	TopicKline        TopicName = "candle"
	TopicLiquidation  TopicName = "liquidation"
	// private:
	TopicPosition  TopicName = "position"
	TopicExecution TopicName = "execution"
	TopicOrder     TopicName = "order"
	TopicStopOrder TopicName = "stop_order"
	TopicWallet    TopicName = "wallet"
)

type Topic[T any] struct {
	Name string `json:"topic"`
	Data T      `json:"data"`
}

type OrderBookSnapshot struct {
	Price  string  `json:"price"`
	Symbol string  `json:"symbol"`
	ID     string  `json:"id"`
	Side   Side    `json:"side"`
	Size   float64 `json:"size"`
}

type OrderBookDelta struct {
	Delete []any `json:"delete"`
	Update []any `json:"update"`
	Insert []any `json:"insert"`
}

type TradeSnapshot struct {
	Timestamp     string        `json:"timestamp"`
	TradeTime     string        `json:"trade_time_ms"`
	Symbol        string        `json:"symbol"`
	Side          Side          `json:"side"`
	Size          float64       `json:"size"`
	Price         string        `json:"price"`
	TickDirection TickDirection `json:"tick_direction"`
	TradeID       string        `json:"trade_id"`
	IsBlockTrade  string        `json:"is_block_trade"`
}

type InstrumentSnapshot struct {
	ID                     uint64        `json:"id"`
	Symbol                 string        `json:"symbol"`
	LastPriceE4            string        `json:"last_price_e4"`
	LastPrice              string        `json:"last_price"`
	Bid1PriceE4            string        `json:"bid1_price_e4"`
	Bid1Price              string        `json:"bid1_price"`
	Ask1PriceE4            string        `json:"ask1_price_e4"`
	Ask1Price              string        `json:"ask1_price"`
	LastTickDirection      TickDirection `json:"last_tick_direction"`
	PrevPrice24hE4         string        `json:"prev_price_24h_e4"`
	PrevPrice24h           string        `json:"prev_price_24h"`
	HighPrice24hE4         string        `json:"high_price_24h_e4"`
	HighPrice24h           string        `json:"high_price_24h"`
	LowPrice24hE4          string        `json:"low_price_24h_e4"`
	LowPrice24h            string        `json:"low_price_24h"`
	PrevPrice1hE4          string        `json:"prev_price_1h_e4"`
	PrevPrice1h            string        `json:"prev_price_1h"`
	MarkPriceE4            string        `json:"mark_price_e4"`
	MarkPrice              string        `json:"mark_price"`
	IndexPriceE4           string        `json:"index_price_e4"`
	IndexPrice             string        `json:"index_price"`
	OpenInterest           string        `json:"open_interest"`
	OpenValueE8            string        `json:"open_value_e8"`
	TotalTurnoverE8        string        `json:"total_turnover_e8"`
	Turnover24hE8          string        `json:"turnover_24h_e8"`
	TotalVolume            string        `json:"total_volume"`
	Volume24h              string        `json:"volume_24h"`
	FundingRateE6          string        `json:"funding_rate_e6"`
	PredictedFundingRateE6 string        `json:"predicted_funding_rate_e6"`
	CrossSeq               string        `json:"cross_seq"`
	CreatedAt              string        `json:"created_at"`
	UpdatedAt              string        `json:"updated_at"`
	NextFundingTime        string        `json:"next_funding_time"`
	CountdownHour          string        `json:"countdown_hour"`
	FundingRateInterval    string        `json:"funding_rate_interval"`
	SettleTimeE9           string        `json:"settle_time_e9"`
	DelistingStatus        string        `json:"delisting_status"`
}

type InstrumentDelta struct {
	Delete []any `json:"delete"`
	Update []any `json:"update"`
	Insert []any `json:"insert"`
}

type KlineSnapshot struct {
	Start     uint64        `json:"start"`
	End       uint64        `json:"end"`
	Period    KlineInterval `json:"period"`
	Open      float64       `json:"open"`
	Close     float64       `json:"close"`
	High      float64       `json:"high"`
	Low       float64       `json:"low"`
	Volume    string        `json:"volume"`
	Turnover  string        `json:"turnover"`
	Confirm   bool          `json:"confirm"`
	CrossSeq  float64       `json:"cross_seq"`
	Timestamp uint64        `json:"timestamp"`
}

type LiquidationSnapshot struct {
	Symbol string `json:"symbol"`
	Side   Side   `json:"side"`
	Price  string `json:"price"`
	Qty    string `json:"qty"`
	Time   int64  `json:"time"`
}

type PositionSnapshot struct {
	UserID           int         `json:"user_id"`
	Symbol           string      `json:"symbol"`
	Size             int         `json:"size"`
	Side             Side        `json:"side"`
	PositionValue    string      `json:"position_value"`
	EntryPrice       string      `json:"entry_price"`
	LiqPrice         string      `json:"liq_price"`
	BustPrice        string      `json:"bust_price"`
	Leverage         string      `json:"leverage"`
	OrderMargin      string      `json:"order_margin"`
	PositionMargin   string      `json:"position_margin"`
	AvailableBalance string      `json:"available_balance"`
	TakeProfit       string      `json:"take_profit"`
	StopLoss         string      `json:"stop_loss"`
	RealisedPnl      string      `json:"realised_pnl"`
	TrailingStop     string      `json:"trailing_stop"`
	TrailingActive   string      `json:"trailing_active"`
	WalletBalance    string      `json:"wallet_balance"`
	RiskID           int         `json:"risk_id"`
	OccClosingFee    string      `json:"occ_closing_fee"`
	OccFundingFee    string      `json:"occ_funding_fee"`
	AutoAddMargin    int         `json:"auto_add_margin"`
	CumRealisedPnl   string      `json:"cum_realised_pnl"`
	PositionStatus   string      `json:"position_status"`
	PositionSeq      int         `json:"position_seq"`
	IsIsolated       bool        `json:"is_isolated"`
	Mode             int         `json:"mode"`
	PositionIdx      PositionIdx `json:"position_idx"`
	TpSlMode         TpSlMode    `json:"tp_sl_mode"`
	TpOrderNum       int         `json:"tp_order_num"`
	SlOrderNum       int         `json:"sl_order_num"`
	TpFreeSize       int         `json:"tp_free_size_x"`
	SlFreeSize       int         `json:"sl_free_size_x"`
}

type ExecutionSnapshot struct {
	OrderID     string   `json:"order_id"`
	OrderLinkID string   `json:"order_link_id"`
	Symbol      string   `json:"symbol"`
	Side        Side     `json:"side"`
	ExecID      string   `json:"exec_id"`
	Price       string   `json:"price"`
	OrderQty    float64  `json:"order_qty"`
	ExecType    ExecType `json:"exec_type"`
	ExecQty     int      `json:"exec_qty"`
	ExecFee     string   `json:"exec_fee"`
	LeavesQty   float64  `json:"leaves_qty"`
	IsMaker     bool     `json:"is_maker"`
	TradeTime   string   `json:"trade_time"`
}

type OrderSnapshot struct {
	OrderID        string       `json:"order_id"`
	OrderLinkID    string       `json:"order_link_id"`
	Symbol         string       `json:"symbol"`
	Side           Side         `json:"side"`
	OrderType      OrderType    `json:"order_type"`
	Price          float64      `json:"price"`
	Qty            string       `json:"qty"`
	TimeInForce    TimeInForce  `json:"time_in_force"`
	CreateType     CreateType   `json:"create_type"`
	CancelType     CancelType   `json:"cancel_type"`
	OrderStatus    OrderStatus  `json:"order_status"`
	LeavesQty      float64      `json:"leaves_qty"`
	CumExecQty     float64      `json:"cum_exec_qty"`
	CumExecValue   string       `json:"cum_exec_value"`
	CumExecFee     string       `json:"cum_exec_fee"`
	Timestamp      string       `json:"timestamp"`
	TakeProfit     float64      `json:"take_profit"`
	TpTrigger      TriggerPrice `json:"tp_trigger_by"`
	SlTrigger      TriggerPrice `json:"sl_trigger_by"`
	StopLoss       float64      `json:"stop_loss"`
	TrailingStop   string       `json:"trailing_stop"`
	LastExecPrice  string       `json:"last_exec_price"`
	ReduceOnly     bool         `json:"reduce_only"`
	CloseOnTrigger bool         `json:"close_on_trigger"`
}

type StopOrderSnapshot struct {
	OrderID        string       `json:"order_id"`
	OrderLinkID    string       `json:"order_link_id"`
	UserID         int          `json:"user_id"`
	Symbol         string       `json:"symbol"`
	Side           Side         `json:"side"`
	OrderType      OrderType    `json:"order_type"`
	Price          float64      `json:"price"`
	CreateType     CreateType   `json:"create_type"`
	CancelType     CancelType   `json:"cancel_type"`
	OrderStatus    OrderStatus  `json:"order_status"`
	StopOrderType  StopOrder    `json:"stop_order_type"`
	TriggerBy      TriggerPrice `json:"trigger_by"`
	TriggerPrice   string       `json:"trigger_price"`
	CloseOnTrigger bool         `json:"close_on_trigger"`
	Timestamp      string       `json:"timestamp"`
	TakeProfit     float64      `json:"take_profit"`
	StopLoss       float64      `json:"stop_loss"`
}

type WalletSnapshot struct {
	UserID           uint64 `json:"user_id"`
	Coin             string `json:"coin"`
	AvailableBalance string `json:"available_balance"`
	WalletBalance    string `json:"wallet_balance"`
}
