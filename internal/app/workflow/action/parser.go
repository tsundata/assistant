package action

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
	statements, err := p.StatementList()
	if err != nil {
		return nil, err
	}

	return NewProgram("main", statements), nil
}

func (p *Parser) StatementList() ([]Ast, error) {
	node, err := p.Statement()
	if err != nil {
		return nil, err
	}

	results := []Ast{node}
	for p.CurrentToken.Type == TokenCarriageReturn {
		err = p.Eat(TokenCarriageReturn)
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
	if p.CurrentToken.Type == TokenID {
		return p.OpcodeStatement()
	} else {
		return p.Empty()
	}
}

func (p *Parser) OpcodeStatement() (Ast, error) {
	token := p.CurrentToken
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}
	expressions, err := p.Expression()
	if err != nil {
		return nil, err
	}

	return NewOpcode(token, expressions, token), nil
}

func (p *Parser) Expression() ([]Ast, error) {
	if p.CurrentToken.Type == TokenCarriageReturn {
		return []Ast{}, nil
	}
	if p.CurrentToken.Type == TokenEOF {
		return []Ast{}, nil
	}

	node, err := p.Factor()
	if err != nil {
		return nil, err
	}
	expressions := []Ast{node}

	for p.CurrentToken.Type == TokenIntegerConst || p.CurrentToken.Type == TokenFloatConst ||
		p.CurrentToken.Type == TokenStringConst || p.CurrentToken.Type == TokenMessageConst ||
		p.CurrentToken.Type == TokenTrue || p.CurrentToken.Type == TokenFalse {
		right, err := p.Factor()
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, right)
	}

	return expressions, nil
}

func (p *Parser) Empty() (Ast, error) {
	return NewNoOp(), nil
}

func (p *Parser) Variable() (Ast, error) {
	node := NewVar(p.CurrentToken)
	err := p.Eat(TokenID)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (p *Parser) Factor() (Ast, error) {
	token := p.CurrentToken
	if token.Type == TokenIntegerConst {
		err := p.Eat(TokenIntegerConst)
		if err != nil {
			return nil, err
		}
		return NewIntegerConst(token), nil
	}
	if token.Type == TokenFloatConst {
		err := p.Eat(TokenFloatConst)
		if err != nil {
			return nil, err
		}
		return NewFloatConst(token), nil
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
