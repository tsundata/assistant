package interpreter

import "fmt"

type TokenType string

const (
	TokenPLUS         = "+"
	TokenMINUS        = "-"
	TokenMULTIPLY     = "*"
	TokenFLOATDIV     = "/"
	TokenLPAREN       = "("
	TokenRPAREN       = ")"
	TokenSEMI         = ";"
	TokenDOT          = "."
	TokenCOLON        = ":"
	TokenCOMMA        = ","
	TokenPROGRAM      = "PROGRAM"
	TokenINTEGER      = "INTEGER"
	TokenREAL         = "REAL"
	TokenINTEGERDIV   = "DIV"
	TokenVAR          = "VAR"
	TokenPROCEDURE    = "PROCEDURE"
	TokenBEGIN        = "BEGIN"
	TokenEND          = "END"
	TokenID           = "ID"
	TokenINTEGERCONST = "INTEGER_CONST"
	TokenREALCONST    = "REAL_CONST"
	TokenASSIGN       = ":="
	TokenEOF          = "EOF"
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
	"PROGRAM":   {Type: TokenPROGRAM, Value: TokenPROGRAM},
	"VAR":       {Type: TokenVAR, Value: TokenVAR},
	"DIV":       {Type: TokenINTEGERDIV, Value: TokenINTEGERDIV},
	"INTEGER":   {Type: TokenINTEGER, Value: TokenINTEGER},
	"REAL":      {Type: TokenREAL, Value: TokenREAL},
	"BEGIN":     {Type: TokenBEGIN, Value: TokenBEGIN},
	"END":       {Type: TokenEND, Value: TokenEND},
	"PROCEDURE": {Type: TokenPROGRAM, Value: TokenPROGRAM},
}
