package interpreter

import "fmt"

type TokenType string

const (
	TokenEOF      = "EOF"
	TokenPLUS     = "PLUS"
	TokenMINUS    = "MINUS"
	TokenMULTIPLY = "MUL"
	TokenDIVIDE   = "DIV"
	TokenINTEGER  = "INTEGER"
	TokenLPAREN   = "("
	TokenRPAREN   = ")"
	TokenID       = "ID"
	TokenASSIGN   = "ASSIGN"
	TokenBEGIN    = "BEGIN"
	TokenEND      = "END"
	TokenSEMI     = "SEMI"
	TokenDOT      = "DOT"
)

type Token struct {
	Type  TokenType
	Value interface{}
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", t.Type, t.Value)
}

var ReservedKeywords = map[string]Token{
	"BEGIN": {Type: TokenBEGIN, Value: TokenBEGIN},
	"END":   {Type: TokenEND, Value: TokenEND},
}
