package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/subscribe/spider"
	"github.com/tsundata/assistant/internal/pkg/model"
	"gorm.io/gorm"
)

type Subscribe struct {
	db *gorm.DB
}

func NewSubscribe(db *gorm.DB) *Subscribe {
	return &Subscribe{db: db}
}

func (s *Subscribe) List(ctx context.Context, payload *pb.SubscribeRequest) (*pb.SubscribeReply, error) {
	var list []model.Subscribe
	s.db.Find(&list)

	var result []string

	mb := make(map[string]bool)
	for _, item := range list {
		mb[item.Source] = item.IsSubscribe
	}
	for source := range spider.SubscribeRules {
		if b, ok := mb[source]; ok {
			mb[source] = b
		} else {
			mb[source] = true
		}
	}

	for source, isSubscribe := range mb {
		result = append(result, fmt.Sprintf("%s [Subscribe:%v]", source, isSubscribe))
	}

	return &pb.SubscribeReply{
		Text: result,
	}, nil
}

func (s *Subscribe) Open(ctx context.Context, payload *pb.SubscribeRequest) (*pb.State, error) {
	var subscribe model.Subscribe
	s.db.Where(model.Subscribe{Source: payload.GetText()}).FirstOrCreate(&subscribe)

	if !subscribe.IsSubscribe {
		s.db.Model(&subscribe).Where("id = ?", subscribe.ID).Update("is_subscribe", true)
	}

	return &pb.State{State: true}, nil
}

func (s *Subscribe) Close(ctx context.Context, payload *pb.SubscribeRequest) (*pb.State, error) {
	var subscribe model.Subscribe
	s.db.Where(model.Subscribe{Source: payload.GetText()}).FirstOrCreate(&subscribe)

	if subscribe.IsSubscribe {
		s.db.Model(&subscribe).Where("id = ?", subscribe.ID).Update("is_subscribe", false)
	}

	return &pb.State{State: true}, nil
}
