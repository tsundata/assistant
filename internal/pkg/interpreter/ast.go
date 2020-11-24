package interpreter

type Ast interface{}

type BinOp struct {
	Left  Ast
	Token *Token
	Op    *Token
	Right Ast
}

func NewBinOp(left Ast, op *Token, right Ast) *BinOp {
	return &BinOp{Left: left, Token: op, Op: op, Right: right}
}

type Num struct {
	Token *Token
	Value float64
}

func NewNum(token *Token) *Num {
	ret := &Num{Token: token}

	if v, ok := token.Value.(int); ok {
		ret.Value = float64(v)
	} else {
		ret.Value = token.Value.(float64)
	}

	return ret
}

type UnaryOp struct {
	Op   *Token
	Expr Ast
}

func NewUnaryOp(op *Token, expr Ast) *UnaryOp {
	return &UnaryOp{Op: op, Expr: expr}
}

type Compound struct {
	Children []Ast
}

func NewCompound() *Compound {
	return &Compound{}
}

type Assign struct {
	Left  Ast
	Op    *Token
	Right Ast
}

func NewAssign(left Ast, op *Token, right Ast) *Assign {
	return &Assign{Left: left, Op: op, Right: right}
}

type Var struct {
	Token *Token
	Value interface{}
}

func NewVar(token *Token) *Var {
	return &Var{Token: token, Value: token.Value}
}

type NoOp struct{}

func NewNoOp() *NoOp {
	return &NoOp{}
}

type Program struct {
	Name  string
	Block Ast
}

func NewProgram(name string, block Ast) *Program {
	return &Program{Name: name, Block: block}
}

type Block struct {
	Declarations      [][]Ast
	CompoundStatement Ast
}

func NewBlock(declarations [][]Ast, compoundStatement Ast) *Block {
	return &Block{Declarations: declarations, CompoundStatement: compoundStatement}
}

type VarDecl struct {
	VarNode  Ast
	TypeNode Ast
}

func NewVarDecl(varNode Ast, typeNode Ast) *VarDecl {
	return &VarDecl{VarNode: varNode, TypeNode: typeNode}
}

type Type struct {
	Token *Token
	Value interface{}
}

func NewType(token *Token) *Type {
	return &Type{Token: token, Value: token.Value}
}

type ProcedureDecl struct {
	ProcName  string
	BlockNode Ast
}

func NewProcedureDecl(procName string, blockNode Ast) *ProcedureDecl {
	return &ProcedureDecl{ProcName: procName, BlockNode: blockNode}
}
