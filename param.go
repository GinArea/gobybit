package bybit

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type UrlParam map[string]any

func (this *UrlParam) Add(name string, val any) {
	if this == nil {
		*this = make(map[string]any)
	}
	(*this)[name] = val
}

func (this *UrlParam) Make() url.Values {
	v := url.Values{}
	for name, val := range *this {
		if reflect.TypeOf(val).Kind() == reflect.Slice {
			l := []string{}
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				l = append(l, fmt.Sprint(s.Index(i)))
			}
			v.Add(name, strings.Join(l, ","))
		} else {
			v.Add(name, fmt.Sprint(val))
		}
	}
	return v
}
