package bybitv5

type WsRequest[T any] struct {
	RequestId string    `json:"req_id,omitempty"`
	Header    *WsHeader `json:",omitempty"`
	Operation string    `json:"op"`
	Args      []T       `json:",omitempty"`
}
