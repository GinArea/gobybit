package iperpetual

import (
	"reflect"
)

type Delta struct {
	Delete []any `json:"delete"`
	Update []any `json:"update"`
	Insert []any `json:"insert"`
}

func (o *Delta) HasData() bool {
	return len(o.Delete) > 0 || len(o.Update) > 0 || len(o.Insert) > 0
}

// Find sliсe item (struct) with required field of id for delta
func wsDeltaFindByID(slice []any, id uint64) (any, bool) {
	for _, v := range slice {
		if m, ok := v.(map[string]any); ok {
			if f, ok := m["id"]; ok {
				if i, ok := f.(float64); ok {
					if uint64(i) == id {
						return m, true
					}
				}
			}
		}
	}
	return nil, false
}

// Find sliсe item (struct) with required field of id for item of shapshot
func wsShotFindByID(rv reflect.Value, id uint64) (reflect.Value, bool) {
	for i := 0; i < rv.Len(); i++ {
		item := rv.Index(i)
		if item.FieldByName("ID").Uint() == id {
			return item, true
		}
	}
	return reflect.Value{}, false
}

// Set delta value by name in struct (impl)
func wsDeltaSetValue(vs reflect.Value, name string, v any) {
	for i := 0; i < vs.NumField(); i++ {
		label, ok := vs.Type().Field(i).Tag.Lookup("json")
		if ok && name == label {
			f := vs.Field(i)
			vv := reflect.ValueOf(v)
			if f.CanSet() && vv.CanConvert(f.Type()) {
				f.Set(vv.Convert(f.Type()))
			}
		}
	}
}

// Apply delta
func WsDeltaApply[T any](v *T, delta Delta) {
	rv := reflect.ValueOf(v).Elem()
	if rv.Kind() == reflect.Slice {
		if len(delta.Delete) > 0 {
			slice := reflect.MakeSlice(rv.Type(), 0, rv.Cap())
			for i := 0; i < rv.Len(); i++ {
				item := rv.Index(i)
				if _, ok := wsDeltaFindByID(delta.Delete, item.FieldByName("ID").Uint()); !ok {
					slice = reflect.Append(slice, item)
				}
			}
			rv.Set(slice)
		}
		for _, k := range delta.Insert {
			if m, ok := k.(map[string]any); ok {
				item := reflect.New(rv.Type().Elem()).Elem()
				for name, value := range m {
					wsDeltaSetValue(item, name, value)
				}
				slice := reflect.Append(rv, item)
				rv.Set(slice)
			}
		}
		for _, k := range delta.Update {
			if m, ok := k.(map[string]any); ok {
				if f, ok := m["id"]; ok {
					if i, ok := f.(float64); ok {
						if item, ok := wsShotFindByID(rv, uint64(i)); ok {
							for name, value := range m {
								wsDeltaSetValue(item, name, value)
							}
						}
					}
				}
			}
		}
	} else {
		WsDeltaUpdate(v, delta)
	}
}

// Apply delta to struct
func WsDeltaUpdate[T any](s *T, delta Delta) {
	for _, k := range delta.Update {
		if m, ok := k.(map[string]any); ok {
			for name, value := range m {
				WsDeltaSetValue(s, name, value)
			}
		}
	}
}

// Set delta value by name in struct
func WsDeltaSetValue[T any](s *T, name string, v any) {
	wsDeltaSetValue(reflect.ValueOf(s).Elem(), name, v)
}
