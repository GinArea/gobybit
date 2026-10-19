package bybitv5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

//
// Mock trade-WS server (httptest + gorilla upgrader): auth -> order.create /
// order.amend / order.cancel with a reqId echo, the official guideline ack
// schema. Captures every raw frame the client sends for schema assertions.
//

type tradeWsMock struct {
	srv      *httptest.Server
	upgrader websocket.Upgrader

	mu     sync.Mutex
	frames []map[string]any // parsed client frames in arrival order

	writeMu sync.Mutex
}

func newTradeWsMock(t *testing.T) *tradeWsMock {
	t.Helper()
	m := &tradeWsMock{}
	m.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := m.upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var frame map[string]any
			if json.Unmarshal(data, &frame) != nil {
				continue
			}
			m.mu.Lock()
			m.frames = append(m.frames, frame)
			m.mu.Unlock()
			switch frame["op"] {
			case "ping":
				m.write(conn, `{"op":"pong","success":true,"conn_id":"mock"}`)
			case "auth":
				m.write(conn, `{"op":"auth","success":true,"conn_id":"mock","ret_msg":""}`)
			case "order.create", "order.amend", "order.cancel":
				reqId, _ := frame["reqId"].(string)
				m.write(conn, fmt.Sprintf(`{"retCode":0,"retMsg":"OK","connId":"mock","reqId":%q,"op":%q,`+
					`"data":{"orderId":"oid-1","orderLinkId":"link-1"},`+
					`"header":{"X-Bapi-Limit":"600","X-Bapi-Limit-Status":"599","X-Bapi-Limit-Reset-Timestamp":"1718000000000","Timenow":"1718000000123"}}`,
					reqId, frame["op"]))
			}
		}
	}))
	t.Cleanup(m.srv.Close)
	return m
}

func (m *tradeWsMock) write(conn *websocket.Conn, s string) {
	m.writeMu.Lock()
	defer m.writeMu.Unlock()
	conn.WriteMessage(websocket.TextMessage, []byte(s))
}

func (m *tradeWsMock) url() string { return "ws" + strings.TrimPrefix(m.srv.URL, "http") }

func (m *tradeWsMock) framesOf(op string) []map[string]any {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []map[string]any
	for _, f := range m.frames {
		if f["op"] == op {
			out = append(out, f)
		}
	}
	return out
}

func waitReady(t *testing.T, ws *WsTrade) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if ws.Ready() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("trade ws not ready within 5s")
}

// TestWsTradeTradeRequest covers the additive trade-WS path: auth, then
// order.create with OrderLinkId/ReduceOnly, order.amend and order.cancel via
// TradeRequest; the mock asserts the official guideline request schema
// (reqId + header + a single arg with string price/qty) and echoes reqId in
// the ack parsed back into WsTradeResponse (incl. the rate-limit header).
func TestWsTradeTradeRequest(t *testing.T) {
	mock := newTradeWsMock(t)
	ws := NewWsTrade("key", "secret").WithBaseUrl(mock.url())
	ws.WithOnResponse(func(r WsTradeResponse) {})
	ws.Run()
	defer ws.Close()
	waitReady(t, ws)

	ro := true
	ws.TradeRequest("order.create", "req-create-1", WsPlaceOrder{
		Category:    Linear,
		Symbol:      "BTCUSDT",
		Side:        Buy,
		OrderType:   Limit,
		Qty:         "0.001",
		Price:       "60000.1",
		TimeInForce: PostOnly,
		OrderLinkId: "link-a",
		ReduceOnly:  &ro,
	})
	ws.TradeRequest("order.amend", "req-amend-1", WsAmendOrder{
		Category: Linear,
		Symbol:   "BTCUSDT",
		OrderId:  "oid-1",
		Price:    "60000.2",
	})
	ws.TradeRequest("order.cancel", "req-cancel-1", WsCancelOrder{
		Category:    Linear,
		Symbol:      "BTCUSDT",
		OrderLinkId: "link-a",
	})

	// The mock keeps the frames; assert the wire schema of each operation.
	expectFrames := func(op string, want int) {
		deadline := time.Now().Add(5 * time.Second)
		for time.Now().Before(deadline) {
			if len(mock.framesOf(op)) == want {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
		t.Fatalf("op %s: frames = %d, want %d", op, len(mock.framesOf(op)), want)
	}
	expectFrames("order.create", 1)
	expectFrames("order.amend", 1)
	expectFrames("order.cancel", 1)

	create := mock.framesOf("order.create")[0]
	assertTradeFrameSchema(t, "create", create, "req-create-1")
	args, _ := create["args"].([]any)
	arg, _ := args[0].(map[string]any)
	// String price/qty as the trade-WS requires.
	if arg["price"] != "60000.1" || arg["qty"] != "0.001" {
		t.Fatalf("create arg price/qty = %v/%v, want strings 60000.1/0.001", arg["price"], arg["qty"])
	}
	if arg["orderLinkId"] != "link-a" {
		t.Fatalf("create arg orderLinkId = %v, want link-a", arg["orderLinkId"])
	}
	if arg["reduceOnly"] != true {
		t.Fatalf("create arg reduceOnly = %v, want true", arg["reduceOnly"])
	}
	if arg["timeInForce"] != "PostOnly" {
		t.Fatalf("create arg timeInForce = %v, want PostOnly", arg["timeInForce"])
	}
	if arg["category"] != "linear" || arg["symbol"] != "BTCUSDT" || arg["side"] != "Buy" || arg["orderType"] != "Limit" {
		t.Fatalf("create arg base fields = %v", arg)
	}

	amend := mock.framesOf("order.amend")[0]
	assertTradeFrameSchema(t, "amend", amend, "req-amend-1")
	aArgs, _ := amend["args"].([]any)
	aArg, _ := aArgs[0].(map[string]any)
	if aArg["orderId"] != "oid-1" || aArg["price"] != "60000.2" {
		t.Fatalf("amend arg = %v, want orderId oid-1, price string 60000.2", aArg)
	}
	if _, ok := aArg["qty"]; ok {
		t.Fatalf("amend arg qty must be omitted when empty: %v", aArg)
	}
	if _, ok := aArg["orderLinkId"]; ok {
		t.Fatalf("amend arg orderLinkId must be omitted when empty: %v", aArg)
	}

	cancel := mock.framesOf("order.cancel")[0]
	assertTradeFrameSchema(t, "cancel", cancel, "req-cancel-1")
	cArgs, _ := cancel["args"].([]any)
	cArg, _ := cArgs[0].(map[string]any)
	if cArg["orderLinkId"] != "link-a" {
		t.Fatalf("cancel arg = %v, want orderLinkId link-a", cArg)
	}
	if _, ok := cArg["orderId"]; ok {
		t.Fatalf("cancel arg orderId must be omitted when empty: %v", cArg)
	}
}

// assertTradeFrameSchema pins the official trade-WS request frame shape
// (guideline, docs master branch): the correlation key is reqId
// (camelCase; never req_id), the singular header object carries the
// uppercase X-BAPI-TIMESTAMP / X-BAPI-RECV-WINDOW keys, and args holds
// exactly one element. Hardcoded against the guideline text, not against
// the client serialization.
func assertTradeFrameSchema(t *testing.T, name string, frame map[string]any, wantReqID string) {
	t.Helper()
	if _, ok := frame["req_id"]; ok {
		t.Fatalf("%s frame carries req_id (wrong key, guideline says reqId): %v", name, frame)
	}
	if frame["reqId"] != wantReqID {
		t.Fatalf("%s frame reqId = %v, want %q", name, frame["reqId"], wantReqID)
	}
	header, _ := frame["header"].(map[string]any)
	if header == nil {
		t.Fatalf("%s frame has no header (singular)", name)
	}
	if _, ok := header["X-BAPI-TIMESTAMP"]; !ok {
		t.Fatalf("%s header misses X-BAPI-TIMESTAMP: %v", name, header)
	}
	if _, ok := header["X-BAPI-RECV-WINDOW"]; !ok {
		t.Fatalf("%s header misses X-BAPI-RECV-WINDOW: %v", name, header)
	}
	args, _ := frame["args"].([]any)
	if len(args) != 1 {
		t.Fatalf("%s args len = %d, want exactly 1", name, len(args))
	}
}

// TestWsTradeResponseParse checks the ack parsing against the guideline
// schema: retCode/retMsg, reqId echo, data{orderId,orderLinkId} and the
// rate-limit header fields.
func TestWsTradeResponseParse(t *testing.T) {
	raw := `{"retCode":0,"retMsg":"OK","connId":"abc","reqId":"req-42","op":"order.amend",` +
		`"data":{"orderId":"oid-9","orderLinkId":"link-9"},` +
		`"header":{"X-Bapi-Limit":"600","X-Bapi-Limit-Status":"598","X-Bapi-Limit-Reset-Timestamp":"1718000000999","Timenow":"1718000000123"}}`
	var r WsTradeResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatal(err)
	}
	if !r.Ok() || r.Code != 0 || r.Message != "OK" {
		t.Fatalf("code/msg = %d/%q", r.Code, r.Message)
	}
	if r.Operation != "order.amend" || r.ConnectionId != "abc" {
		t.Fatalf("op/connId = %q/%q", r.Operation, r.ConnectionId)
	}
	if r.ReqId != "req-42" {
		t.Fatalf("reqId = %q, want req-42", r.ReqId)
	}
	d, err := r.OrderData()
	if err != nil {
		t.Fatal(err)
	}
	if d.OrderId != "oid-9" || d.OrderLinkId != "link-9" {
		t.Fatalf("order data = %+v", d)
	}
	if r.Header.Limit != "600" || r.Header.LimitStatus != "598" ||
		r.Header.LimitResetTimestamp != "1718000000999" || r.Header.TimeNow != "1718000000123" {
		t.Fatalf("header = %+v", r.Header)
	}

	var rejected WsTradeResponse
	if err := json.Unmarshal([]byte(`{"retCode":10001,"retMsg":"EC_PostOnlyWillTakeLiquidity","op":"order.create","reqId":"req-43"}`), &rejected); err != nil {
		t.Fatal(err)
	}
	if rejected.Ok() || rejected.Code != 10001 || rejected.ReqId != "req-43" {
		t.Fatalf("rejected = %+v", rejected)
	}
}
