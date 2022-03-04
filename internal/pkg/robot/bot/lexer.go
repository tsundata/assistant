package bot

import (
	"fmt"
	"github.com/pkg/errors"
	"unicode"
)

type Token struct {
	Type   string
	Value  string
	LineNo int
	Column int
}

const (
	CommandToken = "command"
	StringToken  = "string"
	ObjectToken  = "object"
	TagToken     = "tag"
	EOFToken     = "eof"
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
	return errors.New(fmt.Sprintf("Lexer error on '%s' line: %d column: %d", string(l.CurrentChar), l.LineNo, l.Column))
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

func (l *Lexer) Command() (*Token, error) {
	token := &Token{Type: "", Value: "", LineNo: l.LineNo, Column: l.Column}

	l.Advance()

	var result []rune
	for l.CurrentChar > 0 && l.CurrentChar != '/' {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	l.Advance()

	s := string(result)
	token.Type = CommandToken
	token.Value = s

	return token, nil
}

func (l *Lexer) String() (*Token, error) {
	token := &Token{Type: "", Value: "", LineNo: l.LineNo, Column: l.Column}

	var result []rune
	for l.CurrentChar > 0 && !unicode.IsSpace(l.CurrentChar) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	s := string(result)
	token.Type = StringToken
	token.Value = s

	return token, nil
}

func (l *Lexer) Object() (*Token, error) {
	token := &Token{Type: "", Value: "", LineNo: l.LineNo, Column: l.Column}

	l.Advance()

	var result []rune
	for l.CurrentChar > 0 && !unicode.IsSpace(l.CurrentChar) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	l.Advance()

	s := string(result)
	token.Type = ObjectToken
	token.Value = s

	return token, nil
}

func (l *Lexer) Tag() (*Token, error) {
	token := &Token{Type: "", Value: "", LineNo: l.LineNo, Column: l.Column}

	l.Advance()

	var result []rune
	for l.CurrentChar > 0 && !unicode.IsSpace(l.CurrentChar) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	l.Advance()

	s := string(result)
	token.Type = TagToken
	token.Value = s

	return token, nil
}

func (l *Lexer) GetNextToken() (*Token, error) {
	for l.CurrentChar > 0 {
		if unicode.IsSpace(l.CurrentChar) {
			l.SkipWhitespace()
			continue
		}
		if l.CurrentChar == '/' {
			return l.Command()
		}
		if l.CurrentChar == '@' {
			return l.Object()
		}
		if l.CurrentChar == '#' {
			return l.Tag()
		}
		if !unicode.IsSpace(l.CurrentChar) {
			return l.String()
		}

		return nil, l.error()
	}

	return &Token{Type: EOFToken, Value: ""}, nil
}

type Command struct {
	Main string
	Sub  string
	Args []string
}

func ParseText(in string) ([]*Token, error) {
	if in == "" {
		return []*Token{}, nil
	}
	l := NewLexer([]rune(in))
	var tokens []*Token
	token, err := l.GetNextToken()
	if err != nil {
		return nil, err
	}
	tokens = append(tokens, token)
	for token.Type != EOFToken {
		token, err = l.GetNextToken()
		if err != nil {
			return nil, err
		}
		if token.Type != EOFToken {
			tokens = append(tokens, token)
		}
	}

	return tokens, nil
}
