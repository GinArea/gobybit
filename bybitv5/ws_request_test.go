package bybitv5

import (
	"encoding/json"
	"testing"

	"github.com/msw-x/moon/ujson"
)

// TestWsRequestSerialization pins the wire shape of every WsRequest
// sender as it goes through uws SendJson (ujson.MarshalLowerCase, which
// lowercases only the first character of a key and stumbles over "-"):
//   - auth / subscribe / ping / legacy WsTrade.PlaceOrder frames carry
//     NO correlation key at all (RequestId is empty, omitempty drops
//     it) — proof that the reqId tag change is byte-invisible to every
//     pre-existing consumer (ts/ts-bt on the published version);
//   - only TradeRequest frames (order.create/amend/cancel via the
//     additive P5 path) set RequestId, serialized as the official
//     guideline key "reqId" (camelCase, never req_id).
func TestWsRequestSerialization(t *testing.T) {
	marshal := func(v any) map[string]any {
		data, err := ujson.MarshalLowerCase(v)
		if err != nil {
			t.Fatal(err)
		}
		var frame map[string]any
		if err := json.Unmarshal(data, &frame); err != nil {
			t.Fatal(err)
		}
		return frame
	}
	assertNoReqID := func(name string, frame map[string]any) {
		t.Helper()
		for _, key := range []string{"req_id", "reqId"} {
			if _, ok := frame[key]; ok {
				t.Fatalf("%s frame must not carry %q: %v", name, key, frame)
			}
		}
	}

	// Auth frame (WsAuthClient.auth): op + args only.
	auth := marshal(WsRequest[string]{
		Operation: "auth",
		Args:      []string{"key", "1718000000000", "signature"},
	})
	assertNoReqID("auth", auth)
	if auth["op"] != "auth" {
		t.Fatalf("auth op = %v", auth["op"])
	}
	if args, _ := auth["args"].([]any); len(args) != 3 {
		t.Fatalf("auth args = %v", auth["args"])
	}

	// Subscribe frame (WsClient.SubscribeBatch): op + args only.
	sub := marshal(WsRequest[string]{
		Operation: "subscribe",
		Args:      []string{"orderbook.1.BTCUSDT"},
	})
	assertNoReqID("subscribe", sub)
	if sub["op"] != "subscribe" {
		t.Fatalf("subscribe op = %v", sub["op"])
	}
	if args, _ := sub["args"].([]any); len(args) != 1 {
		t.Fatalf("subscribe args = %v", sub["args"])
	}

	// Legacy WsTrade.PlaceOrder frame: header + op + args, RequestId
	// never set by the legacy path.
	place := marshal(WsRequest[WsPlaceOrder]{
		Header: &WsHeader{
			Timestamp:  "1718000000000",
			RecvWindow: "20000",
		},
		Operation: "order.create",
		Args: []WsPlaceOrder{{
			Category:  Linear,
			Symbol:    "BTCUSDT",
			Side:      Buy,
			OrderType: Limit,
			Qty:       "0.001",
			Price:     "60000.1",
		}},
	})
	assertNoReqID("legacy place", place)
	if place["op"] != "order.create" {
		t.Fatalf("place op = %v", place["op"])
	}
	args, _ := place["args"].([]any)
	if len(args) != 1 {
		t.Fatalf("place args = %v", place["args"])
	}

	// TradeRequest frame (the additive P5 path): exactly "reqId".
	trade := marshal(WsRequest[any]{
		RequestId: "req-42",
		Header: &WsHeader{
			Timestamp:  "1718000000000",
			RecvWindow: "20000",
		},
		Operation: "order.cancel",
		Args:      []any{WsCancelOrder{Category: Linear, Symbol: "BTCUSDT", OrderLinkId: "link-a"}},
	})
	if _, ok := trade["req_id"]; ok {
		t.Fatalf("trade frame must not carry req_id: %v", trade)
	}
	if trade["reqId"] != "req-42" {
		t.Fatalf("trade frame reqId = %v, want req-42", trade["reqId"])
	}
	if trade["op"] != "order.cancel" {
		t.Fatalf("trade op = %v", trade["op"])
	}
	if header, _ := trade["header"].(map[string]any); header == nil ||
		header["X-BAPI-TIMESTAMP"] != "1718000000000" || header["X-BAPI-RECV-WINDOW"] != "20000" {
		t.Fatalf("trade header = %v", trade["header"])
	}
	if args, _ := trade["args"].([]any); len(args) != 1 {
		t.Fatalf("trade args = %v", trade["args"])
	}
}
