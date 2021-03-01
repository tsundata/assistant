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

func (b *SemanticAnalyzer) Visit(node Ast) {
	if n, ok := node.(*Program); ok {
		b.VisitProgram(n)
		return
	}
	if n, ok := node.(*Node); ok {
		b.VisitNode(n)
		return
	}
	if n, ok := node.(*Workflow); ok {
		b.VisitWorkflow(n)
		return
	}
	if n, ok := node.(*Block); ok {
		b.VisitBlock(n)
		return
	}
	if n, ok := node.(*VarDecl); ok {
		b.VisitVarDecl(n)
		return
	}
	if n, ok := node.(*Type); ok {
		b.VisitType(n)
		return
	}
	if n, ok := node.(*BinOp); ok {
		b.VisitBinOp(n)
		return
	}
	if n, ok := node.(*NumberConst); ok {
		b.VisitNumberConst(n)
		return
	}
	if n, ok := node.(*StringConst); ok {
		b.VisitStringConst(n)
		return
	}
	if n, ok := node.(*BooleanConst); ok {
		b.VisitBooleanConst(n)
		return
	}
	if n, ok := node.(*MessageConst); ok {
		b.VisitMessageConst(n)
		return
	}
	if n, ok := node.(*NodeConst); ok {
		b.VisitNodeConst(n)
		return
	}
	if n, ok := node.(*List); ok {
		b.VisitList(n)
		return
	}
	if n, ok := node.(*Dict); ok {
		b.VisitDict(n)
		return
	}
	if n, ok := node.(*UnaryOp); ok {
		b.VisitUnaryOp(n)
		return
	}
	if n, ok := node.(*Compound); ok {
		b.VisitCompound(n)
		return
	}
	if n, ok := node.(*Assign); ok {
		b.VisitAssign(n)
		return
	}
	if n, ok := node.(*Var); ok {
		b.VisitVar(n)
		return
	}
	if n, ok := node.(*NoOp); ok {
		b.VisitNoOp(n)
		return
	}
	if n, ok := node.(*Print); ok {
		b.VisitPrint(n)
		return
	}
	if n, ok := node.(*While); ok {
		b.VisitWhile(n)
		return
	}
	if n, ok := node.(*If); ok {
		b.VisitIf(n)
		return
	}
	if n, ok := node.(*Logical); ok {
		b.VisitLogical(n)
		return
	}
	if n, ok := node.(*Flow); ok {
		b.VisitFlow(n)
		return
	}
}

func (b *SemanticAnalyzer) VisitProgram(node *Program) {
	debugLog("ENTER scope: global")
	globalScope := NewScopedSymbolTable("global", 1, b.CurrentScope)
	b.CurrentScope = globalScope

	// nodes
	for _, item := range node.Nodes {
		b.Visit(item)
	}
	// workflows
	if _, ok := node.Workflows["main"]; !ok {
		panic(b.error(NoMainWorkflow, nil))
	}
	for _, item := range node.Workflows {
		b.Visit(item)
	}

	debugLog(globalScope.String())

	b.CurrentScope = b.CurrentScope.EnclosingScope
	debugLog("LEAVE scope: global")
}

func (b *SemanticAnalyzer) VisitNode(node *Node) {
	nodeSymbol := NewNodeSymbol(node.Name)
	b.CurrentScope.Insert(nodeSymbol)
}

func (b *SemanticAnalyzer) VisitWorkflow(node *Workflow) {
	workflowSymbol := NewWorkflowSymbol(node.Name)
	workflowSymbol.Scenarios = node.Scenarios
	b.CurrentScope.Insert(workflowSymbol)
	b.Visit(node.Scenarios)
}

func (b *SemanticAnalyzer) VisitFlow(node *Flow) {
	for _, item := range node.Nodes {
		s := b.CurrentScope.Lookup(item.(*Token).Value.(string), false)
		if s == nil {
			panic(b.error(IdNotFound, item.(*Token)))
		}
	}
}

func (b *SemanticAnalyzer) VisitBlock(node *Block) {
	for _, declaration := range node.Declarations {
		for _, decl := range declaration {
			b.Visit(decl)
		}
	}
	b.Visit(node.CompoundStatement)
}

func (b *SemanticAnalyzer) VisitVarDecl(node *VarDecl) {
	typeName := node.TypeNode.(*Type).Value.(string)
	typeSymbol := b.CurrentScope.Lookup(typeName, false)
	varName := node.VarNode.(*Var).Value.(string)
	varSymbol := NewVarSymbol(varName, typeSymbol)
	if b.CurrentScope.Lookup(varName, true) != nil {
		panic(b.error(DuplicateId, node.VarNode.(*Var).Token))
	}
	b.CurrentScope.Insert(varSymbol)
}

func (b *SemanticAnalyzer) VisitType(_ *Type) {
	// pass
}

func (b *SemanticAnalyzer) VisitBinOp(node *BinOp) {
	b.Visit(node.Left)
	b.Visit(node.Right)
}

func (b *SemanticAnalyzer) VisitNumberConst(_ *NumberConst) {
	// pass
}

func (b *SemanticAnalyzer) VisitStringConst(_ *StringConst) {
	// pass
}
func (b *SemanticAnalyzer) VisitMessageConst(_ *MessageConst) {
	// pass
}

func (b *SemanticAnalyzer) VisitBooleanConst(_ *BooleanConst) {
	// pass
}

func (b *SemanticAnalyzer) VisitNodeConst(_ *NodeConst) {
	// pass
}

func (b *SemanticAnalyzer) VisitList(node *List) {
	for _, item := range node.Value {
		b.Visit(item)
	}
}

func (b *SemanticAnalyzer) VisitDict(node *Dict) {
	for _, item := range node.Value {
		b.Visit(item)
	}
}

func (b *SemanticAnalyzer) VisitUnaryOp(_ *UnaryOp) {
	// pass
}

func (b *SemanticAnalyzer) VisitCompound(node *Compound) {
	for _, child := range node.Children {
		b.Visit(child)
	}
}

func (b *SemanticAnalyzer) VisitAssign(node *Assign) {
	b.Visit(node.Right)
	b.Visit(node.Left)
}

func (b *SemanticAnalyzer) VisitVar(node *Var) {
	varName := node.Value.(string)
	varSymbol := b.CurrentScope.Lookup(varName, false)

	if varSymbol == nil {
		panic(b.error(IdNotFound, node.Token))
	}
}

func (b *SemanticAnalyzer) VisitNoOp(_ *NoOp) {
	// pass
}

func (b *SemanticAnalyzer) VisitPrint(node *Print) {
	b.Visit(node.Statement)
}

func (b *SemanticAnalyzer) VisitWhile(node *While) {
	// TODO scope
	for _, node := range node.DoBranch {
		b.Visit(node)
	}
}

func (b *SemanticAnalyzer) VisitIf(node *If) {
	// TODO scope
	for _, node := range node.ThenBranch {
		b.Visit(node)
	}
	for _, node := range node.ElseBranch {
		b.Visit(node)
	}
}

func (b *SemanticAnalyzer) VisitLogical(node *Logical) {
	b.Visit(node.Left)
	b.Visit(node.Right)
}
