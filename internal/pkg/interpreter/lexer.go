package interpreter

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	Text        []rune
	Pos         int
	CurrentChar rune
}

func NewLexer(text []rune) *Lexer {
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

func (l *Lexer) Peek() rune {
	peekPos := l.Pos + 1
	if peekPos > len(l.Text)-1 {
		return 0
	} else {
		return l.Text[peekPos]
	}
}

func (l *Lexer) Id() (*Token, error) {
	var result []rune
	for l.CurrentChar > 0 && (unicode.IsLetter(l.CurrentChar) || unicode.IsDigit(l.CurrentChar)) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	if v, ok := ReservedKeywords[strings.ToUpper(string(result))]; ok {
		return &v, nil
	}

	return &Token{TokenID, result}, nil
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

		return &Token{TokenREALCONST, f}, nil
	} else {
		i, err := strconv.Atoi(string(result))
		if err != nil {
			return nil, err
		}
		return &Token{TokenINTEGERCONST, i}, nil
	}
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
			return &Token{TokenASSIGN, ":="}, nil
		}
		if l.CurrentChar == ';' {
			l.Advance()
			return &Token{TokenSEMI, ';'}, nil
		}
		if l.CurrentChar == ':' {
			l.Advance()
			return &Token{TokenCOLON, ':'}, nil
		}
		if l.CurrentChar == ',' {
			l.Advance()
			return &Token{TokenCOMMA, ','}, nil
		}
		if l.CurrentChar == '+' {
			l.Advance()
			return &Token{TokenPLUS, '+'}, nil
		}
		if l.CurrentChar == '-' {
			l.Advance()
			return &Token{TokenMINUS, '-'}, nil
		}
		if l.CurrentChar == '*' {
			l.Advance()
			return &Token{TokenMULTIPLY, '*'}, nil
		}
		if l.CurrentChar == '/' {
			l.Advance()
			return &Token{TokenFLOATDIV, '/'}, nil
		}
		if l.CurrentChar == '(' {
			l.Advance()
			return &Token{TokenLPAREN, '('}, nil
		}
		if l.CurrentChar == ')' {
			l.Advance()
			return &Token{TokenRPAREN, ')'}, nil
		}
		if l.CurrentChar == '.' {
			l.Advance()
			return &Token{TokenDOT, '.'}, nil
		}
		return nil, errors.New("lexer error get next token")
	}

	return &Token{TokenEOF, nil}, nil
}
