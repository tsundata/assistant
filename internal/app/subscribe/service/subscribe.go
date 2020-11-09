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

func (s *Subscribe) List(ctx context.Context, arg string, reply *int) error {
	log.Println("Subscribe.List ............")

	var list []models.Subscribe

	s.db.AutoMigrate(&models.Subscribe{})
	s.db.Create(&models.Subscribe{})
	s.db.Find(&list)
	log.Println(list)

	var errCode int
	errCode = 2232
	*reply = errCode

	return nil
}

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

func (s *Subscribe) Close(source string, reply *int) error {
	return nil
}
