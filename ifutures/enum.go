// [Enums Definitions] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-enums
package ifutures

// [Side (side)]: https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#side-side
type Side string

const (
	Buy  Side = "Buy"
	Sell      = "Sell"
)

// [Symbol (symbol)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#symbol-symbol
// Symbol

// [Currency (currency/coin)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#currency-currency-coin
type Coin string

const (
	BTC  Coin = "Buy"
	ETH       = "ETH"
	EOS       = "EOS"
	XRP       = "XRP"
	USDT      = "USDT"
)

// [Contract Type (contract_type)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#contract-type-contract_type
type ContractType string

const (
	InversePerpetual ContractType = "InversePerpetual"
	LinearPerpetual               = "LinearPerpetual"
	InverseFutures                = "InverseFutures"
)

// [Contract Status (status)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#contract-status-status
type ContractStatus string

const (
	Trading  ContractStatus = "Trading"
	Settling                = "Settling"
	Closed                  = "Closed"
)

// [Wallet fund type (wallet_fund_type/type)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#wallet-fund-type-wallet_fund_type-type
type WalletFund string

const (
	Deposit               WalletFund = "Deposit"
	Withdraw                         = "Withdraw"
	RealisedPNL                      = "RealisedPNL"
	Commission                       = "Commission"
	Refund                           = "Refund"
	Prize                            = "Prize"
	ExchangeOrderWithdraw            = "ExchangeOrderWithdraw"
	ExchangeOrderDeposit             = "ExchangeOrderDeposit"
)

// [Withdraw status (status)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#withdraw-status-status
type Withdraw string

const (
	ToBeConfirmed Withdraw = "ToBeConfirmed"
	UnderReview            = "UnderReview"
	Pending                = "Pending"
	Success                = "Success"
	CancelByUser           = "CancelByUser"
	Reject                 = "Reject"
	Expire                 = "Expire"
)

// [Order type (order_type)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#order-type-order_type
type OrderType string

const (
	Limit  OrderType = "Limit"
	Market           = "Market"
)

// [Quantity (qty)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#quantity-qty
type Qty uint64

// [Price (price)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#price-price
type Price float64

// [Time in force (time_in_force)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#time-in-force-time_in_force
type TimeInForce string

const (
	GoodTillCancel    TimeInForce = "GoodTillCancel"
	ImmediateOrCancel             = "ImmediateOrCancel"
	FillOrKill                    = "FillOrKill"
	PostOnly                      = "FillOrKill"
)

// [Trigger price type (trigger_by)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#trigger-price-type-trigger_by
type TriggerPrice string

const (
	LastPrice  TriggerPrice = "LastPrice"
	IndexPrice              = "IndexPrice"
	MarkPrice               = "MarkPrice"
)

// [Order (order)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#order-order
// This is used for sorting orders/trades in a specified direction.
type Order string

const (
	Desc Order = "Desc"
	Asc        = "Asc"
)

// [Order status (order_status/stop_order_status)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#order-status-order_status-stop_order_status
// Created - order has been accepted by the system but not yet put through the matching engine
// New - order has been placed successfully
// PendingCancel - matching engine has received the cancelation request but it may not be canceled successfully
// Only for conditional orders:
// Untriggered - order yet to be triggered
// Deactivated - order has been canceled by the user before being triggered
// Triggered - order has been triggered by last traded price
// Active - order has been triggered and the new active order has been successfully placed. Is the final state of a successful conditional order
type OrderStatus string

const (
	Created         OrderStatus = "Created"
	New                         = "New"
	Rejected                    = "Rejected"
	PartiallyFilled             = "PartiallyFilled"
	Filled                      = "Filled"
	PendingCancel               = "PendingCancel"
	Cancelled                   = "Cancelled"
	Untriggered                 = "Untriggered"
	Deactivated                 = "Deactivated"
	Triggered                   = "Triggered"
	Active                      = "Active"
)

// [Cancel type (cancel_type)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#cancel-type-cancel_type
// CancelByPrepareLiq, CancelAllBeforeLiq - canceled due to liquidation
// CancelByPrepareAdl, CancelAllBeforeAdl - canceled due to ADL
// CancelByTpSlTsClear - TP/SL order canceled successfully
// CancelByPzSideCh - order has been canceled after TP/SL is triggered
type CancelType string

const (
	CancelByUser        CancelType = "CancelByUser"
	CancelByReduceOnly             = "CancelByReduceOnly"
	CancelByPrepareLiq             = "CancelByPrepareLiq"
	CancelAllBeforeLiq             = "CancelAllBeforeLiq"
	CancelByPrepareAdl             = "CancelByPrepareAdl"
	CancelAllBeforeAdl             = "CancelAllBeforeAdl"
	CancelByAdmin                  = "CancelByAdmin"
	CancelByTpSlTsClear            = "CancelByTpSlTsClear"
	CancelByPzSideCh               = "CancelByPzSideCh"
)

// [Create type (create_type)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#create-type-create_type
// CreateByLiq - Created by partial liquidation
// CreateByAdl_PassThrough - Created by ADL
// CreateByTakeOver_PassThrough - Created by liquidation takeover
type CreateType string

const (
	CreateByUser                CreateType = "CreateByUser"
	CreateByClosing                        = "CreateByClosing"
	CreateByAdminClosing                   = "CreateByAdminClosing"
	CreateByStopOrder                      = "CreateByStopOrder"
	CreateByTakeProfit                     = "CreateByTakeProfit"
	CreateByStopLoss                       = "CreateByStopLoss"
	CreateByPartialTakeProfit              = "CreateByPartialTakeProfit"
	CreateByPartialStopLoss                = "CreateByPartialStopLoss"
	CreateByTrailingStop                   = "CreateByTrailingStop"
	CreateByLiq                            = "CreateByLiq"
	CreateByAdlPassThrough                 = "CreateByAdl_PassThrough"
	CreateByTakeOverPassThrough            = "CreateByTakeOver_PassThrough"
)

// [Exec type (exec_type)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#exec-type-exec_type
type ExecType string

const (
	Trade     ExecType = "Trade"
	AdlTrade           = "AdlTrade"
	Funding            = "Funding"
	BustTrade          = "BustTrade"
	Settle             = "Settle"
)

// [Liquidity type (last_liquidity_ind)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#liquidity-type-last_liquidity_ind
// AddedLiquidity - liquidity maker
// RemovedLiquidity - liquidity Taker
type Liquidity string

const (
	LiquidityAdded   Liquidity = "AddedLiquidity"
	LiquidityRemoved           = "RemovedLiquidity"
)

// [Tick direction type (tick_direction)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#tick-direction-type-tick_direction
// It indicates price fluctuation relative to the last trade.
// PlusTick - price rise
// ZeroPlusTick - trade occurs at the same price as the previous trade, which occurred at a price higher than that for the trade preceding it
// MinusTick - price drop
// ZeroMinusTick - trade occurs at the same price as the previous trade, which occurred at a price lower than that for the trade preceding it
type TickDirection string

const (
	TickPlus      TickDirection = "TickPlus"
	TickZeroPlus                = "TickZeroPlus"
	TickMinus                   = "TickMinus"
	TickZeroMinus               = "TickZeroMinus"
)

// [TP/SL mode (tp_sl_mode)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#tp-sl-mode-tp_sl_mode
// Take profit/stop loss mode
// Full - Full take profit/stop loss mode (a single TP order and a single SL order can be placed, covering the entire position)
// Partial - Partial take profit/stop loss mode (multiple TP and SL orders can be placed, covering portions of the position)
type TpSlMode string

const (
	TpSlFull    TpSlMode = "Full"
	TpSlPartial          = "Partial"
)

// [Kline interval (interval)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#kline-interval-interval
// 1 - 1 minute
// 3 - 3 minutes
// 5 - 5 minutes
// 15 - 15 minutes
// 30 - 30 minutes
// 60 - 1 hour
// 120 - 2 hours
// 240 - 4 hours
// 360 - 6 hours
// 720 - 12 hours
// D - 1 day
// W - 1 week
// M - 1 month
type KlineInterval string

// [Stop order type (stop_order_type)] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#stop-order-type-stop_order_type
type StopOrder string

const (
	TakeProfit   StopOrder = "TakeProfit"
	StopLoss               = "StopLoss"
	TrailingStop           = "TrailingStop"
	Stop                   = "Stop"
)
