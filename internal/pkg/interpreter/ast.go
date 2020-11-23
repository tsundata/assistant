package interpreter

type Ast struct{}

type BinOp struct {
	Ast
	Left  interface{}
	Token *Token
	Op    *Token
	Right interface{}
}

func NewBinOp(left interface{}, op *Token, right interface{}) *BinOp {
	return &BinOp{Left: left, Token: op, Op: op, Right: right}
}

type Num struct {
	Ast
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
	Ast
	Op   *Token
	Expr interface{}
}

func NewUnaryOp(op *Token, expr interface{}) *UnaryOp {
	return &UnaryOp{Op: op, Expr: expr}
}

type Compound struct {
	Ast
	Children []interface{}
}

func NewCompound() *Compound {
	return &Compound{}
}

type Assign struct {
	Ast
	Left  interface{}
	Op    *Token
	Right interface{}
}

func NewAssign(left interface{}, op *Token, right interface{}) *Assign {
	return &Assign{Left: left, Op: op, Right: right}
}

type Var struct {
	Ast
	Token *Token
	Value interface{}
}

func NewVar(token *Token) *Var {
	return &Var{Token: token, Value: token.Value}
}

type NoOp struct {
	Ast
}

func NewNoOp() *NoOp {
	return &NoOp{}
}

type Program struct {
	Name  string
	Block interface{}
}

func NewProgram(name string, block interface{}) *Program {
	return &Program{Name: name, Block: block}
}

type Block struct {
	Declarations      [][]interface{}
	CompoundStatement interface{}
}

func NewBlock(declarations [][]interface{}, compoundStatement interface{}) *Block {
	return &Block{Declarations: declarations, CompoundStatement: compoundStatement}
}

type VarDecl struct {
	VarNode  interface{}
	TypeNode interface{}
}

func NewVarDecl(varNode interface{}, typeNode interface{}) *VarDecl {
	return &VarDecl{VarNode: varNode, TypeNode: typeNode}
}

type Type struct {
	Token *Token
	Value interface{}
}

func NewType(token *Token) *Type {
	return &Type{Token: token, Value: token.Value}
}
