package bybitv5

type WsRequest[T any] struct {
	RequestId string    `json:"reqId,omitempty"`
	Header    *WsHeader `json:",omitempty"`
	Operation string    `json:"op"`
	Args      []T       `json:",omitempty"`
}
