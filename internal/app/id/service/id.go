package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/tsundata/assistant/api/pb"
)

type Id struct {}

func NewId() *Id {
	return &Id{}
}

func (s *Id) GetGlobalId(_ context.Context, _ *pb.IdRequest) (*pb.IdReply, error) {
	sNode, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	return &pb.IdReply{Id: sNode.Generate().Int64()}, nil
}
