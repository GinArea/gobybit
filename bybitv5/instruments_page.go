package bybitv5

// InstrumentsPage is a single page of instruments-info that keeps the
// next-page cursor. The base getInstruments transform drops the cursor,
// which makes a full walk of the (long) linear instrument list impossible
// for tools that must see every symbol. Additive: existing
// GetInstruments* methods stay untouched.
type InstrumentsPage struct {
	List           []Instrument
	NextPageCursor string
}

// DoPage queries one page of instruments-info keeping the cursor. Chain
// calls by passing InstrumentsPage.NextPageCursor back as
// GetInstruments.Cursor until it comes back empty.
func (o GetInstruments) DoPage(c *Client) Response[InstrumentsPage] {
	return GetPub(c.market(), "instruments-info", o, func(r instrumentsResult[Instrument]) (InstrumentsPage, error) {
		return InstrumentsPage{List: r.List, NextPageCursor: r.NextPageCursor}, nil
	})
}

func (o *Client) GetInstrumentsPage(v GetInstruments) Response[InstrumentsPage] {
	return v.DoPage(o)
}
