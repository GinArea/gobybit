package v5

type Response[T any] struct {
	RetCode int
	RetMsg  string
	Time    uint64
	Result  T
}
