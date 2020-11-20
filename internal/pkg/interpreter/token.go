package interpreter

import "fmt"

const (
	EOF      = "EOF"
	PLUS     = "PLUS"
	SUBTRACT = "SUBTRACT"
	INTEGER  = "INTEGER"
)

type Token struct {
	Type  string
	Value interface{}
}

func NewToken(t string, v interface{}) *Token {
	return &Token{Type: t, Value: v}
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", t.Type, t.Value)
}
