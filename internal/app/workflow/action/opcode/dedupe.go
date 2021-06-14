package opcode

import (
	"fmt"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/util/collection"
	"reflect"
	"strconv"
)

type Dedupe struct{}

func NewDedupe() *Dedupe {
	return &Dedupe{}
}

func (o *Dedupe) Type() int {
	return TypeOp
}

func (o *Dedupe) Doc() string {
	return "dedupe [string]? : (any -> any)"
}

func (o *Dedupe) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) == 0 {
		v := reflect.ValueOf(ctx.Value)
		if v.Kind() == reflect.Slice {
			var result []interface{}
			m := make(map[interface{}]struct{})
			for i := 0; i < v.Len(); i++ {
				if _, ok := m[v.Index(i).Interface()]; !ok {
					m[v.Index(i).Interface()] = struct{}{}
					result = append(result, v.Index(i).Interface())
				}
			}
			ctx.SetValue(result)
			return result, nil
		}
	}
	if len(params) == 1 {
		if key, ok := params[0].(string); ok {
			bf := collection.NewBloomFilter(ctx.RDB, fmt.Sprintf("workflow:dedupe:%s", key), 100000, 7)

			if str, ok := ctx.Value.(string); ok {
				if !bf.Lookup(str) {
					bf.Add(str)
					ctx.SetValue(str)
					return str, nil
				} else {
					ctx.SetValue(nil)
					return nil, nil
				}
			}
			if num, ok := ctx.Value.(int64); ok {
				i := strconv.FormatInt(num, 10)
				if !bf.Lookup(i) {
					bf.Add(i)
					ctx.SetValue(num)
					return num, nil
				} else {
					ctx.SetValue(nil)
					return nil, nil
				}
			}
			if arrays, ok := ctx.Value.([]string); ok {
				var result []string
				for _, item := range arrays {
					if !bf.Lookup(item) {
						bf.Add(item)
						result = append(result, item)
					}
				}
				ctx.SetValue(result)
				return result, nil
			}
		}
	}

	return nil, nil
}
