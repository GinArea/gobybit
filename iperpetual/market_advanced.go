// Advanced Data (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-advanceddata)
package iperpetual

// Open Interest (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketopeninterest)
//
// Gets the total amount of unsettled contracts. In other words, the total number of contracts held in open positions.
//
//	symbol Required string Symbol
//	period Required string Data recording period. 5min, 15min, 30min, 1h, 4h, 1d
//	limit           int    Limit for data size per page, max size is 200. Default as showing 50 pieces of data per page
type OpenInterest struct {
	Symbol string `param:"symbol"`
	Period string `param:"period"`
	Limit  *int   `param:"limit"`
}

func (this OpenInterest) Do(client *Client) ([]InterestItem, bool) {
	return GetPublic[[]InterestItem](client, "open-interest", this)
}

type InterestItem struct {
	Symbol       string `json:"symbol"`
	Timestamp    uint64 `json:"timestamp"`
	OpenInterest uint64 `json:"open_interest"`
}

func (this *Client) OpenInterest(v OpenInterest) ([]InterestItem, bool) {
	return v.Do(this)
}

// Latest Big Deal (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketbigdeal)
//
// Obtain filled orders worth more than 500,000 USD within the last 24h.
// This endpoint may return orders which are over the maximum order qty for the symbol you call.
// For instance, the maximum order qty for BTCUSD is 1 million contracts, but in the event of the
// liquidation of a position larger than 1 million this endpoint returns this "impossible" order size.
//
//	symbol Required string Symbol
//	limit           int    Limit for data size per page, max size is 1000. Default as showing 500 pieces of data per page
type LatestBigDeal struct {
	Symbol string `param:"symbol"`
	Limit  *int   `param:"limit"`
}

func (this LatestBigDeal) Do(client *Client) ([]LatestBigDealItem, bool) {
	return GetPublic[[]LatestBigDealItem](client, "big-deal", this)
}

type LatestBigDealItem struct {
	Symbol    string  `json:"symbol"`
	Side      Side    `json:"side"`
	Timestamp uint64  `json:"timestamp"`
	Value     float64 `json:"value"`
}

func (this *Client) LatestBigDeal(v LatestBigDeal) ([]LatestBigDealItem, bool) {
	return v.Do(this)
}

// Long-Short Ratio (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marketaccountratio)
//
// Gets the Bybit user accounts' long-short ratio.
//
//	symbol Required string Symbol
//	period Required string Data recording period. 5min, 15min, 30min, 1h, 4h, 1d
//	limit           int    Limit for data size per page, max size is 500. Default as showing 50 pieces of data per page
type LongShortRatio struct {
	Symbol string `param:"symbol"`
	Period string `param:"period"`
	Limit  *int   `param:"limit"`
}

func (this LongShortRatio) Do(client *Client) ([]LongShortRatioItem, bool) {
	return GetPublic[[]LongShortRatioItem](client, "big-deal", this)
}

type LongShortRatioItem struct {
	Symbol    string  `json:"symbol"`
	BuyRatio  float64 `json:"buy_ratio"`
	SellRatio float64 `json:"sell_ratio"`
	Timestamp uint64  `json:"timestamp"`
}

func (this *Client) LongShortRatio(v LongShortRatio) ([]LongShortRatioItem, bool) {
	return v.Do(this)
}
