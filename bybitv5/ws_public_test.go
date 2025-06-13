package bybitv5

import (
	"fmt"
	"testing"
	"time"
)

func TestOrderBook(t *testing.T) {
	tests := []struct {
		name   string
		symbol string
		client *WsPublic
	}{
		{
			name:   "Orderbook",
			symbol: "XRPUSD",
			client: NewWsPublic().WithCategory(Inverse),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.client.Run()
			time.Sleep(time.Duration(time.Second * 3))
			tt.client.OrderBook(tt.symbol, 1).Subscribe(func(t Topic[Orderbook]) {
				fmt.Printf("%v", t)
			})
			time.Sleep(time.Duration(time.Second * 3600))
		})
	}
}
