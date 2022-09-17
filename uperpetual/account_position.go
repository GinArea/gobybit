// [Position] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-position
package uperpetual

import "github.com/tranquiil/bybit/iperpetual"

// [My Position] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-myposition
type GetPositionAll struct {
}

func (this GetPositionAll) Do(client *Client) ([]PositionItem, bool) {
	return Get[[]PositionItem](client, "position/list", this)
}

type GetPosition struct {
	Symbol iperpetual.Symbol `param:"symbol"`
}

func (this GetPosition) Do(client *Client) ([]PositionData, bool) {
	return Get[[]PositionData](client, "position/list", this)
}

type PositionData struct {
	UserID              int               `json:"user_id"`
	Symbol              iperpetual.Symbol `json:"symbol"`
	Side                Side              `json:"side"`
	Size                int               `json:"size"`
	PositionValue       float64           `json:"position_value"`
	EntryPrice          float64           `json:"entry_price"`
	LiqPrice            float64           `json:"liq_price"`
	BustPrice           float64           `json:"bust_price"`
	Leverage            float64           `json:"leverage"`
	AutoAddMargin       int               `json:"auto_add_margin"`
	IsIsolated          bool              `json:"is_isolated"`
	PositionMargin      float64           `json:"position_margin"`
	OccClosingFee       float64           `json:"occ_closing_fee"`
	RealisedPnl         float64           `json:"realised_pnl"`
	CumRealisedPnl      float64           `json:"cum_realised_pnl"`
	FreeQty             float64           `json:"free_qty"`
	TpSlMode            TpSlMode          `json:"tp_sl_mode"`
	UnrealisedPnl       int               `json:"unrealised_pnl"`
	DeleverageIndicator int               `json:"deleverage_indicator"`
	RiskID              int               `json:"risk_id"`
	StopLoss            float64           `json:"stop_loss"`
	TakeProfit          float64           `json:"take_profit"`
	TrailingStop        float64           `json:"trailing_stop"`
	PositionIdx         PositionIdx       `json:"position_idx"`
	Mode                string            `json:"mode"`
}

type PositionItem struct {
	Data    PositionData `json:"data"`
	IsValid bool         `json:"is_valid"`
}

func (this *Client) GetPositionAll() ([]PositionItem, bool) {
	return GetPositionAll{}.Do(this)
}

func (this *Client) GetPosition(symbol iperpetual.Symbol) ([]PositionData, bool) {
	return GetPosition{Symbol: symbol}.Do(this)
}

// [Set Auto Add Margin] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-setautoaddmargin
type SetAutoAddMargin struct {
	Symbol        iperpetual.Symbol `param:"symbol"`
	Side          Side              `param:"side"`
	AutoAddMargin bool              `param:"auto_add_margin"`
	PositionIdx   *PositionIdx      `param:"position_idx"`
}

func (this SetAutoAddMargin) Do(client *Client) bool {
	_, ok := Post[struct{}](client, "position/set-auto-add-margin", this)
	return ok
}

func (this *Client) SetAutoAddMargin(v SetAutoAddMargin) bool {
	return v.Do(this)
}

// [Cross/Isolated Margin Switch] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-marginswitch
// Switch Cross/Isolated; must set leverage value when switching from Cross to Isolated
type MarginSwitch struct {
	Symbol       iperpetual.Symbol `param:"symbol"`
	IsIsolated   bool              `param:"is_isolated"`
	BuyLeverage  int               `param:"buy_leverage"`
	SellLeverage int               `param:"sell_leverage"`
}

func (this MarginSwitch) Do(client *Client) bool {
	_, ok := Post[struct{}](client, "position/switch-isolated", this)
	return ok
}

func (this *Client) MarginSwitch(v MarginSwitch) bool {
	return v.Do(this)
}

// [Position Mode Switch] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-switchpositionmode
type PositionModeSwitch struct {
	Mode     PositionMode       `param:"mode"`
	Symbol   *iperpetual.Symbol `param:"symbol"`
	Currency *Currency          `param:"coin"`
}

type PositionMode string

const (
	MergedSingle PositionMode = "MergedSingle"
	BothSide     PositionMode = "BothSide"
)

func (this PositionModeSwitch) Do(client *Client) bool {
	_, ok := Post[struct{}](client, "position/switch-mode", this)
	return ok
}

func (this *Client) PositionModeSwitch(v PositionModeSwitch) bool {
	return v.Do(this)
}

// [Full/Partial Position TP/SL Switch] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-switchmode
// Switch mode between Full or Partial
type TpSlModeSwitch struct {
	Symbol   iperpetual.Symbol `param:"symbol"`
	TpSlMode TpSlMode          `param:"tp_sl_mode"`
}

func (this TpSlModeSwitch) Do(client *Client) (TpSlMode, bool) {
	type result struct {
		TpSlMode TpSlMode `json:"tp_sl_mode"`
	}
	r, ok := Post[result](client, "tpsl/switch-mode", this)
	return r.TpSlMode, ok
}

func (this *Client) TpSlModeSwitch(v TpSlModeSwitch) (TpSlMode, bool) {
	return v.Do(this)
}

// [Add/Reduce Margin] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-addmargin
type AddReduceMargin struct {
	Symbol      iperpetual.Symbol `param:"symbol"`
	Side        Side              `param:"side"`
	Margin      int               `param:"margin"`
	PositionIdx *PositionIdx      `param:"position_idx"`
}

func (this AddReduceMargin) Do(client *Client) (AddReduceMarginResult, bool) {
	return Post[AddReduceMarginResult](client, "position/add-margin", this)
}

type AddReduceMarginResult struct {
	Position         PositionListResult `json:"PositionListResult"`
	WalletBalance    float64            `json:"wallet_balance"`
	AvailableBalance float64            `json:"available_balance"`
}

type PositionListResult struct {
	PositionData
	TpTrigger TriggerPrice `json:"tp_trigger_by"`
	SlTrigger TriggerPrice `json:"sl_trigger_by"`
}

func (this *Client) AddReduceMargin(v AddReduceMargin) (AddReduceMarginResult, bool) {
	return v.Do(this)
}

// [Set Leverage] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-setleverage
type SetLeverage struct {
	Symbol       iperpetual.Symbol `param:"symbol"`
	BuyLeverage  int               `param:"buy_leverage"`
	SellLeverage int               `param:"sell_leverage"`
}

func (this SetLeverage) Do(client *Client) bool {
	_, ok := Post[struct{}](client, "position/set-leverage", this)
	return ok
}

func (this *Client) SetLeverage(v SetLeverage) bool {
	return v.Do(this)
}

// [Set Trading-Stop] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-tradingstop
type SetTradingStop struct {
	Symbol       iperpetual.Symbol `param:"symbol"`
	Side         Side              `param:"side"`
	TakeProfit   *int              `param:"take_profit"`
	StopLoss     *int              `param:"stop_loss"`
	TrailingStop *int              `param:"trailing_stop"`
	TpTrigger    *TriggerPrice     `param:"tp_trigger_by"`
	SlTrigger    *TriggerPrice     `param:"sl_trigger_by"`
	SlSize       *int              `param:"sl_size"`
	TpSize       *int              `param:"tp_size"`
	PositionIdx  *PositionIdx      `param:"position_idx"`
}

func (this SetTradingStop) Do(client *Client) bool {
	_, ok := Post[struct{}](client, "position/trading-stop", this)
	return ok
}

func (this *Client) SetTradingStop(v SetTradingStop) bool {
	return v.Do(this)
}

// [Get User Trade Records] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-usertraderecords
// Get user's trading records.
type GetTradeRecords struct {
	Symbol    iperpetual.Symbol `param:"symbol"`
	StartTime *int              `param:"start_time"`
	EndTime   *int              `param:"end_time"`
	ExecType  *ExecType         `json:"exec_type"`
	Page      *int              `param:"page"`
	Limit     *int              `param:"limit"`
}

func (this GetTradeRecords) Do(client *Client) (TradeRecords, bool) {
	return Get[TradeRecords](client, "trade/execution/list", this)
}

type TradeRecord struct {
	OrderID       string            `json:"order_id"`
	OrderLinkID   string            `json:"order_link_id"`
	Side          Side              `json:"side"`
	Symbol        iperpetual.Symbol `json:"symbol"`
	ExecID        string            `json:"exec_id"`
	OrderPrice    string            `json:"order_price"`
	OrderQty      int               `json:"order_qty"`
	OrderType     OrderType         `json:"order_type"`
	FeeRate       string            `json:"fee_rate"`
	ExecPrice     string            `json:"exec_price"`
	ExecType      ExecType          `json:"exec_type"`
	ExecQty       int               `json:"exec_qty"`
	ExecFee       string            `json:"exec_fee"`
	ExecValue     string            `json:"exec_value"`
	LeavesQty     int               `json:"leaves_qty"`
	ClosedSize    int               `json:"closed_size"`
	LastLiquidity string            `json:"last_liquidity_ind"`
	TradeTime     string            `json:"trade_time"`
	TradeTimeMs   uint64            `json:"trade_time_ms"`
}

type TradeRecords struct {
	CurrentPage int           `json:"current_page"`
	Data        []TradeRecord `json:"data"`
}

func (this *Client) GetTradeRecords(v GetTradeRecords) (TradeRecords, bool) {
	return v.Do(this)
}

// [Extended User Trade Records] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-userhistorytraderecords
// Get user's trading records.
type GetExtendedTradeRecords struct {
	Symbol    iperpetual.Symbol `param:"symbol"`
	StartTime *int              `param:"start_time"`
	EndTime   *int              `param:"end_time"`
	ExecType  *ExecType         `json:"exec_type"`
	pageToken *string           `param:"page_token"`
	Limit     *int              `param:"limit"`
}

func (this GetExtendedTradeRecords) Do(client *Client) (ExtendedTradeRecords, bool) {
	return Get[ExtendedTradeRecords](client, "trade/execution/history-list", this)
}

type ExtendedTradeRecords struct {
	PageToken string        `json:"page_token"`
	Data      []TradeRecord `json:"data"`
}

func (this *Client) GetExtendedTradeRecords(v GetExtendedTradeRecords) (ExtendedTradeRecords, bool) {
	return v.Do(this)
}

// [Closed Profit and Loss] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-closedprofitandloss
type ClosedProfitLoss struct {
	Symbol    iperpetual.Symbol `param:"symbol"`
	StartTime *int              `param:"start_time"`
	EndTime   *int              `param:"end_time"`
	ExecType  *ExecType         `param:"exec_type"`
	Page      *int              `param:"page"`
	Limit     *int              `param:"limit"`
}

func (this ClosedProfitLoss) Do(client *Client) (ClosedProfitLossResult, bool) {
	return Get[ClosedProfitLossResult](client, "trade/closed-pnl/list", this)
}

type ClosedData struct {
	ID            int               `json:"id"`
	UserID        int               `json:"user_id"`
	Symbol        iperpetual.Symbol `json:"symbol"`
	OrderID       string            `json:"order_id"`
	Side          Side              `json:"side"`
	Qty           float64           `json:"qty"`
	OrderPrice    float64           `json:"order_price"`
	OrderType     OrderType         `json:"order_type"`
	ExecType      ExecType          `json:"exec_type"`
	ClosedSize    float64           `json:"closed_size"`
	CumEntryValue float64           `json:"cum_entry_value"`
	AvgEntryPrice float64           `json:"avg_entry_price"`
	CumExitValue  float64           `json:"cum_exit_value"`
	AvgExitPrice  float64           `json:"avg_exit_price"`
	ClosedPnl     float64           `json:"closed_pnl"`
	FillCount     int               `json:"fill_count"`
	Leverage      int               `json:"leverage"`
	CreatedAt     uint64            `json:"created_at"`
}

type ClosedProfitLossResult struct {
	CurrentPage int          `json:"current_page"`
	Data        []ClosedData `json:"data"`
}

func (this *Client) ClosedProfitLoss(v ClosedProfitLoss) (ClosedProfitLossResult, bool) {
	return v.Do(this)
}
