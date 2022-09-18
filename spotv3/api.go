// [API Data Endpoints] https://bybit-exchange.github.io/docs/futuresV2/linear/#t-api
// The following API data endpoints do not require authentication.
package spotv3

// [Server Time] https://bybit-exchange.github.io/docs/spot/v3/#t-servertime
func (this *Client) ServerTime() (string, bool) {
	type result struct {
		ServerTime string `json:"serverTime"`
	}
	r, ok := GetPublic[result](this, "server-time", nil)
	return r.ServerTime, ok
}
