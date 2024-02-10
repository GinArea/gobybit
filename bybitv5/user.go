package bybitv5

import (
	"time"

	"github.com/msw-x/moon/ujson"
)

// Get API Key Information
// https://bybit-exchange.github.io/docs/v5/user/apikey-info
type GetKeyInfo struct {
}

type KeyInfo struct {
	Id            ujson.Float64
	Note          string
	ApiKey        string
	ReadOnly      ujson.Bool
	Secret        string
	Permissions   Permissions
	Ips           []string
	Type          int
	DeadlineDay   int
	CreatedAt     time.Time
	ExpiredAt     time.Time
	Unified       int
	Uta           int
	UserID        int
	InviterID     int
	VipLevel      string
	MktMakerLevel string
	AffiliateID   int
	RsaPublicKey  string
	IsMaster      bool
	ParentUid     string
	KycLevel      string
	KycRegion     string
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

func (o GetKeyInfo) Do(c *Client) Response[KeyInfo] {
	return Get(c.user(), "query-api", o, forward[KeyInfo])
}

func (o *Client) GetKeyInfo() Response[KeyInfo] {
	return GetKeyInfo{}.Do(o)
}
