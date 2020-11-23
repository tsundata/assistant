package interpreter

import (
	"errors"
)

var ErrInterpreter = errors.New("interpreter error")

type Interpreter struct {
	Parser *Parser
}

func NewInterpreter(parser *Parser) *Interpreter {
	return &Interpreter{Parser: parser}
}

func (i *Interpreter) Visit(node interface{}) int {
	if n, ok := node.(*BinOp); ok {
		return i.VisitBinOp(n)
	}
	if n, ok := node.(*Num); ok {
		return i.VisitNum(n)
	}
	return 0
}

func (i *Interpreter) VisitBinOp(node *BinOp) int {
	if node.Op.Type == PLUS {
		return i.Visit(node.Left) + i.Visit(node.Right)
	}
	if node.Op.Type == MINUS {
		return i.Visit(node.Left) - i.Visit(node.Right)
	}
	if node.Op.Type == MULTIPLY {
		return i.Visit(node.Left) * i.Visit(node.Right)
	}
	if node.Op.Type == DIVIDE {
		return i.Visit(node.Left) / i.Visit(node.Right)
	}
	return 0
}

func (i *Interpreter) VisitNum(node *Num) int {
	return node.Value.(int)
}

func (i *Interpreter) Interpret() (int, error) {
	tree, err := i.Parser.Parse()
	if err != nil {
		return 0, err
	}
	return i.Visit(tree), nil
}
