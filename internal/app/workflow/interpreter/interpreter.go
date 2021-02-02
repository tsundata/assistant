package interpreter

import (
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/interpreter/nodes"
	"log"
	"strings"
)

type ARType string

const (
	ARTypeWorkflow ARType = "WORKFLOW"
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
		lines = append(lines, s.records[i].String())
	}
	return fmt.Sprintf("CALL STACK\n%s\n\n", strings.Join(lines, "\n"))
}

type ActivationRecord struct {
	Name         string
	Type         ARType
	NestingLevel int
	Members      map[string]interface{}
	ReturnValue  interface{}
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
	lines = append(lines, fmt.Sprintf("%d: %s %s %v", r.NestingLevel, r.Type, r.Name, r.ReturnValue))
	for name, val := range r.Members {
		lines = append(lines, fmt.Sprintf("  %s : %v", name, val))
	}

	return strings.Join(lines, "\n")
}

type Interpreter struct {
	tree      Ast
	callStack *CallStack
	stdout    []interface{}
	nodes     map[string]Ast
	workflow  map[string]Ast

	midClient pb.MiddleClient
}

func NewInterpreter(tree Ast, midClient pb.MiddleClient) *Interpreter {
	return &Interpreter{tree: tree, callStack: NewCallStack(), midClient: midClient}
}

func (i *Interpreter) Visit(node Ast) interface{} {
	if n, ok := node.(*Program); ok {
		return i.VisitProgram(n)
	}
	if n, ok := node.(*Node); ok {
		return i.VisitNode(n)
	}
	if n, ok := node.(*Workflow); ok {
		return i.VisitWorkflow(n)
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
	if n, ok := node.(*NumberConst); ok {
		return i.VisitNumberConst(n)
	}
	if n, ok := node.(*StringConst); ok {
		return i.VisitStringConst(n)
	}
	if n, ok := node.(*BooleanConst); ok {
		return i.VisitBooleanConst(n)
	}
	if n, ok := node.(*MessageConst); ok {
		return i.VisitMessageConst(n)
	}
	if n, ok := node.(*NodeConst); ok {
		return i.VisitNodeConst(n)
	}
	if n, ok := node.(*List); ok {
		return i.VisitList(n)
	}
	if n, ok := node.(*Dict); ok {
		return i.VisitDict(n)
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
	if n, ok := node.(*Flow); ok {
		return i.VisitFlow(n)
	}

	return 0
}

func (i *Interpreter) VisitProgram(node *Program) float64 {
	programName := node.Name
	log.Printf("ENTER: PROGRAM %s\n", programName)

	i.nodes = node.Nodes
	i.workflow = node.Workflows

	// main workflow
	var result float64
	if item, ok := node.Workflows["main"]; ok {
		result = i.Visit(item).(float64)
	}

	log.Printf("LEAVE: PROGRAM %s\n", programName)
	log.Println(i.callStack)

	i.callStack.Pop()

	return result
}

func (i *Interpreter) VisitNode(node *Node) map[string]interface{} {
	return i.Visit(node.With).(map[string]interface{})
}

func (i *Interpreter) VisitWorkflow(node *Workflow) float64 {
	ar := NewActivationRecord(node.Name, ARTypeWorkflow, 1)
	i.callStack.Push(ar)
	log.Printf("ENTER: WORKFLOW %s\n", node.Name)

	r := i.Visit(node.Scenarios).(float64)

	log.Printf("LEAVE: WORKFLOW %s\n", node.Name)
	return r
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

func (i *Interpreter) VisitFlow(node *Flow) float64 {
	var err error
	var input []map[string]interface{}
	for _, item := range node.Nodes {
		nodeName := item.(*Token).Value.(string)
		if item, ok := i.nodes[nodeName]; ok {
			// execute
			input, err = nodes.Construct(i.midClient, nodeName, item.(*Node).Regular, i.Visit(item).(map[string]interface{}), item.(*Node).Secret).Execute(input)
			if err != nil {
				panic(err)
			}
		}
	}

	return 0
}

func (i *Interpreter) VisitVarDecl(node *VarDecl) float64 {
	return 0
}

func (i *Interpreter) VisitType(node *Type) float64 {
	return 0
}

func (i *Interpreter) VisitBinOp(node *BinOp) float64 {
	var left float64
	var right float64
	if v, ok := i.Visit(node.Left).(int); ok {
		left = float64(v)
	} else if v, ok := i.Visit(node.Left).(float64); ok {
		left = v
	}
	if v, ok := i.Visit(node.Right).(int); ok {
		right = float64(v)
	} else if v, ok := i.Visit(node.Right).(float64); ok {
		right = v
	}
	if node.Op.Type == TokenPlus {
		return left + right
	}
	if node.Op.Type == TokenMinus {
		return left - right
	}
	if node.Op.Type == TokenMultiply {
		return left * right
	}
	if node.Op.Type == TokenIntegerDiv {
		return left / right
	}
	return 0
}

func (i *Interpreter) VisitNumberConst(node *NumberConst) float64 {
	return node.Value
}

func (i *Interpreter) VisitStringConst(node *StringConst) string {
	return node.Value
}

func (i *Interpreter) VisitBooleanConst(node *BooleanConst) bool {
	return node.Value
}

func (i *Interpreter) VisitMessageConst(node *MessageConst) interface{} {
	return node.Value
}

func (i *Interpreter) VisitNodeConst(node *NodeConst) interface{} {
	return node.Value
}

func (i *Interpreter) VisitList(node *List) []interface{} {
	var result []interface{}
	for _, item := range node.Value {
		result = append(result, i.Visit(item))
	}

	return result
}

func (i *Interpreter) VisitDict(node *Dict) map[string]interface{} {
	result := make(map[string]interface{})
	for key, item := range node.Value {
		result[key] = i.Visit(item)
	}

	return result
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
			// TODO var nil
			val := ar.Get(varName)
			if val != nil {
				return val
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

func (i *Interpreter) VisitPrint(node *Print) interface{} {
	i.stdout = append(i.stdout, i.Visit(node.Statement))
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

func (i *Interpreter) Stdout() string {
	var out []string
	for _, line := range i.stdout {
		out = append(out, fmt.Sprintf("> %v", line))
	}
	return strings.Join(out, "\n")
}
