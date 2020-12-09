package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/model"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type Web struct {
	db     *gorm.DB
	webURL string
}

func NewWeb(db *gorm.DB, webURL string) *Web {
	return &Web{db: db, webURL: webURL}
}

func (s *Web) CreatePage(ctx context.Context, payload *model.Page, reply *string) error {
	var err error
	payload.UUID, err = model.GenerateMessageUUID()
	if err != nil {
		return err
	}
	payload.Time = time.Now()
	s.db.Create(&payload)

	*reply = fmt.Sprintf("%s/page/%s", s.webURL, payload.UUID)

	return nil
}

func (s *Web) GetPage(ctx context.Context, payload *model.Page, reply *model.Page) error {
	var find model.Page
	s.db.Where("uuid = ?", payload.UUID).Take(&find)
	*reply = find

	return nil
}

func (s *Web) Qr(ctx context.Context, payload *string, reply *string) error {
	*reply = fmt.Sprintf("%s/qr/%s", s.webURL, url.QueryEscape(*payload))

	return nil
}
