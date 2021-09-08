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

func Inject(target interface{}, data map[string]interface{}) {
	t := reflect.TypeOf(target)
	v := reflect.ValueOf(target)
	if t.Kind() != reflect.Ptr {
		return
	}

	tElem := t.Elem()
	vElem := v.Elem()
	for i := 0; i < tElem.NumField(); i++ {
		field := tElem.Field(i)
		if db, ok := field.Tag.Lookup("db"); ok {
			if vv, ok := data[db]; ok {
				f := vElem.Field(i)
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
						if vv != nil {
							f.SetString(vv.(string))
						} else {
							f.SetString("")
						}
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

func StringToBool(i string) bool {
	if i == "true" {
		return true
	}
	if i == "false" {
		return false
	}
	return false
}

func BoolToString(i bool) string {
	if i {
		return "true"
	} else {
		return "false"
	}
}
