package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type Middle struct {
	db     *gorm.DB
	webURL string
}

func NewMiddle(db *gorm.DB, webURL string) *Middle {
	return &Middle{db: db, webURL: webURL}
}

func (s *Middle) CreatePage(ctx context.Context, payload *pb.PageRequest) (*pb.Text, error) {
	uuid, err := model.GenerateMessageUUID()
	if err != nil {
		return nil, err
	}

	page := model.Page{
		UUID:    uuid,
		Title:   payload.GetTitle(),
		Content: payload.GetContent(), // TODO
		Time:    time.Now(),
	}
	s.db.Create(&page)

	return &pb.Text{
		Text: fmt.Sprintf("%s/page/%s", s.webURL, page.UUID),
	}, nil
}

func (s *Middle) GetPage(ctx context.Context, payload *pb.PageRequest) (*pb.PageReply, error) {
	// TODO cache
	var find model.Page
	s.db.Select("uuid", "title", "content").Where("uuid = ?", payload.GetUuid()).Take(&find)

	return &pb.PageReply{
		Uuid:    find.UUID,
		Title:   find.Title,
		Content: find.Content,
	}, nil
}

func (s *Middle) Qr(ctx context.Context, payload *pb.Text) (*pb.Text, error) {
	return &pb.Text{
		Text: fmt.Sprintf("%s/qr/%s", s.webURL, url.QueryEscape(payload.GetText())),
	}, nil
}
