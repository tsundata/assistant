package interpreter

import "fmt"

const (
	EOF      = "EOF"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	MULTIPLY = "MULTIPLY"
	DIVIDE   = "DIVIDE"
	INTEGER  = "INTEGER"
	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
)

type Token struct {
	Type  string
	Value interface{}
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", t.Type, t.Value)
}
