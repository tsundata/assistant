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

type Number struct {
	Token *Token
	Value float64
}

func NewNumber(token *Token) *Number {
	ret := &Number{Token: token}

	if v, ok := token.Value.(int); ok {
		ret.Value = float64(v)
	} else {
		ret.Value = token.Value.(float64)
	}

	return ret
}

type String struct {
	Token *Token
	Value string
}

func NewString(token *Token) *String {
	return &String{Token: token, Value: token.Value.(string)}
}

type Boolean struct {
	Token *Token
	Value bool
}

func NewBoolean(token *Token) *Boolean {
	return &Boolean{Token: token, Value: token.Value.(bool)}
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

type Param struct {
	VarNode  Ast
	TypeNode Ast
}

func NewParam(varNode Ast, typeNode Ast) *Param {
	return &Param{VarNode: varNode, TypeNode: typeNode}
}

type FunctionDecl struct {
	FuncName     string
	FormalParams []Ast
	BlockNode    Ast
	ReturnType   Ast
}

func NewFunctionDecl(funcName string, formalParams []Ast, blockNode Ast, returnType Ast) *FunctionDecl {
	return &FunctionDecl{FuncName: funcName, FormalParams: formalParams, BlockNode: blockNode, ReturnType: returnType}
}

type FunctionCall struct {
	FuncName     string
	ActualParams []Ast
	Token        *Token
	FuncSymbol   Symbol
}

func NewFunctionCall(funcName string, actualParams []Ast, token *Token) *FunctionCall {
	return &FunctionCall{FuncName: funcName, ActualParams: actualParams, Token: token}
}

type If struct {
	Condition  Ast
	ThenBranch []Ast
	ElseBranch []Ast
}

func NewIf(condition Ast, thenBranch []Ast, elseBranch []Ast) *If {
	return &If{Condition: condition, ThenBranch: thenBranch, ElseBranch: elseBranch}
}

type While struct {
	Condition Ast
	DoBranch  []Ast
}

func NewWhile(condition Ast, doBranch []Ast) *While {
	return &While{Condition: condition, DoBranch: doBranch}
}

type Logical struct {
	Left  Ast
	Op    *Token
	Right Ast
}

func NewLogical(left Ast, op *Token, right Ast) *Logical {
	return &Logical{Left: left, Op: op, Right: right}
}

type Print struct {
	Statement Ast
}

func NewPrint(statement Ast) *Print {
	return &Print{Statement: statement}
}

type Return struct {
	Statement Ast
}

func NewReturn(statement Ast) *Return {
	return &Return{Statement: statement}
}
