package transport

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type ParamTag struct {
	Name   string
	IsPtr  bool
	IsJson bool
	Ok     bool
}

func GetParamTag(f reflect.StructField) (tag ParamTag) {
	tag.IsPtr = f.Type.Kind() == reflect.Ptr
	tag.Name, tag.Ok = f.Tag.Lookup("param")
	if !tag.Ok {
		tag.Name, tag.Ok = f.Tag.Lookup("json")
		tag.IsJson = tag.Ok
	}
	return
}

type Param struct {
	IsJson     bool
	HeaderSign bool
	m          map[string]any
}

func NewParam() Param {
	return Param{m: make(map[string]any)}
}

func (o Param) From(v any) Param {
	if v == nil {
		return o
	}
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Struct {
		panic("url param from: object is not struct")
	}
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			if f.Type.Name() == reflect.TypeOf(HeaderSign{}).Name() {
				o.HeaderSign = true
			} else {
				for k, v := range NewParam().From(rv.Field(i).Interface()).m {
					o.Add(k, v)
				}
			}
			continue
		}
		tag := GetParamTag(rv.Type().Field(i))
		if tag.Ok {
			vl := rv.Field(i)
			if tag.IsPtr {
				if rv.Field(i).IsNil() {
					continue
				}
				vl = rv.Field(i).Elem()
			}
			o.Add(tag.Name, vl.Interface())
			o.IsJson = tag.IsJson
		}
	}
	return o
}

func (o *Param) Add(name string, val any) {
	o.m[name] = val
}

func (o Param) Make() url.Values {
	v := url.Values{}
	for name, val := range o.m {
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

type HeaderSign struct {
}
