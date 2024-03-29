package iperpetual

type Response[T any] struct {
	RetCode         int    `json:"ret_code"`
	RetMsg          string `json:"ret_msg"`
	ExtCode         string `json:"ext_code"`
	ExtInfo         string `json:"ext_info"`
	Result          T      `json:"result"`
	TimeNow         string `json:"time_now"`
	RateLimitStatus int    `json:"rate_limit_status"`
	RateLimit       int    `json:"rate_limit"`
}
