package bybit

type Response[T any] struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	Result  T      `json:"result"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	TimeNow string `json:"time_now"`
}
