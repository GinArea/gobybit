package spot

type Response[T any] struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	TimeNow string `json:"time_now"`
	Result  T      `json:"result"`
}
