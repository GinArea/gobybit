// API Data Endpoints (https://bybit-exchange.github.io/docs/spot/v1/#t-api)
package spot

// Server Time (https://bybit-exchange.github.io/docs/spot/v1/#t-servertime)
func (this *Client) ServerTime() (uint64, error) {
	type result struct {
		Time uint64 `json:"serverTime"`
	}
	r, err := GetPublic[result](this, "time", nil)
	return r.Time, err
}
