package transport

import "fmt"

type RateLimitError struct {
	Limit  int
	Status int
}

func (o *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit[%d] status[%d]", o.Limit, o.Status)
}
