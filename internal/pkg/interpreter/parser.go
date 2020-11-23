package interpreter

import "errors"

var ErrParser = errors.New("parser error")

type Ast struct {}

type BinOp struct {
	*Ast
	Left  interface{}
	Token *Token
	Op    *Token
	Right interface{}
}

func NewBinOp(left interface{}, op *Token, right interface{}) *BinOp {
	return &BinOp{Left: left, Token: op, Op: op, Right: right}
}

type Num struct {
	*Ast
	Token *Token
	Value interface{}
}

func NewNum(token *Token) *Num {
	return &Num{Token: token, Value: token.Value}
}

type Parser struct {
	Lexer        *Lexer
	CurrentToken *Token
}

func NewParser(lexer *Lexer) (*Parser, error) {
	token, err := lexer.GetNextToken()
	if err != nil {
		return nil, err
	}
	return &Parser{Lexer: lexer, CurrentToken: token}, nil
}

func (p *Parser) Eat(tokenType string) (err error) {
	if p.CurrentToken.Type == tokenType {
		p.CurrentToken, err = p.Lexer.GetNextToken()
		return
	}

	return ErrInterpreter
}

func (p *Parser) Factor() (interface{}, error) {
	token := p.CurrentToken
	if token.Type == INTEGER {
		err := p.Eat(INTEGER)
		if err != nil {
			return nil, err
		}
		return NewNum(token), nil
	} else if token.Type == LPAREN {
		err := p.Eat(LPAREN)
		if err != nil {
			return nil, err
		}
		node, err := p.Expr()
		if err != nil {
			return nil, err
		}
		err = p.Eat(RPAREN)
		if err != nil {
			return nil, err
		}
		return node, nil
	}

	return nil, ErrParser
}

func (p *Parser) Term() (interface{}, error) {
	node, err := p.Factor()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == MULTIPLY || p.CurrentToken.Type == DIVIDE {
		token := p.CurrentToken
		if token.Type == MULTIPLY {
			err = p.Eat(MULTIPLY)
			if err != nil {
				return nil, err
			}
		}
		if token.Type == DIVIDE {
			err = p.Eat(DIVIDE)
			if err != nil {
				return nil, err
			}
		}

		right, err := p.Factor()
		if err != nil {
			return nil, err
		}
		node = NewBinOp(node, token, right)
	}

	return node, nil
}

func (p *Parser) Expr() (interface{}, error) {
	node, err := p.Term()
	if err != nil {
		return 0, err
	}

	for p.CurrentToken.Type == PLUS || p.CurrentToken.Type == MINUS {
		token := p.CurrentToken
		if token.Type == PLUS {
			err = p.Eat(PLUS)
			if err != nil {
				return nil, err
			}
		}
		if token.Type == MINUS {
			err = p.Eat(MINUS)
			if err != nil {
				return nil, err
			}
		}

		right, err := p.Term()
		if err != nil {
			return nil, err
		}
		node = NewBinOp(node, token, right)
	}

	return node, nil
}

// expr   : term   ((PLUS | MINUS) term)*
// term   : factor ((MUL | DIV) factor)*
// factor : INTEGER | LPAREN expr RPAREN
func (p *Parser) Parse() (interface{}, error) {
	return p.Expr()
}
