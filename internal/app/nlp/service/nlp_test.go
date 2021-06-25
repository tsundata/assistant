package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/tsundata/assistant/api/pb"
)

func TestNLP_Pinyin(t *testing.T) {
	s := NewNLP()
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
	s := NewNLP()
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
