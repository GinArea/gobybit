package bybitv5

type WsHeader struct {
	Timestamp  string `json:"X-BAPI-TIMESTAMP"`
	RecvWindow string `json:"X-BAPI-RECV-WINDOW"`
	Referer    string `json:",omitempty"`
}
