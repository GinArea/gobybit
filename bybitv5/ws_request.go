package bybitv5

type WsRequest struct {
	RequestId string   `json:"req_id,omitempty"`
	Operation string   `json:"op"`
	Args      []string `json:",omitempty"`
}
