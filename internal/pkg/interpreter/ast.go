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
	Value interface{}
}

func NewNum(token *Token) *Num {
	return &Num{Token: token, Value: token.Value}
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
