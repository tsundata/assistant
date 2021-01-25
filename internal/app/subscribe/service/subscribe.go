package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.etcd.io/etcd/clientv3"
	"strings"
)

type Subscribe struct {
	etcd *clientv3.Client
}

func NewSubscribe(etcd *clientv3.Client) *Subscribe {
	return &Subscribe{etcd: etcd}
}

func (s *Subscribe) List(ctx context.Context, payload *pb.SubscribeRequest) (*pb.SubscribeReply, error) {
	resp, err := s.etcd.Get(context.Background(), "subscribe_",
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}

	var result []string

	mb := make(map[string]string)
	for _, ev := range resp.Kvs {
		mb[strings.ReplaceAll(utils.ByteToString(ev.Key), "subscribe_", "")] = utils.ByteToString(ev.Value)
	}

	for source, isSubscribe := range mb {
		result = append(result, fmt.Sprintf("%s [Subscribe:%v]", source, isSubscribe))
	}

	return &pb.SubscribeReply{
		Text: result,
	}, nil
}

func (s *Subscribe) Register(ctx context.Context, payload *pb.SubscribeRequest) (*pb.State, error) {
	key := "subscribe_" + payload.GetText()
	resp, err := s.etcd.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	hasKey := false
	for _, ev := range resp.Kvs {
		if key == utils.ByteToString(ev.Key) {
			hasKey = true
		}
	}

	if !hasKey {
		_, err = s.etcd.Put(context.Background(), key, "true")
		if err != nil {
			return nil, err
		}
	}

	return &pb.State{State: true}, nil
}

func (s *Subscribe) Open(ctx context.Context, payload *pb.SubscribeRequest) (*pb.State, error) {
	_, err := s.etcd.Put(context.Background(), "subscribe_"+payload.GetText(), "true")
	if err != nil {
		return nil, err
	}

	return &pb.State{State: true}, nil
}

func (s *Subscribe) Close(ctx context.Context, payload *pb.SubscribeRequest) (*pb.State, error) {
	_, err := s.etcd.Put(context.Background(), "subscribe_"+payload.GetText(), "false")
	if err != nil {
		return nil, err
	}

	return &pb.State{State: true}, nil
}

func (s *Subscribe) Status(ctx context.Context, payload *pb.SubscribeRequest) (*pb.State, error) {
	key := "subscribe_" + payload.GetText()
	resp, err := s.etcd.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}
	for _, ev := range resp.Kvs {
		if utils.ByteToString(ev.Key) == key {
			return &pb.State{
				State: bytes.Equal(ev.Value, []byte("true")),
			}, nil
		}
	}
	return &pb.State{
		State: false,
	}, nil
}
