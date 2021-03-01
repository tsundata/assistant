package script

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

type NumberConst struct {
	Token *Token
	Value float64
}

func NewNumberConst(token *Token) *NumberConst {
	ret := &NumberConst{Token: token}

	if v, ok := token.Value.(int); ok {
		ret.Value = float64(v)
	} else {
		ret.Value = token.Value.(float64)
	}

	return ret
}

type StringConst struct {
	Token *Token
	Value string
}

func NewStringConst(token *Token) *StringConst {
	return &StringConst{Token: token, Value: token.Value.(string)}
}

type BooleanConst struct {
	Token *Token
	Value bool
}

func NewBooleanConst(token *Token) *BooleanConst {
	return &BooleanConst{Token: token, Value: token.Value.(bool)}
}

type List struct {
	Token *Token
	Value []Ast
}

func NewList(token *Token) *List {
	return &List{Token: token, Value: token.Value.([]Ast)}
}

type Dict struct {
	Token *Token
	Value map[string]Ast
}

func NewDict(token *Token) *Dict {
	return &Dict{Token: token, Value: token.Value.(map[string]Ast)}
}

type MessageConst struct {
	Token *Token
	Value interface{}
}

func NewMessageConst(token *Token) *MessageConst {
	return &MessageConst{Token: token, Value: token.Value.(int)}
}

type NodeConst struct {
	Token *Token
	Value interface{}
}

func NewNodeConst(token *Token) *NodeConst {
	return &NodeConst{Token: token, Value: token.Value.(string)}
}

type Flow struct {
	Nodes []Ast
}

func NewFlow(nodes []Ast) *Flow {
	return &Flow{Nodes: nodes}
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
	Name      string
	Nodes     map[string]Ast
	Workflows map[string]Ast
}

func NewProgram(name string, nodes map[string]Ast, workflows map[string]Ast) *Program {
	return &Program{Name: name, Nodes: nodes, Workflows: workflows}
}

type Node struct {
	Name       string
	Regular    string
	With       Ast
	Parameters map[string]interface{}
	Secret     string
}

func NewNode(name string, regular string, with Ast, secret string) *Node {
	return &Node{Name: name, Regular: regular, With: with, Secret: secret}
}

type Workflow struct {
	Name      string
	Scenarios Ast
}

func NewWorkflow(name string, scenarios Ast) *Workflow {
	return &Workflow{Name: name, Scenarios: scenarios}
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
