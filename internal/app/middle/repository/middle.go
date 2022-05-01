package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type MiddleRepository interface {
	CreatePage(ctx context.Context, page *pb.Page) (int64, error)
	GetPageByUUID(ctx context.Context, uuid string) (pb.Page, error)
	ListApps(ctx context.Context, userId int64) ([]*pb.App, error)
	GetAvailableAppByType(ctx context.Context, userId int64, t string) (pb.App, error)
	GetAppByType(ctx context.Context, userId int64, t string) (pb.App, error)
	UpdateAppByID(ctx context.Context, id int64, token, extra string) error
	CreateApp(ctx context.Context, app *pb.App) (int64, error)
	GetCredentialByName(ctx context.Context, userId int64, name string) (pb.Credential, error)
	GetCredentialByType(ctx context.Context, userId int64, t string) (pb.Credential, error)
	ListCredentials(ctx context.Context, userId int64) ([]*pb.Credential, error)
	CreateCredential(ctx context.Context, credential *pb.Credential) (int64, error)
	ListTags(ctx context.Context, userId int64) ([]*pb.Tag, error)
	ListModelTagsByModelId(ctx context.Context, userId int64, modelId []int64) ([]*pb.ModelTag, error)
	ListModelTagsByModel(ctx context.Context, userId int64, model pb.ModelTag) ([]*pb.ModelTag, error)
	ListModelTagsByTag(ctx context.Context, userId int64, tag string) ([]*pb.ModelTag, error)
	GetOrCreateTag(ctx context.Context, tag *pb.Tag) (pb.Tag, error)
	GetOrCreateModelTag(ctx context.Context, tag *pb.ModelTag) (pb.ModelTag, error)
	ListSubscribe(ctx context.Context) ([]*pb.Subscribe, error)
	CreateSubscribe(ctx context.Context, subscribe pb.Subscribe) error
	UpdateSubscribeStatus(ctx context.Context, name string, status int64) error
	ListUserSubscribe(ctx context.Context, userId int64) ([]*pb.KV, error)
	CreateUserSubscribe(ctx context.Context, subscribe pb.UserSubscribe) error
	UpdateUserSubscribeStatus(ctx context.Context, userId, subscribeId int64, status int64) error
	GetUserSubscribe(ctx context.Context, userId, subscribeId int64) (pb.UserSubscribe, error)
	GetSubscribe(ctx context.Context, name string) (pb.Subscribe, error)
	CreateCounter(ctx context.Context, counter *pb.Counter) (int64, error)
	IncreaseCounter(ctx context.Context, id, amount int64) error
	DecreaseCounter(ctx context.Context, id, amount int64) error
	ListCounter(ctx context.Context, userId int64) ([]*pb.Counter, error)
	GetCounter(ctx context.Context, id int64) (pb.Counter, error)
	GetCounterByFlag(ctx context.Context, userId int64, flag string) (pb.Counter, error)
	Search(ctx context.Context, userId int64, filter [][]string) ([]*pb.Metadata, error)
	CollectMetadata(ctx context.Context, models []interface{}) error
}

type MysqlMiddleRepository struct {
	logger log.Logger
	id     *global.ID
	db     *mysql.Conn
}

func NewMysqlMiddleRepository(logger log.Logger, id *global.ID, db *mysql.Conn) MiddleRepository {
	return &MysqlMiddleRepository{logger: logger, id: id, db: db}
}

func (r *MysqlMiddleRepository) CreatePage(ctx context.Context, page *pb.Page) (int64, error) {
	page.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&page).Error
	if err != nil {
		return 0, err
	}
	return page.Id, nil
}

func (r *MysqlMiddleRepository) GetPageByUUID(ctx context.Context, uuid string) (pb.Page, error) {
	var find pb.Page
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&find).Error
	if err != nil {
		return pb.Page{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) ListApps(ctx context.Context, userId int64) ([]*pb.App, error) {
	var items []*pb.App
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("created_at DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) GetAvailableAppByType(ctx context.Context, userId int64, t string) (pb.App, error) {
	var find pb.App
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Where("type = ?", t).
		Where("token <> ?", "").First(&find).Error
	if err != nil {
		return pb.App{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) GetAppByType(ctx context.Context, userId int64, t string) (pb.App, error) {
	var find pb.App
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Where("type = ?", t).Last(&find).Error
	if err != nil {
		return pb.App{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) UpdateAppByID(ctx context.Context, id int64, token, extra string) error {
	return r.db.WithContext(ctx).Model(&pb.App{}).Where("id = ?", id).
		UpdateColumns(map[string]interface{}{
			"token": token,
			"extra": extra,
		}).Error
}

func (r *MysqlMiddleRepository) CreateApp(ctx context.Context, app *pb.App) (int64, error) {
	app.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&app).Error
	if err != nil {
		return 0, err
	}
	return app.Id, nil
}

func (r *MysqlMiddleRepository) GetCredentialByName(ctx context.Context, userId int64, name string) (pb.Credential, error) {
	var find pb.Credential
	err := r.db.WithContext(ctx).Where("user_id = ? AND name = ?", userId, name).First(&find).Error
	if err != nil {
		return pb.Credential{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) GetCredentialByType(ctx context.Context, userId int64, t string) (pb.Credential, error) {
	var find pb.Credential
	err := r.db.WithContext(ctx).Where("user_id = ? AND type = ?", userId, t).First(&find).Error
	if err != nil {
		return pb.Credential{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) ListCredentials(ctx context.Context, userId int64) ([]*pb.Credential, error) {
	var items []*pb.Credential
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) CreateCredential(ctx context.Context, credential *pb.Credential) (int64, error) {
	credential.Id = r.id.Generate(ctx)
	err := r.db.WithContext(ctx).Create(&credential).Error
	if err != nil {
		return 0, err
	}
	return credential.Id, nil
}

func (r *MysqlMiddleRepository) ListTags(ctx context.Context, userId int64) ([]*pb.Tag, error) {
	var items []*pb.Tag
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Order("id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) ListModelTagsByModelId(ctx context.Context, userId int64, modelId []int64) ([]*pb.ModelTag, error) {
	var m []struct {
		ModelId int64
		Name    string
	}

	err := r.db.WithContext(ctx).
		Model(&pb.ModelTag{}).
		Select("model_id, name").
		Where("model_tags.user_id = ? AND model_tags.model_id IN ?", userId, modelId).
		Joins("LEFT JOIN tags ON model_tags.tag_id = tags.id").
		Order("model_tags.id DESC").Find(&m).Error
	if err != nil {
		return nil, err
	}

	var items []*pb.ModelTag
	for _, item := range m {
		items = append(items, &pb.ModelTag{
			ModelId: item.ModelId,
			Name:    item.Name,
		})
	}

	return items, nil
}

func (r *MysqlMiddleRepository) ListModelTagsByModel(ctx context.Context, userId int64, model pb.ModelTag) ([]*pb.ModelTag, error) {
	var items []*pb.ModelTag

	err := r.db.WithContext(ctx).
		Where("model_tags.user_id = ?", userId).
		Where("model_tags.service = ? AND model_tags.model = ? AND tags.name = ?", model.Service, model.Model, model.Name).
		Joins("LEFT JOIN tags ON model_tags.tag_id = tags.id").
		Order("model_tags.id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *MysqlMiddleRepository) ListModelTagsByTag(ctx context.Context, userId int64, tag string) ([]*pb.ModelTag, error) {
	var items []*pb.ModelTag

	err := r.db.WithContext(ctx).
		Where("model_tags.user_id = ?", userId).
		Where("tags.name = ?", tag).
		Joins("LEFT JOIN tags ON model_tags.tag_id = tags.id").
		Order("model_tags.id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *MysqlMiddleRepository) GetOrCreateTag(ctx context.Context, tag *pb.Tag) (pb.Tag, error) {
	var find pb.Tag
	err := r.db.WithContext(ctx).Where("user_id = ? AND name = ?", tag.UserId, tag.Name).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return pb.Tag{}, err
	}

	if find.Id <= 0 {
		find.Id = r.id.Generate(ctx)
		find.UserId = tag.UserId
		find.Name = tag.Name
		find.CreatedAt = time.Now().Unix()
		find.UpdatedAt = time.Now().Unix()
		err = r.db.WithContext(ctx).Create(&find).Error
		if err != nil {
			return pb.Tag{}, err
		}
	}

	return find, nil
}

func (r *MysqlMiddleRepository) GetOrCreateModelTag(ctx context.Context, model *pb.ModelTag) (pb.ModelTag, error) {
	var find pb.ModelTag
	err := r.db.WithContext(ctx).Where("user_id = ? AND service = ? AND model = ? AND model_id = ? AND tag_id = ?",
		model.UserId, model.Service, model.Model, model.ModelId, model.TagId).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return pb.ModelTag{}, err
	}

	if find.Id <= 0 {
		model.Id = r.id.Generate(ctx)
		model.CreatedAt = time.Now().Unix()
		model.UpdatedAt = time.Now().Unix()
		err = r.db.WithContext(ctx).Create(&model).Error
		if err != nil {
			return pb.ModelTag{}, err
		}
	}

	return find, nil
}

func (r *MysqlMiddleRepository) ListSubscribe(ctx context.Context) ([]*pb.Subscribe, error) {
	var items []*pb.Subscribe
	err := r.db.WithContext(ctx).Order("id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) CreateSubscribe(ctx context.Context, subscribe pb.Subscribe) error {
	var find pb.Subscribe
	err := r.db.WithContext(ctx).Where("name = ?", subscribe.Name).Take(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.Id <= 0 {
		subscribe.Id = r.id.Generate(ctx)
		subscribe.CreatedAt = time.Now().Unix()
		subscribe.UpdatedAt = time.Now().Unix()
		err = r.db.WithContext(ctx).Create(&subscribe).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MysqlMiddleRepository) UpdateSubscribeStatus(ctx context.Context, name string, status int64) error {
	return r.db.WithContext(ctx).Model(&pb.Subscribe{}).Where("name = ?", name).
		UpdateColumns(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now().Unix(),
		}).Error
}

func (r *MysqlMiddleRepository) ListUserSubscribe(ctx context.Context, userId int64) ([]*pb.KV, error) {
	var items []struct {
		Name   string
		Status int
	}
	err := r.db.WithContext(ctx).
		Model(&pb.UserSubscribe{}).
		Select("subscribes.name as name, user_subscribes.status as status").
		Where("user_id = ?", userId).
		Joins("LEFT JOIN subscribes ON subscribes.id = user_subscribes.subscribe_id").
		Order("user_subscribes.id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	var result []*pb.KV
	for _, item := range items {
		result = append(result, &pb.KV{
			Key:   item.Name,
			Value: strconv.Itoa(item.Status),
		})
	}

	return result, nil
}

func (r *MysqlMiddleRepository) CreateUserSubscribe(ctx context.Context, subscribe pb.UserSubscribe) error {
	var find pb.UserSubscribe
	err := r.db.WithContext(ctx).Where("user_id = ? AND subscribe_id = ?", subscribe.UserId, subscribe.SubscribeId).Take(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.Id <= 0 {
		subscribe.Id = r.id.Generate(ctx)
		subscribe.CreatedAt = time.Now().Unix()
		subscribe.UpdatedAt = time.Now().Unix()
		err = r.db.WithContext(ctx).Create(&subscribe).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *MysqlMiddleRepository) UpdateUserSubscribeStatus(ctx context.Context, userId, subscribeId int64, status int64) error {
	return r.db.WithContext(ctx).Model(&pb.Subscribe{}).Where("user_id = ? AND subscribe_id = ?", userId, subscribeId).
		UpdateColumns(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now().Unix(),
		}).Error
}

func (r *MysqlMiddleRepository) GetUserSubscribe(ctx context.Context, userId, subscribeId int64) (pb.UserSubscribe, error) {
	var find pb.UserSubscribe
	err := r.db.WithContext(ctx).Where("user_id = ? AND subscribe_id = ?", userId, subscribeId).
		Take(&find).Error
	if err != nil {
		return pb.UserSubscribe{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) GetSubscribe(ctx context.Context, name string) (pb.Subscribe, error) {
	var find pb.Subscribe
	err := r.db.WithContext(ctx).Where("name = ?", name).Take(&find).Error
	if err != nil {
		return pb.Subscribe{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) CreateCounter(ctx context.Context, counter *pb.Counter) (int64, error) {
	counter.Id = r.id.Generate(ctx)
	counter.CreatedAt = time.Now().Unix()
	counter.UpdatedAt = time.Now().Unix()
	err := r.db.WithContext(ctx).Create(&counter)
	if err != nil {
		return 0, nil
	}
	r.record(ctx, counter.Id, counter.Digit)
	return counter.Id, nil
}

func (r *MysqlMiddleRepository) IncreaseCounter(ctx context.Context, id, amount int64) error {
	err := r.db.WithContext(ctx).Model(&pb.Counter{}).
		Where("id = ?", id).
		Update("digit", gorm.Expr("digit + ?", amount)).Error
	if err != nil {
		return err
	}
	r.record(ctx, id, amount)
	return nil
}

func (r *MysqlMiddleRepository) DecreaseCounter(ctx context.Context, id, amount int64) error {
	err := r.db.WithContext(ctx).Model(&pb.Counter{}).
		Where("id = ?", id).
		Update("digit", gorm.Expr("digit - ?", amount)).Error
	if err != nil {
		return err
	}
	r.record(ctx, id, -amount)
	return nil
}

func (r *MysqlMiddleRepository) ListCounter(ctx context.Context, userId int64) ([]*pb.Counter, error) {
	var items []*pb.Counter
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).
		Order("updated_at DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) record(ctx context.Context, id, digit int64) {
	err := r.db.WithContext(ctx).Exec("INSERT INTO `counter_records` (`id`, `counter_id`, `digit`, `created_at`) VALUES (?, ?, ?, ?)",
		r.id.Generate(ctx), id, digit, time.Now().Unix()).Error
	if err != nil {
		r.logger.Error(err)
	}
}

func (r *MysqlMiddleRepository) GetCounter(ctx context.Context, id int64) (pb.Counter, error) {
	var find pb.Counter
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&find).Error
	if err != nil {
		return pb.Counter{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) GetCounterByFlag(ctx context.Context, userId int64, flag string) (pb.Counter, error) {
	var find pb.Counter
	err := r.db.WithContext(ctx).Where("user_id = ? AND flag = ?", userId, flag).First(&find).Error
	if err != nil {
		return pb.Counter{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) Search(ctx context.Context, userId int64, filter [][]string) ([]*pb.Metadata, error) {
	var items []*pb.Metadata
	builder := r.db.WithContext(ctx).Where("user_id = ?", userId)

	for _, item := range filter {
		if len(item) != 2 {
			continue
		}
		switch item[0] {
		case "model":
			builder.Where("model = ?", item[1])
		case "model_id":
			builder.Where("model_id = ?", item[1])
		case "text":
			builder.Where("text = ?", item[1])
		default:
			builder.Where(fmt.Sprintf("`data`->'$.\"%s\"' = ? OR `extra`->'$.\"%s\"' = ?", item[0], item[0]), item[1], item[1])
		}
	}

	err := builder.Limit(10).Order("updated_at DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) CollectMetadata(ctx context.Context, models []interface{}) error {
	for _, m := range models {
		r.db.WithContext(ctx).Find(&m)
		jsonByte, _ := json.Marshal(m)
		arrData := util.ByteToString(jsonByte)
		arrValue := gjson.Get(arrData, "@this")

		for _, result := range arrValue.Array() {
			modelId := gjson.Get(result.Raw, "id").Int()
			userId := gjson.Get(result.Raw, "user_id").Int()
			sequence := gjson.Get(result.Raw, "sequence").Int()
			extra := "{}"
			extraValue := gjson.Get(result.Raw, "payload")
			if extraValue.Exists() {
				extra = extraValue.String()
			}
			model := util.ModelName(m)

			name := gjson.Get(result.Raw, "name").String()
			title := gjson.Get(result.Raw, "title").String()
			text := gjson.Get(result.Raw, "text").String()
			content := gjson.Get(result.Raw, "content").String()
			b := strings.Builder{}
			b.WriteString(name)
			b.WriteString(title)
			b.WriteString(text)
			b.WriteString(content)
			text = util.SubString(b.String(), 0, 100)

			var find pb.Metadata
			r.db.WithContext(ctx).
				Model(&pb.Metadata{}).
				Select("id").
				Where("user_id = ? AND model = ? AND model_id = ?", userId, model, modelId).
				Take(&find)
			if find.Id > 0 {
				r.db.WithContext(ctx).
					Model(&pb.Metadata{}).
					Where("user_id = ? AND model = ? AND model_id = ?", userId, model, modelId).
					UpdateColumns(map[string]interface{}{
						"text":       text,
						"data":       result.Raw,
						"extra":      extra,
						"updated_at": time.Now().Unix(),
					})
			} else {
				r.db.WithContext(ctx).Create(&pb.Metadata{
					Id:        r.id.Generate(ctx),
					UserId:    userId,
					Model:     model,
					ModelId:   modelId,
					Sequence:  sequence,
					Text:      text,
					Data:      result.Raw,
					Extra:     extra,
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				})
			}
		}
	}
	return nil
}
