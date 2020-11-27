package interpreter

import (
	"errors"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/utils/collection"
	"strings"
)

type ARType string

const (
	ARTypeProgram  ARType = "PROGRAM"
	ARTypeFunction ARType = "FUNCTION"
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

func (i *Interpreter) Visit(node Ast) interface{} {
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
	if n, ok := node.(*Number); ok {
		return i.VisitNumber(n)
	}
	if n, ok := node.(*String); ok {
		return i.VisitString(n)
	}
	if n, ok := node.(*Boolean); ok {
		return i.VisitBoolean(n)
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
	if n, ok := node.(*FunctionDecl); ok {
		return i.VisitFunctionDecl(n)
	}
	if n, ok := node.(*FunctionCall); ok {
		return i.VisitFunctionCall(n)
	}
	if n, ok := node.(*Print); ok {
		return i.VisitPrint(n)
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

	return result.(float64)
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
	if node.Op.Type == TokenPlus {
		return i.Visit(node.Left).(float64) + i.Visit(node.Right).(float64)
	}
	if node.Op.Type == TokenMinus {
		return i.Visit(node.Left).(float64) - i.Visit(node.Right).(float64)
	}
	if node.Op.Type == TokenMultiply {
		return i.Visit(node.Left).(float64) * i.Visit(node.Right).(float64)
	}
	if node.Op.Type == TokenIntegerDiv {
		return i.Visit(node.Left).(float64) / i.Visit(node.Right).(float64)
	}
	return i.Visit(node.Left).(float64) / i.Visit(node.Right).(float64)
}

func (i *Interpreter) VisitNumber(node *Number) float64 {
	return node.Value
}

func (i *Interpreter) VisitString(node *String) string {
	return node.Value
}

func (i *Interpreter) VisitBoolean(node *Boolean) bool {
	return node.Value
}

func (i *Interpreter) VisitUnaryOp(node *UnaryOp) float64 {
	op := node.Op.Type
	if op == TokenPlus {
		return +i.Visit(node.Expr).(float64)
	} else if op == TokenMinus {
		return -i.Visit(node.Expr).(float64)
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

func (i *Interpreter) VisitVar(node *Var) interface{} {
	if varName, ok := node.Value.(string); ok {
		ar := i.callStack.Peek()
		if ar != nil {
			val := ar.Get(varName)
			if val != nil {
				if v, ok := val.(float64); ok {
					return v
				}
				if v, ok := val.(string); ok {
					return v
				}
				if v, ok := val.(bool); ok {
					return v
				}
			}
		} else {
			panic(errors.New("interpreter error var name"))
		}
	}
	return nil
}

func (i *Interpreter) VisitNoOp(node *NoOp) float64 {
	return 0
}

func (i *Interpreter) VisitFunctionDecl(node *FunctionDecl) float64 {
	return 0
}

func (i *Interpreter) VisitFunctionCall(node *FunctionCall) float64 {
	procName := node.ProcName
	procSymbol := node.ProcSymbol

	ar := NewActivationRecord(procName, ARTypeFunction, procSymbol.(*FunctionSymbol).ScopeLevel+1)

	var formalParams []Ast
	if procSymbol != nil {
		formalParams = procSymbol.(*FunctionSymbol).FormalParams
	}
	actualParams := node.ActualParams

	for _, item := range collection.Zip(formalParams, actualParams) {
		k := item.Element1.(*VarSymbol).Name
		v := i.Visit(item.Element2)
		ar.Set(k, v)
	}

	i.callStack.Push(ar)

	fmt.Printf("ENTER: FUNCTION %s\n", procName)
	fmt.Println(i.callStack)

	if procSymbol != nil {
		i.Visit(procSymbol.(*FunctionSymbol).BlockAst)
	}

	fmt.Printf("LEAVE: FUNCTION %s\n", procName)
	fmt.Println(i.callStack)

	i.callStack.Pop()

	return 0
}

func (i *Interpreter) VisitPrint(node *Print) interface{} {
	fmt.Println(">>> ", i.Visit(node.Statement))
	return nil
}

func (i *Interpreter) VisitWhile(node *While) float64 {
	for i.Visit(node.Condition).(bool) {
		for _, n := range node.DoBranch {
			i.Visit(n)
		}
	}
	return 0
}

func (i *Interpreter) VisitIf(node *If) interface{} {
	if i.Visit(node.Condition).(bool) {
		for _, n := range node.ThenBranch {
			i.Visit(n)
		}
	} else {
		for _, n := range node.ElseBranch {
			i.Visit(n)
		}
	}
	return nil
}

func (i *Interpreter) VisitLogical(node *Logical) bool {
	if node.Op.Type == TokenOr {
		return (i.Visit(node.Left) != 0) || (i.Visit(node.Right) != 0)
	}
	if node.Op.Type == TokenAnd {
		return (i.Visit(node.Left) != 0) && (i.Visit(node.Right) != 0)
	}
	if node.Op.Type == TokenEqual {
		return i.Visit(node.Left) == i.Visit(node.Right)
	}
	if node.Op.Type == TokenNotEqual {
		return i.Visit(node.Left) != i.Visit(node.Right)
	}
	if node.Op.Type == TokenGreater {
		return i.Visit(node.Left).(float64) > i.Visit(node.Right).(float64)
	}
	if node.Op.Type == TokenGreaterEqual {
		return i.Visit(node.Left).(float64) >= i.Visit(node.Right).(float64)
	}
	if node.Op.Type == TokenLess {
		return i.Visit(node.Left).(float64) < i.Visit(node.Right).(float64)
	}
	if node.Op.Type == TokenLessEqual {
		return i.Visit(node.Left).(float64) <= i.Visit(node.Right).(float64)
	}

	return false
}

func (i *Interpreter) Interpret() (float64, error) {
	if i.tree == nil {
		return 0, errors.New("error ast tree")
	}
	return i.Visit(i.tree).(float64), nil
}
