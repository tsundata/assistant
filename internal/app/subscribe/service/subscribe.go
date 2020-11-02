package service

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/pkg/models"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"gorm.io/gorm"
	"log"
	"time"
)

type Subscribe struct {
	db *gorm.DB
}

func NewSubscribe(db *gorm.DB) *Subscribe {
	return &Subscribe{db: db}
}

func (s *Subscribe) List(arg string, reply *int) error {
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

func (s *Subscribe) Open(payload []byte, reply *[]byte) error {
	var in proto.Detail
	err := utils.ProtoUnmarshal(payload, &in)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(in)

	ct, _ := ptypes.TimestampProto(time.Now())
	detail := &proto.Detail{
		Id:          1,
		Name:        "out =====>",
		Price:       1000,
		CreatedTime: ct,
	}
	*reply, _ = utils.ProtoMarshal(detail)

	return nil
}

func (s *Subscribe) Close(source string, reply *int) error {
	return nil
}
