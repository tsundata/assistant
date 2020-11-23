package interpreter

import "fmt"

type TokenType string

const (
	TokenEOF          = "EOF"
	TokenPLUS         = "PLUS"
	TokenMINUS        = "MINUS"
	TokenMULTIPLY     = "MUL"
	TokenDIVIDE       = "DIV"
	TokenINTEGER      = "INTEGER"
	TokenLPAREN       = "("
	TokenRPAREN       = ")"
	TokenID           = "ID"
	TokenASSIGN       = "ASSIGN"
	TokenBEGIN        = "BEGIN"
	TokenEND          = "END"
	TokenSEMI         = "SEMI"
	TokenDOT          = "DOT"
	TokenINTEGERCONST = "INTEGER_CONST"
	TokenREAL         = "REAL"
	TokenINTEGERDIV   = "INTEGER_DIV"
	TokenFLOATDIV     = "FLOAT_DIV"
	TokenREALCONST    = "REAL_CONST"
	TokenPROGRAM      = "PROGRAM"
	TokenCOMMA        = "COMMA"
	TokenCOLON        = "COLON"
	TokenVAR          = "VAR"
	TokenDIV          = "DIV"
)

type Token struct {
	Type  TokenType
	Value interface{}
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", t.Type, t.Value)
}

var ReservedKeywords = map[string]Token{
	"BEGIN":   {Type: TokenBEGIN, Value: TokenBEGIN},
	"END":     {Type: TokenEND, Value: TokenEND},
	"PROGRAM": {Type: TokenPROGRAM, Value: TokenPROGRAM},
	"VAR":     {Type: TokenVAR, Value: TokenVAR},
	"DIV":     {Type: TokenINTEGERDIV, Value: TokenDIV},
	"INTEGER": {Type: TokenINTEGER, Value: TokenINTEGER},
	"REAL":    {Type: TokenREAL, Value: TokenREAL},
}
