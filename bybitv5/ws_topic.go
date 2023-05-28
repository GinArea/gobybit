package bybitv5

import (
	"encoding/json"

	"github.com/msw-x/moon/ujson"
)

type PublicTopic struct {
	Topic string
	Ts    int64
	Type  string
	Data  json.RawMessage
}

func (o PublicTopic) Name() string {
	return o.Topic
}

func (o PublicTopic) RawData() []byte {
	return o.Data
}

type PrivateTopic struct {
	Topic        string
	CreationTime int64
	Id           string
	Data         json.RawMessage
}

func (o PrivateTopic) Name() string {
	return o.Topic
}

func (o PrivateTopic) RawData() []byte {
	return o.Data
}

type TradeShot struct {
	Timestamp     int64         `json:"ts"`
	Symbol        string        `json:"s"`
	Side          Side          `json:"S"`
	Size          ujson.Float64 `json:"v"`
	Price         ujson.Float64 `json:"p"`
	TickDirection TickDirection `json:"L"` // Unique field for future
	Id            string        `json:"i"`
	Block         bool          `json:"BT"`
}
