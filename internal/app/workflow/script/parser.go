package script

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
	var err error
	nodes := make(map[string]Ast)
	if p.CurrentToken.Type == TokenNode {
		nodes, err = p.Node()
		if err != nil {
			return nil, err
		}
	}

	workflows := make(map[string]Ast)
	if p.CurrentToken.Type == TokenWorkflow {
		workflows, err = p.Workflow()
		if err != nil {
			return nil, err
		}
	}

	return NewProgram("main", nodes, workflows), nil
}

func (p *Parser) Node() (map[string]Ast, error) {
	nodes := make(map[string]Ast)

	for p.CurrentToken.Type == TokenNode {
		err := p.Eat(TokenNode)
		if err != nil {
			return nil, err
		}
		name := p.CurrentToken.Value.(string)
		err = p.Eat(TokenID)
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenLParen)
		if err != nil {
			return nil, err
		}
		regular := p.CurrentToken.Value.(string)
		err = p.Eat(TokenID)
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenRParen)
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenColon)
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenWith)
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenColon)
		if err != nil {
			return nil, err
		}
		with, err := p.Dict()
		if err != nil {
			return nil, err
		}
		var secret string
		if p.CurrentToken.Type == TokenSecret {
			err = p.Eat(TokenSecret)
			if err != nil {
				return nil, err
			}
			err = p.Eat(TokenColon)
			if err != nil {
				return nil, err
			}
			secret = p.CurrentToken.Value.(string)
			err = p.Eat(TokenID)
			if err != nil {
				return nil, err
			}
		}
		err = p.Eat(TokenEnd)
		if err != nil {
			return nil, err
		}

		nodes[name] = NewNode(name, regular, with, secret)
	}

	return nodes, nil
}

func (p *Parser) Workflow() (map[string]Ast, error) {
	workflows := make(map[string]Ast)

	for p.CurrentToken.Type == TokenWorkflow {
		err := p.Eat(TokenWorkflow)
		if err != nil {
			return nil, err
		}
		name := p.CurrentToken.Value.(string)
		err = p.Eat(TokenID)
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenColon)
		if err != nil {
			return nil, err
		}
		scenarios, err := p.Block()
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenEnd)
		if err != nil {
			return nil, err
		}
		workflows[name] = NewWorkflow(name, scenarios)
	}

	return workflows, nil
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

	for {
		if p.CurrentToken.Type == TokenVar {
			err := p.Eat(TokenVar)
			if err != nil {
				return nil, err
			}
			for p.CurrentToken.Type == TokenID {
				varDecl, err := p.VariableDeclaration()
				if err != nil {
					return nil, err
				}
				declarations = append(declarations, varDecl)
				err = p.Eat(TokenSemi)
				if err != nil {
					return nil, err
				}
			}
		} else {
			break
		}
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
	for p.CurrentToken.Type == TokenComma {
		err = p.Eat(TokenComma)
		if err != nil {
			return nil, err
		}
		paramTokens = append(paramTokens, p.CurrentToken)
		err = p.Eat(TokenID)
		if err != nil {
			return nil, err
		}
	}

	err = p.Eat(TokenColon)
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

func (p *Parser) VariableDeclaration() ([]Ast, error) {
	varNodes := []Ast{NewVar(p.CurrentToken)}
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenComma {
		err = p.Eat(TokenComma)
		if err != nil {
			return nil, err
		}
		varNodes = append(varNodes, NewVar(p.CurrentToken))
		err = p.Eat(TokenID)
		if err != nil {
			return nil, err
		}
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

func (p *Parser) TypeSpec() (Ast, error) {
	token := p.CurrentToken
	if p.CurrentToken.Type == TokenInteger {
		err := p.Eat(TokenInteger)
		if err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == TokenFloat {
		err := p.Eat(TokenFloat)
		if err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == TokenString {
		err := p.Eat(TokenString)
		if err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == TokenBoolean {
		err := p.Eat(TokenBoolean)
		if err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == TokenList {
		err := p.Eat(TokenList)
		if err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == TokenDict {
		err := p.Eat(TokenDict)
		if err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == TokenMessage {
		err := p.Eat(TokenMessage)
		if err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == TokenNode {
		err := p.Eat(TokenNode)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, p.error(UnexpectedToken, token)
	}

	return NewType(token), nil
}

func (p *Parser) CompoundStatement() (Ast, error) {
	nodes, err := p.StatementList()
	if err != nil {
		return nil, err
	}

	root := NewCompound()
	root.Children = append(root.Children, nodes...)
	return root, nil
}

func (p *Parser) StatementList() ([]Ast, error) {
	node, err := p.Statement()
	if err != nil {
		return nil, err
	}

	results := []Ast{node}
	for p.CurrentToken.Type == TokenSemi {
		err = p.Eat(TokenSemi)
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
	// FIXME
	if p.CurrentToken.Type == TokenLCurly {
		return p.CompoundStatement()
	} else if p.CurrentToken.Type == TokenNodeConst {
		return p.FlowStatement()
	} else if p.CurrentToken.Type == TokenID {
		return p.AssignmentStatement()
	} else if p.CurrentToken.Type == TokenPrint {
		return p.PrintStatement()
	} else if p.CurrentToken.Type == TokenIf {
		return p.IfStatement()
	} else if p.CurrentToken.Type == TokenWhile {
		return p.WhileStatement()
	} else {
		return p.Empty()
	}
}

func (p *Parser) FlowStatement() (Ast, error) {
	var nodes []Ast
	node := p.CurrentToken
	err := p.Eat(TokenNodeConst)
	if err != nil {
		return nil, err
	}
	nodes = append(nodes, node)

	for p.CurrentToken.Type == TokenFlow {
		err = p.Eat(TokenFlow)
		if err != nil {
			return nil, err
		}

		right := p.CurrentToken
		err := p.Eat(TokenNodeConst)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, right)
	}

	return NewFlow(nodes), nil
}

func (p *Parser) PrintStatement() (Ast, error) {
	err := p.Eat(TokenPrint)
	if err != nil {
		return nil, err
	}
	statement, err := p.Expression()
	if err != nil {
		return nil, err
	}

	return NewPrint(statement), nil
}

func (p *Parser) WhileStatement() (Ast, error) {
	err := p.Eat(TokenWhile)
	if err != nil {
		return nil, err
	}
	condition, err := p.Expression()
	if err != nil {
		return nil, err
	}

	err = p.Eat(TokenColon)
	if err != nil {
		return nil, err
	}

	doBranch, err := p.StatementList()
	if err != nil {
		return nil, err
	}

	err = p.Eat(TokenEnd)
	if err != nil {
		return nil, err
	}

	return NewWhile(condition, doBranch), nil
}

func (p *Parser) IfStatement() (Ast, error) {
	err := p.Eat(TokenIf)
	if err != nil {
		return nil, err
	}
	condition, err := p.Expression()
	if err != nil {
		return nil, err
	}

	err = p.Eat(TokenColon)
	if err != nil {
		return nil, err
	}

	thenBranch, err := p.StatementList()
	if err != nil {
		return nil, err
	}

	var elseBranch []Ast
	if p.CurrentToken.Type == TokenElse {
		err = p.Eat(TokenElse)
		if err != nil {
			return nil, err
		}
		elseBranch, err = p.StatementList()
		if err != nil {
			return nil, err
		}
	}

	err = p.Eat(TokenEnd)
	if err != nil {
		return nil, err
	}

	return NewIf(condition, thenBranch, elseBranch), nil
}

func (p *Parser) Expression() (Ast, error) {
	return p.LogicOr()
}

func (p *Parser) LogicOr() (Ast, error) {
	node, err := p.LogicAnd()
	if err != nil {
		return 0, err
	}

	for p.CurrentToken.Type == TokenOr {
		token := p.CurrentToken
		err = p.Eat(TokenOr)
		if err != nil {
			return nil, err
		}
		right, err := p.LogicAnd()
		if err != nil {
			return nil, err
		}
		node = NewLogical(node, token, right)
	}

	return node, nil
}

func (p *Parser) LogicAnd() (Ast, error) {
	node, err := p.Equality()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenAnd {
		token := p.CurrentToken
		err = p.Eat(TokenAnd)
		if err != nil {
			return nil, err
		}
		right, err := p.Equality()
		if err != nil {
			return nil, err
		}
		node = NewLogical(node, token, right)
	}

	return node, nil
}

func (p *Parser) Equality() (Ast, error) {
	node, err := p.Comparison()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenEqual || p.CurrentToken.Type == TokenNotEqual {
		token := p.CurrentToken

		err = p.Eat(p.CurrentToken.Type)
		if err != nil {
			return nil, err
		}

		right, err := p.Comparison()
		if err != nil {
			return nil, err
		}
		node = NewLogical(node, token, right)
	}

	return node, nil
}

func (p *Parser) Comparison() (Ast, error) {
	node, err := p.Expr()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenGreater || p.CurrentToken.Type == TokenGreaterEqual ||
		p.CurrentToken.Type == TokenLess || p.CurrentToken.Type == TokenLessEqual {
		token := p.CurrentToken

		err = p.Eat(p.CurrentToken.Type)
		if err != nil {
			return nil, err
		}

		right, err := p.Expr()
		if err != nil {
			return nil, err
		}
		node = NewLogical(node, token, right)
	}

	return node, nil
}

func (p *Parser) AssignmentStatement() (Ast, error) {
	left, err := p.Variable()
	if err != nil {
		return nil, err
	}
	token := p.CurrentToken
	err = p.Eat(TokenAssign)
	if err != nil {
		return nil, err
	}
	right, err := p.Expr()
	if err != nil {
		return nil, err
	}

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

func (p *Parser) List() (Ast, error) {
	var result []Ast
	err := p.Eat(TokenLSquare)
	if err != nil {
		return nil, err
	}

	node, err := p.Factor()
	if err != nil {
		return nil, err
	}
	result = append(result, node)

	for p.CurrentToken.Type == TokenComma {
		err = p.Eat(TokenComma)
		if err != nil {
			return nil, err
		}

		node, err = p.Factor()
		if err != nil {
			return nil, err
		}
		result = append(result, node)
	}

	err = p.Eat(TokenRSquare)
	if err != nil {
		return nil, err
	}

	return NewList(&Token{Type: TokenList, Value: result, Column: p.Lexer.Column, LineNo: p.Lexer.LineNo}), nil
}

func (p *Parser) Dict() (Ast, error) {
	result := make(map[string]Ast)
	err := p.Eat(TokenLCurly)
	if err != nil {
		return nil, err
	}

	key := p.CurrentToken.Value.(string)

	err = p.Eat(TokenStringConst)
	if err != nil {
		return nil, err
	}
	err = p.Eat(TokenColon)
	if err != nil {
		return nil, err
	}

	value, err := p.Factor()
	if err != nil {
		return nil, err
	}

	result[key] = value

	for p.CurrentToken.Type == TokenComma {
		err = p.Eat(TokenComma)
		if err != nil {
			return nil, err
		}

		key = p.CurrentToken.Value.(string)

		err = p.Eat(TokenStringConst)
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenColon)
		if err != nil {
			return nil, err
		}

		value, err = p.Factor()
		if err != nil {
			return nil, err
		}

		result[key] = value
	}

	err = p.Eat(TokenRCurly)
	if err != nil {
		return nil, err
	}

	return NewDict(&Token{Type: TokenDict, Value: result, Column: p.Lexer.Column, LineNo: p.Lexer.LineNo}), nil
}

func (p *Parser) Empty() (Ast, error) {
	return NewNoOp(), nil
}

func (p *Parser) Expr() (Ast, error) {
	node, err := p.Term()
	if err != nil {
		return 0, err
	}

	for p.CurrentToken.Type == TokenPlus || p.CurrentToken.Type == TokenMinus {
		token := p.CurrentToken

		err = p.Eat(p.CurrentToken.Type)
		if err != nil {
			return nil, err
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

	for p.CurrentToken.Type == TokenMultiply || p.CurrentToken.Type == TokenIntegerDiv || p.CurrentToken.Type == TokenFloatDiv {
		token := p.CurrentToken

		err = p.Eat(p.CurrentToken.Type)
		if err != nil {
			return nil, err
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
	if token.Type == TokenPlus {
		err := p.Eat(TokenPlus)
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
	if token.Type == TokenMinus {
		err := p.Eat(TokenMinus)
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
	if token.Type == TokenIntegerConst {
		err := p.Eat(TokenIntegerConst)
		if err != nil {
			return nil, err
		}
		return NewNumberConst(token), nil
	}
	if token.Type == TokenFloatConst {
		err := p.Eat(TokenFloatConst)
		if err != nil {
			return nil, err
		}
		return NewNumberConst(token), nil
	}
	if token.Type == TokenStringConst {
		err := p.Eat(TokenStringConst)
		if err != nil {
			return nil, err
		}
		return NewStringConst(token), nil
	}
	if token.Type == TokenMessageConst {
		err := p.Eat(TokenMessageConst)
		if err != nil {
			return nil, err
		}
		return NewMessageConst(token), nil
	}
	if token.Type == TokenTrue {
		err := p.Eat(TokenTrue)
		if err != nil {
			return nil, err
		}
		return NewBooleanConst(token), nil
	}
	if token.Type == TokenFalse {
		err := p.Eat(TokenFalse)
		if err != nil {
			return nil, err
		}
		return NewBooleanConst(token), nil
	}
	if token.Type == TokenLSquare {
		return p.List()
	}
	if token.Type == TokenLCurly {
		return p.Dict()
	}
	if token.Type == TokenLParen {
		err := p.Eat(TokenLParen)
		if err != nil {
			return nil, err
		}
		node, err := p.Expr()
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenRParen)
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
