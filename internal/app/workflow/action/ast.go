package action

type Ast interface{}

type IntegerConst struct {
	Token *Token
	Value int64
}

func NewIntegerConst(token *Token) *IntegerConst {
	return &IntegerConst{Token: token, Value: token.Value.(int64)}
}

type FloatConst struct {
	Token *Token
	Value float64
}

func NewFloatConst(token *Token) *FloatConst {
	return &FloatConst{Token: token, Value: token.Value.(float64)}
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

type MessageConst struct {
	Token *Token
	Value interface{}
}

func NewMessageConst(token *Token) *MessageConst {
	return &MessageConst{Token: token, Value: token.Value.(int64)}
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
	Name       string
	Statements []Ast
}

func NewProgram(name string, statements []Ast) *Program {
	return &Program{Name: name, Statements: statements}
}

type Opcode struct {
	ID          Ast
	Expressions []Ast
}

func NewOpcode(ID Ast, expressions []Ast) *Opcode {
	return &Opcode{ID: ID, Expressions: expressions}
}
