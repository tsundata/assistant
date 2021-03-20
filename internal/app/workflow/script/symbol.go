package script

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/utils/collection"
	"strings"
)

type Symbol interface{}

type VarSymbol struct {
	Name       string
	Type       Symbol
	ScopeLevel int
}

func NewVarSymbol(name string, t Symbol) *VarSymbol {
	return &VarSymbol{Name: name, Type: t, ScopeLevel: 0}
}

func (s *VarSymbol) String() string {
	return fmt.Sprintf("<VarSymbol(name=%s, type=%v)>", s.Name, s.Type)
}

type BuiltinTypeSymbol struct {
	Name       string
	Type       Symbol
	ScopeLevel int
}

func NewBuiltinTypeSymbol(name string) *BuiltinTypeSymbol {
	return &BuiltinTypeSymbol{Name: name, ScopeLevel: 0}
}

func (s *BuiltinTypeSymbol) String() string {
	return fmt.Sprintf("<BuiltinTypeSymbol(name=%s)>", s.Name)
}

type NodeSymbol struct {
	Name       string
	Parameters []Ast
	BlockAst   Ast
	ScopeLevel int
}

func NewNodeSymbol(name string) *NodeSymbol {
	return &NodeSymbol{Name: name, ScopeLevel: 0}
}

func (s *NodeSymbol) String() string {
	return fmt.Sprintf("<NodeSymbol(name=%s, parameters=%v)>", s.Name, s.Parameters)
}

type WorkflowSymbol struct {
	Name       string
	Scenarios  Ast
	ScopeLevel int
}

func NewWorkflowSymbol(name string) *WorkflowSymbol {
	return &WorkflowSymbol{Name: name, ScopeLevel: 0}
}

func (s *WorkflowSymbol) String() string {
	return fmt.Sprintf("<WorkflowSymbol(name=%s)>", s.Name)
}

type ScopedSymbolTable struct {
	symbols        *collection.OrderedDict
	ScopeName      string
	ScopeLevel     int
	EnclosingScope *ScopedSymbolTable
}

func NewScopedSymbolTable(scopeName string, scopeLevel int, enclosingScope *ScopedSymbolTable) *ScopedSymbolTable {
	table := &ScopedSymbolTable{
		symbols:        collection.NewOrderedDict(),
		ScopeName:      scopeName,
		ScopeLevel:     scopeLevel,
		EnclosingScope: enclosingScope,
	}
	table.Insert(NewBuiltinTypeSymbol("INTEGER"))
	table.Insert(NewBuiltinTypeSymbol("FLOAT"))
	table.Insert(NewBuiltinTypeSymbol("BOOL"))
	table.Insert(NewBuiltinTypeSymbol("STRING"))
	table.Insert(NewBuiltinTypeSymbol("LIST"))
	table.Insert(NewBuiltinTypeSymbol("DICT"))
	table.Insert(NewBuiltinTypeSymbol("MESSAGE"))
	table.Insert(NewBuiltinTypeSymbol("NODE"))
	return table
}

func (t *ScopedSymbolTable) String() string {
	if t == nil {
		return ""
	}

	var lines []string

	lines = append(lines, fmt.Sprintf("Scope name : %s", t.ScopeName))
	lines = append(lines, fmt.Sprintf("Scope level : %d", t.ScopeLevel))

	if t.EnclosingScope != nil {
		lines = append(lines, fmt.Sprintf("Enclosing scope : %s", t.EnclosingScope.ScopeName))
	}

	lines = append(lines, "------------------------------------")
	lines = append(lines, "Scope (Scoped symbol table) contents")

	i := 0
	for v := range t.symbols.Iterate() {
		i++
		lines = append(lines, fmt.Sprintf("%6d: %v", i, v))
	}

	return fmt.Sprintf("\nSCOPE (SCOPED SYMBOL TABLE)\n===========================\n%s\n", strings.Join(lines, "\n"))
}

func (t *ScopedSymbolTable) Insert(symbol Symbol) {
	debugLog(fmt.Sprintf("Insert: %s\n", symbol))

	var name string
	if s, ok := symbol.(*VarSymbol); ok {
		name = s.Name
		s.ScopeLevel = t.ScopeLevel
		t.symbols.Set(name, s)
		return
	}
	if s, ok := symbol.(*BuiltinTypeSymbol); ok {
		name = s.Name
		s.ScopeLevel = t.ScopeLevel
		t.symbols.Set(name, s)
		return
	}
	if s, ok := symbol.(*NodeSymbol); ok {
		name = s.Name
		s.ScopeLevel = t.ScopeLevel
		t.symbols.Set(name, s)
		return
	}
	if s, ok := symbol.(*WorkflowSymbol); ok {
		name = s.Name
		s.ScopeLevel = t.ScopeLevel
		t.symbols.Set(name, s)
		return
	}
}

func (t *ScopedSymbolTable) Lookup(name string, currentScopeOnly bool) Symbol {
	debugLog(fmt.Sprintf("Lookup: %s. (Scope name: %s)\n", name, t.ScopeName))

	s := t.symbols.Get(name)
	if s != nil {
		return s.(Symbol)
	}
	if currentScopeOnly {
		return nil
	}

	if t.EnclosingScope != nil {
		return t.EnclosingScope.Lookup(name, false)
	}
	return nil
}

type SemanticAnalyzer struct {
	CurrentScope *ScopedSymbolTable
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{CurrentScope: nil}
}

func (b *SemanticAnalyzer) error(errorCode ErrorCode, token *Token) error {
	return Error{
		ErrorCode: errorCode,
		Token:     token,
		Message:   fmt.Sprintf("%s -> %v", errorCode, token),
		Type:      SemanticErrorType,
	}
}

func (b *SemanticAnalyzer) Visit(node Ast) error {
	if n, ok := node.(*Program); ok {
		return b.VisitProgram(n)
	}
	if n, ok := node.(*Node); ok {
		return b.VisitNode(n)
	}
	if n, ok := node.(*Workflow); ok {
		return b.VisitWorkflow(n)
	}
	if n, ok := node.(*Block); ok {
		return b.VisitBlock(n)
	}
	if n, ok := node.(*VarDecl); ok {
		return b.VisitVarDecl(n)
	}
	if n, ok := node.(*Type); ok {
		return b.VisitType(n)
	}
	if n, ok := node.(*BinOp); ok {
		return b.VisitBinOp(n)
	}
	if n, ok := node.(*NumberConst); ok {
		return b.VisitNumberConst(n)
	}
	if n, ok := node.(*StringConst); ok {
		return b.VisitStringConst(n)
	}
	if n, ok := node.(*BooleanConst); ok {
		return b.VisitBooleanConst(n)
	}
	if n, ok := node.(*MessageConst); ok {
		return b.VisitMessageConst(n)
	}
	if n, ok := node.(*NodeConst); ok {
		return b.VisitNodeConst(n)
	}
	if n, ok := node.(*List); ok {
		return b.VisitList(n)
	}
	if n, ok := node.(*Dict); ok {
		return b.VisitDict(n)
	}
	if n, ok := node.(*UnaryOp); ok {
		return b.VisitUnaryOp(n)
	}
	if n, ok := node.(*Compound); ok {
		return b.VisitCompound(n)
	}
	if n, ok := node.(*Assign); ok {
		return b.VisitAssign(n)
	}
	if n, ok := node.(*Var); ok {
		return b.VisitVar(n)
	}
	if n, ok := node.(*NoOp); ok {
		return b.VisitNoOp(n)
	}
	if n, ok := node.(*Print); ok {
		return b.VisitPrint(n)
	}
	if n, ok := node.(*While); ok {
		return b.VisitWhile(n)
	}
	if n, ok := node.(*If); ok {
		return b.VisitIf(n)
	}
	if n, ok := node.(*Logical); ok {
		return b.VisitLogical(n)
	}
	if n, ok := node.(*Flow); ok {
		return b.VisitFlow(n)
	}
	return nil
}

func (b *SemanticAnalyzer) VisitProgram(node *Program) error {
	debugLog("ENTER scope: global")
	globalScope := NewScopedSymbolTable("global", 1, b.CurrentScope)
	b.CurrentScope = globalScope

	// nodes
	for _, item := range node.Nodes {
		err := b.Visit(item)
		if err != nil {
			return err
		}
	}
	// workflows
	if _, ok := node.Workflows["main"]; !ok {
		panic(b.error(NoMainWorkflow, nil))
	}
	for _, item := range node.Workflows {
		err := b.Visit(item)
		if err != nil {
			return err
		}
	}

	debugLog(globalScope.String())

	b.CurrentScope = b.CurrentScope.EnclosingScope
	debugLog("LEAVE scope: global")
	return nil
}

func (b *SemanticAnalyzer) VisitNode(node *Node) error {
	nodeSymbol := NewNodeSymbol(node.Name)
	b.CurrentScope.Insert(nodeSymbol)
	return nil
}

func (b *SemanticAnalyzer) VisitWorkflow(node *Workflow) error {
	workflowSymbol := NewWorkflowSymbol(node.Name)
	workflowSymbol.Scenarios = node.Scenarios
	b.CurrentScope.Insert(workflowSymbol)
	return b.Visit(node.Scenarios)
}

func (b *SemanticAnalyzer) VisitFlow(node *Flow) error {
	for _, item := range node.Nodes {
		s := b.CurrentScope.Lookup(item.(*Token).Value.(string), false)
		if s == nil {
			return b.error(IdNotFound, item.(*Token))
		}
	}
	return nil
}

func (b *SemanticAnalyzer) VisitBlock(node *Block) error {
	for _, declaration := range node.Declarations {
		for _, decl := range declaration {
			err := b.Visit(decl)
			if err != nil {
				return err
			}
		}
	}
	return b.Visit(node.CompoundStatement)
}

func (b *SemanticAnalyzer) VisitVarDecl(node *VarDecl) error {
	typeName := node.TypeNode.(*Type).Value.(string)
	typeSymbol := b.CurrentScope.Lookup(typeName, false)
	varName := node.VarNode.(*Var).Value.(string)
	varSymbol := NewVarSymbol(varName, typeSymbol)
	if b.CurrentScope.Lookup(varName, true) != nil {
		return b.error(DuplicateId, node.VarNode.(*Var).Token)
	}
	b.CurrentScope.Insert(varSymbol)
	return nil
}

func (b *SemanticAnalyzer) VisitType(_ *Type) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitBinOp(node *BinOp) error {
	err := b.Visit(node.Left)
	if err != nil {
		return err
	}
	return b.Visit(node.Right)
}

func (b *SemanticAnalyzer) VisitNumberConst(_ *NumberConst) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitStringConst(_ *StringConst) error {
	// pass
	return nil
}
func (b *SemanticAnalyzer) VisitMessageConst(_ *MessageConst) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitBooleanConst(_ *BooleanConst) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitNodeConst(_ *NodeConst) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitList(node *List) error {
	for _, item := range node.Value {
		err := b.Visit(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *SemanticAnalyzer) VisitDict(node *Dict) error {
	for _, item := range node.Value {
		err := b.Visit(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *SemanticAnalyzer) VisitUnaryOp(_ *UnaryOp) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitCompound(node *Compound) error {
	for _, child := range node.Children {
		err := b.Visit(child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *SemanticAnalyzer) VisitAssign(node *Assign) error {
	err := b.Visit(node.Right)
	if err != nil {
		return err
	}
	return b.Visit(node.Left)
}

func (b *SemanticAnalyzer) VisitVar(node *Var) error {
	varName := node.Value.(string)
	varSymbol := b.CurrentScope.Lookup(varName, false)

	if varSymbol == nil {
		return b.error(IdNotFound, node.Token)
	}
	return nil
}

func (b *SemanticAnalyzer) VisitNoOp(_ *NoOp) error {
	// pass
	return nil
}

func (b *SemanticAnalyzer) VisitPrint(node *Print) error {
	err := b.Visit(node.Statement)
	if err != nil {
		return err
	}
	return nil
}

func (b *SemanticAnalyzer) VisitWhile(node *While) error {
	for _, node := range node.DoBranch {
		err := b.Visit(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *SemanticAnalyzer) VisitIf(node *If) error {
	for _, node := range node.ThenBranch {
		err := b.Visit(node)
		if err != nil {
			return err
		}
	}
	for _, node := range node.ElseBranch {
		err := b.Visit(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *SemanticAnalyzer) VisitLogical(node *Logical) error {
	err := b.Visit(node.Left)
	if err != nil {
		return err
	}
	return b.Visit(node.Right)
}
