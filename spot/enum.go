// [Enums Definitions] https://bybit-exchange.github.io/docs/spot/v1/#t-enums
package spot

// [Side (side)]: https://bybit-exchange.github.io/docs/spot/v1/#side-side
type Side string

const (
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

// [Symbol (symbol)] https://bybit-exchange.github.io/docs/spot/v1/#symbol-symbol
type Symbol string

// [Order type (type/orderTypes)] https://bybit-exchange.github.io/docs/spot/v1/#order-type-type-ordertypes
type OrderType string

const (
	Limit       OrderType = "LIMIT"
	Market      OrderType = "MARKET"
	LimitMarket OrderType = "LIMIT_MAKER"
)

// [Currency (currency/coin)] https://bybit-exchange.github.io/docs/spot/v1/#currency-currency-coin
// The transfer API also includes:
// DOT
// DOGE
// LTC
// XLM
// USD
// Cross Margin Trading Endpoints support below currency:
// BTC
// ETH
// XRP
// SOL
// LTC
// EOS
// LINK
// XLM
// USDC
// USDT
type Currency string

const (
	BTC  Currency = "BTC"
	ETH  Currency = "ETH"
	EOS  Currency = "EOS"
	XRP  Currency = "XRP"
	USDT Currency = "USDT"

	DOT  Currency = "DOT"
	DOGE Currency = "DOGE"
	LTC  Currency = "LTC"
	XLM  Currency = "XLM"
	USD  Currency = "USD"

	SOL  Currency = "SOL"
	USDC Currency = "USDC"
)

// [Order status (status)] https://bybit-exchange.github.io/docs/spot/v1/#order-status-status
// ORDER_NEW Untriggered
// ORDER_FILLED Triggered
// ORDER_FAILED fail to trigger
type OrderStatus string

const (
	New             OrderStatus = "NEW"
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	Filled          OrderStatus = "FILLED"
	Canceled        OrderStatus = "CANCELED"
	PendingCancel   OrderStatus = "PENDING_CANCEL"
	PendingNew      OrderStatus = "PENDING_NEW"
	Rejected        OrderStatus = "REJECTED"
	// The below status for tp/sl order only
	OrderNew      OrderStatus = "ORDER_NEW"
	OrderCanceled OrderStatus = "ORDER_CANCELED"
	OrderFilled   OrderStatus = "ORDER_FILLED"
	OrderFailed   OrderStatus = "ORDER_FAILED"
)

// [Quantity (qty)] https://bybit-exchange.github.io/docs/spot/v1/#quantity-qty
type Qty uint64

// [Price (price)] https://bybit-exchange.github.io/docs/spot/v1/#price-price
type Price float64

// [Time in force (time_in_force)] https://bybit-exchange.github.io/docs/spot/v1/#time-in-force-time_in_force
// GTC - Good Till Canceled
// FOK - Fill or Kill
// IOC - Immediate or Cancel
type TimeInForce string

const (
	GoodTillCanceled  TimeInForce = "GTC"
	FillOrKill        TimeInForce = "FOK"
	ImmediateOrCancel TimeInForce = "IOC"
)

// [Kline interval (interval)] https://bybit-exchange.github.io/docs/spot/v1/#kline-interval-interval
// 1m - 1 minute
// 3m - 3 minutes
// 5m - 5 minutes
// 15m - 15 minutes
// 30m - 30 minutes
// 1h - 1 hour
// 2h - 2 hours
// 4h - 4 hours
// 6h - 6 hours
// 12h - 12 hours
// 1d - 1 day
// 1w - 1 week
// 1M - 1 month
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

// [LT status (status)] https://bybit-exchange.github.io/docs/spot/v1/#lt-status-status
// 1 - ETP can be purchased and redeemed
// 2 - ETP can be purchased, but not redeemed
// 3 - ETP can be redeemed, but not purchased
// 4 - ETP cannot be purchased nor redeemed
type LtStatus string

// [LT order status (orderStatus)] https://bybit-exchange.github.io/docs/spot/v1/#lt-order-status-orderstatus
// 1 - Completed
// 2 - In progress
// 3 - Failed
type LtOrderStatus string

const (
	Completed  LtOrderStatus = "1"
	InProgress LtOrderStatus = "2"
	Failed     LtOrderStatus = "3"
)

// [LT order type (orderType)] https://bybit-exchange.github.io/docs/spot/v1/#lt-order-type-ordertype
// 1 - Purchase
// 2 - Redemption
type LtOrderType string

const (
	Purchase   LtOrderType = "1"
	Redemption LtOrderType = "2"
)

// [TP/Sl order status (status)] https://bybit-exchange.github.io/docs/spot/v1/#tp-sl-order-status-status