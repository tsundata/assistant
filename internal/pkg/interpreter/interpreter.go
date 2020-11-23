package interpreter

import (
	"errors"
	"fmt"
)

type Interpreter struct {
	Parser      *Parser
	GlobalScope map[string]interface{}
}

func NewInterpreter(parser *Parser) *Interpreter {
	return &Interpreter{Parser: parser, GlobalScope: make(map[string]interface{})}
}

func (i *Interpreter) Visit(node interface{}) int {
	if n, ok := node.(*BinOp); ok {
		return i.VisitBinOp(n)
	}
	if n, ok := node.(*Num); ok {
		return i.VisitNum(n)
	}
	if n, ok := node.(*UnaryOp); ok {
		return i.VisitUnaryOp(n)
	}
	if n, ok := node.(*Compound); ok {
		return i.VisitCompound(n)
	}
	if n, ok := node.(*Assign); ok {
		return i.VisitAssign(n)
	}
	if n, ok := node.(*Var); ok {
		return i.VisitVar(n)
	}
	if n, ok := node.(*NoOp); ok {
		return i.VisitNoOp(n)
	}

	return 0
}

func (i *Interpreter) VisitBinOp(node *BinOp) int {
	if node.Op.Type == TokenPLUS {
		return i.Visit(node.Left) + i.Visit(node.Right)
	}
	if node.Op.Type == TokenMINUS {
		return i.Visit(node.Left) - i.Visit(node.Right)
	}
	if node.Op.Type == TokenMULTIPLY {
		return i.Visit(node.Left) * i.Visit(node.Right)
	}
	if node.Op.Type == TokenDIVIDE {
		return i.Visit(node.Left) / i.Visit(node.Right)
	}
	return 0
}

func (i *Interpreter) VisitNum(node *Num) int {
	return node.Value.(int)
}

func (i *Interpreter) VisitUnaryOp(node *UnaryOp) int {
	op := node.Op.Type
	if op == TokenPLUS {
		return +i.Visit(node.Expr)
	} else if op == TokenMINUS {
		return -i.Visit(node.Expr)
	}
	return 0
}

func (i *Interpreter) VisitCompound(node *Compound) int {
	for _, child := range node.Children {
		i.Visit(child)
	}
	return 0
}

func (i *Interpreter) VisitAssign(node *Assign) int {
	fmt.Println(node)
	if left, ok := node.Left.(*Var); ok {
		varName := left.Value
		if value, ok := varName.([]rune); ok {
			i.GlobalScope[string(value)] = i.Visit(node.Right)
		}
	}
	return 0
}

func (i *Interpreter) VisitVar(node *Var) int {
	if varName, ok := node.Value.([]rune); ok {
		if val, ok := i.GlobalScope[string(varName)]; ok {
			return val.(int)
		} else {
			panic(errors.New("interpreter error var name"))
		}
	}
	return 0
}

func (i *Interpreter) VisitNoOp(node *NoOp) int {
	return 0
}

func (i *Interpreter) Interpret() (int, error) {
	tree, err := i.Parser.Parse()
	if err != nil {
		return 0, err
	}
	return i.Visit(tree), nil
}
