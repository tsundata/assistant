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
	for l.CurrentChar != '\n' && l.Pos != len(l.Text) {
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

		token.Type = TokenFloatConst
		token.Value = f
	} else {
		i, err := strconv.Atoi(string(result))
		if err != nil {
			return nil, err
		}

		token.Type = TokenIntegerConst
		token.Value = i
	}

	return token, nil
}

func (l *Lexer) String() (*Token, error) {
	l.Advance()

	// TODO Escape
	var result []rune
	for l.CurrentChar != '"' {
		result = append(result, l.CurrentChar)
		l.Advance()
	}
	l.Advance()

	return &Token{Type: TokenStringConst, Value: string(result), LineNo: l.LineNo, Column: l.Column}, nil
}

func (l *Lexer) Message() (*Token, error) {
	l.Advance()

	var result []rune
	for l.CurrentChar > 0 && unicode.IsDigit(l.CurrentChar) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	i, err := strconv.Atoi(string(result))
	if err != nil {
		return nil, err
	}

	return &Token{Type: TokenMessageConst, Value: i, LineNo: l.LineNo, Column: l.Column}, nil
}

func (l *Lexer) Id() (*Token, error) {
	token := &Token{Type: "", Value: nil, LineNo: l.LineNo, Column: l.Column}

	var result []rune
	for l.CurrentChar > 0 && (unicode.IsLetter(l.CurrentChar) || unicode.IsDigit(l.CurrentChar)) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	s := string(result)
	if v, ok := ReservedKeywords[strings.ToUpper(s)]; ok {
		token.Type = v.Type
		token.Value = v.Value
	} else {
		token.Type = TokenID
		token.Value = s
	}

	return token, nil
}

func (l *Lexer) GetNextToken() (*Token, error) {
	for l.CurrentChar > 0 {
		if unicode.IsSpace(l.CurrentChar) {
			l.SkipWhitespace()
			continue
		}
		if l.CurrentChar == '/' && l.Peek() == '/' {
			l.Advance()
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
		if l.CurrentChar == '"' {
			return l.String()
		}
		if l.CurrentChar == '#' {
			return l.Message()
		}
		if l.CurrentChar == ':' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenAssign, Value: TokenAssign, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '=' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenEqual, Value: TokenEqual, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '!' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenNotEqual, Value: TokenNotEqual, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '>' {
			l.Advance()
			return &Token{Type: TokenGreater, Value: TokenGreater, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '>' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenGreaterEqual, Value: TokenNotEqual, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '<' {
			l.Advance()
			return &Token{Type: TokenLess, Value: TokenLess, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '<' && l.Peek() == '=' {
			l.Advance()
			l.Advance()
			return &Token{Type: TokenLessEqual, Value: TokenLessEqual, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ';' {
			l.Advance()
			return &Token{Type: TokenSemi, Value: TokenSemi, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ':' {
			l.Advance()
			return &Token{Type: TokenColon, Value: TokenColon, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ',' {
			l.Advance()
			return &Token{Type: TokenComma, Value: TokenComma, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '+' {
			l.Advance()
			return &Token{Type: TokenPlus, Value: TokenPlus, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '-' {
			l.Advance()
			return &Token{Type: TokenMinus, Value: TokenMinus, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '*' {
			l.Advance()
			return &Token{Type: TokenMultiply, Value: TokenMultiply, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '/' {
			l.Advance()
			return &Token{Type: TokenFloatDiv, Value: TokenFloatDiv, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '(' {
			l.Advance()
			return &Token{Type: TokenLParen, Value: TokenLParen, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ')' {
			l.Advance()
			return &Token{Type: TokenRParen, Value: TokenRParen, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '.' {
			l.Advance()
			return &Token{Type: TokenDot, Value: TokenDot, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '[' {
			l.Advance()
			return &Token{Type: TokenLSquare, Value: TokenLSquare, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == ']' {
			l.Advance()
			return &Token{Type: TokenRSquare, Value: TokenRSquare, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '{' {
			l.Advance()
			return &Token{Type: TokenLCurly, Value: TokenLCurly, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '}' {
			l.Advance()
			return &Token{Type: TokenRCurly, Value: TokenRCurly, LineNo: l.LineNo, Column: l.Column}, nil
		}
		if l.CurrentChar == '@' {
			l.Advance()
			return &Token{Type: TokenAt, Value: TokenAt, LineNo: l.LineNo, Column: l.Column}, nil
		}
		return nil, l.error()
	}

	return &Token{Type: TokenEOF, Value: nil}, nil
}
