package service

import (
	"github.com/tsundata/assistant/internal/pkg/models"
	"gorm.io/gorm"
	"log"
)

type Subscribe struct {
	DB *gorm.DB
}

func (s *Subscribe) List(arg string, reply *int) error {
	log.Println("Subscribe.List ............")

	var list []models.Subscribe

	s.DB.AutoMigrate(&models.Subscribe{})
	s.DB.Create(&models.Subscribe{})
	s.DB.Find(&list)
	log.Println(list)

	var errCode int
	errCode = 2232
	*reply = errCode

	return nil
}

func (s *Subscribe) Open(source string, reply *int) error {
	return nil
}

func (s *Subscribe) Close(source string, reply *int) error {
	return nil
}
