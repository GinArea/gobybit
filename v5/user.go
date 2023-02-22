package v5

import (
	"time"

	"github.com/ginarea/gobybit/transport"
)

type GetKeyInfo struct {
}

func (o GetKeyInfo) Do(client *Client) (KeyInfo, error) {
	return Get[KeyInfo](client, "user/query-api", o)
}

type KeyInfo struct {
	Id            transport.Float64
	Note          string
	ApiKey        string
	ReadOnly      int
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

func (o *Client) GetKeyInfo() (KeyInfo, error) {
	return GetKeyInfo{}.Do(o)
}
