package transport

import "fmt"

type PreRateLimitError struct {
	Limit  int
	Status int
}

func (o *PreRateLimitError) Error() string {
	return fmt.Sprintf("rate limit[%d] status[%d]", o.Limit, o.Status)
}
