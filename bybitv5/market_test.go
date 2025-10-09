package bybitv5

import (
	"fmt"
	"testing"
)

func Test_GetInstrumentInfo_Spot(t *testing.T) {
	tests := []struct {
		name    string
		client  *Client
		request GetInstruments
	}{
		{
			name:   "Spot instrument info",
			client: NewClient(),
			request: GetInstruments{
				Category: Spot,
				Symbol:   "BTCUSDT",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.client.GetInstrumentsSpot(tt.request)
			fmt.Printf("%v", got)
		})
	}
}
