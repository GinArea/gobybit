package iperpetual

import "github.com/msw-x/moon/ufmt"

type Subscription struct {
	Topic    TopicName
	Interval string
	Symbol   string
}

func (this *Subscription) String() string {
	s := []string{string(this.Topic)}
	if this.Interval != "" {
		s = append(s, this.Interval)
	}
	if this.Symbol != "" {
		s = append(s, this.Symbol)
	}
	return ufmt.JoinSliceWith(".", s)
}
