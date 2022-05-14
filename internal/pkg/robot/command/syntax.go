package command

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/util"
	"regexp"
	"strconv"
	"unicode"
)

type Syntax struct {
	Text        []rune
	Pos         int
	CurrentChar rune
	LineNo      int
	Column      int
}

func NewSyntax(text []rune) *Syntax {
	return &Syntax{Text: text, Pos: 0, CurrentChar: text[0], LineNo: 1, Column: 1}
}

func (l *Syntax) error() error {
	return errors.New(fmt.Sprintf("Syntax error on '%s' line: %d column: %d", string(l.CurrentChar), l.LineNo, l.Column))
}

func (l *Syntax) Advance() {
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

func (l *Syntax) Peek() rune {
	peekPos := l.Pos + 1
	if peekPos > len(l.Text)-1 {
		return 0
	} else {
		return l.Text[peekPos]
	}
}

func (l *Syntax) SkipWhitespace() {
	for l.CurrentChar > 0 && unicode.IsSpace(l.CurrentChar) {
		l.Advance()
	}
}

func (l *Syntax) Parameter() (*Token, error) {
	token := &Token{Type: "", Value: util.Variable(""), LineNo: l.LineNo, Column: l.Column}

	var result []rune
	if l.CurrentChar == '[' {
		l.Advance()
		result = append(result, l.CurrentChar)

		l.Advance()
		for l.CurrentChar > 0 && l.CurrentChar != ']' {
			result = append(result, l.CurrentChar)
			l.Advance()
		}
		l.Advance()

		token.Type = ParameterToken
		token.Value = util.Variable(string(result))
	}

	return token, nil
}

func (l *Syntax) Character() (*Token, error) {
	token := &Token{Type: "", Value: util.Variable(""), LineNo: l.LineNo, Column: l.Column}

	var result []rune
	for l.CurrentChar > 0 && !unicode.IsSpace(l.CurrentChar) {
		result = append(result, l.CurrentChar)
		l.Advance()
	}

	s := string(result)
	token.Type = CharacterToken
	token.Value = util.Variable(s)

	return token, nil
}

func (l *Syntax) GetNextToken() (*Token, error) {
	for l.CurrentChar > 0 {
		if unicode.IsSpace(l.CurrentChar) {
			l.SkipWhitespace()
			continue
		}
		if l.CurrentChar == '[' {
			return l.Parameter()
		}
		if !unicode.IsSpace(l.CurrentChar) {
			return l.Character()
		}

		return nil, l.error()
	}

	return &Token{Type: EOFToken, Value: util.Variable("")}, nil
}

func SyntaxCheck(define string, actual []*Token) (bool, error) {
	s := NewSyntax([]rune(define))
	var tokens []*Token
	token, err := s.GetNextToken()
	if err != nil {
		return false, err
	}
	tokens = append(tokens, token)
	for token.Type != EOFToken {
		token, err = s.GetNextToken()
		if err != nil {
			return false, err
		}
		if token.Type != EOFToken {
			tokens = append(tokens, token)
		}
	}

	if len(tokens) != len(actual) {
		return false, nil
	}

	res := true
	for i, t := range tokens {
		if t.Type == CharacterToken {
			if t.Value != actual[i].Value {
				res = false
				continue
			}
		}
		if t.Type == ParameterToken {
			switch t.Value.Source {
			case "number":
				n, _ := actual[i].Value.String()
				re := regexp.MustCompile(`\d+`)
				if !re.MatchString(n) {
					res = false
					continue
				}
				num, err := strconv.ParseInt(n, 10, 64)
				if err == nil {
					actual[i].Value = util.Variable(num)
				}
			case "bool":
				if !(actual[i].Value.Source == "true" || actual[i].Value.Source == "false") {
					res = false
					continue
				}
				if actual[i].Value.Source == "true" {
					actual[i].Value = util.Variable(true)
				}
				if actual[i].Value.Source == "false" {
					actual[i].Value = util.Variable(false)
				}
			case "string":
			case "any":
			}
		}
	}
	return res, nil
}
