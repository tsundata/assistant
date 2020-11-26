package interpreter

import (
	"errors"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/utils/collection"
	"strings"
)

type ARType string

const (
	ARTypeProgram   ARType = "PROGRAM"
	ARTypeProcedure ARType = "PROCEDURE"
)

type CallStack struct {
	records []*ActivationRecord
}

func NewCallStack() *CallStack {
	return &CallStack{}
}

func (s *CallStack) Push(ar *ActivationRecord) {
	s.records = append(s.records, ar)
}

func (s *CallStack) Pop() *ActivationRecord {
	if len(s.records) > 0 {
		top := s.records[len(s.records)-1]
		s.records = s.records[:len(s.records)-1]
		return top
	}
	return nil
}

func (s *CallStack) Peek() *ActivationRecord {
	if len(s.records) > 0 {
		return s.records[len(s.records)-1]
	}
	return nil
}

func (s *CallStack) String() string {
	var lines []string
	for i := len(s.records) - 1; i >= 0; i-- {
		lines = append(lines, fmt.Sprintf("%s", s.records[i]))
	}
	return fmt.Sprintf("CALL STACK\n%s\n\n", strings.Join(lines, "\n"))
}

type ActivationRecord struct {
	Name         string
	Type         ARType
	NestingLevel int
	Members      map[string]interface{}
}

func NewActivationRecord(name string, t ARType, nestingLevel int) *ActivationRecord {
	return &ActivationRecord{Name: name, Type: t, NestingLevel: nestingLevel, Members: make(map[string]interface{})}
}

func (r *ActivationRecord) Get(key string) interface{} {
	return r.Members[key]
}

func (r *ActivationRecord) Set(key string, value interface{}) {
	r.Members[key] = value
}

func (r *ActivationRecord) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("%d: %s %s", r.NestingLevel, r.Type, r.Name))
	for name, val := range r.Members {
		lines = append(lines, fmt.Sprintf("  %s : %v", name, val))
	}

	return strings.Join(lines, "\n")
}

type Interpreter struct {
	tree      Ast
	callStack *CallStack
}

func NewInterpreter(tree Ast) *Interpreter {
	return &Interpreter{tree: tree, callStack: NewCallStack()}
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
	if n, ok := node.(*While); ok {
		return i.VisitWhile(n)
	}
	if n, ok := node.(*If); ok {
		return i.VisitIf(n)
	}
	if n, ok := node.(*Logical); ok {
		return i.VisitLogical(n)
	}

	return 0
}

func (i *Interpreter) VisitProgram(node *Program) float64 {
	programName := node.Name
	fmt.Printf("ENTER: PROGRAM %s\n", programName)

	ar := NewActivationRecord(programName, ARTypeProgram, 1)
	i.callStack.Push(ar)
	fmt.Println(i.callStack)

	result := i.Visit(node.Block)

	fmt.Printf("LEAVE: PROGRAM %s\n", programName)
	fmt.Println(i.callStack)

	i.callStack.Pop()

	return result
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
		if value, ok := varName.(string); ok {
			ar := i.callStack.Peek()
			if ar != nil {
				ar.Set(value, i.Visit(node.Right))
			}
		}
	}
	return 0
}

func (i *Interpreter) VisitVar(node *Var) float64 {
	if varName, ok := node.Value.(string); ok {
		ar := i.callStack.Peek()
		if ar != nil {
			val := ar.Get(varName)
			if val != nil {
				return val.(float64)
			} else {
				// TODO Uninitialized
				return 0
			}
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
	procName := node.ProcName
	procSymbol := node.ProcSymbol

	ar := NewActivationRecord(procName, ARTypeProcedure, procSymbol.(*ProcedureSymbol).ScopeLevel+1)

	var formalParams []Ast
	if procSymbol != nil {
		formalParams = procSymbol.(*ProcedureSymbol).FormalParams
	}
	actualParams := node.ActualParams

	for _, item := range collection.Zip(formalParams, actualParams) {
		k := item.Element1.(*VarSymbol).Name
		v := i.Visit(item.Element2)
		ar.Set(k, v)
	}

	i.callStack.Push(ar)

	fmt.Printf("ENTER: PROCEDURE %s\n", procName)
	fmt.Println(i.callStack)

	if procSymbol != nil {
		i.Visit(procSymbol.(*ProcedureSymbol).BlockAst)
	}

	fmt.Printf("LEAVE: PROCEDURE %s\n", procName)
	fmt.Println(i.callStack)

	i.callStack.Pop()

	return 0
}

func (i *Interpreter) VisitWhile(node *While) float64 {
	for i.Visit(node.Condition) != 0 {
		i.Visit(node.DoBranch)
	}
	return 0
}

func (i *Interpreter) VisitIf(node *If) float64 {
	if i.Visit(node.Condition) != 0 {
		return i.Visit(node.ThenBranch)
	} else {
		return i.Visit(node.ElseBranch)
	}
}

func (i *Interpreter) VisitLogical(node *Logical) float64 {
	var b bool
	if node.Op.Type == TokenOR {
		b = (i.Visit(node.Left) != 0) || (i.Visit(node.Right) != 0)
	}
	if node.Op.Type == TokenAND {
		b = (i.Visit(node.Left) != 0) && (i.Visit(node.Right) != 0)
	}
	if node.Op.Type == TokenEQUAL {
		b = i.Visit(node.Left) == i.Visit(node.Right)
	}
	if node.Op.Type == TokenNOTEQUAL {
		b = i.Visit(node.Left) != i.Visit(node.Right)
	}
	if node.Op.Type == TokenGREATER {
		b = i.Visit(node.Left) > i.Visit(node.Right)
	}
	if node.Op.Type == TokenGREATEREQUAL {
		b = i.Visit(node.Left) >= i.Visit(node.Right)
	}
	if node.Op.Type == TokenLESS {
		b = i.Visit(node.Left) < i.Visit(node.Right)
	}
	if node.Op.Type == TokenLESSEQUAL {
		b = i.Visit(node.Left) <= i.Visit(node.Right)
	}

	if b {
		return 1
	} else {
		return 0
	}
}

func (i *Interpreter) Interpret() (float64, error) {
	if i.tree == nil {
		return 0, errors.New("error ast tree")
	}
	return i.Visit(i.tree), nil
}
