package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/interpreter"
	"go.etcd.io/etcd/clientv3"
)

type Workflow struct {
	etcd *clientv3.Client
	midClient pb.MiddleClient
}

func NewWorkflow(etcd *clientv3.Client, midClient pb.MiddleClient) *Workflow {
	return &Workflow{etcd: etcd, midClient: midClient}
}

func (s *Workflow) Run(ctx context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	p, err := interpreter.NewParser(interpreter.NewLexer([]rune(payload.GetText())))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	sa := interpreter.NewSemanticAnalyzer()
	sa.Visit(tree)

	i := interpreter.NewInterpreter(tree, s.midClient)
	r, err := i.Interpret()
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{
		Text: fmt.Sprintf("%f", r),
	}, nil
}
