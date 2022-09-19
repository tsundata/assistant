package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/repository"
	"gorm.io/gorm"
)

type IdSvcClient interface {
	GetGlobalId(context.Context, *pb.GetGlobalIdRequest) (*pb.GetGlobalIdReply, error)
}

type Id struct {
	repo repository.IdRepository
}

func NewId(repo repository.IdRepository) IdSvcClient {
	return &Id{repo: repo}
}

func (s *Id) GetGlobalId(ctx context.Context, payload *pb.GetGlobalIdRequest) (*pb.GetGlobalIdReply, error) {
	node, err := s.repo.GetOrCreateNode(ctx, &pb.Node{Ip: payload.Ip, Port: payload.Port})
	if err != nil {
		return nil, err
	}
	if node.Id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	sNode, err := snowflake.NewNode(node.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetGlobalIdReply{Id: sNode.Generate().Int64()}, nil
}
