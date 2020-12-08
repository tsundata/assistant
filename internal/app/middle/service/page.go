package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/model"
	"gorm.io/gorm"
	"time"
)

type Page struct {
	db     *gorm.DB
	webURL string
}

func NewPage(db *gorm.DB, webURL string) *Page {
	return &Page{db: db, webURL: webURL}
}

func (s *Page) Create(ctx context.Context, payload *model.Page, reply *string) error {
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

func (s *Page) Get(ctx context.Context, payload *model.Page, reply *model.Page) error {
	var find model.Page
	s.db.Where("uuid = ?", payload.UUID).Take(&find)
	*reply = find

	return nil
}
