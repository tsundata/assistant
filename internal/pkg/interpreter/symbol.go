package interpreter

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/utils/collection"
	"strings"
)

type Symbol interface{}

type VarSymbol struct {
	Name string
	Type Symbol
}

func NewVarSymbol(name string, t Symbol) *VarSymbol {
	s := &VarSymbol{}
	s.Name = name
	s.Type = t
	return s
}

func (s *VarSymbol) String() string {
	return fmt.Sprintf("<VarSymbol(name=%s:type=%v)>", s.Name, s.Type)
}

type BuiltinTypeSymbol struct {
	Name string
	Type Symbol
}

func NewBuiltinTypeSymbol(name string) *BuiltinTypeSymbol {
	s := &BuiltinTypeSymbol{}
	s.Name = name
	return s
}

func (s *BuiltinTypeSymbol) String() string {
	return fmt.Sprintf("<BuiltinTypeSymbol(name=%s)>", s.Name)
}

type ProcedureSymbol struct {
	Name   string
	Type   Symbol
	Params []Ast
}

func NewProcedureSymbol(name string) *ProcedureSymbol {
	return &ProcedureSymbol{Name: name}
}

func (s *ProcedureSymbol) String() string {
	return fmt.Sprintf("<ProcedureSymbol(name=%s, parameters=%v)>", s.Name, s.Params)
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
	table.Insert(NewBuiltinTypeSymbol("REAL"))
	return table
}

func (t *ScopedSymbolTable) String() string {
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
	fmt.Printf("Insert: %s\n", symbol)
	var name string
	if s, ok := symbol.(*VarSymbol); ok {
		name = s.Name
	}
	if s, ok := symbol.(*BuiltinTypeSymbol); ok {
		name = s.Name
	}
	t.symbols.Set(name, symbol)
}

func (t *ScopedSymbolTable) Lookup(name string, currentScopeOnly bool) Symbol {
	fmt.Printf("Lookup: %s. (Scope name: %s)\n", name, t.ScopeName)
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

func (b *SemanticAnalyzer) Visit(node Ast) {
	if n, ok := node.(*Block); ok {
		b.VisitBlock(n)
		return
	}
	if n, ok := node.(*Program); ok {
		b.VisitProgram(n)
		return
	}
	if n, ok := node.(*Compound); ok {
		b.VisitCompound(n)
		return
	}
	if n, ok := node.(*NoOp); ok {
		b.VisitNoOp(n)
		return
	}
	if n, ok := node.(*ProcedureDecl); ok {
		b.VisitProcedureDecl(n)
		return
	}
	if n, ok := node.(*BinOp); ok {
		b.VisitBinOp(n)
		return
	}
	if n, ok := node.(*VarDecl); ok {
		b.VisitVarDecl(n)
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
}

func (b *SemanticAnalyzer) VisitBlock(node *Block) {
	for _, declaration := range node.Declarations {
		for _, decl := range declaration {
			b.Visit(decl)
		}
	}
	b.Visit(node.CompoundStatement)
}

func (b *SemanticAnalyzer) VisitProgram(node *Program) {
	fmt.Println("ENTER scope: global")
	globalScope := NewScopedSymbolTable("global", 1, b.CurrentScope)
	b.CurrentScope = globalScope

	// visit subtree
	b.Visit(node.Block)

	fmt.Println(globalScope.String())

	b.CurrentScope = b.CurrentScope.EnclosingScope
	fmt.Println("LEAVE scope: global")
}

func (b *SemanticAnalyzer) VisitCompound(node *Compound) {
	for _, child := range node.Children {
		b.Visit(child)
	}
}

func (b *SemanticAnalyzer) VisitNoOp(node *NoOp) {
	// pass
}

func (b *SemanticAnalyzer) VisitProcedureDecl(node *ProcedureDecl) {
	procName := node.ProcName
	procSymbol := NewProcedureSymbol(procName)
	b.CurrentScope.Insert(procSymbol)

	fmt.Printf("ENTER scope: %s\n", procName)
	procedureScope := NewScopedSymbolTable(procName, b.CurrentScope.ScopeLevel+1, b.CurrentScope)
	b.CurrentScope = procedureScope

	for _, param := range node.Params {
		paramType := b.CurrentScope.Lookup(param.(*Param).TypeNode.(*Type).Value.(string), false)
		paramName := param.(*Param).VarNode.(*Var).Value.([]rune)
		varSymbol := NewVarSymbol(string(paramName), paramType)
		b.CurrentScope.Insert(varSymbol)
		procSymbol.Params = append(procSymbol.Params, varSymbol)
	}

	b.Visit(node.BlockNode)

	fmt.Println(procedureScope.String())

	b.CurrentScope = b.CurrentScope.EnclosingScope
	fmt.Printf("LEAVE scope: %s\n", procName)
}

func (b *SemanticAnalyzer) VisitBinOp(node *BinOp) {
	b.Visit(node.Left)
	b.Visit(node.Right)
}

func (b *SemanticAnalyzer) VisitVarDecl(node *VarDecl) {
	typeName := node.TypeNode.(*Type).Value.(string)
	typeSymbol := b.CurrentScope.Lookup(typeName, false)
	varName := node.VarNode.(*Var).Value.([]rune)
	varSymbol := NewVarSymbol(string(varName), typeSymbol)
	if b.CurrentScope.Lookup(string(varName), true) != nil {
		panic(fmt.Sprintf("Error: Duplicate identifier '%s' found", string(varName)))
	}
	b.CurrentScope.Insert(varSymbol)
}

func (b *SemanticAnalyzer) VisitAssign(node *Assign) {
	b.Visit(node.Right)
	b.Visit(node.Left)
}

func (b *SemanticAnalyzer) VisitVar(node *Var) {
	varName := node.Value.([]rune)
	varSymbol := b.CurrentScope.Lookup(string(varName), false)

	if varSymbol == nil {
		panic(fmt.Sprintf("Error: Symbol(identifier) not found '%s'", string(varName)))
	}
}
