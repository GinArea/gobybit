package bybitv5

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestClient builds a client pointed at the mock server. NewClient
// pins the /v5 path, so handlers see /v5/... paths like production.
func newTestClient(t *testing.T, srv *httptest.Server, auth bool) *Client {
	t.Helper()
	c := NewClient().WithBaseUrl(srv.URL)
	if auth {
		c.WithAuth("test-key", "test-secret")
	}
	return c
}

// TestFeeRateParse walks the real transport chain (signed GET through the
// mock) against the documented response schema: symbol rows with string
// rates; derivatives responses omit the category field.
func TestFeeRateParse(t *testing.T) {
	body := `{"retCode":0,"retMsg":"OK","result":{"list":[{"symbol":"BTCUSDT","takerFeeRate":"0.0015","makerFeeRate":"-0.0005","rpiMakerFeeRate":"-0.005"}]},"time":1787901023267}`
	var path, query string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, query = r.URL.Path, r.URL.RawQuery
		w.Write([]byte(body))
	}))
	defer srv.Close()

	c := newTestClient(t, srv, true)
	got := c.GetFeeRate(GetFeeRate{Category: Linear})
	if err := got.Error; err != nil {
		t.Fatalf("error: %v", err)
	}
	if path != "/v5/account/fee-rate" {
		t.Fatalf("request path: %s", path)
	}
	if !strings.Contains(query, "category=linear") {
		t.Fatalf("query missing category=linear: %s", query)
	}
	if len(got.Data.List) != 1 {
		t.Fatalf("list size: %d", len(got.Data.List))
	}
	f := got.Data.List[0]
	if f.Symbol != "BTCUSDT" || f.TakerFeeRate != "0.0015" || f.MakerFeeRate != "-0.0005" || f.RpiMakerFeeRate != "-0.005" {
		t.Fatalf("row mismatch: %+v", f)
	}
	if got.Data.Category != "" {
		t.Fatalf("derivatives category must stay empty, got %q", got.Data.Category)
	}
}

// TestFeeRateSymbolParam checks the per-symbol top-up form of the request.
func TestFeeRateSymbolParam(t *testing.T) {
	var query string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.RawQuery
		w.Write([]byte(`{"retCode":0,"retMsg":"OK","result":{"list":[{"symbol":"ICPUSDT","takerFeeRate":"0.0015","makerFeeRate":"-0.0006","rpiMakerFeeRate":""}]}}`))
	}))
	defer srv.Close()

	c := newTestClient(t, srv, true)
	got := c.GetFeeRate(GetFeeRate{Category: Linear, Symbol: "ICPUSDT"})
	if err := got.Error; err != nil {
		t.Fatalf("error: %v", err)
	}
	if !strings.Contains(query, "symbol=ICPUSDT") {
		t.Fatalf("query missing symbol: %s", query)
	}
	if len(got.Data.List) != 1 || got.Data.List[0].Symbol != "ICPUSDT" {
		t.Fatalf("row mismatch: %+v", got.Data.List)
	}
}

// TestFeeRateError maps a venue-level error (retCode != 0) to Response.Error.
func TestFeeRateError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"retCode":10001,"retMsg":"params error","result":{},"time":1}`))
	}))
	defer srv.Close()

	c := newTestClient(t, srv, true)
	got := c.GetFeeRate(GetFeeRate{Category: Linear})
	if got.Error == nil {
		t.Fatal("expected an error for retCode != 0")
	}
	if !strings.Contains(got.Error.Error(), "params error") {
		t.Fatalf("error text: %v", got.Error)
	}
}

// TestInstrumentsPageCursor checks that DoPage keeps NextPageCursor across
// pages — the property the base getInstruments loses.
func TestInstrumentsPageCursor(t *testing.T) {
	pages := map[string]string{
		"":      `{"retCode":0,"retMsg":"OK","result":{"category":"linear","list":[{"symbol":"BTCUSDT","contractType":"LinearPerpetual","settleCoin":"USDT"}],"nextPageCursor":"PAGE2"},"time":1}`,
		"PAGE2": `{"retCode":0,"retMsg":"OK","result":{"category":"linear","list":[{"symbol":"ETHUSDT","contractType":"LinearPerpetual","settleCoin":"USDT"}],"nextPageCursor":""},"time":2}`,
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(pages[r.URL.Query().Get("cursor")]))
	}))
	defer srv.Close()

	c := newTestClient(t, srv, false)
	p1 := c.GetInstrumentsPage(GetInstruments{Category: Linear, Status: StatusTrading, Limit: 1000})
	if err := p1.Error; err != nil {
		t.Fatalf("page1 error: %v", err)
	}
	if len(p1.Data.List) != 1 || p1.Data.List[0].Symbol != "BTCUSDT" {
		t.Fatalf("page1 list: %+v", p1.Data.List)
	}
	if p1.Data.NextPageCursor != "PAGE2" {
		t.Fatalf("page1 cursor: %q", p1.Data.NextPageCursor)
	}

	p2 := c.GetInstrumentsPage(GetInstruments{Category: Linear, Status: StatusTrading, Limit: 1000, Cursor: p1.Data.NextPageCursor})
	if err := p2.Error; err != nil {
		t.Fatalf("page2 error: %v", err)
	}
	if len(p2.Data.List) != 1 || p2.Data.List[0].Symbol != "ETHUSDT" {
		t.Fatalf("page2 list: %+v", p2.Data.List)
	}
	if p2.Data.NextPageCursor != "" {
		t.Fatalf("page2 cursor must be empty, got %q", p2.Data.NextPageCursor)
	}
}
