package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

type NLP struct{}

func NewNLP() *NLP {
	return &NLP{}
}

func (s *NLP) Pinyin(ctx context.Context, req *pb.TextRequest) (*pb.WordsReply, error) {
	panic("implement me")
}

func (s *NLP) Segmentation(ctx context.Context, req *pb.TextRequest) (*pb.WordsReply, error) {
	panic("implement me")
}
