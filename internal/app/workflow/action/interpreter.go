package action

import (
	ctx "context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/context"
	"github.com/tsundata/assistant/internal/app/workflow/action/opcode"
	"log"
	"strings"
)

type Opcoder interface {
	Run(*context.Context, []interface{}) (interface{}, error)
}

func runOpcode(ctx *context.Context, name string, params []interface{}) (interface{}, error) {
	var o Opcoder
	switch strings.ToLower(name) {
	case "get":
		o = opcode.NewGet()
	case "count":
		o = opcode.NewCount()
	default:
		return nil, errors.New("not opcode")
	}
	return o.Run(ctx, params)
}

type Interpreter struct {
	tree Ast
	ctx  *context.Context
}

func NewInterpreter(tree Ast) *Interpreter {
	return &Interpreter{tree: tree, ctx: context.NewContext()}
}

func (i *Interpreter) SetClient(midClient pb.MiddleClient, msgClient pb.MessageClient) {
	i.ctx.MidClient = midClient
	i.ctx.MsgClient = msgClient
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

	// Run
	debugLog(fmt.Sprintf("Run: Opecode %v", node.ID))
	debugLog(fmt.Sprintf("%+v", params))
	res, err := runOpcode(i.ctx, node.ID.(*Token).Value.(string), params)
	if err != nil {
		log.Println(err)
		return -1
	}
	debugLog(fmt.Sprintf("%+v\n", res))

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
	if i.ctx.MsgClient != nil {
		reply, err := i.ctx.MsgClient.Get(ctx.Background(), &pb.MessageRequest{Id: node.Value.(int64)})
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
