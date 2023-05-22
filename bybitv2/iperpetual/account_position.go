// Position (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-position)
package iperpetual

import (
	"errors"
	"time"

	"github.com/ginarea/gobybit/bybitv2/transport"
)

// My Position (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-myposition)
type GetPosition struct {
	Symbol *string `param:"symbol"`
}

func (o GetPosition) Do(client *Client) ([]PositionItem, error) {
	if o.Symbol == nil {
		return Get[[]PositionItem](client, "position/list", o)
	}
	r, err := Get[PositionItem](client, "position/list", o)
	return []PositionItem{r}, err
}

type PositionBase struct {
	ID                  int               `json:"id"`
	UserID              int               `json:"user_id"`
	RiskID              int               `json:"risk_id"`
	Symbol              string            `json:"symbol"`
	Side                Side              `json:"side"`
	Size                int               `json:"size"`
	PositionValue       transport.Float64 `json:"position_value"`
	EntryPrice          transport.Float64 `json:"entry_price"`
	IsIsolated          bool              `json:"is_isolated"`
	AutoAddMargin       int               `json:"auto_add_margin"`
	Leverage            transport.Float64 `json:"leverage"`
	EffectiveLeverage   transport.Float64 `json:"effective_leverage"`
	PositionMargin      transport.Float64 `json:"position_margin"`
	LiqPrice            transport.Float64 `json:"liq_price"`
	BustPrice           transport.Float64 `json:"bust_price"`
	OccClosingFee       transport.Float64 `json:"occ_closing_fee"`
	OccFundingFee       transport.Float64 `json:"occ_funding_fee"`
	TakeProfit          transport.Float64 `json:"take_profit"`
	StopLoss            transport.Float64 `json:"stop_loss"`
	TrailingStop        transport.Float64 `json:"trailing_stop"`
	PositionStatus      string            `json:"position_status"`
	DeleverageIndicator int               `json:"deleverage_indicator"`
	OcCalcData          string            `json:"oc_calc_data"`
	OrderMargin         transport.Float64 `json:"order_margin"`
	WalletBalance       transport.Float64 `json:"wallet_balance"`
	RealisedPnl         transport.Float64 `json:"realised_pnl"`
	CumRealisedPnl      transport.Float64 `json:"cum_realised_pnl"`
	CrossSeq            int               `json:"cross_seq"`
	PositionSeq         int               `json:"position_seq"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

type PositionData struct {
	PositionBase
	PositionIdx   PositionIdx `json:"position_idx"`
	Mode          int         `json:"mode"`
	UnrealisedPnl float64     `json:"unrealised_pnl"`
	TpSlMode      TpSlMode    `json:"tp_sl_mode"`
}

type PositionItem struct {
	Data    PositionData `json:"data"`
	IsValid bool         `json:"is_valid"`
}

func (o *Client) GetPosition(symbol *string) ([]PositionItem, error) {
	return GetPosition{Symbol: symbol}.Do(o)
}

func (o *Client) GetOnePosition(symbol string) (i PositionItem, err error) {
	ret, err := GetPosition{Symbol: &symbol}.Do(o)
	if err == nil {
		if len(ret) == 1 {
			i = ret[0]
		} else {
			err = errors.New("position result len != 1")
		}
	}
	return
}

func (o *Client) GetAllPositions() ([]PositionItem, error) {
	return GetPosition{}.Do(o)
}

// Change Margin (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-changemargin)
type ChangeMargin struct {
	Symbol string `param:"symbol"`
	Margin string `param:"margin"`
}

func (o ChangeMargin) Do(client *Client) (float64, error) {
	return Post[float64](client, "position/change-position-margin", o)
}

func (o *Client) ChangeMargin(v ChangeMargin) (float64, error) {
	return v.Do(o)
}

// Set Trading-Stop (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-tradingstop)
type SetTradingStop struct {
	Symbol            string        `param:"symbol"`
	TakeProfit        *int          `param:"take_profit"`
	StopLoss          *int          `param:"stop_loss"`
	TrailingStop      *int          `param:"trailing_stop"`
	TpTrigger         *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger         *TriggerPrice `param:"sl_trigger_by"`
	NewTrailingActive *int          `param:"new_trailing_active"`
	SlSize            *int          `param:"sl_size"`
	TpSize            *int          `param:"tp_size"`
}

func (o SetTradingStop) Do(client *Client) (SetTradingStopResult, error) {
	return Post[SetTradingStopResult](client, "position/trading-stop", o)
}

type SetTradingStopExt struct {
	TrailingActive string `json:"trailing_active"`
	SlTrigger      string `json:"sl_trigger_by"`
	TpTrigger      string `json:"tp_trigger_by"`
	V              int    `json:"v"`
	Mm             int    `json:"mm"`
}

type SetTradingStopResult struct {
	PositionBase
	CumCommission int               `json:"cum_commission"`
	ExtFields     SetTradingStopExt `json:"ext_fields"`
}

func (o *Client) SetTradingStop(v SetTradingStop) (SetTradingStopResult, error) {
	return v.Do(o)
}

// Set Leverage (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-setleverage)
type SetLeverage struct {
	Symbol       string `param:"symbol"`
	Leverage     int    `param:"leverage"`
	LeverageOnly *bool  `param:"leverage_only"`
}

func (o SetLeverage) Do(client *Client) (int, error) {
	return Post[int](client, "position/leverage/save", o)
}

func (o *Client) SetLeverage(v SetLeverage) (int, error) {
	return v.Do(o)
}

// Full/Partial Position TP/SL Switch (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-switchmode)
//
// Switch mode between Full or Partial
type TpSlModeSwitch struct {
	Symbol   string    `param:"symbol"`
	TpSlMode *TpSlMode `param:"tp_sl_mode"`
}

func (o TpSlModeSwitch) Do(client *Client) (TpSlMode, error) {
	type result struct {
		TpSlMode TpSlMode `json:"tp_sl_mode"`
	}
	r, err := Post[result](client, "tpsl/switch-mode", o)
	return r.TpSlMode, err
}

func (o *Client) TpSlModeSwitch(v TpSlModeSwitch) (TpSlMode, error) {
	return v.Do(o)
}

// Cross/Isolated Margin Switch (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marginswitch)
//
// Switch Cross/Isolated; must set leverage value when switching from Cross to Isolated
type MarginSwitch struct {
	Symbol       string `param:"symbol"`
	IsIsolated   bool   `param:"is_isolated"`
	BuyLeverage  int    `param:"buy_leverage"`
	SellLeverage int    `param:"sell_leverage"`
}

func (o MarginSwitch) Do(client *Client) error {
	_, err := Post[struct{}](client, "position/switch-isolated", o)
	return err
}

func (o *Client) MarginSwitch(v MarginSwitch) error {
	return v.Do(o)
}

// Get User Trade Records (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-usertraderecords)
//
// Get user's trading records. The results are ordered in ascending order (the first item is the oldest)
type GetTradeRecords struct {
	Symbol    string     `param:"symbol"`
	OrderID   *string    `param:"order_id"`
	StartTime *int       `param:"start_time"`
	Page      *int       `param:"page"`
	Limit     *int       `param:"limit"`
	Order     *SortOrder `param:"order"`
}

func (o GetTradeRecords) Do(client *Client) (TradeRecords, error) {
	return Get[TradeRecords](client, "execution/list", o)
}

type TradeRecord struct {
	ClosedSize    int       `json:"closed_size"`
	CrossSeq      int       `json:"cross_seq"`
	ExecFee       string    `json:"exec_fee"`
	ExecID        string    `json:"exec_id"`
	ExecPrice     string    `json:"exec_price"`
	ExecQty       int       `json:"exec_qty"`
	ExecTime      int64     `json:"exec_time"`
	ExecType      ExecType  `json:"exec_type"`
	ExecValue     string    `json:"exec_value"`
	FeeRate       string    `json:"fee_rate"`
	LastLiquidity string    `json:"last_liquidity_ind"`
	LeavesQty     int       `json:"leaves_qty"`
	NthFill       int       `json:"nth_fill"`
	OrderID       string    `json:"order_id"`
	OrderLinkID   string    `json:"order_link_id"`
	OrderPrice    string    `json:"order_price"`
	OrderQty      int       `json:"order_qty"`
	OrderType     OrderType `json:"order_type"`
	Side          Side      `json:"side"`
	Symbol        string    `json:"symbol"`
	UserID        int       `json:"user_id"`
	TradeTime     uint64    `json:"trade_time_ms"`
}

type TradeRecords struct {
	OrderID      string        `json:"order_id"`
	TradeRecords []TradeRecord `json:"trade_list"`
}

func (o *Client) GetTradeRecords(v GetTradeRecords) (TradeRecords, error) {
	return v.Do(o)
}

// Closed Profit and Loss (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-closedprofitandloss)
type ClosedProfitLoss struct {
	Symbol    string    `param:"symbol"`
	StartTime *int      `param:"start_time"`
	EndTime   *int      `param:"end_time"`
	ExecType  *ExecType `param:"exec_type"`
	Page      *int      `param:"page"`
	Limit     *int      `param:"limit"`
}

func (o ClosedProfitLoss) Do(client *Client) (ClosedProfitLossResult, error) {
	return Get[ClosedProfitLossResult](client, "trade/closed-pnl/list", o)
}

type ClosedData struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Symbol        string    `json:"symbol"`
	OrderID       string    `json:"order_id"`
	Side          Side      `json:"side"`
	Qty           float64   `json:"qty"`
	OrderPrice    float64   `json:"order_price"`
	OrderType     OrderType `json:"order_type"`
	ExecType      ExecType  `json:"exec_type"`
	ClosedSize    float64   `json:"closed_size"`
	CumEntryValue float64   `json:"cum_entry_value"`
	AvgEntryPrice float64   `json:"avg_entry_price"`
	CumExitValue  float64   `json:"cum_exit_value"`
	AvgExitPrice  float64   `json:"avg_exit_price"`
	ClosedPnl     float64   `json:"closed_pnl"`
	FillCount     int       `json:"fill_count"`
	Leverage      int       `json:"leverage"`
	CreatedAt     uint64    `json:"created_at"`
}

type ClosedProfitLossResult struct {
	CurrentPage int          `json:"current_page"`
	Data        []ClosedData `json:"data"`
}

func (o *Client) ClosedProfitLoss(v ClosedProfitLoss) (ClosedProfitLossResult, error) {
	return v.Do(o)
}
