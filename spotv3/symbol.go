package spotv3

func SymbolPtr(s string) *Symbol {
	v := Symbol(s)
	return &v
}
