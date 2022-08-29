package bybit

type Coin string

const (
	CoinBTC  Coin = "BTC"
	CoinETH       = "ETH"
	CoinEOS       = "EOS"
	CoinXRP       = "XRP"
	CoinUSDT      = "USDT"
)

type SymbolInverse string

const (
	SymbolInverseBTCUSD SymbolInverse = "BTCUSD"
	SymbolInverseETHUSD               = "ETHUSD"
	SymbolInverseEOSUSD               = "EOSUSD"
	SymbolInverseXRPUSD               = "XRPUSD"
)

type SymbolUSDT string

const (
	SymbolUSDTBTC  SymbolUSDT = "BTCUSDT"
	SymbolUSDTETH             = "ETHUSDT"
	SymbolUSDTLTC             = "LTCUSDT"
	SymbolUSDTLINK            = "LINKUSDT"
	SymbolUSDTXTZ             = "XTZUSDT"
	SymbolUSDTBCH             = "BCHUSDT"
	SymbolUSDTADA             = "ADAUSDT"
	SymbolUSDTDOT             = "DOTUSDT"
	SymbolUSDTUNI             = "UNIUSDT"
)

type SymbolSpot string

const (
	SymbolSpotBTCUSDT    SymbolSpot = "BTCUSDT"
	SymbolSpotETHUSDT               = "ETHUSDT"
	SymbolSpotEOSUSDT               = "EOSUSDT"
	SymbolSpotXRPUSDT               = "XRPUSDT"
	SymbolSpotUNIUSDT               = "UNIUSDT"
	SymbolSpotBTCETH                = "BTCETH"
	SymbolSpotDOGEXRP               = "DOGEXRP"
	SymbolSpotXLMUSDT               = "XLMUSDT"
	SymbolSpotLTCUSDT               = "LTCUSDT"
	SymbolSpotXRPBTC                = "XRPBTC"
	SymbolSpotDOGEUSDT              = "DOGEUSDT"
	SymbolSpotBITUSDT               = "BITUSDT"
	SymbolSpotMANAUSDT              = "MANAUSDT"
	SymbolSpotAXSUSDT               = "AXSUSDT"
	SymbolSpotDYDXUSDT              = "DYDXUSDT"
	SymbolSpotPMTEST5BTC            = "PMTEST5BTC"
	SymbolSpotGENEUSDT              = "GENEUSDT"
)

type Side string

const (
	SideNone Side = "NONE"
	SideBuy       = "BUY"
	SideSell      = "SELL"
)

type OrderType string

const (
	OrderTypeLimit      OrderType = "LIMIT"
	OrderTypeMarket               = "MARKET"
	OrderTypeLimitMaker           = "LIMIT_MAKER"
)

type OrderStatus string

const (
	OrderStatusCreated         OrderStatus = "Created"
	OrderStatusRejected                    = "Rejected"
	OrderStatusNew                         = "New"
	OrderStatusPartiallyFilled             = "PartiallyFilled"
	OrderStatusFilled                      = "Filled"
	OrderStatusCancelled                   = "Cancelled"
	OrderStatusPendingCancel               = "PendingCancel"
)

type OrderStatusSpot string

const (
	OrderStatusSpotNew             OrderStatusSpot = "NEW"
	OrderStatusSpotPartiallyFilled                 = "PARTIALLY_FILLED"
	OrderStatusSpotFilled                          = "FILLED"
	OrderStatusSpotCanceled                        = "CANCELED"
	OrderStatusSpotPendingCancel                   = "PENDING_CANCEL"
	OrderStatusSpotPendingNew                      = "PENDING_NEW"
	OrderStatusSpotRejected                        = "REJECTED"
)

type TimeInForce string

const (
	TimeInForceGoodTillCancel    TimeInForce = "GoodTillCancel"
	TimeInForceImmediateOrCancel             = "ImmediateOrCancel"
	TimeInForceFillOrKill                    = "FillOrKill"
	TimeInForcePostOnly                      = "PostOnly"
)

type TimeInForceSpot string

const (
	TimeInForceSpotGTC TimeInForceSpot = "GTC"
	TimeInForceSpotFOK                 = "FOK"
	TimeInForceSpotIOC                 = "IOC"
)

type Interval string

const (
	Interval1   Interval = "1"
	Interval3            = "3"
	Interval5            = "5"
	Interval15           = "15"
	Interval30           = "30"
	Interval60           = "60"
	Interval120          = "120"
	Interval240          = "240"
	Interval360          = "360"
	Interval720          = "720"
	IntervalD            = "D"
	IntervalW            = "W"
	IntervalM            = "M"
)

type SpotInterval string

const (
	SpotInterval1m  SpotInterval = "1m"
	SpotInterval3m               = "3m"
	SpotInterval5m               = "5m"
	SpotInterval15m              = "15m"
	SpotInterval30m              = "30m"
	SpotInterval1h               = "1h"
	SpotInterval2h               = "2h"
	SpotInterval4h               = "4h"
	SpotInterval6h               = "6h"
	SpotInterval12h              = "12h"
	SpotInterval1d               = "1d"
	SpotInterval1w               = "1w"
	SpotInterval1M               = "1M"
)

type TickDirection string

const (
	TickDirectionPlusTick      TickDirection = "PlusTick"
	TickDirectionZeroPlusTick                = "ZeroPlusTick"
	TickDirectionMinusTick                   = "MinusTick"
	TickDirectionZeroMinusTick               = "ZeroMinusTick"
)

type Period string

const (
	Period5min  Period = "5min"
	Period15min        = "15min"
	Period30min        = "30min"
	Period1h           = "1h"
	Period4h           = "4h"
	Period1d           = "1d"
)

type TpSlMode string

const (
	TpSlModeFull    TpSlMode = "Full"
	TpSlModePartial          = "Partial"
)

type ExecType string

const (
	ExecTypeTrade     ExecType = "Trade"
	ExecTypeAdlTrade           = "AdlTrade"
	ExecTypeFunding            = "Funding"
	ExecTypeBustTrade          = "BustTrade"
)

func MinimumVolumeUSDT(symbol SymbolUSDT) float64 {
	switch symbol {
	case SymbolUSDTBTC:
		return 0.001
	case SymbolUSDTETH:
		return 0.01
	case SymbolUSDTBCH:
		return 0.01
	case SymbolUSDTLTC:
		return 0.1
	case SymbolUSDTLINK:
		return 0.1
	case SymbolUSDTXTZ:
		return 0.1
	case SymbolUSDTDOT:
		return 0.1
	case SymbolUSDTUNI:
		return 0.1
	case SymbolUSDTADA:
		return 1
	default:
		panic("nothing")
	}
}

type Direction string

const (
	DirectionPrev Direction = "prev"
	DirectionNext           = "next"
)

type TopicName string

const (
	TopicDepth      TopicName = "depth"
	TopicKline                = "kline"
	TopicTrade                = "trade"
	TopicBookTicker           = "bookTicker"
	TopicRealtimes            = "realtimes"
)

type TopicEvent string

const (
	TopicEventSub    TopicEvent = "sub"
	TopicEventCancel            = "cancel"
)
