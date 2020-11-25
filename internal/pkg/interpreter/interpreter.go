package interpreter

import (
	"errors"
)

type Interpreter struct {
	tree         interface{}
	GlobalMemory map[string]interface{}
}

func NewInterpreter(tree Ast) *Interpreter {
	return &Interpreter{tree: tree, GlobalMemory: make(map[string]interface{})}
}

func (i *Interpreter) Visit(node Ast) float64 {
	if n, ok := node.(*Program); ok {
		return i.VisitProgram(n)
	}
	if n, ok := node.(*Block); ok {
		return i.VisitBlock(n)
	}
	if n, ok := node.(*VarDecl); ok {
		return i.VisitVarDecl(n)
	}
	if n, ok := node.(*Type); ok {
		return i.VisitType(n)
	}
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
	if n, ok := node.(*ProcedureDecl); ok {
		return i.VisitProcedureDecl(n)
	}
	if n, ok := node.(*ProcedureCall); ok {
		return i.VisitProcedureCall(n)
	}

	return 0
}

func (i *Interpreter) VisitProgram(node *Program) float64 {
	return i.Visit(node.Block)
}

func (i *Interpreter) VisitBlock(node *Block) float64 {
	for _, declaration := range node.Declarations {
		for _, decl := range declaration {
			i.Visit(decl)
		}
	}
	i.Visit(node.CompoundStatement)
	return 0
}

func (i *Interpreter) VisitVarDecl(node *VarDecl) float64 {
	return 0
}

func (i *Interpreter) VisitType(node *Type) float64 {
	return 0
}

func (i *Interpreter) VisitBinOp(node *BinOp) float64 {
	if node.Op.Type == TokenPLUS {
		return i.Visit(node.Left) + i.Visit(node.Right)
	}
	if node.Op.Type == TokenMINUS {
		return i.Visit(node.Left) - i.Visit(node.Right)
	}
	if node.Op.Type == TokenMULTIPLY {
		return i.Visit(node.Left) * i.Visit(node.Right)
	}
	if node.Op.Type == TokenINTEGERDIV {
		return i.Visit(node.Left) / i.Visit(node.Right)
	}
	return i.Visit(node.Left) / i.Visit(node.Right)
}

func (i *Interpreter) VisitNum(node *Num) float64 {
	return node.Value
}

func (i *Interpreter) VisitUnaryOp(node *UnaryOp) float64 {
	op := node.Op.Type
	if op == TokenPLUS {
		return +i.Visit(node.Expr)
	} else if op == TokenMINUS {
		return -i.Visit(node.Expr)
	}
	return 0
}

func (i *Interpreter) VisitCompound(node *Compound) float64 {
	for _, child := range node.Children {
		i.Visit(child)
	}
	return 0
}

func (i *Interpreter) VisitAssign(node *Assign) float64 {
	if left, ok := node.Left.(*Var); ok {
		varName := left.Value
		if value, ok := varName.([]rune); ok {
			i.GlobalMemory[string(value)] = i.Visit(node.Right)
		}
	}
	return 0
}

func (i *Interpreter) VisitVar(node *Var) float64 {
	if varName, ok := node.Value.([]rune); ok {
		if val, ok := i.GlobalMemory[string(varName)]; ok {
			return val.(float64)
		} else {
			panic(errors.New("interpreter error var name"))
		}
	}
	return 0
}

func (i *Interpreter) VisitNoOp(node *NoOp) float64 {
	return 0
}

func (i *Interpreter) VisitProcedureDecl(node *ProcedureDecl) float64 {
	return 0
}

func (i *Interpreter) VisitProcedureCall(node *ProcedureCall) float64 {
	return 0
}

func (i *Interpreter) Interpret() (float64, error) {
	if i.tree == nil {
		return 0, errors.New("error ast tree")
	}
	return i.Visit(i.tree), nil
}
