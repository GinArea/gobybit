// [API Key Info] https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-key
package iperpetual

type GetKeyInfo struct {
}

func (this GetKeyInfo) Do(client *Client) ([]KeyInfo, bool) {
	return GetPrivate[[]KeyInfo](client, "account/api-key", this)
}

type KeyInfo struct {
	ApiKey        string   `json:"api_key"`
	Type          string   `json:"type"`
	UserID        int      `json:"user_id"`
	InviterID     int      `json:"inviter_id"`
	Ips           []string `json:"ips"`
	Note          string   `json:"note"`
	Permissions   []string `json:"permissions"`
	CreatedAt     string   `json:"created_at"`
	ExpiredAt     string   `json:"expired_at"`
	ReadOnly      bool     `json:"read_only"`
	VipLevel      string   `json:"vip_level"`
	MktMakerLevel string   `json:"mkt_maker_level"`
}

func (this *Client) GetKeyInfo() ([]KeyInfo, bool) {
	return GetKeyInfo{}.Do(this)
}