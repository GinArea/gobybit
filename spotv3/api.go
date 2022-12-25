// API Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-api)
package spotv3

// Server Time (https://bybit-exchange.github.io/docs/spot/v3/#t-servertime)
func (this *Client) ServerTime() (string, error) {
	type result struct {
		ServerTime string `json:"serverTime"`
	}
	r, err := GetPublic[result](this, "server-time", nil)
	return r.ServerTime, err
}
