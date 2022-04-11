package action

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/action/opcode"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"strings"
)

type Interpreter struct {
	Debug  bool
	tree   Ast
	stdout []interface{}
	ctx    context.Context
	inCtx  *inside.Context
	comp   component.Component
}

func NewInterpreter(ctx context.Context, tree Ast) *Interpreter {
	return &Interpreter{ctx: ctx, tree: tree, inCtx: inside.NewComponent()}
}

func (i *Interpreter) SetMessage(message pb.Message) {
	i.inCtx.SetMessage(message)
}

func (i *Interpreter) SetComponent(comp component.Component) {
	i.comp = comp
}

func (i *Interpreter) Visit(node Ast) interface{} {
	if n, ok := node.(*Program); ok {
		return i.VisitProgram(n)
	}
	if n, ok := node.(*Opcode); ok {
		return i.VisitOpcode(n)
	}
	if n, ok := node.(*IntegerConst); ok {
		return i.VisitIntegerConst(n)
	}
	if n, ok := node.(*FloatConst); ok {
		return i.VisitFloatConst(n)
	}
	if n, ok := node.(*StringConst); ok {
		return i.VisitStringConst(n)
	}
	if n, ok := node.(*BooleanConst); ok {
		return i.VisitBooleanConst(n)
	}
	if n, ok := node.(*MessageConst); ok {
		return i.VisitMessageConst(n)
	}
	if n, ok := node.(*Var); ok {
		return i.VisitVar(n)
	}
	if n, ok := node.(*NoOp); ok {
		return i.VisitNoOp(n)
	}

	return 0
}

func (i *Interpreter) VisitProgram(node *Program) interface{} {
	// main
	var result interface{}
	for _, item := range node.Statements {
		result = i.Visit(item)
	}

	return result
}

func (i *Interpreter) VisitOpcode(node *Opcode) float64 {
	var params []interface{}
	for _, item := range node.Expressions {
		params = append(params, i.Visit(item))
	}

	name := node.ID.(*Token).Value.(string)
	debugLog(fmt.Sprintf("Run Opecode: %v", node.ID))
	debugLog(fmt.Sprintf("params: %+v", params))
	input := i.inCtx.Value
	debugLog(fmt.Sprintf("context: %+v", input))
	op := opcode.NewOpcode(name)
	if op == nil {
		return 0
	}

	// Async opcode
	if op.Type() == opcode.TypeAsync {
		return 0
	}
	// Cond opcode
	if op.Type() != opcode.TypeCond && !i.inCtx.Continue {
		debugLog(fmt.Sprintf("skip: %s", name))
		return 0
	}

	// Run
	res, err := op.Run(i.ctx, i.inCtx, i.comp, params)
	i.stdout = append(i.stdout, opcodeLog(name, params, input, res, err))
	if err != nil {
		if i.comp.GetLogger() != nil {
			i.comp.GetLogger().Error(err)
		}
		return -1
	}
	debugLog(fmt.Sprintf("result: %+v\n", res))
	i.Debug = i.inCtx.Debug

	return 0
}

func (i *Interpreter) VisitIntegerConst(node *IntegerConst) int64 {
	return node.Value
}

func (i *Interpreter) VisitFloatConst(node *FloatConst) float64 {
	return node.Value
}

func (i *Interpreter) VisitStringConst(node *StringConst) string {
	return node.Value
}

func (i *Interpreter) VisitBooleanConst(node *BooleanConst) bool {
	return node.Value
}

func (i *Interpreter) VisitMessageConst(node *MessageConst) interface{} {
	if i.comp.Message() != nil {
		reply, err := i.comp.Message().GetBySequence(context.Background(), &pb.MessageRequest{
			Message: &pb.Message{
				UserId:   i.inCtx.Message.UserId,
				Sequence: node.Value.(int64),
			},
		})
		if err != nil {
			i.comp.GetLogger().Error(err)
			return ""
		}
		return reply.Message.GetText()
	}
	return ""
}

func (i *Interpreter) VisitVar(_ *Var) interface{} {
	return nil
}

func (i *Interpreter) VisitNoOp(_ *NoOp) interface{} {
	return 0
}

func (i *Interpreter) Interpret() (interface{}, error) {
	if i.tree == nil {
		return 0, errors.New("error ast tree")
	}
	return i.Visit(i.tree), nil
}

func (i *Interpreter) Stdout() string {
	var out strings.Builder
	for _, line := range i.stdout {
		out.WriteString(fmt.Sprintf("%v\n", line))
	}
	return out.String()
}
