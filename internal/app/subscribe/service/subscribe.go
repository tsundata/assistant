package service

import (
	"context"
	"fmt"
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

func (s *Subscribe) List(ctx context.Context, payload *model.Subscribe, reply *[]string) error {
	var list []model.Subscribe
	s.db.Find(&list)

	var result []string

	mb := make(map[string]bool)
	for _, item := range list {
		mb[item.Source] = item.IsSubscribe
	}
	for source, _ := range spider.SubscribeRules {
		if b, ok := mb[source]; ok {
			mb[source] = b
		} else {
			mb[source] = true
		}
	}

	for source, isSubscribe := range mb {
		result = append(result, fmt.Sprintf("%s [Subscribe:%v]", source, isSubscribe))
	}

	*reply = result

	return nil
}

func (s *Subscribe) Open(ctx context.Context, payload *string, reply *bool) error {
	var subscribe model.Subscribe
	s.db.Where(model.Subscribe{Source: *payload}).FirstOrCreate(&subscribe)

	if subscribe.IsSubscribe != true {
		s.db.Model(&subscribe).Where("id = ?", subscribe.ID).Update("is_subscribe", true)
	}

	*reply = true

	return nil
}

func (s *Subscribe) Close(ctx context.Context, payload *string, reply *bool) error {
	var subscribe model.Subscribe
	s.db.Where(model.Subscribe{Source: *payload}).FirstOrCreate(&subscribe)

	if subscribe.IsSubscribe != false {
		s.db.Model(&subscribe).Where("id = ?", subscribe.ID).Update("is_subscribe", false)
	}

	*reply = true

	return nil
}
