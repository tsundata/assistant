package interpreter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	Text        []rune
	Pos         int
	CurrentChar rune
	LineNo      int
	Column      int
}

func NewLexer(text []rune) *Lexer {
	return &Lexer{Text: text, Pos: 0, CurrentChar: text[0], LineNo: 1, Column: 1}
}

func (l *Lexer) error() error {
	return Error{
		Message: fmt.Sprintf("Lexer error on '%s' line: %d column: %d", string(l.CurrentChar), l.LineNo, l.Column),
		Type:    LexerErrorType,
	}
}

func (l *Lexer) Advance() {
	if l.CurrentChar == '\n' {
		l.LineNo += 1
		l.Column = 0
	}
	l.Pos++
	if l.Pos > len(l.Text)-1 {
		l.CurrentChar = 0
	} else {
		l.CurrentChar = l.Text[l.Pos]
		l.Column += 1
	}
}

func (l *Lexer) Peek() rune {
	peekPos := l.Pos + 1
	if peekPos > len(l.Text)-1 {
		return 0
	} else {
		return l.Text[peekPos]
	}
}

func (l *Lexer) SkipWhitespace() {
	for l.CurrentChar > 0 && unicode.IsSpace(l.CurrentChar) {
		l.Advance()
	}
}

func (l *Lexer) SkipComment() {
	for l.CurrentChar != '}' {
		l.Advance()
	}
	l.Advance()
}

func (l *Lexer) Number() (*Token, error) {
	token := &Token{Type: "", Value: nil, LineNo: l.LineNo, Column: l.Column}

	var result []rune
	for l.CurrentChar > 0 && unicode.IsDigit(l.CurrentChar) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	if l.CurrentChar == '.' {
		result = append(result, l.CurrentChar)
		l.Advance()

		for l.CurrentChar > 0 && unicode.IsDigit(l.CurrentChar) {
			result = append(result, l.CurrentChar)
			l.Advance()
		}

		f, err := strconv.ParseFloat(string(result), 64)
		if err != nil {
			return nil, err
		}

		token.Type = TokenREALCONST
		token.Value = f
	} else {
		i, err := strconv.Atoi(string(result))
		if err != nil {
			return nil, err
		}

		token.Type = TokenINTEGERCONST
		token.Value = i
	}

	return token, nil
}

func (l *Lexer) Id() (*Token, error) {
	token := &Token{Type: "", Value: nil, LineNo: l.LineNo, Column: l.Column}

	var result []rune
	for l.CurrentChar > 0 && (unicode.IsLetter(l.CurrentChar) || unicode.IsDigit(l.CurrentChar)) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	if v, ok := ReservedKeywords[strings.ToUpper(string(result))]; ok {
		token.Type = v.Type
		token.Value = strings.ToUpper(string(result))
	} else {
		token.Type = TokenID
		token.Value = string(result)
	}

	return token, nil
}

func (l *Lexer) GetNextToken() (*Token, error) {
	for l.CurrentChar > 0 {
		if unicode.IsSpace(l.CurrentChar) {
			l.SkipWhitespace()
			continue
		}
		if l.CurrentChar == '{' {
			l.Advance()
			l.SkipComment()
			continue
		}
		if unicode.IsDigit(l.CurrentChar) {
			return l.Number()
		}
		if unicode.IsLetter(l.CurrentChar) {
			return l.Id()
		}
		if l.CurrentChar == ':' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenASSIGN, Value: TokenASSIGN, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '=' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenEQUAL, Value: TokenEQUAL, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '!' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenNOTEQUAL, Value: TokenNOTEQUAL, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '>' {
			l.Advance()
			return &Token{Type: TokenGREATER, Value: TokenGREATER, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '>' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenGREATEREQUAL, Value: TokenNOTEQUAL, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '<' {
			l.Advance()
			return &Token{Type: TokenLESS, Value: TokenLESS, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '<' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenLESSEQUAL, Value: TokenLESSEQUAL, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ';' {
			l.Advance()
			return &Token{Type: TokenSEMI, Value: TokenSEMI, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ':' {
			l.Advance()
			return &Token{Type: TokenCOLON, Value: TokenCOLON, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ',' {
			l.Advance()
			return &Token{Type: TokenCOMMA, Value: TokenCOMMA, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '+' {
			l.Advance()
			return &Token{Type: TokenPLUS, Value: TokenPLUS, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '-' {
			l.Advance()
			return &Token{Type: TokenMINUS, Value: TokenMINUS, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '*' {
			l.Advance()
			return &Token{Type: TokenMULTIPLY, Value: TokenMULTIPLY, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '/' {
			l.Advance()
			return &Token{Type: TokenFLOATDIV, Value: TokenFLOATDIV, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '(' {
			l.Advance()
			return &Token{Type: TokenLPAREN, Value: TokenLPAREN, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ')' {
			l.Advance()
			return &Token{Type: TokenRPAREN, Value: TokenRPAREN, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '.' {
			l.Advance()
			return &Token{Type: TokenDOT, Value: TokenDOT, LineNo: l.LineNo, Column: l.Column}, nil
		}
		return nil, l.error()
	}

	return &Token{Type: TokenEOF, Value: nil}, nil
}
