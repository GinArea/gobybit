// [API Data Endpoints] https://bybit-exchange.github.io/docs/spot/v1/#t-api
// The following API data endpoints do not require authentication.
package spot

// [Server Time] https://bybit-exchange.github.io/docs/spot/v1/#t-servertime
func (this *Client) ServerTime() (uint64, bool) {
	type result struct {
		Time uint64 `json:"serverTime"`
	}
	r, ok := GetPublic[result](this, "time", nil)
	return r.Time, ok
}
