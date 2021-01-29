package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.etcd.io/etcd/clientv3"
	"net/url"
	"strings"
	"time"
)

type Middle struct {
	db     *sqlx.DB
	etcd   *clientv3.Client
	webURL string
}

func NewMiddle(db *sqlx.DB, etcd *clientv3.Client, webURL string) *Middle {
	return &Middle{db: db, etcd: etcd, webURL: webURL}
}

func (s *Middle) CreatePage(ctx context.Context, payload *pb.PageRequest) (*pb.TextReply, error) {
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}

	page := model.Page{
		UUID:    uuid,
		Title:   payload.GetTitle(),
		Content: payload.GetContent(),
		Time:    time.Now(),
	}

	_, err = s.db.NamedExec("INSERT INTO `pages` (`uuid`, `title`, `content`, `time`) VALUES (:uuid, :title, :content, :time)", page)
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{
		Text: fmt.Sprintf("%s/page/%s", s.webURL, page.UUID),
	}, nil
}

func (s *Middle) GetPage(ctx context.Context, payload *pb.PageRequest) (*pb.PageReply, error) {
	var find model.Page
	err := s.db.Get(&find, "SELECT * FROM `pages` WHERE `uuid` = ?", payload.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.PageReply{
		Uuid:    find.UUID,
		Title:   find.Title,
		Content: find.Content,
	}, nil
}

func (s *Middle) Qr(ctx context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	return &pb.TextReply{
		Text: fmt.Sprintf("%s/qr/%s", s.webURL, url.QueryEscape(payload.GetText())),
	}, nil
}

func (s *Middle) Apps(ctx context.Context, payload *pb.TextRequest) (*pb.AppReply, error) {
	return nil, nil
}

func (s *Middle) StoreAppOAuth(ctx context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	return nil, nil
}

func (s *Middle) GetCredentials(ctx context.Context, payload *pb.TextRequest) (*pb.CredentialReply, error) {
	var items []model.Credential
	err := s.db.Select(&items, "SELECT * FROM `credentials` ORDER BY `id` DESC")
	if err != nil {
		return nil, err
	}

	var kvs []*pb.KV
	for _, item := range items {
		kvs = append(kvs, &pb.KV{
			Key:   item.Name,
			Value: item.Content,
		})
	}

	return &pb.CredentialReply{
		Items: kvs,
	}, nil
}

func (s *Middle) CreateCredential(ctx context.Context, payload *pb.KVsRequest) (*pb.TextReply, error) {
	name := ""
	m := make(map[string]string)
	for _, item := range payload.Kvs {
		if item.Key == "name" {
			name = item.Value
		} else {
			m[item.Key] = item.Value
		}
	}
	if name == "" {
		return nil, errors.New("name key error")
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec("INSERT INTO `credentials` (`name`, `type`, `content`, `time`) VALUES (?, ?, ?, ?)", name, "", utils.ByteToString(data), time.Now())
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{}, nil
}

func (s *Middle) GetSetting(ctx context.Context, payload *pb.TextRequest) (*pb.SettingReply, error) {
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

func (s *Middle) CreateSetting(ctx context.Context, payload *pb.KVRequest) (*pb.TextReply, error) {
	_, err := s.etcd.Put(context.Background(), "setting/"+payload.Key, payload.Value)
	if err != nil {
		return nil, err
	}
	return &pb.TextReply{
		Text: "ok",
	}, nil
}

func (s *Middle) GetMenu(ctx context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	uuid, err := authUUID(s.etcd)
	if err != nil {
		return nil, err
	}
	return &pb.TextReply{
		Text: fmt.Sprintf(`
Memo
%s/memo/%s

Apps
%s/apps/%s

Credentials
%s/credentials/%s

Setting
%s/setting/%s
`, s.webURL, uuid, s.webURL, uuid, s.webURL, uuid, s.webURL, uuid),
	}, nil
}

func (s *Middle) Authorization(ctx context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
	resp, err := s.etcd.Get(context.Background(), "user/auth_uuid")
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return &pb.StateReply{
			State: false,
		}, nil
	}

	return &pb.StateReply{
		State: payload.Text == utils.ByteToString(resp.Kvs[0].Value),
	}, nil
}

func authUUID(etcd *clientv3.Client) (string, error) {
	var uuid string
	resp, err := etcd.Get(context.Background(), "user/auth_uuid")
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		uuid, err = utils.GenerateUUID()
		if err != nil {
			return "", err
		}

		lease, err := etcd.Grant(context.Background(), 3600)
		if err != nil {
			return "", err
		}
		_, err = etcd.Put(context.Background(), "user/auth_uuid", uuid, clientv3.WithLease(lease.ID))
		if err != nil {
			return "", err
		}
	} else {
		uuid = utils.ByteToString(resp.Kvs[0].Value)
	}

	return uuid, nil
}
