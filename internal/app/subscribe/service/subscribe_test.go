package service

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"reflect"
	"testing"

	"github.com/tsundata/assistant/api/pb"
)

func TestSubscribe_List(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Subscribe)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.Del(context.Background(), RuleKey).Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test1", "true").Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test2", "true").Result()
	if err != nil {
		t.Fatal(err)
	}

	s := NewSubscribe(rdb)

	type args struct {
		ctx context.Context
		in1 *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Subscribe
		args    args
		want    int
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{}},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.List(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Text) != tt.want {
				t.Errorf("Subscribe.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Register(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Subscribe)
	if err != nil {
		t.Fatal(err)
	}

	s := NewSubscribe(rdb)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Subscribe
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Register(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Open(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Subscribe)
	if err != nil {
		t.Fatal(err)
	}

	s := NewSubscribe(rdb)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Subscribe
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Open(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Close(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Subscribe)
	if err != nil {
		t.Fatal(err)
	}

	s := NewSubscribe(rdb)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Subscribe
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Close(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Status(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Subscribe)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.Del(context.Background(), RuleKey).Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test1", "true").Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test2", "false").Result()
	if err != nil {
		t.Fatal(err)
	}

	s := NewSubscribe(rdb)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Subscribe
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test1"}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test2"}},
			&pb.StateReply{State: false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Status(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}
