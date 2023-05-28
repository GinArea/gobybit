// Enums Definitions (https://bybit-exchange.github.io/docs/v5/enum)
package bybitv5

// accountType (https://bybit-exchange.github.io/docs/v5/enum#accounttype)
type AccountType string

const (
	AccountContract   AccountType = "CONTRACT"
	AccountSpot       AccountType = "SPOT"
	AccountInvestment AccountType = "INVESTMENT"
	AccountOption     AccountType = "OPTION"
	AccountUnified    AccountType = "UNIFIED"
	AccountFund       AccountType = "FUND"
)

// category (https://bybit-exchange.github.io/docs/v5/enum#category)
type Category string

const (
	Spot    Category = "spot"
	Linear  Category = "linear"
	Inverse Category = "inverse"
	Option  Category = "option"
)

// interval (https://bybit-exchange.github.io/docs/v5/enum#interval)
//
//	1   - 1 minute
//	3   - 3 minutes
//	5   - 5 minutes
//	15  - 15 minutes
//	30  - 30 minutes
//	60  - 1 hour
//	120 - 2 hours
//	240 - 4 hours
//	360 - 6 hours
//	720 - 12 hours
//	D   - 1 day
//	W   - 1 week
//	M   - 1 month
type Interval string

const (
	Interval1m  Interval = "1"
	Interval3m  Interval = "3"
	Interval5m  Interval = "5"
	Interval15m Interval = "15"
	Interval30m Interval = "30"
	Interval1h  Interval = "60"
	Interval2h  Interval = "120"
	Interval4h  Interval = "240"
	Interval6h  Interval = "360"
	Interval12h Interval = "720"
	Interval1d  Interval = "D"
	Interval1w  Interval = "W"
	Interval1M  Interval = "M"
)

type IntervalTime string

const (
	IntervalTime5min  IntervalTime = "5min"
	IntervalTime15min IntervalTime = "15min"
	IntervalTime30min IntervalTime = "30min"
	IntervalTime1h    IntervalTime = "1h"
	IntervalTime4h    IntervalTime = "4h"
	IntervalTime1d    IntervalTime = "1d"
)

// period (https://bybit-exchange.github.io/docs/v5/enum#period)
type Period int

const (
	Period7days   Period = 7
	Period14days  Period = 14
	Period21days  Period = 21
	Period30days  Period = 30
	Period60days  Period = 60
	Period90days  Period = 90
	Period180days Period = 180
	Period270days Period = 270
)

// contractType (https://bybit-exchange.github.io/docs/v5/enum#contracttype)
type ContractType string

const (
	InversePerpetual ContractType = "InversePerpetual"
	LinearPerpetual  ContractType = "LinearPerpetual"
	LinearFutures    ContractType = "LinearFutures"
	InverseFutures   ContractType = "InverseFutures"
)

// status (https://bybit-exchange.github.io/docs/v5/enum#status)
type Status string

const (
	StatusPreLaunch  Status = "PreLaunch"
	StatusTrading    Status = "Trading"
	StatusSettling   Status = "Settling"
	StatusDelivering Status = "Delivering"
	StatusClosed     Status = "Closed"
)

// triggerBy (https://bybit-exchange.github.io/docs/v5/enum#triggerby)
type TriggerBy string

const (
	TriggerByLastPrice  TriggerBy = "LastPrice"
	TriggerByIndexPrice TriggerBy = "IndexPrice"
	TriggerByMarkPrice  TriggerBy = "MarkPrice"
)

// timeInForce (https://bybit-exchange.github.io/docs/v5/enum#timeinforce)
type TimeInForce string

const (
	GoodTillCancel    TimeInForce = "GTC"
	ImmediateOrCancel TimeInForce = "IOC"
	FillOrKill        TimeInForce = "FOK"
)

// smpType (https://bybit-exchange.github.io/docs/v5/enum#smptype)
type SmpType string

const (
	SmpTypeNone SmpType = "None"
	CancelMaker SmpType = "CancelMaker"
	CancelTaker SmpType = "CancelTaker"
	CancelBoth  SmpType = "CancelBoth"
)

// positionIdx (https://bybit-exchange.github.io/docs/v5/enum#positionidx)
type PositionIdx int

const (
	OneWayMode          PositionIdx = 0
	BuySideOfHedgeMode  PositionIdx = 1
	SellSideOfHedgeMode PositionIdx = 2
)

// orderStatus https://bybit-exchange.github.io/docs/v5/enum#orderstatus
type OrderStatus string

const (
	OrderStatusCreated                 OrderStatus = "Created"
	OrderStatusNew                     OrderStatus = "New"
	OrderStatusRejected                OrderStatus = "Rejected"
	OrderStatusPartiallyFilled         OrderStatus = "PartiallyFilled"
	OrderStatusPartiallyFilledCanceled OrderStatus = "PartiallyFilledCanceled"
	OrderStatusFilled                  OrderStatus = "Filled"
	OrderStatusCancelled               OrderStatus = "Cancelled"
	OrderStatusUntriggered             OrderStatus = "Untriggered"
	OrderStatusTriggered               OrderStatus = "Triggered"
	OrderStatusDeactivated             OrderStatus = "Deactivated"
	OrderStatusActive                  OrderStatus = "Active"
)

// cancelType https://bybit-exchange.github.io/docs/v5/enum#canceltype
type CancelType string

const (
	CancelByUser                  CancelType = "CancelByUser"
	CancelByReduceOnly            CancelType = "CancelByReduceOnly"
	CancelByPrepareLiq            CancelType = "CancelByPrepareLiq"
	CancelByPrepareAdl            CancelType = "CancelByPrepareAdl"
	CancelByAdmin                 CancelType = "CancelByAdmin"
	CancelByTpSlTsClear           CancelType = "CancelByTpSlTsClear"
	CancelByPzSideCh              CancelType = "CancelByPzSideCh"
	CancelBySmp                   CancelType = "CancelBySmp"
	CancelAllBeforeLiq            CancelType = "CancelAllBeforeLiq"
	CancelAllBeforeAdl            CancelType = "CancelAllBeforeAdl"
	CancelBySettle                CancelType = "CancelBySettle"
	CancelByCannotAffordOrderCost CancelType = "CancelByCannotAffordOrderCost"
	CancelByPmTrialMmOverEquity   CancelType = "CancelByPmTrialMmOverEquity"
	CancelByAccountBlocking       CancelType = "CancelByAccountBlocking"
	CancelByDelivery              CancelType = "CancelByDelivery"
	CancelByMmpTriggered          CancelType = "CancelByMmpTriggered"
	CancelByCrossSelfMuch         CancelType = "CancelByCrossSelfMuch"
	CancelByCrossReachMaxTradeNum CancelType = "CancelByCrossReachMaxTradeNum"
	CancelByDCP                   CancelType = "CancelByDCP"
)

// rejectReason https://bybit-exchange.github.io/docs/v5/enum#rejectreason
type RejectReason string

const (
	RejectNoError                      RejectReason = "EC_NoError"
	RejectOthers                       RejectReason = "EC_Others"
	RejectUnknownMessageType           RejectReason = "EC_UnknownMessageType"
	RejectMissingClOrdID               RejectReason = "EC_MissingClOrdID"
	RejectMissingOrigClOrdID           RejectReason = "EC_MissingOrigClOrdID"
	RejectClOrdIDOrigClOrdIDAreTheSame RejectReason = "EC_ClOrdIDOrigClOrdIDAreTheSame"
	RejectDuplicatedClOrdID            RejectReason = "EC_DuplicatedClOrdID"
	RejectOrigClOrdIDDoesNotExist      RejectReason = "EC_OrigClOrdIDDoesNotExist"
	RejectTooLateToCancel              RejectReason = "EC_TooLateToCancel"
	RejectUnknownOrderType             RejectReason = "EC_UnknownOrderType"
	RejectUnknownSide                  RejectReason = "EC_UnknownSide"
	RejectUnknownTimeInForce           RejectReason = "EC_UnknownTimeInForce"
	RejectWronglyRouted                RejectReason = "EC_WronglyRouted"
	RejectMarketOrderPriceIsNotZero    RejectReason = "EC_MarketOrderPriceIsNotZero"
	RejectLimitOrderInvalidPrice       RejectReason = "EC_LimitOrderInvalidPrice"
	RejectNoEnoughQtyToFill            RejectReason = "EC_NoEnoughQtyToFill"
	RejectNoImmediateQtyToFill         RejectReason = "EC_NoImmediateQtyToFill"
	RejectPerCancelRequest             RejectReason = "EC_PerCancelRequest"
	RejectMarketOrderCannotBePostOnly  RejectReason = "EC_MarketOrderCannotBePostOnly"
	RejectPostOnlyWillTakeLiquidity    RejectReason = "EC_PostOnlyWillTakeLiquidity"
	RejectCancelReplaceOrder           RejectReason = "EC_CancelReplaceOrder"
	RejectInvalidSymbolStatus          RejectReason = "EC_InvalidSymbolStatus"
)

// stopOrderType https://bybit-exchange.github.io/docs/v5/enum#stopordertype
type StopOrderType string

const (
	StopOrderTakeProfit        StopOrderType = "TakeProfit"
	StopOrderStopLoss          StopOrderType = "StopLoss"
	StopOrderTrailingStop      StopOrderType = "TrailingStop"
	StopOrderStop              StopOrderType = "Stop"
	StopOrderPartialTakeProfit StopOrderType = "PartialTakeProfit"
	StopOrderPartialStopLoss   StopOrderType = "PartialStopLoss"
	StopOrderTpslOrder         StopOrderType = "TpslOrder"
)

// positionStatus https://bybit-exchange.github.io/docs/v5/enum#positionstatus
type PositionStatus string

const (
	PositionStatusNormal PositionStatus = "Normal"
	PositionStatusLiq    PositionStatus = "Liq"
	PositionStatusAdl    PositionStatus = "Adl"
)

// tickDirection https://bybit-exchange.github.io/docs/v5/enum#tickdirection
type TickDirection string

const (
	TickPlus      TickDirection = "PlusTick"
	TickZeroPlus  TickDirection = "ZeroPlusTick"
	TickMinus     TickDirection = "MinusTick"
	TickZeroMinus TickDirection = "ZeroMinusTick"
)

type Side string

const (
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

type OrderType string

const (
	Limit  OrderType = "Limit"
	Market OrderType = "Market"
)

type TpSlMode string

const (
	TpSlModeFull    TpSlMode = "Full"
	TpSlModePartial TpSlMode = "Partial"
)

type TradeMode int

const (
	TradeCrossMargin    TradeMode = 0
	TradeIsolatedMargin TradeMode = 1
)
