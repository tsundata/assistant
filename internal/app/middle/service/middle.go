package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.etcd.io/bbolt"
	"go.etcd.io/etcd/clientv3"
	"net/url"
	"strings"
	"time"
)

type Middle struct {
	db     *bbolt.DB
	etcd   *clientv3.Client
	webURL string
}

func NewMiddle(db *bbolt.DB, etcd *clientv3.Client, webURL string) *Middle {
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
		Content: payload.GetContent(),
		Time:    time.Now(),
	}

	tx, err := s.db.Begin(true)
	if err != nil {
		return nil, err
	}
	b, err := tx.CreateBucketIfNotExists([]byte("middle"))
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(page)
	if err != nil {
		return nil, err
	}
	err = b.Put([]byte(page.UUID), data)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &pb.Text{
		Text: fmt.Sprintf("%s/page/%s", s.webURL, page.UUID),
	}, nil
}

func (s *Middle) GetPage(ctx context.Context, payload *pb.PageRequest) (*pb.PageReply, error) {
	// TODO cache
	tx, err := s.db.Begin(true)
	if err != nil {
		return nil, err
	}
	b, err := tx.CreateBucketIfNotExists([]byte("middle"))
	if err != nil {
		return nil, err
	}
	v := b.Get([]byte(payload.Uuid))

	var find model.Page
	err = json.Unmarshal(v, &find)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

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

func (s *Middle) Apps(ctx context.Context, payload *pb.Text) (*pb.AppReply, error) {
	return nil, nil
}

func (s *Middle) StoreAppOAuth(ctx context.Context, payload *pb.Text) (*pb.Text, error) {
	return nil, nil
}

func (s *Middle) Credentials(ctx context.Context, payload *pb.Text) (*pb.Text, error) {
	return nil, nil
}

func (s *Middle) GetCredentials(ctx context.Context, payload *pb.Text) (*pb.Text, error) {
	return nil, nil
}

func (s *Middle) GetCredential(ctx context.Context, payload *pb.Text) (*pb.Text, error) {
	return nil, nil
}

func (s *Middle) CreateCredential(ctx context.Context, payload *pb.Text) (*pb.Text, error) {
	return nil, nil
}

func (s *Middle) Setting(ctx context.Context, payload *pb.Text) (*pb.SettingReply, error) {
	resp, err := s.etcd.Get(context.Background(), "setting/",
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}
	var reply pb.SettingReply
	for _, ev := range resp.Kvs {
		reply.Items = append(reply.Items, &pb.KV{
			Key:   strings.ReplaceAll(utils.ByteToString(ev.Key), "setting/", ""),
			Value: utils.ByteToString(ev.Value),
		})
	}
	return &reply, nil
}

func (s *Middle) CreateSetting(ctx context.Context, payload *pb.KV) (*pb.Text, error) {
	_, err := s.etcd.Put(context.Background(), "setting/"+payload.Key, payload.Value)
	if err != nil {
		return nil, err
	}
	return &pb.Text{
		Text: "ok",
	}, nil
}
