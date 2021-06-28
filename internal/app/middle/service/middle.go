package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/middle/repository"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"net/url"
	"strings"
	"time"
)

type Middle struct {
	conf *config.AppConfig
	rdb  *redis.Client
	repo repository.MiddleRepository
}

func NewMiddle(conf *config.AppConfig, rdb *redis.Client, repo repository.MiddleRepository) *Middle {
	return &Middle{rdb: rdb, repo: repo, conf: conf}
}

func (s *Middle) GetMenu(_ context.Context, _ *pb.TextRequest) (*pb.TextReply, error) {
	uuid, err := authUUID(s.rdb)
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

Action
%s/action/%s
`, s.conf.Web.Url, uuid, s.conf.Web.Url, uuid, s.conf.Web.Url, uuid, s.conf.Web.Url, uuid, s.conf.Web.Url, uuid),
	}, nil
}

func (s *Middle) GetQrUrl(_ context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	return &pb.TextReply{
		Text: fmt.Sprintf("%s/qr/%s", s.conf.Web.Url, url.QueryEscape(payload.GetText())),
	}, nil
}

func (s *Middle) CreatePage(_ context.Context, payload *pb.PageRequest) (*pb.TextReply, error) {
	uuid, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}

	page := model.Page{
		UUID:    uuid,
		Type:    payload.GetType(),
		Title:   payload.GetTitle(),
		Content: payload.GetContent(),
		Time:    time.Now(),
	}

	_, err = s.repo.CreatePage(page)
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{
		Text: fmt.Sprintf("%s/page/%s", s.conf.Web.Url, page.UUID),
	}, nil
}

func (s *Middle) GetPage(_ context.Context, payload *pb.PageRequest) (*pb.PageReply, error) {
	find, err := s.repo.GetPageByUUID(payload.GetUuid())
	if err != nil {
		return nil, err
	}

	return &pb.PageReply{
		Uuid:    find.UUID,
		Type:    find.Type,
		Title:   find.Title,
		Content: find.Content,
	}, nil
}

func (s *Middle) GetApps(_ context.Context, _ *pb.TextRequest) (*pb.AppsReply, error) {
	apps, err := s.repo.ListApps()
	if err != nil {
		return nil, err
	}

	providerApps := map[string]bool{}
	for _, provider := range vendors.OAuthProviderApps {
		providerApps[provider] = true
	}

	haveApps := make(map[string]bool)
	var res []*pb.App
	for _, app := range apps {
		haveApps[app.Type] = true
		res = append(res, &pb.App{
			Title:        fmt.Sprintf("%s (%s)", app.Name, app.Type),
			IsAuthorized: app.Token != "",
			Type:         app.Type,
			Name:         app.Name,
			Token:        app.Token,
			Extra:        app.Extra,
			Time:         app.Time.Format("2006-01-02 15:04:05"),
		})
	}

	for k := range providerApps {
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
	find, err := s.repo.GetAvailableAppByType(payload.GetText())
	if err != nil {
		return nil, err
	}

	var kvs []*pb.KV

	if find.ID > 0 {
		var extra map[string]string
		err = json.Unmarshal(util.StringToByte(find.Extra), &extra)
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

	app, err := s.repo.GetAppByType(payload.GetType())
	if err != nil {
		return nil, err
	}

	if app.ID > 0 {
		err = s.repo.UpdateAppByID(int64(app.ID), payload.GetToken(), payload.GetExtra())
		if err != nil {
			return nil, err
		}
	} else {
		_, err = s.repo.CreateApp(model.App{
			Name:  payload.GetName(),
			Type:  payload.GetType(),
			Token: payload.GetToken(),
			Extra: payload.GetExtra(),
		})
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
	var err error
	if payload.GetName() != "" {
		find, err = s.repo.GetCredentialByName(payload.GetName())
	} else if payload.GetType() != "" {
		find, err = s.repo.GetCredentialByType(payload.GetType())
	}
	if err != nil {
		return nil, err
	}

	var kvs []*pb.KV
	if find.ID > 0 {
		var data map[string]string
		err := json.Unmarshal(util.StringToByte(find.Content), &data)
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
	items, err := s.repo.ListCredentials()
	if err != nil {
		return nil, err
	}

	var credentials []*pb.Credential
	for _, item := range items {
		credentials = append(credentials, &pb.Credential{
			Name:    item.Name,
			Type:    item.Type,
			Content: item.Content,
			Time:    item.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.CredentialsReply{
		Credentials: credentials,
	}, nil
}

func (s *Middle) GetMaskingCredentials(_ context.Context, _ *pb.TextRequest) (*pb.MaskingReply, error) {
	items, err := s.repo.ListCredentials()
	if err != nil {
		return nil, err
	}

	var kvs []*pb.KV
	for _, item := range items {
		// Data masking
		var data map[string]string
		err := json.Unmarshal(util.StringToByte(item.Content), &data)
		if err != nil {
			return nil, err
		}
		for k, v := range data {
			if k != "name" && k != "type" {
				data[k] = util.DataMasking(v)
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
			Value: util.ByteToString(content),
		})
	}

	return &pb.MaskingReply{
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

	_, err = s.repo.CreateCredential(model.Credential{Name: name, Type: category, Content: util.ByteToString(data), Time: time.Now()})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetSettings(_ context.Context, _ *pb.TextRequest) (*pb.SettingsReply, error) {
	result, err := s.conf.GetSettings()
	if err != nil {
		return nil, err
	}

	var reply pb.SettingsReply
	for k, v := range result {
		reply.Items = append(reply.Items, &pb.KV{
			Key:   k,
			Value: v,
		})
	}
	return &reply, nil
}

func (s *Middle) GetSetting(_ context.Context, payload *pb.TextRequest) (*pb.SettingReply, error) {
	result, err := s.conf.GetSetting(payload.GetText())
	if err != nil {
		return nil, err
	}

	return &pb.SettingReply{
		Key:   payload.GetText(),
		Value: result,
	}, nil
}

func (s *Middle) CreateSetting(_ context.Context, payload *pb.KVRequest) (*pb.StateReply, error) {
	err := s.conf.SetSetting(payload.GetKey(), payload.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetStats(ctx context.Context, _ *pb.TextRequest) (*pb.TextReply, error) {
	var result []string

	// count
	keys, _, err := s.rdb.Scan(ctx, 0, "stats:count:*", 1000).Result()
	if err != nil {
		return nil, err
	}
	if len(keys) <= 0 {
		return &pb.TextReply{Text: "not stats"}, nil
	}
	values, err := s.rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		result = append(result, fmt.Sprintf("%s: %s", strings.ReplaceAll(keys[i], "stats:count:", ""), values[i]))
	}

	// month
	keys, _, err = s.rdb.Scan(ctx, 0, "stats:month:*", 1000).Result()
	if err != nil {
		return nil, err
	}
	values, err = s.rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		binString := ""
		for _, c := range values[i].(string) {
			binString = fmt.Sprintf("%s%.8b", binString, c)
		}
		result = append(result, fmt.Sprintf("%s: %s", strings.ReplaceAll(keys[i], "stats:month:", ""), binString))
	}

	return &pb.TextReply{Text: strings.Join(result, "\n")}, nil
}

const AuthKey = "user:auth:token"

func authUUID(rdb *redis.Client) (string, error) {
	var uuid string
	uuid, err := rdb.Get(context.Background(), AuthKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}
	if errors.Is(err, redis.Nil) {
		uuid, err = util.GenerateUUID()
		if err != nil {
			return "", err
		}

		status := rdb.Set(context.Background(), AuthKey, uuid, 60*time.Minute)
		if status.Err() != nil {
			return "", err
		}
	}

	return uuid, nil
}
