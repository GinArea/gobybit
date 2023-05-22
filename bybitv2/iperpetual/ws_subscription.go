package iperpetual

import "github.com/msw-x/moon/ufmt"

type Subscription struct {
	Topic    TopicName
	Interval string
	Symbol   string
}

func (o *Subscription) String() string {
	s := []string{string(o.Topic)}
	if o.Interval != "" {
		s = append(s, o.Interval)
	}
	if o.Symbol != "" {
		s = append(s, o.Symbol)
	}
	return ufmt.JoinSliceWith(".", s)
}
