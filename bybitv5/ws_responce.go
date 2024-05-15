package bybitv5

import (
	"encoding/json"

	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/ulog"
)

type WsResponse interface {
	IsOperateion() bool
	OperationIs(string) bool
	Ok() bool
	Log(*ulog.Log)
}

type WsBaseResponse struct {
	Operation    string `json:"op"`
	ConnectionId string `json:"conn_id"`
	Success      bool
	Message      string `json:"ret_msg"`
	Args         []string
}

func (o WsBaseResponse) IsOperateion() bool {
	return o.Operation != ""
}

func (o WsBaseResponse) OperationIs(v string) bool {
	return o.Operation == v
}

func (o WsBaseResponse) Ok() bool {
	return o.Success
}

func (o WsBaseResponse) Log(log *ulog.Log) {
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

type WsOrderData struct {
	OrderId     string
	OrderLinkId string
}

type WsResponseHeader struct {
	Limit               string `json:"X-Bapi-Limit"`
	LimitStatus         string `json:"X-Bapi-Limit-Status"`
	LimitResetTimestamp string `json:"X-Bapi-Limit-Reset-Timestamp"`
	TraceId             string
	TimeNow             string
}

type WsTradeResponse struct {
	Operation    string `json:"op"`
	ConnectionId string `json:"connId"`
	Code         int    `json:"retCode"`
	Message      string `json:"retMsg"`
	Data         json.RawMessage
	Header       WsResponseHeader
}

func (o WsTradeResponse) OrderData() (v WsOrderData, err error) {
	err = json.Unmarshal(o.Data, &v)
	return
}

func (o WsTradeResponse) IsOperateion() bool {
	return o.Operation != ""
}

func (o WsTradeResponse) OperationIs(v string) bool {
	return o.Operation == v
}

func (o WsTradeResponse) Ok() bool {
	return o.Code == 0
}

func (o WsTradeResponse) Log(log *ulog.Log) {
	switch o.Operation {
	case "ping":
	case "pong":
	case "order.create":
		if o.Code == 0 {
			log.Info("order-create")
		} else {
			log.Error("order-create:", o.Message)
		}
	default:
		log.Error("invalid response:", o.Operation)
	}
}
