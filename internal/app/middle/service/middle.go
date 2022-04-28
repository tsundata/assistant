package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ego/gse"
	"github.com/go-redis/redis/v8"
	"github.com/mozillazg/go-pinyin"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/middle/classifier"
	"github.com/tsundata/assistant/internal/app/middle/repository"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"gorm.io/gorm"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Middle struct {
	conf    *config.AppConfig
	rdb     *redis.Client
	repo    repository.MiddleRepository
	storage pb.StorageSvcClient
}

func NewMiddle(conf *config.AppConfig, rdb *redis.Client, repo repository.MiddleRepository, storage pb.StorageSvcClient) *Middle {
	return &Middle{rdb: rdb, repo: repo, conf: conf, storage: storage}
}

func (s *Middle) GetQrUrl(_ context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	return &pb.TextReply{
		Text: fmt.Sprintf("%s/qr/%s", s.conf.Web.Url, url.QueryEscape(payload.GetText())),
	}, nil
}

func (s *Middle) GetChartUrl(_ context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	uuid := payload.GetText()

	return &pb.TextReply{
		Text: fmt.Sprintf("%s/chart/%s", s.conf.Web.Url, uuid),
	}, nil
}

func (s *Middle) CreatePage(ctx context.Context, payload *pb.PageRequest) (*pb.TextReply, error) {
	uuid := util.UUID()

	page := pb.Page{
		Uuid:    uuid,
		Type:    payload.Page.GetType(),
		Title:   payload.Page.GetTitle(),
		Content: payload.Page.GetContent(),
	}

	_, err := s.repo.CreatePage(ctx, &page)
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{
		Text: fmt.Sprintf("%s/page/%s", s.conf.Web.Url, page.Uuid),
	}, nil
}

func (s *Middle) GetPage(ctx context.Context, payload *pb.PageRequest) (*pb.PageReply, error) {
	find, err := s.repo.GetPageByUUID(ctx, payload.Page.GetUuid())
	if err != nil {
		return nil, err
	}

	return &pb.PageReply{
		Page: &pb.Page{
			Uuid:    find.Uuid,
			Type:    find.Type,
			Title:   find.Title,
			Content: find.Content,
		},
	}, nil
}

func (s *Middle) GetApps(ctx context.Context, _ *pb.TextRequest) (*pb.AppsReply, error) {
	id, _ := md.FromIncoming(ctx)

	apps, err := s.repo.ListApps(ctx, id)
	if err != nil {
		return nil, err
	}

	providerApps := map[string]bool{}
	for _, provider := range vendors.OAuthProviderApps {
		providerApps[provider] = true
	}

	haveApps := make(map[string]bool)
	var res []*pb.App
	for _, item := range apps {
		haveApps[item.Type] = true
		res = append(res, &pb.App{
			Title:        fmt.Sprintf("%s (%s)", item.Name, item.Type),
			IsAuthorized: item.Token != "",
			Type:         item.Type,
			Name:         item.Name,
			Token:        item.Token,
			Extra:        item.Extra,
			CreatedAt:    item.CreatedAt,
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

func (s *Middle) GetAvailableApp(ctx context.Context, payload *pb.TextRequest) (*pb.AppReply, error) {
	id, _ := md.FromIncoming(ctx)
	find, err := s.repo.GetAvailableAppByType(ctx, id, payload.GetText())
	if err != nil {
		return nil, err
	}

	var kvs []*pb.KV
	if find.Id > 0 {
		var extra map[string]string
		if find.Extra != "" {
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
	}

	return &pb.AppReply{
		Name:  find.Name,
		Type:  find.Type,
		Token: find.Token,
		Extra: kvs,
	}, nil
}

func (s *Middle) StoreAppOAuth(ctx context.Context, payload *pb.AppRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	if payload.App.GetToken() == "" {
		return &pb.StateReply{
			State: false,
		}, nil
	}

	item, err := s.repo.GetAppByType(ctx, id, payload.App.GetType())
	if err != nil {
		return nil, err
	}

	if item.Id > 0 {
		err = s.repo.UpdateAppByID(ctx, item.Id, payload.App.GetToken(), payload.App.GetExtra())
		if err != nil {
			return nil, err
		}
	} else {
		_, err = s.repo.CreateApp(ctx, &pb.App{
			UserId: id,
			Name:   payload.App.GetName(),
			Type:   payload.App.GetType(),
			Token:  payload.App.GetToken(),
			Extra:  payload.App.GetExtra(),
		})
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{
		State: true,
	}, nil
}

func (s *Middle) GetCredential(ctx context.Context, payload *pb.CredentialRequest) (*pb.CredentialReply, error) {
	id, _ := md.FromIncoming(ctx)
	var err error
	var find pb.Credential
	if payload.GetName() != "" {
		find, err = s.repo.GetCredentialByName(ctx, id, payload.GetName())
	} else if payload.GetType() != "" {
		find, err = s.repo.GetCredentialByType(ctx, id, payload.GetType())
	}
	if err != nil {
		return nil, err
	}

	var kvs []*pb.KV
	if find.Id > 0 {
		var data map[string]string
		if find.Content != "" {
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
	}

	return &pb.CredentialReply{
		Name:    find.Name,
		Type:    find.Type,
		Content: kvs,
	}, nil
}

func (s *Middle) GetCredentials(ctx context.Context, _ *pb.TextRequest) (*pb.CredentialsReply, error) {
	id, _ := md.FromIncoming(ctx)
	items, err := s.repo.ListCredentials(ctx, id)
	if err != nil {
		return nil, err
	}

	var credentials []*pb.Credential
	for _, item := range items {
		credentials = append(credentials, &pb.Credential{
			Name:      item.Name,
			Type:      item.Type,
			Content:   item.Content,
			CreatedAt: item.CreatedAt,
		})
	}

	return &pb.CredentialsReply{
		Credentials: credentials,
	}, nil
}

func (s *Middle) GetMaskingCredentials(ctx context.Context, _ *pb.TextRequest) (*pb.MaskingReply, error) {
	id, _ := md.FromIncoming(ctx)
	items, err := s.repo.ListCredentials(ctx, id)
	if err != nil {
		return nil, err
	}

	var kvs []*pb.KV
	for _, item := range items {
		// Data masking
		var data map[string]string
		if item.Content == "" {
			continue
		}
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

func (s *Middle) CreateCredential(ctx context.Context, payload *pb.KVsRequest) (*pb.StateReply, error) {
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

	_, err = s.repo.CreateCredential(ctx, &pb.Credential{Name: name, Type: category, Content: util.ByteToString(data)})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetSettings(ctx context.Context, _ *pb.TextRequest) (*pb.SettingsReply, error) {
	result, err := s.conf.GetSettings(ctx)
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

func (s *Middle) GetSetting(ctx context.Context, payload *pb.TextRequest) (*pb.SettingReply, error) {
	result, err := s.conf.GetSetting(ctx, payload.GetText())
	if err != nil {
		return nil, err
	}

	return &pb.SettingReply{
		Key:   payload.GetText(),
		Value: result,
	}, nil
}

func (s *Middle) CreateSetting(ctx context.Context, payload *pb.KVRequest) (*pb.StateReply, error) {
	err := s.conf.SetSetting(ctx, payload.GetKey(), payload.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

type countKV struct {
	key   string
	value interface{}
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
	var kvs []countKV
	for i := 0; i < len(keys); i++ {
		kvs = append(kvs, countKV{
			key:   keys[i],
			value: values[i],
		})
	}
	sort.Slice(kvs, func(i, j int) bool {
		l := 0
		if v, ok := kvs[i].value.(string); ok {
			l, _ = strconv.Atoi(v)
		}
		r := 0
		if v, ok := kvs[j].value.(string); ok {
			r, _ = strconv.Atoi(v)
		}
		return l > r
	})

	for _, i := range kvs {
		result = append(result, fmt.Sprintf("%s: %s", strings.ReplaceAll(i.key, "stats:count:", ""), i.value))
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

func (s *Middle) ListSubscribe(ctx context.Context, _ *pb.SubscribeRequest) (*pb.SubscribeReply, error) {
	subscribes, err := s.repo.ListSubscribe(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.SubscribeReply{
		Subscribe: subscribes,
	}, nil
}

func (s *Middle) RegisterSubscribe(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	err := s.repo.CreateSubscribe(ctx, pb.Subscribe{
		Name:   payload.Text,
		Status: enum.SubscribeEnableStatus,
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) OpenSubscribe(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateSubscribeStatus(ctx, payload.Text, enum.SubscribeEnableStatus)
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) CloseSubscribe(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateSubscribeStatus(ctx, payload.Text, enum.SubscribeDisableStatus)
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetSubscribeStatus(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	subscribes, err := s.repo.ListSubscribe(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range subscribes {
		if item.Name == payload.GetText() {
			return &pb.StateReply{
				State: item.Status == enum.SubscribeEnableStatus,
			}, nil
		}
	}
	return &pb.StateReply{
		State: false,
	}, nil
}

func (s *Middle) GetUserSubscribe(ctx context.Context, _ *pb.TextRequest) (*pb.GetUserSubscribeReply, error) {
	id, _ := md.FromIncoming(ctx)
	kv := make(map[string]string)

	systemSubs, err := s.repo.ListSubscribe(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range systemSubs {
		kv[item.Name] = strconv.Itoa(int(item.Status))
	}

	userSubs, err := s.repo.ListUserSubscribe(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, item := range userSubs {
		kv[item.Key] = item.Value
	}

	var result []*pb.KV
	for k, v := range kv {
		result = append(result, &pb.KV{
			Key:   k,
			Value: v,
		})
	}

	return &pb.GetUserSubscribeReply{Subscribe: result}, nil
}

func (s *Middle) SwitchUserSubscribe(ctx context.Context, payload *pb.SwitchUserSubscribeRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	for _, item := range payload.Subscribe {
		subs, err := s.repo.GetSubscribe(ctx, item.Key)
		if err != nil {
			return nil, err
		}

		find, err := s.repo.GetUserSubscribe(ctx, id, subs.Id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		var status int64
		if item.Value == "1" {
			status = enum.SubscribeEnableStatus
		} else {
			status = enum.SubscribeDisableStatus
		}
		if find.Id > 0 {
			err = s.repo.UpdateUserSubscribeStatus(ctx, id, subs.Id, status)
			if err != nil {
				return nil, err
			}
		} else {
			err = s.repo.CreateUserSubscribe(ctx, pb.UserSubscribe{
				UserId:      id,
				SubscribeId: subs.Id,
				Status:      status,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetUserSubscribeStatus(ctx context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	state := false

	subscribes, err := s.repo.ListSubscribe(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range subscribes {
		if item.Name == payload.GetText() {
			state = item.Status == enum.SubscribeEnableStatus
			break
		}
	}

	userSubs, err := s.repo.ListUserSubscribe(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, item := range userSubs {
		if item.Key == payload.GetText() {
			state = item.Value == "1"
			break
		}
	}

	return &pb.StateReply{State: state}, nil
}

func (s *Middle) ListCron(ctx context.Context, _ *pb.CronRequest) (*pb.CronReply, error) {
	res, err := s.rdb.HGetAll(ctx, enum.CronKey).Result()
	if err != nil {
		return nil, err
	}

	var result []*pb.Cron
	for source, isCron := range res {
		result = append(result, &pb.Cron{
			Name:  source,
			State: util.StringToBool(isCron),
		})
	}

	return &pb.CronReply{
		Cron: result,
	}, nil
}

func (s *Middle) RegisterCron(ctx context.Context, payload *pb.CronRequest) (*pb.StateReply, error) {
	resp, err := s.rdb.HMGet(ctx, enum.CronKey, payload.GetText()).Result()
	if err != nil {
		return nil, err
	}

	exist := true
	if len(resp) == 0 || (len(resp) == 1 && resp[0] == nil) {
		exist = false
	}

	if !exist {
		_, err = s.rdb.HMSet(ctx, enum.CronKey, payload.GetText(), "true").Result()
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) StartCron(ctx context.Context, payload *pb.CronRequest) (*pb.StateReply, error) {
	_, err := s.rdb.HMSet(ctx, enum.CronKey, payload.GetText(), "true").Result()
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) StopCron(ctx context.Context, payload *pb.CronRequest) (*pb.StateReply, error) {
	_, err := s.rdb.HMSet(ctx, enum.CronKey, payload.GetText(), "false").Result()
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Middle) GetCronStatus(ctx context.Context, payload *pb.CronRequest) (*pb.StateReply, error) {
	resp, err := s.rdb.HGetAll(ctx, enum.CronKey).Result()
	if err != nil {
		return nil, err
	}
	for k, v := range resp {
		if k == payload.GetText() {
			return &pb.StateReply{
				State: v == "true",
			}, nil
		}
	}
	return &pb.StateReply{
		State: false,
	}, nil
}

func (s *Middle) GetOrCreateTag(ctx context.Context, payload *pb.TagRequest) (*pb.TagReply, error) {
	tag, err := s.repo.GetOrCreateTag(ctx, &pb.Tag{
		Name: payload.Tag.GetName(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.TagReply{Tag: &pb.Tag{
		Id:        tag.Id,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
	}}, nil
}

func (s *Middle) GetTags(ctx context.Context, _ *pb.TagRequest) (*pb.TagsReply, error) {
	id, _ := md.FromIncoming(ctx)
	items, err := s.repo.ListTags(ctx, id)
	if err != nil {
		return nil, err
	}

	var tags []*pb.Tag
	for _, item := range items {
		tags = append(tags, &pb.Tag{
			Id:        item.Id,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
		})
	}

	return &pb.TagsReply{
		Tags: tags,
	}, nil
}

func (s *Middle) SaveModelTag(ctx context.Context, payload *pb.ModelTagRequest) (*pb.ModelTagReply, error) {
	id, _ := md.FromIncoming(ctx)
	tag, err := s.repo.GetOrCreateTag(ctx, &pb.Tag{
		UserId: id,
		Name:   payload.Tag,
	})
	if err != nil {
		return nil, err
	}
	payload.Model.UserId = id
	payload.Model.TagId = tag.Id
	_, err = s.repo.GetOrCreateModelTag(ctx, payload.Model)
	if err != nil {
		return nil, err
	}
	return &pb.ModelTagReply{Model: payload.Model}, nil
}

func (s *Middle) GetTagsByModelId(ctx context.Context, payload *pb.ModelIdRequest) (*pb.GetTagsReply, error) {
	id, _ := md.FromIncoming(ctx)
	tags, err := s.repo.ListModelTagsByModelId(ctx, id, payload.ModelId)
	if err != nil {
		return nil, err
	}
	return &pb.GetTagsReply{Tags: tags}, nil
}

func (s *Middle) GetModelTags(ctx context.Context, payload *pb.ModelTagRequest) (*pb.GetTagsReply, error) {
	id, _ := md.FromIncoming(ctx)
	var err error
	var tags []*pb.ModelTag
	if payload.Model != nil {
		tags, err = s.repo.ListModelTagsByModel(ctx, id, *payload.Model)
	} else {
		tags, err = s.repo.ListModelTagsByTag(ctx, id, payload.Tag)
	}
	if err != nil {
		return nil, err
	}

	return &pb.GetTagsReply{Tags: tags}, nil
}

func (s *Middle) GetChartData(ctx context.Context, payload *pb.ChartDataRequest) (*pb.ChartDataReply, error) {
	resp, err := s.rdb.Get(ctx, fmt.Sprintf("middle:chart:%s", payload.ChartData.GetUuid())).Result()
	if err != nil {
		return nil, err
	}

	var data pb.ChartData
	err = json.Unmarshal(util.StringToByte(resp), &data)
	if err != nil {
		return nil, err
	}
	return &pb.ChartDataReply{ChartData: &pb.ChartData{
		Uuid:     data.Uuid,
		Title:    data.Title,
		SubTitle: data.SubTitle,
		XAxis:    data.XAxis,
		Series:   data.Series,
	}}, nil
}

func (s *Middle) SetChartData(ctx context.Context, payload *pb.ChartDataRequest) (*pb.ChartDataReply, error) {
	uuid := util.UUID()
	data := pb.ChartData{
		Uuid:     uuid,
		Title:    payload.ChartData.GetTitle(),
		SubTitle: payload.ChartData.GetSubTitle(),
		XAxis:    payload.ChartData.GetXAxis(),
		Series:   payload.ChartData.GetSeries(),
	}
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	_, err = s.rdb.Set(ctx, fmt.Sprintf("middle:chart:%s", uuid), util.ByteToString(d), 90*24*time.Hour).Result()
	if err != nil {
		return nil, err
	}
	return &pb.ChartDataReply{ChartData: &pb.ChartData{Uuid: uuid}}, nil
}

func (s *Middle) Pinyin(_ context.Context, payload *pb.TextRequest) (*pb.WordsReply, error) {
	if payload.GetText() == "" {
		return &pb.WordsReply{Text: []string{}}, nil
	}
	a := pinyin.NewArgs()
	py := pinyin.Pinyin(payload.GetText(), a)
	var result []string
	for _, i := range py {
		result = append(result, strings.Join(i, " "))
	}
	return &pb.WordsReply{Text: result}, nil
}

func (s *Middle) Segmentation(_ context.Context, payload *pb.TextRequest) (*pb.WordsReply, error) {
	if payload.GetText() == "" {
		return &pb.WordsReply{Text: []string{}}, nil
	}
	// gse preload dict
	seg, err := gse.New("zh", "alpha")
	if err != nil {
		return nil, err
	}
	result := seg.Cut(payload.GetText(), true)
	return &pb.WordsReply{Text: result}, nil
}

func (s *Middle) Classifier(_ context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	rules, err := classifier.ReadRulesConfig(s.conf)
	if err != nil {
		return nil, err
	}

	c := classifier.NewClassifier()
	err = c.SetRules(rules)
	if err != nil {
		return nil, err
	}

	if payload.GetText() == "" {
		return nil, app.ErrInvalidParameter
	}

	res, err := c.Do(payload.GetText())
	if err != nil {
		if errors.Is(err, app.ErrInvalidParameter) {
			return &pb.TextReply{Text: ""}, nil
		}
		return nil, err
	}
	return &pb.TextReply{Text: string(res)}, nil
}

func (s *Middle) CreateAvatar(ctx context.Context, payload *pb.TextRequest) (*pb.TextReply, error) {
	avatarImage, err := util.Avatar(payload.Text)
	if err != nil {
		return nil, err
	}

	data := bytes.NewReader(avatarImage)
	buf := make([]byte, 1024)
	uc, err := s.storage.UploadFile(ctx)
	if err != nil {
		return nil, err
	}
	err = uc.Send(&pb.FileRequest{
		Data: &pb.FileRequest_Info{Info: &pb.FileInfo{FileType: "png"}},
	})
	if err != nil {
		return nil, err
	}
	for {
		n, err := data.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		err = uc.Send(&pb.FileRequest{Data: &pb.FileRequest_Chuck{Chuck: buf[:n]}})
		if err != nil {
			return nil, err
		}
	}
	fileReply, err := uc.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{Text: fileReply.Path}, nil
}
