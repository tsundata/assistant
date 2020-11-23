package interpreter

import (
	"errors"
)

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
	err := p.Eat(TokenPROGRAM)
	if err != nil {
		return nil, err
	}
	varNode, err := p.Variable()
	if err != nil {
		return nil, err
	}
	programName := varNode.(*Var).Value.([]rune)
	err = p.Eat(TokenSEMI)
	if err != nil {
		return nil, err
	}
	blockNode, err := p.Block()
	if err != nil {
		return nil, err
	}
	programNode := NewProgram(string(programName), blockNode)
	err = p.Eat(TokenDOT)
	if err != nil {
		return nil, err
	}
	return programNode, nil
}

func (p *Parser) Block() (interface{}, error) {
	declarationNodes, err := p.Declarations()
	if err != nil {
		return nil, err
	}
	compoundStatementNode, err := p.CompoundStatement()
	if err != nil {
		return nil, err
	}
	return NewBlock(declarationNodes, compoundStatementNode), nil
}

func (p *Parser) Declarations() ([][]interface{}, error) {
	var declarations [][]interface{}
	if p.CurrentToken.Type == TokenVAR {
		err := p.Eat(TokenVAR)
		if err != nil {
			return nil, err
		}
		for p.CurrentToken.Type == TokenID {
			varDecl, err := p.VariableDeclaration()
			if err != nil {
				return nil, err
			}
			declarations = append(declarations, varDecl)
			err = p.Eat(TokenSEMI)
			if err != nil {
				return nil, err
			}
		}
	}
	return declarations, nil
}

func (p *Parser) VariableDeclaration() ([]interface{}, error) {
	varNodes := []interface{}{NewVar(p.CurrentToken)}
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenCOMMA {
		err := p.Eat(TokenCOMMA)
		if err != nil {
			return nil, err
		}
		varNodes = append(varNodes, NewVar(p.CurrentToken))
		err = p.Eat(TokenID)
	}

	err = p.Eat(TokenCOLON)
	if err != nil {
		return nil, err
	}

	typeNode, err := p.TypeSpec()
	if err != nil {
		return nil, err
	}
	var varDeclarations []interface{}
	for _, varNode := range varNodes {
		varDeclarations = append(varDeclarations, NewVarDecl(varNode, typeNode))
	}
	return varDeclarations, nil
}

func (p *Parser) TypeSpec() (interface{}, error) {
	token := p.CurrentToken
	if p.CurrentToken.Type == TokenINTEGER {
		err := p.Eat(TokenINTEGER)
		if err != nil {
			return nil, err
		}
	} else {
		err := p.Eat(TokenREAL)
		if err != nil {
			return nil, err
		}
	}

	return NewType(token), nil
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

	for p.CurrentToken.Type == TokenMULTIPLY || p.CurrentToken.Type == TokenINTEGERDIV || p.CurrentToken.Type == TokenFLOATDIV {
		token := p.CurrentToken
		if token.Type == TokenMULTIPLY {
			err = p.Eat(TokenMULTIPLY)
			if err != nil {
				return nil, err
			}
		}
		if token.Type == TokenINTEGERDIV {
			err = p.Eat(TokenINTEGERDIV)
			if err != nil {
				return nil, err
			}
		}
		if token.Type == TokenFLOATDIV {
			err = p.Eat(TokenFLOATDIV)
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
	if token.Type == TokenINTEGERCONST {
		err := p.Eat(TokenINTEGERCONST)
		if err != nil {
			return nil, err
		}
		return NewNum(token), nil
	}
	if token.Type == TokenREALCONST {
		err := p.Eat(TokenREALCONST)
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

// program : PROGRAM variable SEMI block DOT
//
// block : declarations compound_statement
//
// declarations : VAR (variable_declaration SEMI)+
//	            | empty
//
// variable_declaration : ID (COMMA ID)* COLON type_spec
//
// type_spec : INTEGER
//	         | REAL
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
// term : factor ((MUL | INTEGER_DIV | FLOAT_DIV) factor)*
//
// factor : PLUS factor
//	      | MINUS factor
//	      | INTEGER_CONST
//	      | REAL_CONST
//	      | LPAREN expr RPAREN
//	      | variable
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
