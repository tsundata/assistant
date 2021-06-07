package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

type Finance struct{}

func NewFinance() *Finance {
	return &Finance{}
}

func (f Finance) CreateBill(ctx context.Context, req *pb.BillRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (f Finance) GetBill(ctx context.Context, req *pb.BillRequest) (*pb.BillReply, error) {
	panic("implement me")
}

func (f Finance) GetBills(ctx context.Context, req *pb.BillRequest) (*pb.BillsReply, error) {
	panic("implement me")
}

func (f Finance) DeleteBill(ctx context.Context, req *pb.BillRequest) (*pb.StateReply, error) {
	panic("implement me")
}
