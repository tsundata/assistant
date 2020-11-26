package interpreter

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
	return fmt.Sprintf("<VarSymbol(name=%s:type=%v)>", s.Name, s.Type)
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

type ProcedureSymbol struct {
	Name         string
	Type         Symbol
	FormalParams []Ast
	BlockAst     Ast
	ScopeLevel   int
}

func NewProcedureSymbol(name string) *ProcedureSymbol {
	return &ProcedureSymbol{Name: name, ScopeLevel: 0}
}

func (s *ProcedureSymbol) String() string {
	return fmt.Sprintf("<ProcedureSymbol(name=%s, parameters=%v)>", s.Name, s.FormalParams)
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
	fmt.Printf("Insert: %s\n", symbol)
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
	if s, ok := symbol.(*ProcedureSymbol); ok {
		name = s.Name
		s.ScopeLevel = t.ScopeLevel
		t.symbols.Set(name, s)
		return
	}
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

func (b *SemanticAnalyzer) error(errorCode ErrorCode, token *Token) error {
	return Error{
		ErrorCode: errorCode,
		Token:     token,
		Message:   fmt.Sprintf("%s -> %v", errorCode, token),
		Type:      SemanticErrorType,
	}
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
	if n, ok := node.(*ProcedureCall); ok {
		b.VisitProcedureCall(n)
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

	for _, param := range node.FormalParams {
		paramType := b.CurrentScope.Lookup(param.(*Param).TypeNode.(*Type).Value.(string), false)
		paramName := param.(*Param).VarNode.(*Var).Value.(string)
		varSymbol := NewVarSymbol(paramName, paramType)
		b.CurrentScope.Insert(varSymbol)
		procSymbol.FormalParams = append(procSymbol.FormalParams, varSymbol)
	}

	b.Visit(node.BlockNode)

	fmt.Println(procedureScope.String())

	b.CurrentScope = b.CurrentScope.EnclosingScope
	fmt.Printf("LEAVE scope: %s\n", procName)

	procSymbol.BlockAst = node.BlockNode
}

func (b *SemanticAnalyzer) VisitBinOp(node *BinOp) {
	b.Visit(node.Left)
	b.Visit(node.Right)
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

func (b *SemanticAnalyzer) VisitProcedureCall(node *ProcedureCall) {
	procSymbol := b.CurrentScope.Lookup(node.ProcName, false)
	var formalParams []Ast
	if procSymbol != nil {
		formalParams = procSymbol.(*ProcedureSymbol).FormalParams
	}
	actualParams := node.ActualParams

	if len(actualParams) != len(formalParams) {
		panic(b.error(WrongParamsNum, node.Token))
	}

	for _, paramNode := range node.ActualParams {
		b.Visit(paramNode)
	}

	node.ProcSymbol = procSymbol
}

func (b *SemanticAnalyzer) VisitWhile(node *While) {
	// TODO scope
	b.Visit(node.DoBranch)
}

func (b *SemanticAnalyzer) VisitIf(node *If) {
	// TODO scope
	b.Visit(node.ThenBranch)
	b.Visit(node.ElseBranch)
}

func (b *SemanticAnalyzer) VisitLogical(node *Logical) {
	b.Visit(node.Left)
	b.Visit(node.Right)
}
