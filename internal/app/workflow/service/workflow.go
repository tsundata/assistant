package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action"
	"github.com/tsundata/assistant/internal/app/workflow/script"
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

func (s *Workflow) RunScript(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	p, err := script.NewParser(script.NewLexer([]rune(payload.GetText())))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	sa := script.NewSemanticAnalyzer()
	sa.Visit(tree)

	i := script.NewInterpreter(tree)
	i.SetClient(s.midClient)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{
		Text: fmt.Sprintf("Tracing\n-------\n %s", i.Stdout()),
	}, nil
}

func (s *Workflow) RunAction(_ context.Context, payload *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	p, err := action.NewParser(action.NewLexer([]rune(payload.GetText())))
	if err != nil {
		return nil, err
	}
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	i := action.NewInterpreter(tree)
	i.SetClient(s.midClient, s.msgClient)
	_, err = i.Interpret()
	if err != nil {
		return nil, err
	}

	return &pb.WorkflowReply{
		Text: fmt.Sprintf("Tracing\n-------\n %s", i.Stdout()),
	}, nil
}
