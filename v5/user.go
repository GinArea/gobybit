package v5

import (
	"time"

	"github.com/ginarea/gobybit/transport"
)

type GetKeyInfo struct {
}

func (this GetKeyInfo) Do(client *Client) (KeyInfo, error) {
	return Get[KeyInfo](client, "user/query-api", this)
}

type KeyInfo struct {
	Id            transport.Float64 `json:"id"`
	Note          string            `json:"note"`
	ApiKey        string            `json:"apiKey"`
	ReadOnly      int               `json:"readOnly"`
	Secret        string            `json:"secret"`
	Permissions   Permissions       `json:"permissions"`
	Ips           []string          `json:"ips"`
	Type          int               `json:"type"`
	DeadlineDay   int               `json:"deadlineDay"`
	CreatedAt     time.Time         `json:"createdAt"`
	ExpiredAt     time.Time         `json:"expiredAt"`
	Unified       int               `json:"unified"`
	Uta           int               `json:"uta"`
	UserID        int               `json:"userID"`
	InviterID     int               `json:"inviterID"`
	VipLevel      string            `json:"vipLevel"`
	MktMakerLevel string            `json:"mktMakerLevel"`
	AffiliateID   int               `json:"affiliateID"`
	RsaPublicKey  string            `json:"rsaPublicKey"`
}

type Permissions struct {
	ContractTrade []string
	Spot          []string
	Wallet        []string
	Options       []string
	Derivatives   []string
	CopyTrading   []string
	BlockTrade    []string
	Exchange      []string
	NFT           []string
}

func (this *Client) GetKeyInfo() (KeyInfo, error) {
	return GetKeyInfo{}.Do(this)
}
