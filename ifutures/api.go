// [API Data Endpoints] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-api
// The following API data endpoints do not require authentication.
package ifutures

// [Server Time] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-servertime
func (this *Client) ServerTime() (string, bool) {
	resp := &Response[struct{}]{}
	err := this.GetPublic("time", nil, resp)
	return resp.TimeNow, err == nil

}

// [Announcement] https://bybit-exchange.github.io/docs/futuresV2/inverse_futures/#t-announcement
// Get Bybit OpenAPI announcements in the last 30 days in reverse order.
type Announcement struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Linkg     string `json:"link"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
}

func (this *Client) Announcement() ([]Announcement, bool) {
	resp := &Response[[]Announcement]{}
	err := this.GetPublic("announcement", nil, resp)
	return resp.Result, err == nil

}
