package spotv3

type Response[T any] struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Time    uint64 `json:"time"`
	Result  T      `json:"result"`
}
