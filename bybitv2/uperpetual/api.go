// API Data Endpoints (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-api)
package uperpetual

import "github.com/ginarea/gobybit/bybitv2/iperpetual"

// Server Time (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-servertime)
func (this *Client) ServerTime() (string, error) {
	return this.iperpetual().ServerTime()
}

// Announcement (https://bybit-exchange.github.io/docs/futuresV2/linear/#t-announcement)
//
// Get Bybit OpenAPI announcements in the last 30 days in reverse order
func (this *Client) Announcement() ([]iperpetual.Announcement, error) {
	return this.iperpetual().Announcement()
}
