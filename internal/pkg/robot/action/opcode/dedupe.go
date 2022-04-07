package opcode

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
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

func (o *Dedupe) Run(_ context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) == 0 {
		v := reflect.ValueOf(comp.Value)
		if v.Kind() == reflect.Slice {
			var result []interface{}
			m := make(map[interface{}]struct{})
			for i := 0; i < v.Len(); i++ {
				if _, ok := m[v.Index(i).Interface()]; !ok {
					m[v.Index(i).Interface()] = struct{}{}
					result = append(result, v.Index(i).Interface())
				}
			}
			comp.SetValue(result)
			return result, nil
		}
	}
	if len(params) == 1 {
		if key, ok := params[0].(string); ok {
			bf := collection.NewBloomFilter(comp.Rdb, fmt.Sprintf("chatbot:opcode:dedupe:%s", key), 100000, 7)

			if str, ok := comp.Value.(string); ok {
				if !bf.Lookup(str) {
					bf.Add(str)
					comp.SetValue(str)
					return str, nil
				} else {
					comp.SetValue(nil)
					return nil, nil
				}
			}
			if num, ok := comp.Value.(int64); ok {
				i := strconv.FormatInt(num, 10)
				if !bf.Lookup(i) {
					bf.Add(i)
					comp.SetValue(num)
					return num, nil
				} else {
					comp.SetValue(nil)
					return nil, nil
				}
			}
			if arrays, ok := comp.Value.([]string); ok {
				var result []string
				for _, item := range arrays {
					if !bf.Lookup(item) {
						bf.Add(item)
						result = append(result, item)
					}
				}
				comp.SetValue(result)
				return result, nil
			}
		}
	}

	return nil, nil
}
