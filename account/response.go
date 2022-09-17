package account

type Response[T any] struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  T      `json:"result"`
	TimeNow int64  `json:"time_now"`
}
