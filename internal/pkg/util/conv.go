package util

import (
	"reflect"
	"time"
)

func BoolInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Inject(obj interface{}, data map[string]interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() != reflect.Ptr {
		return
	}

	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		db, ok := field.Tag.Lookup("db")
		if ok {
			if vv, ok := data[db]; ok {
				f := v.Elem().Field(i)
				if !f.IsValid() || !f.CanSet() {
					continue
				}
				switch field.Type.Kind() {
				case reflect.Int64:
					if num, ok := vv.(float64); ok {
						iNum := int64(num)
						if !f.OverflowInt(iNum) {
							f.SetInt(iNum)
						}
					}
					if _, ok := vv.(int64); ok {
						if !f.OverflowInt(vv.(int64)) {
							f.SetInt(vv.(int64))
						}
					}
				case reflect.String:
					if tt, ok := vv.(time.Time); ok {
						f.SetString(tt.Format("2006-01-02 15:04:05"))
					} else {
						f.SetString(vv.(string))
					}
				case reflect.Float64:
					if !f.OverflowFloat(vv.(float64)) {
						f.SetFloat(vv.(float64))
					}
				case reflect.Bool:
					if num, ok := vv.(int64); ok {
						if num > 0 {
							f.SetBool(true)
						}
					}
				}
			}
		}
	}
}
