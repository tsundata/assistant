package interpreter

import (
	"errors"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"strconv"
)

var ErrParsingInput = errors.New("error parsing input")

type Interpreter struct {
	Text         string
	Pos          int
	CurrentToken *Token
	CurrentChar  byte
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{Text: text, Pos: 0, CurrentToken: nil, CurrentChar: text[0]}
}

func (i *Interpreter) Advance() {
	i.Pos++
	if i.Pos >= len(i.Text) {
		i.CurrentChar = 0
	} else {
		i.CurrentChar = i.Text[i.Pos]
	}
}

func (i *Interpreter) SkipWhitespace() {
	for i.CurrentChar > 0 && i.CurrentChar == ' ' {
		i.Advance()
	}
}

func (i *Interpreter) Integer() int {
	var result []byte
	for i.CurrentChar > 0 && utils.IsDigit(i.CurrentChar) {
		result = append(result, i.CurrentChar)
		i.Advance()
	}
	num, _ := strconv.Atoi(string(result))
	return num
}

func (i *Interpreter) GetNextToken() (*Token, error) {
	for i.CurrentChar > 0 {
		if i.CurrentChar == ' ' {
			i.SkipWhitespace()
			continue
		}
		if utils.IsDigit(i.CurrentChar) {
			return NewToken(INTEGER, i.Integer()), nil
		}
		if i.CurrentChar == '+' {
			i.Advance()
			return NewToken(PLUS, '+'), nil
		}
		if i.CurrentChar == '-' {
			i.Advance()
			return NewToken(MINUS, '-'), nil
		}
		return nil, ErrParsingInput
	}

	return NewToken(EOF, nil), nil
}

func (i *Interpreter) Eat(tokenType string) (err error) {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken, err = i.GetNextToken()
		return
	}

	return ErrParsingInput
}

func (i Interpreter) Expr() (int, error) {
	var err error
	i.CurrentToken, err = i.GetNextToken()
	if err != nil {
		return 0, err
	}

	left := i.CurrentToken
	err = i.Eat(INTEGER)
	if err != nil {
		return 0, err
	}

	op := i.CurrentToken
	err = i.Eat(op.Type)
	if err != nil {
		return 0, err
	}

	right := i.CurrentToken
	err = i.Eat(INTEGER)
	if err != nil {
		return 0, err
	}

	if op.Type == PLUS {
		result := left.Value.(int) + right.Value.(int)
		return result, nil
	}

	if op.Type == MINUS {
		result := left.Value.(int) - right.Value.(int)
		return result, nil
	}

	return 0, ErrParsingInput
}
