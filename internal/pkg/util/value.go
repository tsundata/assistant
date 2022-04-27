package util

type Value struct {
	Source interface{}
}

func (v Value) String() (string, bool) {
	s, ok := v.Source.(string)
	return s, ok
}

func (v Value) Int64() (int64, bool) {
	s, ok := v.Source.(int64)
	return s, ok
}

func (v Value) Bool() (bool, bool) {
	s, ok := v.Source.(bool)
	return s, ok
}

func Variable(i interface{}) Value {
	return Value{Source: i}
}
