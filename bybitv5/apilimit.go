package bybitv5

// Get All Rate Limits (UID-level pools, rules for Pros)
// https://bybit-exchange.github.io/docs/v5/rate-limit/rules-for-pros/apilimit-query-all
type GetApiLimits struct {
	Limit  int    `url:",omitempty"`
	Cursor string `url:",omitempty"`
	Uids   string `url:",omitempty"`
}

type ApiLimits struct {
	NextPageCursor string           `json:"nextPageCursor"`
	List           []ApiLimitsEntry `json:"list"`
}

type ApiLimitsEntry struct {
	Uids    string `json:"uids"`
	BizType string `json:"bizType"`
	Rate    int    `json:"rate"`
}

func (o GetApiLimits) Do(c *Client) Response[ApiLimits] {
	return Get(c.apilimit(), "query-all", o, forward[ApiLimits])
}

func (o *Client) GetApiLimits() Response[ApiLimits] {
	return GetApiLimits{}.Do(o)
}
