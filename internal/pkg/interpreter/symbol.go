package interpreter

import (
	"fmt"
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
	return fmt.Sprintf("<%s:%v>", s.Name, s.Type)
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
	return s.Name
}

type SymbolTable struct {
	symbols map[string]Symbol
}

func NewSymbolTable() *SymbolTable {
	table := &SymbolTable{symbols: make(map[string]Symbol)}
	table.Define(NewBuiltinTypeSymbol("INTEGER"))
	table.Define(NewBuiltinTypeSymbol("REAL"))
	return table
}

func (t *SymbolTable) String() string {
	return fmt.Sprintf("Symbols: %v\n", t.symbols)
}

func (t *SymbolTable) Define(symbol Symbol) {
	fmt.Printf("Define: %v\n", symbol)
	var name string
	if s, ok := symbol.(*VarSymbol); ok {
		name = s.Name
	}
	if s, ok := symbol.(*BuiltinTypeSymbol); ok {
		name = s.Name
	}
	t.symbols[name] = symbol
}

func (t *SymbolTable) Lookup(name string) Symbol {
	fmt.Printf("Lookup: %s\n", name)

	if _, ok := t.symbols[name]; ok {
		return t.symbols[name].(Symbol)
	}
	return nil
}

type SymbolTableBuilder struct {
	symbolTable *SymbolTable
}

func NewSymbolTableBuilder() *SymbolTableBuilder {
	return &SymbolTableBuilder{symbolTable: NewSymbolTable()}
}

func (b *SymbolTableBuilder) Visit(node Ast) {
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
	if n, ok := node.(*Num); ok {
		b.VisitNum(n)
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
}

func (b *SymbolTableBuilder) VisitBlock(node *Block) {
	for _, declaration := range node.Declarations {
		for _, decl := range declaration {
			b.Visit(decl)
		}
	}
	b.Visit(node.CompoundStatement)
}

func (b *SymbolTableBuilder) VisitProgram(node *Program) {
	b.Visit(node.Block)
}

func (b *SymbolTableBuilder) VisitBinOp(node *BinOp) {
	b.Visit(node.Left)
	b.Visit(node.Right)
}

func (b *SymbolTableBuilder) VisitNum(node *Num) {
	// pass
}

func (b *SymbolTableBuilder) VisitUnaryOp(node *UnaryOp) {
	b.Visit(node.Expr)
}

func (b *SymbolTableBuilder) VisitCompound(node *Compound) {
	for _, child := range node.Children {
		b.Visit(child)
	}
}

func (b *SymbolTableBuilder) VisitNoOp(node *NoOp) {
	// pass
}

func (b *SymbolTableBuilder) VisitVarDecl(node *VarDecl) {
	typeName := node.TypeNode.(*Type).Value.(string)
	typeSymbol := b.symbolTable.Lookup(typeName)
	varName := node.VarNode.(*Var).Value.([]rune)
	varSymbol := NewVarSymbol(string(varName), typeSymbol)
	b.symbolTable.Define(varSymbol)
}

func (b *SymbolTableBuilder) VisitAssign(node *Assign) {
	varName := node.Left.(*Var).Value.([]rune)
	varSymbol := b.symbolTable.Lookup(string(varName))

	if varSymbol == nil {
		panic(fmt.Sprintf("error var symbol %s %v", string(varName), varSymbol))
	}

	b.Visit(node.Right)
}

func (b *SymbolTableBuilder) VisitVar(node *Var) {
	varName := node.Value.([]rune)
	varSymbol := b.symbolTable.Lookup(string(varName))

	if varSymbol == nil {
		panic(fmt.Sprintf("error var symbol %s %v", string(varName), varSymbol))
	}
}

func (b *SymbolTableBuilder) VisitProcedureDecl() {
	// pass
}
