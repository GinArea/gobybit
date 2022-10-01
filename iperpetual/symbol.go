package iperpetual

func SymbolPtr(s string) *Symbol {
	v := Symbol(s)
	return &v
}
