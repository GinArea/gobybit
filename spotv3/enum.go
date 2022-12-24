// Enums Definitions (https://bybit-exchange.github.io/docs/spot/v3/#t-enums)
package spotv3

// Side (side) (https://bybit-exchange.github.io/docs/spot/v3/#side-side)
type Side string

const (
	None Side = "None"
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

// Time in force (time_in_force) (https://bybit-exchange.github.io/docs/spot/v3/#time-in-force-timeinforce)
//
//	GTC - Good Till Canceled
//	FOK - Fill or Kill
//	IOC - Immediate or Cancel
type TimeInForce string

const (
	GoodTillCanceled  TimeInForce = "GTC"
	FillOrKill        TimeInForce = "FOK"
	ImmediateOrCancel TimeInForce = "IOC"
)

// Symbol (symbol) (https://bybit-exchange.github.io/docs/spot/v3/#symbol-symbol)

// Order type (type/orderTypes) (https://bybit-exchange.github.io/docs/spot/v3/#order-type-type-ordertypes)
type OrderType string

const (
	Limit       OrderType = "LIMIT"
	Market      OrderType = "MARKET"
	LimitMarket OrderType = "LIMIT_MAKER"
)

// Currency (currency/coin) https://bybit-exchange.github.io/docs/spot/v1/#currency-currency-coin

// Order status (status) (https://bybit-exchange.github.io/docs/spot/v3/#order-status-status)
//
//	ORDER_NEW Untriggered
//	ORDER_FILLED Triggered
//	ORDER_FAILED fail to trigger
type OrderStatus string

const (
	New             OrderStatus = "NEW"
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	Filled          OrderStatus = "FILLED"
	Canceled        OrderStatus = "CANCELED"
	PendingCancel   OrderStatus = "PENDING_CANCEL"
	PendingNew      OrderStatus = "PENDING_NEW"
	Rejected        OrderStatus = "REJECTED"
)

// Quantity (qty) (https://bybit-exchange.github.io/docs/spot/v3/#quantity-qty)
type Qty uint64

// Price (price) (https://bybit-exchange.github.io/docs/spot/v3/#price-price)
type Price float64

// Kline interval (interval) (https://bybit-exchange.github.io/docs/spot/v3/#kline-interval-interval)
//
//	1m - 1 minute
//	3m - 3 minutes
//	5m - 5 minutes
//	15m - 15 minutes
//	30m - 30 minutes
//	1h - 1 hour
//	2h - 2 hours
//	4h - 4 hours
//	6h - 6 hours
//	12h - 12 hours
//	1d - 1 day
//	1w - 1 week
//	1M - 1 month
type KlineInterval string

const (
	Interval1m  KlineInterval = "1m"
	Interval3m  KlineInterval = "3m"
	Interval5m  KlineInterval = "5m"
	Interval15m KlineInterval = "15m"
	Interval30m KlineInterval = "30m"
	Interval1h  KlineInterval = "1h"
	Interval2h  KlineInterval = "2h"
	Interval4h  KlineInterval = "4h"
	Interval6h  KlineInterval = "6h"
	Interval12h KlineInterval = "12h"
	Interval1d  KlineInterval = "1d"
	Interval1w  KlineInterval = "1w"
	Interval1M  KlineInterval = "1M"
)

// LT status (status) (https://bybit-exchange.github.io/docs/spot/v3/#lt-status-status)
//
//	1 - ETP can be purchased and redeemed
//	2 - ETP can be purchased, but not redeemed
//	3 - ETP can be redeemed, but not purchased
//	4 - ETP cannot be purchased nor redeemed
type LtStatus string

// LT order status (orderStatus) (https://bybit-exchange.github.io/docs/spot/v3/#lt-order-status-orderstatus)
//
//	1 - Completed
//	2 - In progress
//	3 - Failed
type LtOrderStatus string

const (
	Completed  LtOrderStatus = "1"
	InProgress LtOrderStatus = "2"
	Failed     LtOrderStatus = "3"
)

// LT order type (orderType) (https://bybit-exchange.github.io/docs/spot/v3/#lt-order-type-ordertype)
//
//	1 - Purchase
//	2 - Redemption
type LtOrderType string

const (
	Purchase   LtOrderType = "1"
	Redemption LtOrderType = "2"
)

// TP/Sl order status (status) (https://bybit-exchange.github.io/docs/spot/v3/#tp-sl-order-status-status)
