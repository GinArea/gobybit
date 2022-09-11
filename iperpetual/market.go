// [Market Data Endpoints] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketdata
package iperpetual

// [Query Symbol] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-querysymbol
type SymbolInfo struct {
	Name            string         `json:"name"`
	Alias           string         `json:"alias"`
	Status          ContractStatus `json:"status"`
	BaseCurrency    string         `json:"base_currency"`
	QuoteCurrency   string         `json:"quote_currency"`
	PriceScale      int            `json:"price_scale"`
	TakerFee        string         `json:"taker_fee"`
	MakerFee        string         `json:"maker_fee"`
	FundingInterval int            `json:"funding_interval"`
	LeverageFilter  LeverageFilter `json:"leverage_filter"`
	PriceFilter     PriceFilter    `json:"price_filter"`
	LotSizeFilter   LotSizeFilter  `json:"lot_size_filter"`
}

type LeverageFilter struct {
	Min  int    `json:"min_leverage"`
	Max  int    `json:"max_leverage"`
	Step string `json:"leverage_step"`
}

type PriceFilter struct {
	Min      string `json:"min_price"`
	Max      string `json:"max_price"`
	TickSize string `json:"tick_size"`
}

type LotSizeFilter struct {
	MaxTradingQty         float64 `json:"max_trading_qty"`
	MinTradingQty         float64 `json:"min_trading_qty"`
	QtyStep               float64 `json:"qty_step"`
	PostOnlyMaxTradingQty string  `json:"post_only_max_trading_qty"`
}

func (this *Client) QuerySymbol() ([]SymbolInfo, bool) {
	return GetPublic[[]SymbolInfo](this, "symbols", nil)
}

func (this *Client) QuerySymbolNames() ([]string, bool) {
	result, ok := this.QuerySymbol()
	names := make([]string, len(result))
	for n, s := range result {
		names[n] = s.Name
	}
	return names, ok
}
