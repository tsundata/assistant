package interpreter

import (
	"errors"
)

var ErrParsingInput = errors.New("error parsing input")

type Interpreter struct {
	Lexer        *Lexer
	CurrentToken *Token
}

func NewInterpreter(lexer *Lexer) (*Interpreter, error) {
	token, err := lexer.GetNextToken()
	if err != nil {
		return nil, err
	}
	return &Interpreter{Lexer: lexer, CurrentToken: token}, nil
}

func (i *Interpreter) Eat(tokenType string) (err error) {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken, err = i.Lexer.GetNextToken()
		return
	}

	return ErrParsingInput
}

func (i *Interpreter) Factor() (int, error) {
	token := i.CurrentToken
	err := i.Eat(INTEGER)
	if err != nil {
		return 0, nil
	}
	return token.Value.(int), nil
}

// expr   : factor ((MUL | DIV) factor)*
// factor : INTEGER
func (i *Interpreter) Expr() (int, error) {
	value, err := i.Factor()
	if err != nil {
		return 0, err
	}

	for i.CurrentToken.Type == MULTIPLY || i.CurrentToken.Type == DIVIDE {
		tokenType := i.CurrentToken.Type
		if tokenType == MULTIPLY {
			err = i.Eat(MULTIPLY)
			if err != nil {
				return 0, err
			}
			num, err := i.Factor()
			if err != nil {
				return 0, err
			}
			value *= num
		}
		if tokenType == DIVIDE {
			err = i.Eat(DIVIDE)
			if err != nil {
				return 0, err
			}
			num, err := i.Factor()
			if err != nil {
				return 0, err
			}
			value /= num
		}
	}

	return value, nil
}

func (i *Interpreter) Parse() (int, error) {
	return i.Expr()
}
