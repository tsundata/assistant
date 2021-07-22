package service

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"reflect"
	"testing"
)

func TestNLP_Pinyin(t *testing.T) {
	s := NewNLP(nil)
	type args struct {
		in0 context.Context
		req *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *NLP
		args    args
		want    *pb.WordsReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{Text: "测试"}},
			&pb.WordsReply{Text: []string{"ce", "shi"}},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.TextRequest{Text: ""}},
			&pb.WordsReply{Text: []string{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pinyin(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NLP.Pinyin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NLP.Pinyin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNLP_Segmentation(t *testing.T) {
	s := NewNLP(nil)
	type args struct {
		in0 context.Context
		req *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *NLP
		args    args
		want    *pb.WordsReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{Text: "现在进行单元测试"}},
			&pb.WordsReply{Text: []string{"现在", "进行", "单元测试"}},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.TextRequest{Text: ""}},
			&pb.WordsReply{Text: []string{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Segmentation(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NLP.Segmentation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NLP.Segmentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNLP_Classifier(t *testing.T) {
	conf, err := config.CreateAppConfig(enum.NLP)
	if err != nil {
		t.Fatal(err)
	}
	s := NewNLP(conf)
	type args struct {
		in0 context.Context
		req *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *NLP
		args    args
		want    *pb.TextReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{Text: "test"}},
			&pb.TextReply{Text: string(enum.StrengthAttr)},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.TextRequest{Text: "demo2"}},
			&pb.TextReply{Text: string(enum.CultureAttr)},
			false,
		},
		{
			"case3",
			s,
			args{context.Background(), &pb.TextRequest{Text: "demo8"}},
			&pb.TextReply{Text: ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Classifier(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NLP.Classifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NLP.Classifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
