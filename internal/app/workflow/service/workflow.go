package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/interpreter"
	"go.etcd.io/etcd/clientv3"
)

type Workflow struct {
	etcd      *clientv3.Client
	midClient pb.MiddleClient
	msgClient pb.MessageClient
}

func NewWorkflow(etcd *clientv3.Client, midClient pb.MiddleClient, msgClient pb.MessageClient) *Workflow {
	return &Workflow{etcd: etcd, midClient: midClient, msgClient: msgClient}
}

func (s *Workflow) Run(ctx context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	var script []rune
	if payload.GetId() > 0 {
		reply, err := s.msgClient.Get(context.Background(), &pb.MessageRequest{Id: payload.GetId()})
		if err != nil {
			return nil, err
		}
		script = []rune(reply.GetText())
	} else if payload.GetText() != "" {
		script = []rune(payload.GetText())
	}

	p, err := interpreter.NewParser(interpreter.NewLexer(script))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	sa := interpreter.NewSemanticAnalyzer()
	sa.Visit(tree)

	i := interpreter.NewInterpreter(tree)
	i.SetClient(s.midClient)
	r, err := i.Interpret()
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{
		Text: fmt.Sprintf("%f", r),
	}, nil
}
