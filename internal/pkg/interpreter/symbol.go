package interpreter

import (
	"fmt"
)

type Symbol interface {
	String() string
}

type VarSymbol struct {
	Name string
	Type interface{}
	Symbol
}

func NewVarSymbol(name string, t interface{}) *VarSymbol {
	s := &VarSymbol{}
	s.Name = name
	s.Type = t
	return s
}

func (s *VarSymbol) String() string {
	// return fmt.Sprintf("<%s:%v>", s.Name, s.Type)
	return s.Name
}

type BuiltinTypeSymbol struct {
	Name string
	Type interface{}
	Symbol
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
	symbols map[string]interface{}
}

func NewSymbolTable() *SymbolTable {
	table := &SymbolTable{symbols: make(map[string]interface{})}
	table.Define(NewBuiltinTypeSymbol("INTEGER"))
	table.Define(NewBuiltinTypeSymbol("REAL"))
	return table
}

func (t *SymbolTable) String() string {
	return fmt.Sprintf("Symbols: %v\n", t.symbols)
}

func (t *SymbolTable) Define(symbol Symbol) {
	fmt.Printf("Define: %v\n", symbol)
	t.symbols[symbol.String()] = symbol
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

func (b *SymbolTableBuilder) Visit(node interface{}) float64 {
	if n, ok := node.(*Program); ok {
		return b.VisitProgram(n)
	}
	if n, ok := node.(*Block); ok {
		return b.VisitBlock(n)
	}
	if n, ok := node.(*VarDecl); ok {
		return b.VisitVarDecl(n)
	}
	if n, ok := node.(*BinOp); ok {
		return b.VisitBinOp(n)
	}
	if n, ok := node.(*Num); ok {
		return b.VisitNum(n)
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

	return 0
}

func (b *SymbolTableBuilder) VisitBlock(node *Block) float64 {
	for _, declaration := range node.Declarations {
		for _, decl := range declaration {
			b.Visit(decl)
		}
	}
	b.Visit(node.CompoundStatement)
	return 0
}

func (b *SymbolTableBuilder) VisitProgram(node *Program) float64 {
	return b.Visit(node.Block)
}

func (b *SymbolTableBuilder) VisitBinOp(node *BinOp) float64 {
	b.Visit(node.Left)
	b.Visit(node.Right)
	return 0
}

func (b *SymbolTableBuilder) VisitNum(node *Num) float64 {
	return 0
}

func (b *SymbolTableBuilder) VisitUnaryOp(node *UnaryOp) float64 {
	b.Visit(node.Expr)
	return 0
}

func (b *SymbolTableBuilder) VisitCompound(node *Compound) float64 {
	for _, child := range node.Children {
		b.Visit(child)
	}
	return 0
}

func (b *SymbolTableBuilder) VisitNoOp(node *NoOp) float64 {
	return 0
}

func (b *SymbolTableBuilder) VisitVarDecl(node *VarDecl) float64 {
	typeName := node.TypeNode.(*Type).Value.(string)
	typeSymbol := b.symbolTable.Lookup(typeName)
	varName := node.VarNode.(*Var).Value.([]rune)
	varSymbol := NewVarSymbol(string(varName), typeSymbol)
	b.symbolTable.Define(varSymbol)
	return 0
}

func (b *SymbolTableBuilder) VisitAssign(node *Assign) float64 {
	varName := node.Left.(*Var).Value.([]rune)
	varSymbol := b.symbolTable.Lookup(string(varName))

	if varSymbol == nil {
		panic(fmt.Sprintf("error var symbol %s %v", string(varName), varSymbol))
	}

	b.Visit(node.Right)
	return 0
}

func (b *SymbolTableBuilder) VisitVar(node *Var) float64 {
	varName := node.Value.([]rune)
	varSymbol := b.symbolTable.Lookup(string(varName))

	if varSymbol == nil {
		panic(fmt.Sprintf("error var symbol %s %v", string(varName), varSymbol))
	}

	return 0
}
