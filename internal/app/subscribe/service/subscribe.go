package service

import (
	"context"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/pkg/models"
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
func (s *Subscribe) List(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	log.Println("Subscribe.List ............")

	var list []models.Subscribe

	s.db.AutoMigrate(&models.Subscribe{})
	s.db.Create(&models.Subscribe{})
	s.db.Find(&list)
	log.Println(list)

	return nil
}

// TODO
func (s *Subscribe) Open(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	log.Println(payload)

	*reply = proto.Detail{
		Id:          1,
		Name:        "out =====>",
		Price:       1000,
		CreatedTime: nil,
	}

	return nil
}

// TODO
func (s *Subscribe) View(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	return nil
}

// TODO
func (s *Subscribe) Close(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	return nil
}
