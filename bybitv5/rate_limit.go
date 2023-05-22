package bybitv5

type RateLimit struct {
	Limit          int   `http:"X-Bapi-Limit"`
	Status         int   `http:"X-Bapi-Limit-Status"`
	ResetTimestamp int64 `http:"X-Bapi-Limit-Reset-Timestamp"`
}
