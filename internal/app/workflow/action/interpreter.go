package action

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/app/workflow/action/opcode"
	"log"
	"strings"
)

type Interpreter struct {
	tree   Ast
	stdout []interface{}
	Ctx    *inside.Context
}

func NewInterpreter(tree Ast) *Interpreter {
	return &Interpreter{tree: tree, Ctx: inside.NewContext()}
}

func (i *Interpreter) SetClient(midClient pb.MiddleClient, msgClient pb.MessageClient, wfClient pb.WorkflowClient, taskClient pb.TaskClient) {
	i.Ctx.MidClient = midClient
	i.Ctx.MsgClient = msgClient
	i.Ctx.WfClient = wfClient
	i.Ctx.TaskClient = taskClient
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
	// Special opcode
	if name == "webhook" || name == "cron" {
		return 0
	}

	// Run
	debugLog(fmt.Sprintf("Run Opecode: %v", node.ID))
	debugLog(fmt.Sprintf("params: %+v", params))
	input := i.Ctx.Value
	debugLog(fmt.Sprintf("context: %+v", input))
	res, err := opcode.RunOpcode(i.Ctx, name, params)
	i.stdout = append(i.stdout, opcodeLog(name, params, input, res, err))
	if err != nil {
		log.Println(err)
		return -1
	}
	debugLog(fmt.Sprintf("result: %+v\n", res))

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
	if i.Ctx.MsgClient != nil {
		reply, err := i.Ctx.MsgClient.Get(context.Background(), &pb.MessageRequest{Id: node.Value.(int64)})
		if err != nil {
			log.Println(err)
			return ""
		}
		return reply.GetText()
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
