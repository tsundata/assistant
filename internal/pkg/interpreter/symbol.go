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

type SymbolTable struct {
	symbols *collection.OrderedDict
}

func NewSymbolTable() *SymbolTable {
	table := &SymbolTable{symbols: collection.NewOrderedDict()}
	table.Insert(NewBuiltinTypeSymbol("INTEGER"))
	table.Insert(NewBuiltinTypeSymbol("REAL"))
	return table
}

func (t *SymbolTable) String() string {
	var lines []string
	i := 0
	for v := range t.symbols.Iterate() {
		i++
		lines = append(lines, fmt.Sprintf("%6d: %v", i, v))
	}

	return fmt.Sprintf("Symbol table contents\n%s\n", strings.Join(lines, "\n"))
}

func (t *SymbolTable) Lookup(name string) Symbol {
	fmt.Printf("Lookup: %s\n", name)
	s := t.symbols.Get(name)
	if s != nil {
		return s.(Symbol)
	}
	return nil
}

func (t *SymbolTable) Insert(symbol Symbol) {
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

type SemanticAnalyzer struct {
	symbolTable *SymbolTable
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{symbolTable: NewSymbolTable()}
}

func (b *SemanticAnalyzer) Visit(node Ast) {
	if n, ok := node.(*Program); ok {
		b.VisitProgram(n)
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
	if n, ok := node.(*BinOp); ok {
		b.VisitBinOp(n)
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
	b.Visit(node.Block)
}

func (b *SemanticAnalyzer) VisitCompound(node *Compound) {
	for _, child := range node.Children {
		b.Visit(child)
	}
}

func (b *SemanticAnalyzer) VisitNoOp(node *NoOp) {
	// pass
}

func (b *SemanticAnalyzer) VisitBinOp(node *BinOp) {
	b.Visit(node.Left)
	b.Visit(node.Right)
}

func (b *SemanticAnalyzer) VisitVarDecl(node *VarDecl) {
	typeName := node.TypeNode.(*Type).Value.(string)
	typeSymbol := b.symbolTable.Lookup(typeName)
	varName := node.VarNode.(*Var).Value.([]rune)
	varSymbol := NewVarSymbol(string(varName), typeSymbol)
	if b.symbolTable.Lookup(string(varName)) != nil {
		panic(fmt.Sprintf("Error: Duplicate identifier '%s' found", string(varName)))
	}
	b.symbolTable.Insert(varSymbol)
}

func (b *SemanticAnalyzer) VisitAssign(node *Assign) {
	varName := node.Left.(*Var).Value.([]rune)
	varSymbol := b.symbolTable.Lookup(string(varName))

	if varSymbol == nil {
		panic(fmt.Sprintf("error var symbol %s %v", string(varName), varSymbol))
	}

	b.Visit(node.Right)
}

func (b *SemanticAnalyzer) VisitVar(node *Var) {
	varName := node.Value.([]rune)
	varSymbol := b.symbolTable.Lookup(string(varName))

	if varSymbol == nil {
		panic(fmt.Sprintf("Error: Symbol(identifier) not found '%s'", string(varName)))
	}
}
