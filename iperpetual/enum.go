// Enums Definitions (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-enums)
package iperpetual

// Side (side) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#side-side)
type Side string

const (
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

// Symbol (symbol) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#symbol-symbol)
type Symbol string

// Currency (currency/coin) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#currency-currency-coin)
type Currency string

const (
	BTC  Currency = "BTC"
	ETH  Currency = "ETH"
	EOS  Currency = "EOS"
	XRP  Currency = "XRP"
	DOT  Currency = "DOT"
	USDT Currency = "USDT"
)

// Contract Type (contract_type) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#contract-type-contract_type)
type ContractType string

const (
	InversePerpetual ContractType = "InversePerpetual"
	LinearPerpetual  ContractType = "LinearPerpetual"
	InverseFutures   ContractType = "InverseFutures"
)

// Contract Status (status) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#contract-status-status)
type ContractStatus string

const (
	Trading  ContractStatus = "Trading"
	Settling ContractStatus = "Settling"
	Closed   ContractStatus = "Closed"
)

// Wallet fund type (wallet_fund_type/type) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#wallet-fund-type-wallet_fund_type-type)
//   ReturnServiceCash - Refund of handling fee bonus
//   Insurance - Insurance account transfer
//   SubMember - Parent-child account transfer
//   Coupon - Coupon interest
//   AccountTransfer - Account transfer
//   CashBack - Cash back
type WalletFund string

const (
	FundDeposit               WalletFund = "Deposit"
	FundWithdraw              WalletFund = "Withdraw"
	FundRealisedPNL           WalletFund = "RealisedPNL"
	FundCommission            WalletFund = "Commission"
	FundRefund                WalletFund = "Refund"
	FundPrize                 WalletFund = "Prize"
	FundExchangeOrderWithdraw WalletFund = "ExchangeOrderWithdraw"
	FundExchangeOrderDeposit  WalletFund = "ExchangeOrderDeposit"
	FundReturnServiceCash     WalletFund = "ReturnServiceCash"
	FundInsurance             WalletFund = "Insurance"
	FundSubMember             WalletFund = "SubMember"
	FundCoupon                WalletFund = "Coupon"
	FundAccountTransfer       WalletFund = "AccountTransfer"
	FundCashBack              WalletFund = "CashBack"
)

// Withdraw status (status) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#withdraw-status-status)
type Withdraw string

const (
	WithdrawToBeConfirmed Withdraw = "ToBeConfirmed"
	WithdrawUnderReview   Withdraw = "UnderReview"
	WithdrawPending       Withdraw = "Pending"
	WithdrawSuccess       Withdraw = "Success"
	WithdrawCancelByUser  Withdraw = "CancelByUser"
	WithdrawReject        Withdraw = "Reject"
	WithdrawExpire        Withdraw = "Expire"
)

// Order type (order_type) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#order-type-order_type)
type OrderType string

const (
	Limit  OrderType = "Limit"
	Market OrderType = "Market"
)

// Quantity (qty) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#quantity-qty)
type Qty uint64

// Price (price) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#price-price)
type Price float64

// Time in force (time_in_force) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#time-in-force-time_in_force)
type TimeInForce string

const (
	GoodTillCancel    TimeInForce = "GoodTillCancel"
	ImmediateOrCancel TimeInForce = "ImmediateOrCancel"
	FillOrKill        TimeInForce = "FillOrKill"
	PostOnly          TimeInForce = "FillOrKill"
)

// Trigger price type (trigger_by) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#trigger-price-type-trigger_by)
type TriggerPrice string

const (
	LastPrice  TriggerPrice = "LastPrice"
	IndexPrice TriggerPrice = "IndexPrice"
	MarkPrice  TriggerPrice = "MarkPrice"
)

// Order (order) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#order-order)
//
// This is used for sorting orders/trades in a specified direction.
type SortOrder string

const (
	Desc SortOrder = "desc"
	Asc  SortOrder = "asc"
)

// Order status (order_status/stop_order_status) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#order-status-order_status-stop_order_status)
//   Created - order has been accepted by the system but not yet put through the matching engine
//   New - order has been placed successfully
//   PendingCancel - matching engine has received the cancelation request but it may not be canceled successfully
//   Only for conditional orders:
//   Untriggered - order yet to be triggered
//   Deactivated - order has been canceled by the user before being triggered
//   Triggered - order has been triggered by last traded price
//   Active - order has been triggered and the new active order has been successfully placed. Is the final state of a successful conditional order
type OrderStatus string

const (
	Created         OrderStatus = "Created"
	New             OrderStatus = "New"
	Rejected        OrderStatus = "Rejected"
	PartiallyFilled OrderStatus = "PartiallyFilled"
	Filled          OrderStatus = "Filled"
	PendingCancel   OrderStatus = "PendingCancel"
	Cancelled       OrderStatus = "Cancelled"
	Untriggered     OrderStatus = "Untriggered"
	Deactivated     OrderStatus = "Deactivated"
	Triggered       OrderStatus = "Triggered"
	Active          OrderStatus = "Active"
)

// Cancel type (cancel_type) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#cancel-type-cancel_type)
//   CancelByPrepareLiq, CancelAllBeforeLiq - canceled due to liquidation
//   CancelByPrepareAdl, CancelAllBeforeAdl - canceled due to ADL
//   CancelByTpSlTsClear - TP/SL order canceled successfully
//   CancelByPzSideCh - order has been canceled after TP/SL is triggered
type CancelType string

const (
	CancelByUser        CancelType = "CancelByUser"
	CancelByReduceOnly  CancelType = "CancelByReduceOnly"
	CancelByPrepareLiq  CancelType = "CancelByPrepareLiq"
	CancelAllBeforeLiq  CancelType = "CancelAllBeforeLiq"
	CancelByPrepareAdl  CancelType = "CancelByPrepareAdl"
	CancelAllBeforeAdl  CancelType = "CancelAllBeforeAdl"
	CancelByAdmin       CancelType = "CancelByAdmin"
	CancelByTpSlTsClear CancelType = "CancelByTpSlTsClear"
	CancelByPzSideCh    CancelType = "CancelByPzSideCh"
)

// Create type (create_type) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#create-type-create_type)
//   CreateByLiq - Created by partial liquidation
//   CreateByAdl_PassThrough - Created by ADL
//   CreateByTakeOver_PassThrough - Created by liquidation takeover
type CreateType string

const (
	CreateByUser                CreateType = "CreateByUser"
	CreateByClosing             CreateType = "CreateByClosing"
	CreateByAdminClosing        CreateType = "CreateByAdminClosing"
	CreateByStopOrder           CreateType = "CreateByStopOrder"
	CreateByTakeProfit          CreateType = "CreateByTakeProfit"
	CreateByStopLoss            CreateType = "CreateByStopLoss"
	CreateByPartialTakeProfit   CreateType = "CreateByPartialTakeProfit"
	CreateByPartialStopLoss     CreateType = "CreateByPartialStopLoss"
	CreateByTrailingStop        CreateType = "CreateByTrailingStop"
	CreateByLiq                 CreateType = "CreateByLiq"
	CreateByAdlPassThrough      CreateType = "CreateByAdl_PassThrough"
	CreateByTakeOverPassThrough CreateType = "CreateByTakeOver_PassThrough"
)

// Exec type (exec_type) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#exec-type-exec_type)
type ExecType string

const (
	Trade     ExecType = "Trade"
	AdlTrade  ExecType = "AdlTrade"
	Funding   ExecType = "Funding"
	BustTrade ExecType = "BustTrade"
	Settle    ExecType = "Settle"
)

// Liquidity type (last_liquidity_ind) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#liquidity-type-last_liquidity_ind)
//   AddedLiquidity - liquidity maker
//   RemovedLiquidity - liquidity Taker
type Liquidity string

const (
	LiquidityAdded   Liquidity = "AddedLiquidity"
	LiquidityRemoved Liquidity = "RemovedLiquidity"
)

// Tick direction type (tick_direction) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#tick-direction-type-tick_direction)
//
// It indicates price fluctuation relative to the last trade.
//   PlusTick - price rise
//   ZeroPlusTick - trade occurs at the same price as the previous trade, which occurred at a price higher than that for the trade preceding it
//   MinusTick - price drop
//   ZeroMinusTick - trade occurs at the same price as the previous trade, which occurred at a price lower than that for the trade preceding it
type TickDirection string

const (
	TickPlus      TickDirection = "TickPlus"
	TickZeroPlus  TickDirection = "TickZeroPlus"
	TickMinus     TickDirection = "TickMinus"
	TickZeroMinus TickDirection = "TickZeroMinus"
)

// TP/SL mode (tp_sl_mode) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#tp-sl-mode-tp_sl_mode)
//
// Take profit/stop loss mode
//   Full - Full take profit/stop loss mode (a single TP order and a single SL order can be placed, covering the entire position)
//   Partial - Partial take profit/stop loss mode (multiple TP and SL orders can be placed, covering portions of the position)
type TpSlMode string

const (
	TpSlFull    TpSlMode = "Full"
	TpSlPartial TpSlMode = "Partial"
)

// Kline interval (interval) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#kline-interval-interval)
//   1 - 1 minute
//   3 - 3 minutes
//   5 - 5 minutes
//   15 - 15 minutes
//   30 - 30 minutes
//   60 - 1 hour
//   120 - 2 hours
//   240 - 4 hours
//   360 - 6 hours
//   720 - 12 hours
//   D - 1 day
//   W - 1 week
//   M - 1 month
type KlineInterval string

const (
	Interval1m  KlineInterval = "1"
	Interval3m  KlineInterval = "3"
	Interval5m  KlineInterval = "5"
	Interval15m KlineInterval = "15"
	Interval30m KlineInterval = "30"
	Interval1h  KlineInterval = "60"
	Interval2h  KlineInterval = "120"
	Interval4h  KlineInterval = "240"
	Interval6h  KlineInterval = "360"
	Interval12h KlineInterval = "720"
	Interval1d  KlineInterval = "D"
	Interval1w  KlineInterval = "W"
	Interval1M  KlineInterval = "M"
)

// Stop order type (stop_order_type) (https://bybit-exchange.github.io/docs/futuresV2/inverse/#stop-order-type-stop_order_type)
type StopOrder string

const (
	TakeProfit   StopOrder = "TakeProfit"
	StopLoss     StopOrder = "StopLoss"
	TrailingStop StopOrder = "TrailingStop"
	Stop         StopOrder = "Stop"
)
