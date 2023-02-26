package spotv3

import (
	"github.com/ginarea/gobybit/transport"
)

type TopicName string

const (
	// public
	TopicDepth      TopicName = "orderbook"
	TopicTrade      TopicName = "trade"
	TopicKline      TopicName = "kline"
	TopicTickers    TopicName = "tickers"
	TopicBookTicker TopicName = "bookticker"
	// private
	TopicOrder     TopicName = "order"
	TopicStopOrder TopicName = "stopOrder"
	TopicTicket    TopicName = "ticketInfo"
	TopicOutbound  TopicName = "outboundAccountInfo"
)

type Topic[T any] struct {
	Name      string `json:"topic"`
	Type      string `json:"type"`
	Data      T      `json:"data"`
	Timestamp any    `json:"ts"`
}

type DepthShot struct {
	Timestamp uint64     `json:"t"` // Timestamp (last update time of the order book)
	Symbol    string     `json:"s"` // Trading pair
	Bids      [][]string `json:"b"` // Best bid price, quantity
	Asks      [][]string `json:"a"` // Best ask price, quantity
}

type TradeShot struct {
	TradeID   string `json:"v"` // Trade ID
	Timestamp uint64 `json:"t"` // Timestamp (trading time in the match box)
	Price     string `json:"p"` // Price
	Quantity  string `json:"q"` // Quantity
	M         bool   `json:"m"` // True indicates buy side is taker, false indicates sell side is taker
}

type KlineShot struct {
	Timestamp     uint64 `json:"t"` // Starting time
	Symbol        string `json:"s"` // Trading pair
	ClosePrice    string `json:"c"` // Close price
	HighPrice     string `json:"h"` // High price
	LowPrice      string `json:"l"` // Low price
	OpenPrice     string `json:"o"` // Open price
	TradingVolume string `json:"v"` // Trading volume
}

type TickersShot struct {
	Timestamp          uint64 `json:"t"`  // Starting time
	Symbol             string `json:"s"`  // Trading pair
	OpenPrice          string `json:"o"`  // Open price
	HighPrice          string `json:"h"`  // High price
	LowPrice           string `json:"l"`  // Low price
	ClosePrice         string `json:"c"`  // Close price
	TradingVolume      string `json:"v"`  // Trading volume
	TradingQuoteVolume string `json:"qv"` // Trading quote volume
	Change             string `json:"m"`  // Change
}

type BookTickerShot struct {
	Symbol       string            `json:"s"`  // Trading pair
	BestBidPrice transport.Float64 `json:"bp"` // Best bid price
	BidQuantity  transport.Float64 `json:"bq"` // Bid quantity
	BestAskPrice transport.Float64 `json:"qp"` // Best ask price
	AskQuantity  transport.Float64 `json:"qq"` // Ask quantity
	Timestamp    uint64            `json:"t"`  // The time that message is sent out
}

type OrderShot struct {
	EventType           string              `json:"e"` // Event type
	EventTime           transport.Timestamp `json:"E"` // Event time
	Symbol              string              `json:"s"` // Trading pair
	UserOrderID         string              `json:"c"` // User-generated order ID
	Side                Side                `json:"S"` // BUY indicates buy order, SELL indicates sell order
	OrderType           OrderType           `json:"o"` // Order type, LIMIT/MARKET_OF_QUOTE/MARKET_OF_BASE
	TimeInForce         TimeInForce         `json:"f"` // Time in force
	Quantity            transport.Float64   `json:"q"` // Quantity
	Price               transport.Float64   `json:"p"` // Price
	OrderStatus         OrderStatus         `json:"X"` // Order status
	OrderID             string              `json:"i"` // Order ID
	OrderIDofOpponent   string              `json:"M"` // Order ID of the opponent trader
	LastFilledQuantity  string              `json:"l"` // Last filled quantity
	TotalFilledQuantity string              `json:"z"` // Total filled quantity
	LastTradedPrice     string              `json:"L"` // Last traded price
	TradingFee          transport.Float64   `json:"n"` // Trading fee (for a single fill)
	AssetType           string              `json:"N"` // Asset type in which fee is paid
	IsNormalTrade       bool                `json:"u"` // Is normal trade. False if self-trade.
	IsWorking           bool                `json:"w"` // Is working
	IsLimitMaker        bool                `json:"m"` // Is LIMIT_MAKER
	OrderCreationTime   transport.Timestamp `json:"O"` // Order creation time
	TotalFilledValue    string              `json:"Z"` // Total filled value
	AccountID           string              `json:"A"` // Account ID of the opponent trader
	IsClose             bool                `json:"C"` // Is close
	Leverage            string              `json:"v"` // Leverage
	Liquidation         string              `json:"d"` // NO_LIQ indicates that it is not a liquidation order. IOC indicates that it is a liquidation order.
	TradeID             string              `json:"t"` // Trade ID
}

type StopOrderShot struct {
	EventType          string `json:"e"` // Event type
	EventTime          string `json:"E"` // Event time
	Symbol             string `json:"s"` // Trading pair
	UserOrderID        string `json:"c"` // User-generated order ID
	Side               string `json:"S"` // BUY indicates buy order, SELL indicates sell order
	OrderType          string `json:"o"` // Order type, LIMIT/MARKET_OF_QUOTE/MARKET_OF_BASE
	TimeInForce        string `json:"f"` // Time in force
	Quantity           string `json:"q"` // Quantity
	Price              string `json:"p"` // Price
	OrderStatus        string `json:"X"` // Order status
	OrderID            string `json:"i"` // Order ID
	OrderCreationTime  string `json:"T"` // Order creation time
	OrderTriggeredTime string `json:"t"` // Order triggered time
	OrderUpdatedTime   string `json:"C"` // Order updated time
}

type TicketShot struct {
	EventType           string `json:"e"` // Event type
	EventTime           string `json:"E"` // Event time
	Symbol              string `json:"s"` // Trading pair
	Quantity            string `json:"q"` // Quantity
	Timestamp           string `json:"t"` // Timestamp
	Price               string `json:"p"` // Price
	TradeID             string `json:"T"` // Trade ID
	OrderID             string `json:"o"` // Order ID
	OrderIDofOpponent   string `json:"O"` // Order ID of the opponent trader
	AccountID           string `json:"a"` // Account ID
	AccountIDofOpponent string `json:"A"` // Account ID of the opponent trader
	IsLimitMaker        bool   `json:"m"` // Is LIMIT_MAKER
	Side                string `json:"S"` // BUY indicates buy order, SELL indicates sell order
}

type OutboundShot struct {
	EventType           string         `json:"e"` // Event type
	Timestamp           string         `json:"E"` // Timestamp
	AllowTrade          bool           `json:"T"` // Allow trade
	AllowWithdraw       bool           `json:"W"` // Allow withdraw
	AllowDeposit        bool           `json:"D"` // Allow deposit
	WalletBalanceChange []OutboundItem `json:"B"` // Wallet balance change
}

type OutboundItem struct {
	Coin      string            `json:"a"` // Coin name
	Available transport.Float64 `json:"f"` // Available balance
	Reserved  transport.Float64 `json:"l"` // Reserved for orders
}
