package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"testing"
)

func TestIdRepository_GetOrCreateNode(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Id)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}

	type args struct {
		node *pb.Node
	}
	tests := []struct {
		name    string
		r       IdRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{node: &pb.Node{Ip: "127.0.0.1", Port: 5000}}, false},
		{"case2", sto, args{node: &pb.Node{Ip: "127.0.0.1", Port: 5001}}, false},
		{"case3", sto, args{node: &pb.Node{Ip: "127.0.0.1", Port: 5001}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetOrCreateNode(context.Background(), tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetOrCreateNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
