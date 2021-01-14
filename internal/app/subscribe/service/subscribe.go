package service

import (
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
	resp, err := s.etcd.Get(context.Background(), "subscribe_", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}

	var result []string

	mb := make(map[string]string)
	for _, item := range resp.Kvs {
		mb[strings.Replace(utils.ByteToString(item.Key), "subscribe_", "", -1)] = utils.ByteToString(item.Value)
	}

	for source, isSubscribe := range mb {
		result = append(result, fmt.Sprintf("%s [Subscribe:%v]", source, isSubscribe))
	}

	return &pb.SubscribeReply{
		Text: result,
	}, nil
}

func (s *Subscribe) Register(ctx context.Context, payload *pb.SubscribeRequest) (*pb.State, error) {
	_, err := s.etcd.Put(context.Background(), "subscribe_"+payload.GetText(), "true")
	if err != nil {
		return nil, err
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
