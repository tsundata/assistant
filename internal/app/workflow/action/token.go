package action

import "fmt"

type TokenType string

const (
	TokenPlus           = "+"
	TokenMinus          = "-"
	TokenMultiply       = "*"
	TokenFloatDiv       = "/"
	TokenLParen         = "("
	TokenRParen         = ")"
	TokenLSquare        = "["
	TokenRSquare        = "]"
	TokenLCurly         = "{"
	TokenRCurly         = "}"
	TokenSemi           = ";"
	TokenDot            = "."
	TokenColon          = ":"
	TokenComma          = ","
	TokenInteger        = "INTEGER"
	TokenFloat          = "FLOAT"
	TokenString         = "STRING"
	TokenBoolean        = "BOOLEAN"
	TokenMessage        = "MESSAGE"
	TokenID             = "ID"
	TokenIntegerConst   = "INTEGER_CONST"
	TokenFloatConst     = "FLOAT_CONST"
	TokenStringConst    = "STRING_CONST"
	TokenMessageConst   = "MESSAGE_CONST"
	TokenTrue           = "TRUE"
	TokenFalse          = "FALSE"
	TokenCarriageReturn = "\n"
	TokenEOF            = "EOF"
)

type Token struct {
	Type   TokenType
	Value  interface{}
	LineNo int
	Column int
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %v, position=%d:%d)", t.Type, t.Value, t.LineNo, t.Column)
}

var ReservedKeywords = map[string]Token{
	"INT":     {Type: TokenInteger, Value: TokenInteger},
	"FLOAT":   {Type: TokenFloat, Value: TokenFloat},
	"STRING":  {Type: TokenString, Value: TokenString},
	"BOOL":    {Type: TokenBoolean, Value: TokenBoolean},
	"TRUE":    {Type: TokenTrue, Value: true},
	"FALSE":   {Type: TokenFalse, Value: false},
	"MESSAGE": {Type: TokenMessage, Value: TokenMessage},
}
