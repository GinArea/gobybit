package bybitv5

import "testing"

func TestInstrumentsSUI(t *testing.T) {
	ins := GetInstruments{Category: Linear, Symbol: "SUIUSDT"}.Do(NewClient())
	if !ins.Ok() {
		t.Fatalf("ins: %v", ins.Error)
	}
	for _, i := range ins.Data {
		if i.Symbol != "SUIUSDT" {
			continue
		}
		t.Logf("tick=%v qtyStep=%v minOrderQty=%v minNotional=%v",
			i.PriceFilter.TickSize.Value(), i.LotSizeFilter.QtyStep.Value(),
			i.LotSizeFilter.MinOrderQty.Value(), i.LotSizeFilter.MinNotionalValue.Value())
	}
}
