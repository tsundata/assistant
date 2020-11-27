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
	err := p.Eat(TokenProgram)
	if err != nil {
		return nil, err
	}
	varNode, err := p.Variable()
	if err != nil {
		return nil, err
	}
	programName := varNode.(*Var).Value.(string)
	err = p.Eat(TokenSemi)
	if err != nil {
		return nil, err
	}
	blockNode, err := p.Block()
	if err != nil {
		return nil, err
	}
	programNode := NewProgram(programName, blockNode)
	err = p.Eat(TokenDot)
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

	for p.CurrentToken.Type == TokenFunction {
		funcDecl, err := p.FunctionDeclaration()
		if err != nil {
			return nil, err
		}
		declarations = append(declarations, []Ast{funcDecl})
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

func (p *Parser) FormalParameterList() ([]Ast, error) {
	if p.CurrentToken.Type != TokenID {
		return []Ast{}, nil
	}

	paramNodes, err := p.FormalParameters()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == TokenSemi {
		err := p.Eat(TokenSemi)
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

	for p.CurrentToken.Type == TokenComma {
		err := p.Eat(TokenComma)
		if err != nil {
			return nil, err
		}
		varNodes = append(varNodes, NewVar(p.CurrentToken))
		err = p.Eat(TokenID)
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

func (p *Parser) FunctionDeclaration() (Ast, error) {
	err := p.Eat(TokenFunction)
	if err != nil {
		return nil, err
	}
	funcName := p.CurrentToken.Value.(string)
	err = p.Eat(TokenID)
	if err != nil {
		return nil, err
	}

	var formalParams []Ast
	if p.CurrentToken.Type == TokenLParen {
		err = p.Eat(TokenLParen)
		if err != nil {
			return nil, err
		}
		formalParams, err = p.FormalParameterList()
		if err != nil {
			return nil, err
		}
		err = p.Eat(TokenRParen)
		if err != nil {
			return nil, err
		}
	}

	var returnType Ast
	if p.CurrentToken.Type == TokenColon {
		err = p.Eat(TokenColon)
		if err != nil {
			return nil, err
		}
		returnType, err = p.TypeSpec()
		if err != nil {
			return nil, err
		}
	}

	err = p.Eat(TokenSemi)
	if err != nil {
		return nil, err
	}
	blockNode, err := p.Block()
	if err != nil {
		return nil, err
	}
	funcDecl := NewFunctionDecl(funcName, formalParams, blockNode, returnType)
	err = p.Eat(TokenSemi)
	if err != nil {
		return nil, err
	}

	return funcDecl, nil
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
	} else {
		err := p.Eat(TokenFloat)
		if err != nil {
			return nil, err
		}
	}

	return NewType(token), nil
}

func (p *Parser) CompoundStatement() (Ast, error) {
	err := p.Eat(TokenBegin)
	if err != nil {
		return nil, err
	}
	nodes, err := p.StatementList()
	err = p.Eat(TokenEnd)
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
	if p.CurrentToken.Type == TokenBegin {
		return p.CompoundStatement()
	} else if p.CurrentToken.Type == TokenID && p.Lexer.CurrentChar == '(' {
		return p.FunctionCallStatement()
	} else if p.CurrentToken.Type == TokenID {
		return p.AssignmentStatement()
	} else if p.CurrentToken.Type == TokenReturn {
		return p.ReturnStatement()
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

func (p *Parser) ReturnStatement() (Ast, error) {
	err := p.Eat(TokenReturn)
	if err != nil {
		return nil, err
	}
	statement, err := p.Expression()
	if err != nil {
		return nil, err
	}

	return NewReturn(statement), nil
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

	err = p.Eat(TokenDo)
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

	err = p.Eat(TokenThen)
	if err != nil {
		return nil, err
	}

	thenBranch, err := p.StatementList()
	if err != nil {
		return nil, err
	}

	var elseBranch []Ast
	if p.CurrentToken.Type == TokenElse {
		err := p.Eat(TokenElse)
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

func (p *Parser) FunctionCallStatement() (Ast, error) {
	token := p.CurrentToken

	funcName := p.CurrentToken.Value.(string)
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}
	err = p.Eat(TokenLParen)
	if err != nil {
		return nil, err
	}
	var actualParams []Ast
	if p.CurrentToken.Type != TokenRParen {
		node, err := p.Expr()
		if err != nil {
			return nil, err
		}
		actualParams = append(actualParams, node)
	}

	for p.CurrentToken.Type == TokenComma {
		err := p.Eat(TokenComma)
		if err != nil {
			return nil, err
		}
		node, err := p.Expr()
		if err != nil {
			return nil, err
		}
		actualParams = append(actualParams, node)
	}

	err = p.Eat(TokenRParen)
	if err != nil {
		return nil, err
	}

	return NewFunctionCall(funcName, actualParams, token), nil
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
		return NewNumber(token), nil
	}
	if token.Type == TokenFloatConst {
		err := p.Eat(TokenFloatConst)
		if err != nil {
			return nil, err
		}
		return NewNumber(token), nil
	}
	if token.Type == TokenStringConst {
		err := p.Eat(TokenStringConst)
		if err != nil {
			return nil, err
		}
		return NewString(token), nil
	}
	if token.Type == TokenTrue {
		err := p.Eat(TokenTrue)
		if err != nil {
			return nil, err
		}
		return NewBoolean(token), nil
	}
	if token.Type == TokenFalse {
		err := p.Eat(TokenFalse)
		if err != nil {
			return nil, err
		}
		return NewBoolean(token), nil
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
	if token.Type == TokenID && p.Lexer.CurrentChar == '(' {
		return p.FunctionCallStatement()
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
