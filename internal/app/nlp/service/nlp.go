package service

import (
	"context"
	"github.com/mozillazg/go-pinyin"
	"github.com/tsundata/assistant/api/pb"
	"strings"
)

type NLP struct{}

func NewNLP() *NLP {
	return &NLP{}
}

func (s *NLP) Pinyin(_ context.Context, req *pb.TextRequest) (*pb.WordsReply, error) {
	if req.GetText() == "" {
		return &pb.WordsReply{Text: []string{}}, nil
	}
	a := pinyin.NewArgs()
	py := pinyin.Pinyin(req.GetText(), a)
	var result []string
	for _, i := range py {
		result = append(result, strings.Join(i, " "))
	}
	return &pb.WordsReply{Text: result}, nil
}

func (s *NLP) Segmentation(ctx context.Context, req *pb.TextRequest) (*pb.WordsReply, error) {
	panic("implement me")
}
