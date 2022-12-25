// API Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-api)
package iperpetual

// Server Time (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-servertime)
func (this *Client) ServerTime() (string, error) {
	resp := &Response[struct{}]{}
	err := this.GetPublic("time", nil, resp)
	return resp.TimeNow, err
}

// Announcement (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-announcement)
//
// Get Bybit OpenAPI announcements in the last 30 days in reverse order.
type Announcement struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Linkg     string `json:"link"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
}

func (this *Client) Announcement() ([]Announcement, error) {
	resp := &Response[[]Announcement]{}
	err := this.GetPublic("announcement", nil, resp)
	return resp.Result, err
}
