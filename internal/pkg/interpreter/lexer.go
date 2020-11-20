package interpreter

import (
	"github.com/tsundata/assistant/internal/pkg/utils"
	"strconv"
)

type Lexer struct {
	Text        string
	Pos         int
	CurrentChar byte
}

func NewLexer(text string) *Lexer {
	return &Lexer{Text: text, Pos: 0, CurrentChar: text[0]}
}

func (l *Lexer) Advance() {
	l.Pos++
	if l.Pos > len(l.Text)-1 {
		l.CurrentChar = 0
	} else {
		l.CurrentChar = l.Text[l.Pos]
	}
}

func (l *Lexer) SkipWhitespace() {
	for l.CurrentChar > 0 && l.CurrentChar == ' ' {
		l.Advance()
	}
}

func (l *Lexer) Integer() int {
	var result []byte
	for l.CurrentChar > 0 && utils.IsDigit(l.CurrentChar) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}
	num, _ := strconv.Atoi(string(result))
	return num
}

func (l *Lexer) GetNextToken() (*Token, error) {
	for l.CurrentChar > 0 {
		if l.CurrentChar == ' ' {
			l.SkipWhitespace()
			continue
		}
		if utils.IsDigit(l.CurrentChar) {
			return NewToken(INTEGER, l.Integer()), nil
		}
		if l.CurrentChar == '*' {
			l.Advance()
			return NewToken(MULTIPLY, '*'), nil
		}
		if l.CurrentChar == '/' {
			l.Advance()
			return NewToken(DIVIDE, '/'), nil
		}
		return nil, ErrParsingInput
	}

	return NewToken(EOF, nil), nil
}
