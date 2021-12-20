package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/mock"
	"testing"

	"github.com/tsundata/assistant/api/pb"
)

func TestGetGlobalId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockIdRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetOrCreateNode(gomock.Any(), gomock.Any()).Return(pb.Node{
			Id: 1,
		}, nil),
		repo.EXPECT().GetOrCreateNode(gomock.Any(), gomock.Any()).Return(pb.Node{
			Id: 2,
		}, nil),
	)

	s := NewId(repo)

	type args struct {
		ctx context.Context
		in1 *pb.IdRequest
	}
	tests := []struct {
		name    string
		s       *Id
		args    args
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.IdRequest{Ip: "127.0.0.1", Port: 5000}},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.IdRequest{Ip: "127.0.0.1", Port: 5001}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.GetGlobalId(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cron.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
