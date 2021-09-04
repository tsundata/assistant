package service

import (
	"context"
	"github.com/go-ego/gse"
	"github.com/mozillazg/go-pinyin"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/nlp/classifier"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"strings"
)

type NLP struct {
	conf   *config.AppConfig
	logger log.Logger
	seg    *gse.Segmenter
}

func NewNLP(conf *config.AppConfig, logger log.Logger) *NLP {
	// gse preload dict
	seg, err := gse.New("zh", "alpha")
	if err != nil {
		logger.Error(err)
	}
	return &NLP{
		conf:   conf,
		logger: logger,
		seg:    &seg,
	}
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

func (s *NLP) Classifier(_ context.Context, req *pb.TextRequest) (*pb.TextReply, error) {
	rules, err := classifier.ReadRulesConfig(s.conf)
	if err != nil {
		return nil, err
	}

	c := classifier.NewClassifier()
	err = c.SetRules(rules)
	if err != nil {
		return nil, err
	}

	if req.GetText() == "" {
		return nil, errors.New("error text")
	}

	res, err := c.Do(req.GetText())
	if err != nil {
		if errors.Is(err, classifier.ErrEmpty) {
			return &pb.TextReply{Text: ""}, nil
		}
		return nil, err
	}
	return &pb.TextReply{Text: string(res)}, nil
}
