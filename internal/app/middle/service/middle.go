package service

import (
	"context"
	"database/sql"
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

func (s *Middle) CreatePage(_ context.Context, payload *pb.PageRequest) (*pb.TextReply, error) {
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

func (s *Middle) GetPage(_ context.Context, payload *pb.PageRequest) (*pb.PageReply, error) {
	var find model.Page
	err := s.db.Get(&find, "SELECT * FROM `pages` WHERE `uuid` = ?", payload.GetUuid())
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &pb.PageReply{
		Uuid:    find.UUID,
		Title:   find.Title,
		Content: find.Content,
	}, nil
}

func (s *Middle) Qr(_ context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	return &pb.TextReply{
		Text: fmt.Sprintf("%s/qr/%s", s.webURL, url.QueryEscape(payload.GetText())),
	}, nil
}

func (s *Middle) Apps(_ context.Context, _ *pb.TextRequest) (*pb.AppsReply, error) {
	var apps []model.App
	err := s.db.Select(&apps, "SELECT * FROM `apps` ORDER BY `time` DESC")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	systemApps := map[string]bool{
		"pocket": true,
		"github": true,
	}

	haveApps := make(map[string]bool)
	var res []*pb.App
	for _, app := range apps {
		haveApps[app.Type] = true
		res = append(res, &pb.App{
			Title:        fmt.Sprintf("%s (%s)", app.Name, app.Type),
			IsAuthorized: app.Token != "",
			Type:         app.Type,
		})
	}

	for k := range systemApps {
		if _, ok := haveApps[k]; !ok {
			res = append(res, &pb.App{
				Title:        fmt.Sprintf("%s (%s)", k, k),
				IsAuthorized: false,
				Type:         k,
			})
		}
	}

	return &pb.AppsReply{
		Apps: res,
	}, nil
}

func (s *Middle) GetAvailableApp(_ context.Context, payload *pb.TextRequest) (*pb.AppReply, error) {
	var find model.App
	err := s.db.Get(&find, "SELECT * FROM apps WHERE `type` = ? AND `token` <> '' LIMIT 1", payload.GetText())
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var kvs []*pb.KV

	if find.ID > 0 {
		var extra map[string]string
		err = json.Unmarshal(utils.StringToByte(find.Extra), &extra)
		if err != nil {
			return nil, err
		}
		for k, v := range extra {
			kvs = append(kvs, &pb.KV{
				Key:   k,
				Value: v,
			})
		}
	}

	return &pb.AppReply{
		Name:  find.Name,
		Type:  find.Type,
		Token: find.Token,
		Extra: kvs,
	}, nil
}

func (s *Middle) StoreAppOAuth(_ context.Context, payload *pb.AppRequest) (*pb.StateReply, error) {
	if payload.GetToken() == "" {
		return &pb.StateReply{
			State: false,
		}, nil
	}

	var app model.App
	err := s.db.Get(&app, "SELECT * FROM apps WHERE type = ? ORDER BY id DESC LIMIT 1", payload.GetType())
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if app.ID > 0 {
		_, err = s.db.Exec("UPDATE apps SET `token` = ?, `extra` = ?, `time` = ? WHERE id = ?", payload.GetToken(), payload.GetExtra(), time.Now(), app.ID)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = s.db.Exec("INSERT INTO `apps` (`name`, `type`, `token`, `extra`) VALUES (?, ?, ?, ?)",
			payload.GetName(), payload.GetType(), payload.GetToken(), payload.GetExtra())
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{
		State: true,
	}, nil
}

func (s *Middle) GetCredential(_ context.Context, payload *pb.CredentialRequest) (*pb.CredentialReply, error) {
	var find model.Credential
	if payload.GetName() != "" {
		err := s.db.Get(&find, "SELECT * FROM credentials WHERE name = ? LIMIT 1", payload.GetName())
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
	} else if payload.GetType() != "" {
		err := s.db.Get(&find, "SELECT * FROM credentials WHERE type = ? LIMIT 1", payload.GetType())
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
	}

	var kvs []*pb.KV
	if find.ID > 0 {
		var data map[string]string
		err := json.Unmarshal(utils.StringToByte(find.Content), &data)
		if err != nil {
			return nil, err
		}
		for k, v := range data {
			kvs = append(kvs, &pb.KV{
				Key:   k,
				Value: v,
			})
		}
	}

	return &pb.CredentialReply{
		Name:    find.Name,
		Type:    find.Type,
		Content: kvs,
	}, nil
}

func (s *Middle) GetCredentials(_ context.Context, _ *pb.TextRequest) (*pb.CredentialsReply, error) {
	var items []model.Credential
	err := s.db.Select(&items, "SELECT * FROM `credentials` ORDER BY `id` DESC")
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var kvs []*pb.KV
	for _, item := range items {
		// Data masking
		var data map[string]string
		err := json.Unmarshal(utils.StringToByte(item.Content), &data)
		if err != nil {
			return nil, err
		}
		for k, v := range data {
			if k != "name" && k != "type" {
				data[k] = utils.DataMasking(v)
			} else {
				data[k] = v
			}
		}
		content, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		kvs = append(kvs, &pb.KV{
			Key:   item.Name,
			Value: utils.ByteToString(content),
		})
	}

	return &pb.CredentialsReply{
		Items: kvs,
	}, nil
}

func (s *Middle) CreateCredential(_ context.Context, payload *pb.KVsRequest) (*pb.StateReply, error) {
	name := ""
	category := ""
	m := make(map[string]string)
	for _, item := range payload.GetKvs() {
		if item.Key == "name" {
			name = item.Value
		} else if item.Key == "type" {
			category = item.Value
		}
		m[item.Key] = item.Value
	}
	if name == "" {
		return nil, errors.New("name key error")
	}

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec("INSERT INTO `credentials` (`name`, `type`, `content`, `time`) VALUES (?, ?, ?, ?)",
		name, category, utils.ByteToString(data), time.Now())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetSetting(_ context.Context, _ *pb.TextRequest) (*pb.SettingReply, error) {
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

func (s *Middle) CreateSetting(_ context.Context, payload *pb.KVRequest) (*pb.StateReply, error) {
	_, err := s.etcd.Put(context.Background(), "setting/"+payload.GetKey(), payload.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetScripts(_ context.Context, _ *pb.TextRequest) (*pb.ScriptsReply, error) {
	var items []model.Message
	err := s.db.Select(&items, "SELECT * FROM `messages` WHERE `type` = ? ORDER BY `id` DESC", model.MessageTypeScript)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var kvs []*pb.Script
	for _, item := range items {
		kvs = append(kvs, &pb.Script{
			Id:   int64(item.ID),
			Text: item.Text,
		})
	}

	return &pb.ScriptsReply{
		Items: kvs,
	}, nil
}

func (s *Middle) CreateScript(_ context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
	if payload.GetText() == "" {
		return &pb.StateReply{State: false}, nil
	}
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}
	_, err = s.db.Exec("INSERT INTO `messages` (`uuid`, `type`, `text`, `time`) VALUES (?, ?, ?, ?)",
		uuid, model.MessageTypeScript, payload.GetText(), time.Now())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetMenu(_ context.Context, _ *pb.TextRequest) (*pb.TextReply, error) {
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

Scripts
%s/scripts/%s
`, s.webURL, uuid, s.webURL, uuid, s.webURL, uuid, s.webURL, uuid, s.webURL, uuid),
	}, nil
}

func (s *Middle) Authorization(_ context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
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
		State: payload.GetText() == utils.ByteToString(resp.Kvs[0].Value),
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
