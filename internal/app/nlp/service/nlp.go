package service

import (
	"context"
	"github.com/go-ego/gse"
	"github.com/mozillazg/go-pinyin"
	"github.com/tsundata/assistant/api/pb"
	"strings"
)

type NLP struct {
	seg *gse.Segmenter
}

func NewNLP(seg *gse.Segmenter) *NLP {
	return &NLP{seg: seg}
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

func (s *NLP) Segmentation(_ context.Context, req *pb.TextRequest) (*pb.WordsReply, error) {
	if req.GetText() == "" {
		return &pb.WordsReply{Text: []string{}}, nil
	}
	result := s.seg.Cut(req.GetText(), true)
	return &pb.WordsReply{Text: result}, nil
}
