package bybitv5

import (
	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
)

type WsResponse struct {
	Operation    string `json:"op"`
	ConnectionId string `json:"conn_id"`
	Success      bool
	Message      string `json:"ret_msg"`
	Args         []string
}

func (o WsResponse) Valid() bool {
	return o.Operation != ""
}

func (o WsResponse) Log(log *ulog.Log) {
	switch o.Operation {
	case "ping":
	case "pong":
	case "subscribe":
		if o.Success {
			log.Info("subscribe: success")
		} else {
			log.Error("subscribe:", o.Message)
		}
	case "unsubscribe":
		log.Info("unsubscribe:", ufmt.SuccessFailure(o.Success))
	default:
		log.Error("invalid response:", o.Operation)
	}
}
