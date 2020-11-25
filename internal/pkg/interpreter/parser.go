package interpreter

import (
	"fmt"
)

type Parser struct {
	Lexer        *Lexer
	CurrentToken *Token
}

func NewParser(lexer *Lexer) (*Parser, error) {
	p := &Parser{Lexer: lexer, CurrentToken: nil}
	token, err := p.getNextToken()
	if err != nil {
		return nil, err
	}
	p.CurrentToken = token
	return p, nil
}

func (p *Parser) getNextToken() (*Token, error) {
	return p.Lexer.GetNextToken()
}

func (p *Parser) error(errorCode ErrorCode, token *Token) error {
	return Error{
		Token:   token,
		Message: fmt.Sprintf("%s -> %v", errorCode, token),
		Type:    ParserErrorType,
	}
}

func (p *Parser) Eat(tokenType TokenType) (err error) {
	if p.CurrentToken.Type == tokenType {
		p.CurrentToken, err = p.Lexer.GetNextToken()
		return
	}

	return p.error(UnexpectedToken, p.CurrentToken)
}

func (p *Parser) Program() (Ast, error) {
	err := p.Eat(TokenPROGRAM)
	if err != nil {
		return nil, err
	}
	varNode, err := p.Variable()
	if err != nil {
		return nil, err
	}
	programName := varNode.(*Var).Value.(string)
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

func (p *Parser) Block() (Ast, error) {
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

func (p *Parser) Declarations() ([][]Ast, error) {
	var declarations [][]Ast

	for true {
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
		} else {
			break
		}
	}

	for p.CurrentToken.Type == TokenPROGRAM {
		procDecl, err := p.ProcedureDeclaration()
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, []Ast{procDecl})
	}

	return declarations, nil
}

func (p *Parser) FormalParameters() ([]Ast, error) {
	var paramNodes []Ast

	paramTokens := []*Token{p.CurrentToken}
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}
	for p.CurrentToken.Type == TokenCOMMA {
		err = p.Eat(TokenCOMMA)
		if err != nil {
			return nil, err
		}
		paramTokens = append(paramTokens, p.CurrentToken)
		err = p.Eat(TokenID)
		if err != nil {
			return nil, err
		}
	}

	err = p.Eat(TokenCOLON)
	if err != nil {
		return nil, err
	}
	typeNode, err := p.TypeSpec()
	if err != nil {
		return nil, err
	}

	for _, paramToken := range paramTokens {
		paramNode := NewParam(NewVar(paramToken), typeNode)
		paramNodes = append(paramNodes, paramNode)
	}

	return paramNodes, nil
}

func (p *Parser) FormalParameterList() ([]Ast, error) {
	if p.CurrentToken.Type != TokenID {
		return []Ast{}, nil
	}

	paramNodes, err := p.FormalParameters()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenSEMI {
		err := p.Eat(TokenSEMI)
		if err != nil {
			return nil, err
		}
		params, err := p.FormalParameters()
		if err != nil {
			return nil, err
		}
		paramNodes = append(paramNodes, params...)
	}

	return paramNodes, nil
}

func (p *Parser) VariableDeclaration() ([]Ast, error) {
	varNodes := []Ast{NewVar(p.CurrentToken)}
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
	var varDeclarations []Ast
	for _, varNode := range varNodes {
		varDeclarations = append(varDeclarations, NewVarDecl(varNode, typeNode))
	}
	return varDeclarations, nil
}

func (p *Parser) ProcedureDeclaration() (Ast, error) {
	err := p.Eat(TokenPROGRAM)
	if err != nil {
		return nil, err
	}
	procName := p.CurrentToken.Value.(string)
	err = p.Eat(TokenID)
	if err != nil {
		return nil, err
	}

	var formalParams []Ast
	if p.CurrentToken.Type == TokenLPAREN {
		err = p.Eat(TokenLPAREN)
		if err != nil {
			return nil, err
		}
		formalParams, err = p.FormalParameterList()
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenRPAREN)
		if err != nil {
			return nil, err
		}
	}

	err = p.Eat(TokenSEMI)
	if err != nil {
		return nil, err
	}
	blockNode, err := p.Block()
	if err != nil {
		return nil, err
	}
	procDecl := NewProcedureDecl(string(procName), formalParams, blockNode)
	err = p.Eat(TokenSEMI)
	if err != nil {
		return nil, err
	}

	return procDecl, nil
}

func (p *Parser) TypeSpec() (Ast, error) {
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

func (p *Parser) CompoundStatement() (Ast, error) {
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

func (p *Parser) StatementList() ([]Ast, error) {
	node, err := p.Statement()
	if err != nil {
		return nil, err
	}

	results := []Ast{node}

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

func (p *Parser) Statement() (Ast, error) {
	if p.CurrentToken.Type == TokenBEGIN {
		return p.CompoundStatement()
	} else if p.CurrentToken.Type == TokenID && p.Lexer.CurrentChar == '(' {
		return p.ProcallStatement()
	} else if p.CurrentToken.Type == TokenID {
		return p.AssignmentStatement()
	} else {
		return p.Empty()
	}
}

func (p *Parser) ProcallStatement() (Ast, error) {
	token := p.CurrentToken

	procName := p.CurrentToken.Value.(string)
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}
	err = p.Eat(TokenLPAREN)
	if err != nil {
		return nil, err
	}
	var actualParams []Ast
	if p.CurrentToken.Type != TokenRPAREN {
		node, err := p.Expr()
		if err != nil {
			return nil, err
		}
		actualParams = append(actualParams, node)
	}

	for p.CurrentToken.Type == TokenCOMMA {
		err := p.Eat(TokenCOMMA)
		if err != nil {
			return nil, err
		}
		node, err := p.Expr()
		if err != nil {
			return nil, err
		}
		actualParams = append(actualParams, node)
	}

	err = p.Eat(TokenRPAREN)
	if err != nil {
		return nil, err
	}

	return NewProcedureCall(procName, actualParams, token), nil
}

func (p *Parser) AssignmentStatement() (Ast, error) {
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

func (p *Parser) Variable() (Ast, error) {
	node := NewVar(p.CurrentToken)
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (p *Parser) Empty() (Ast, error) {
	return NewNoOp(), nil
}

func (p *Parser) Expr() (Ast, error) {
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

func (p *Parser) Term() (Ast, error) {
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

func (p *Parser) Factor() (Ast, error) {
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

func (p *Parser) Parse() (Ast, error) {
	node, err := p.Program()
	if err != nil {
		return nil, err
	}
	if p.CurrentToken.Type != TokenEOF {
		return nil, p.error(UnexpectedToken, p.CurrentToken)
	}
	return node, nil
}
