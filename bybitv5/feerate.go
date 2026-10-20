package bybitv5

// Get Fee Rate
// https://bybit-exchange.github.io/docs/v5/account/fee-rate
//
//	category Required string  Product type. spot,linear,inverse,option.
//	                             linear/inverse both return all contracts
//	symbol            string  Symbol name, like BTCUSDT. Valid for
//	                             linear, inverse, spot
type GetFeeRate struct {
	Category Category
	Symbol   string `url:",omitempty"`
}

// FeeRateList is the result payload of the fee-rate endpoint. The
// category field is documented for spot/option only; derivatives
// responses omit it (stays empty).
type FeeRateList struct {
	Category string    `json:"category"`
	List     []FeeRate `json:"list"`
}

// FeeRate is one symbol row of the account fee-rate response. The values
// are the actual rates of the calling account (VIP / Pro / MM tier
// applied by the venue). rpiMakerFeeRate applies to RPI (Retail Price
// Improvement) maker orders.
type FeeRate struct {
	Symbol          string
	TakerFeeRate    string
	MakerFeeRate    string
	RpiMakerFeeRate string
}

func (o GetFeeRate) Do(c *Client) Response[FeeRateList] {
	return Get(c.account(), "fee-rate", o, forward[FeeRateList])
}

func (o *Client) GetFeeRate(v GetFeeRate) Response[FeeRateList] {
	return v.Do(o)
}
