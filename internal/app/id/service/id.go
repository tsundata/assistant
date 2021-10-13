package service

import (
	"context"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/id/repository"
)

type Id struct {
	repo repository.IdRepository
}

func NewId(repo repository.IdRepository) *Id {
	return &Id{repo: repo}
}

func (s *Id) GetGlobalId(ctx context.Context, payload *pb.IdRequest) (*pb.IdReply, error) {
	node, err := s.repo.GetOrCreateNode(ctx, &pb.Node{Ip: payload.Ip, Port: payload.Port})
	if err != nil {
		return nil, err
	}
	if node.Id <= 0 {
		return nil, errors.New("error node")
	}
	sNode, err := snowflake.NewNode(node.Id)
	if err != nil {
		return nil, err
	}
	return &pb.IdReply{Id: sNode.Generate().Int64()}, nil
}
