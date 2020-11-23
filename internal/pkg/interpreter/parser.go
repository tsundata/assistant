package interpreter

import "errors"

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

func (p *Parser) Eat(tokenType TokenType) (err error) {
	if p.CurrentToken.Type == tokenType {
		p.CurrentToken, err = p.Lexer.GetNextToken()
		return
	}

	return errors.New("parser error eat")
}

func (p *Parser) Program() (interface{}, error) {
	node, err := p.CompoundStatement()
	if err != nil {
		return nil, err
	}
	err = p.Eat(TokenDOT)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (p *Parser) CompoundStatement() (interface{}, error) {
	err := p.Eat(TokenBEGIN)
	if err != nil {
		return nil, err
	}
	nodes, err := p.StatementList()
	err = p.Eat(TokenEND)
	if err != nil {
		return nil, err
	}

	root := NewCompound()
	for _, node := range nodes {
		root.Children = append(root.Children, node)
	}
	return root, nil
}

func (p *Parser) StatementList() ([]interface{}, error) {
	node, err := p.Statement()
	if err != nil {
		return nil, err
	}

	results := []interface{}{node}

	for p.CurrentToken.Type == TokenSEMI {
		err = p.Eat(TokenSEMI)
		if err != nil {
			return nil, err
		}
		i, err := p.Statement()
		if err != nil {
			return nil, err
		}
		results = append(results, i)
	}

	if p.CurrentToken.Type == TokenID {
		return nil, errors.New("parser error statement list: id")
	}

	return results, nil
}

func (p *Parser) Statement() (interface{}, error) {
	if p.CurrentToken.Type == TokenBEGIN {
		return p.CompoundStatement()
	} else if p.CurrentToken.Type == TokenID {
		return p.AssignmentStatement()
	} else {
		return p.Empty()
	}
}

func (p *Parser) AssignmentStatement() (interface{}, error) {
	left, err := p.Variable()
	if err != nil {
		return nil, err
	}
	token := p.CurrentToken
	err = p.Eat(TokenASSIGN)
	if err != nil {
		return nil, err
	}
	right, err := p.Expr()

	return NewAssign(left, token, right), nil
}

func (p *Parser) Variable() (interface{}, error) {
	node := NewVar(p.CurrentToken)
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (p *Parser) Empty() (interface{}, error) {
	return NewNoOp(), nil
}

func (p *Parser) Expr() (interface{}, error) {
	node, err := p.Term()
	if err != nil {
		return 0, err
	}

	for p.CurrentToken.Type == TokenPLUS || p.CurrentToken.Type == TokenMINUS {
		token := p.CurrentToken
		if token.Type == TokenPLUS {
			err = p.Eat(TokenPLUS)
			if err != nil {
				return nil, err
			}
		}
		if token.Type == TokenMINUS {
			err = p.Eat(TokenMINUS)
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

func (p *Parser) Term() (interface{}, error) {
	node, err := p.Factor()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenMULTIPLY || p.CurrentToken.Type == TokenDIVIDE {
		token := p.CurrentToken
		if token.Type == TokenMULTIPLY {
			err = p.Eat(TokenMULTIPLY)
			if err != nil {
				return nil, err
			}
		}
		if token.Type == TokenDIVIDE {
			err = p.Eat(TokenDIVIDE)
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

func (p *Parser) Factor() (interface{}, error) {
	token := p.CurrentToken
	if token.Type == TokenPLUS {
		err := p.Eat(TokenPLUS)
		if err != nil {
			return nil, err
		}
		i, err := p.Factor()
		if err != nil {
			return nil, err
		}
		node := NewUnaryOp(token, i)
		return node, nil
	}
	if token.Type == TokenMINUS {
		err := p.Eat(TokenMINUS)
		if err != nil {
			return nil, err
		}
		i, err := p.Factor()
		if err != nil {
			return nil, err
		}
		node := NewUnaryOp(token, i)
		return node, nil
	}
	if token.Type == TokenINTEGER {
		err := p.Eat(TokenINTEGER)
		if err != nil {
			return nil, err
		}
		return NewNum(token), nil
	}
	if token.Type == TokenLPAREN {
		err := p.Eat(TokenLPAREN)
		if err != nil {
			return nil, err
		}
		node, err := p.Expr()
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenRPAREN)
		if err != nil {
			return nil, err
		}
		return node, nil
	}

	return p.Variable()
}

// program : compound_statement DOT
//
// compound_statement : BEGIN statement_list END
//
// statement_list : statement
//				  | statement SEMI statement_list
//
// statement : compound_statement
//				  | assignment_statement
//				  | empty
//
// assignment_statement : variable ASSIGN expr
//
// empty :
//
// expr : term ((PLUS | MINUS) term)*
//
// term : factor ((MUL | DIV) factor)*
//
// factor : PLUS factor
//		  | MINUS factor
//        | INTEGER
//        | LPAREN expr RPAREN
//        | variable
//
// variable : ID
func (p *Parser) Parse() (interface{}, error) {
	node, err := p.Program()
	if err != nil {
		return nil, err
	}
	if p.CurrentToken.Type != TokenEOF {
		return nil, errors.New("parser error not eof")
	}
	return node, nil
}
