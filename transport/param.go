package transport

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type ParamTag struct {
	Name  string
	IsPtr bool
	Ok    bool
}

func GetParamTag(f reflect.StructField) (tag ParamTag) {
	tag.IsPtr = f.Type.Kind() == reflect.Ptr
	tag.Name, tag.Ok = f.Tag.Lookup("param")
	return
}

type UrlParam map[string]any

func NewUrlParam() UrlParam {
	return make(map[string]any)
}

func (this UrlParam) From(v any) UrlParam {
	if v == nil {
		return this
	}
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Struct {
		panic("url param from: object is not struct")
	}
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			for k, v := range NewUrlParam().From(rv.Field(i).Interface()) {
				this.Add(k, v)
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
			this.Add(tag.Name, vl.Interface())
		}
	}
	return this
}

func (this UrlParam) Make() url.Values {
	v := url.Values{}
	for name, val := range this {
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

func (this *UrlParam) Add(name string, val any) {
	(*this)[name] = val
}
