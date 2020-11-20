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
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{Text: text, Pos: 0, CurrentToken: nil}
}

func (i *Interpreter) GetNextToken() (*Token, error) {
	text := i.Text

	if i.Pos > len(text)-1 {
		return NewToken(EOF, nil), nil
	}

	currentChar := text[i.Pos]

	if utils.IsDigit(text[i.Pos]) {
		index := i.Pos
		for ; index < len(text); index++ {
			if !utils.IsDigit(text[index]) {
				break
			}
		}

		num, err := strconv.Atoi(text[i.Pos : index])
		if err != nil {
			return nil, err
		}

		token := NewToken(INTEGER, num)
		i.Pos = index
		return token, nil
	}

	if currentChar == '+' {
		token := NewToken(PLUS, currentChar)
		i.Pos++
		return token, nil
	}

	if currentChar == '-' {
		token := NewToken(SUBTRACT, currentChar)
		i.Pos++
		return token, nil
	}

	if currentChar == ' ' {
		i.Pos++
		return i.GetNextToken()
	}

	return nil, ErrParsingInput
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

	if op.Type == SUBTRACT {
		result := left.Value.(int) - right.Value.(int)
		return result, nil
	}

	return 0, ErrParsingInput
}
