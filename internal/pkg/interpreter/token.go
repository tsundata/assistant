package interpreter

import "fmt"

type TokenType string

const (
	TokenPlus         = "+"
	TokenMinus        = "-"
	TokenMultiply     = "*"
	TokenFloatDiv     = "/"
	TokenLParen       = "("
	TokenRParen       = ")"
	TokenLSquare      = "["
	TokenRSquare      = "]"
	TokenLCurly       = "{"
	TokenRCurly       = "}"
	TokenSemi         = ";"
	TokenDot          = "."
	TokenColon        = ":"
	TokenComma        = ","
	TokenProgram      = "PROGRAM"
	TokenImport       = "IMPORT"
	TokenInteger      = "INTEGER"
	TokenFloat        = "FLOAT"
	TokenString       = "STRING"
	TokenBoolean      = "BOOLEAN"
	TokenList         = "LIST"
	TokenDict         = "DICT"
	TokenMessage      = "MESSAGE"
	TokenIntegerDiv   = "DIV"
	TokenVar          = "VAR"
	TokenFunction     = "FUNCTION"
	TokenBegin        = "BEGIN"
	TokenEnd          = "END"
	TokenReturn       = "RETURN"
	TokenPrint        = "PRINT"
	TokenID           = "ID"
	TokenIntegerConst = "INTEGER_CONST"
	TokenFloatConst   = "FLOAT_CONST"
	TokenStringConst  = "STRING_CONST"
	TokenMessageConst = "MESSAGE_CONST"
	TokenAssign       = ":="
	TokenEOF          = "EOF"
	TokenIf           = "IF"
	TokenThen         = "THEN"
	TokenElse         = "ELSE"
	TokenWhile        = "WHILE"
	TokenDo           = "DO"
	TokenOr           = "OR"
	TokenAnd          = "AND"
	TokenTrue         = "TRUE"
	TokenFalse        = "FALSE"
	TokenEqual        = "=="
	TokenNotEqual     = "!="
	TokenGreater      = ">"
	TokenGreaterEqual = ">="
	TokenLess         = "<"
	TokenLessEqual    = "<="
	TokenAt           = "@"
	TokenHash         = "#"
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
	"PROGRAM": {Type: TokenProgram, Value: TokenProgram},
	"IMPORT":  {Type: TokenImport, Value: TokenImport},
	"FUNC":    {Type: TokenFunction, Value: TokenFunction},
	"VAR":     {Type: TokenVar, Value: TokenVar},
	"DIV":     {Type: TokenIntegerDiv, Value: TokenIntegerDiv},
	"INT":     {Type: TokenInteger, Value: TokenInteger},
	"FLOAT":   {Type: TokenFloat, Value: TokenFloat},
	"STRING":  {Type: TokenString, Value: TokenString},
	"BOOL":    {Type: TokenBoolean, Value: TokenBoolean},
	"BEGIN":   {Type: TokenBegin, Value: TokenBegin},
	"END":     {Type: TokenEnd, Value: TokenEnd},
	"IF":      {Type: TokenIf, Value: TokenIf},
	"THEN":    {Type: TokenThen, Value: TokenThen},
	"ELSE":    {Type: TokenElse, Value: TokenElse},
	"WHILE":   {Type: TokenWhile, Value: TokenWhile},
	"DO":      {Type: TokenDo, Value: TokenDo},
	"OR":      {Type: TokenOr, Value: TokenOr},
	"AND":     {Type: TokenAnd, Value: TokenAnd},
	"TRUE":    {Type: TokenTrue, Value: true},
	"FALSE":   {Type: TokenFalse, Value: false},
	"PRINT":   {Type: TokenPrint, Value: TokenPrint},
	"RETURN":  {Type: TokenReturn, Value: TokenReturn},
	"LIST":    {Type: TokenList, Value: TokenList},
	"DICT":    {Type: TokenDict, Value: TokenDict},
	"MESSAGE": {Type: TokenMessage, Value: TokenMessage},
}
