package spot

type TopicName string

const (
	TopicDepth      TopicName = "depth"
	TopicKline      TopicName = "kline"
	TopicTrade      TopicName = "trade"
	TopicBookTicker TopicName = "bookTicker"
	TopicRealtimes  TopicName = "realtimes"
)

type TopicEvent string

const (
	TopicEventSub    TopicEvent = "sub"
	TopicEventCancel TopicEvent = "cancel"
)

type Topic struct {
	Name   TopicName   `json:"topic"`
	Params TopicParams `json:"params"`
	Event  TopicEvent  `json:"event"`
}

type TopicParams struct {
	Symbol     string        `json:"symbol"`
	Binary     string        `json:"binary"`
	SymbolName string        `json:"symbolName"`
	KlineType  KlineInterval `json:"klineType"`
}

type TopicNotification[T any] struct {
	Topic  TopicName   `json:"topic"`
	Params TopicParams `json:"params"`
	Data   T           `json:"data"`
}

type TopicSubscribtion struct {
	Topic
	Code    string `json:"code"`
	Message string `json:"msg"`
}

func (this *TopicSubscribtion) Ok() bool {
	return this.Code == "0"
}

func (this *TopicSubscribtion) HasCode() bool {
	return this.Code != ""
}

type TopicDataDepth struct {
	Timestamp uint64     `json:"t"` // Timestamp (last update time of the order book)
	Symbol    string     `json:"s"` // Trading pair
	Version   string     `json:"v"` // Version
	Bids      [][]string `json:"b"` // Best bid price, quantity
	Asks      [][]string `json:"a"` // Best ask price, quantity
}

type TopicDataKline struct {
	Timestamp     uint64 `json:"t"`  // Starting time
	Symbol        string `json:"s"`  // Trading pair
	SymbolName    string `json:"sn"` // Trading pair
	ClosePrice    string `json:"c"`  // Close price
	HighPrice     string `json:"h"`  // High price
	LowPrice      string `json:"l"`  // Low price
	OpenPrice     string `json:"o"`  // Open price
	TradingVolume string `json:"v"`  // Trading volume
}

type TopicDataTrade struct {
	TradeID   string `json:"v"` // Trade ID
	Timestamp uint64 `json:"t"` // Timestamp (trading time in the match box)
	Price     string `json:"p"` // Price
	Quantity  string `json:"q"` // Quantity
	M         bool   `json:"m"` // True indicates buy side is taker, false indicates sell side is taker
}

type TopicDataBookTicker struct {
	Symbol    string `json:"s"`        // Trading pair
	BidPrice  string `json:"bidPrice"` // Best bid price
	BidQty    string `json:"bidQty"`   // Bid quantity
	AskPrice  string `json:"askPrice"` // Best ask price
	AskQty    string `json:"askQty"`   // Ask quantity
	Timestamp uint64 `json:"time"`     // Timestamp (last update time of the order book)
}

type TopicDataRealtimes struct {
	Timestamp          uint64 `json:"t"`  // Timestamp (trading time in the match box)
	Symbol             string `json:"s"`  // Trading pair
	ClosePrice         string `json:"c"`  // Close price
	HighPrice          string `json:"h"`  // High price
	LowPrice           string `json:"l"`  // Low price
	OpenPrice          string `json:"o"`  // Open price
	TradingVolume      string `json:"v"`  // Trading volume
	TradingQuoteVolume string `json:"qv"` // Trading quote volume
	Change             string `json:"m"`  // Change
}
