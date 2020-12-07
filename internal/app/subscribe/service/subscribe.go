package service

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/model"
	"gorm.io/gorm"
	"log"
)

type Subscribe struct {
	db *gorm.DB
}

func NewSubscribe(db *gorm.DB) *Subscribe {
	return &Subscribe{db: db}
}

// TODO
func (s *Subscribe) List(ctx context.Context, payload *model.Event, reply *model.Event) error {
	var list []model.Subscribe

	s.db.AutoMigrate(&model.Subscribe{})
	s.db.Create(&model.Subscribe{})
	s.db.Find(&list)
	log.Println(list)

	return nil
}

// TODO
func (s *Subscribe) Open(ctx context.Context, payload *model.Event, reply *model.Event) error {
	log.Println(payload)

	*reply = model.Event{
		UUID: "out --->",
	}

	return nil
}

// TODO
func (s *Subscribe) View(ctx context.Context, payload *model.Event, reply *model.Event) error {
	return nil
}

// TODO
func (s *Subscribe) Close(ctx context.Context, payload *model.Event, reply *model.Event) error {
	return nil
}
