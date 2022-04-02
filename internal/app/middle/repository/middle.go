package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type MiddleRepository interface {
	CreatePage(ctx context.Context, page *pb.Page) (int64, error)
	GetPageByUUID(ctx context.Context, uuid string) (pb.Page, error)
	ListApps(ctx context.Context, userId int64) ([]*pb.App, error)
	GetAvailableAppByType(ctx context.Context, t string) (pb.App, error)
	GetAppByType(ctx context.Context, t string) (pb.App, error)
	UpdateAppByID(ctx context.Context, id int64, token, extra string) error
	CreateApp(ctx context.Context, app *pb.App) (int64, error)
	GetCredentialByName(ctx context.Context, userId int64, name string) (pb.Credential, error)
	GetCredentialByType(ctx context.Context, userId int64, t string) (pb.Credential, error)
	ListCredentials(ctx context.Context, userId int64) ([]*pb.Credential, error)
	CreateCredential(ctx context.Context, credential *pb.Credential) (int64, error)
	ListTags(ctx context.Context) ([]*pb.Tag, error)
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
}

type MysqlMiddleRepository struct {
	id *global.ID
	db *mysql.Conn
}

func NewMysqlMiddleRepository(id *global.ID, db *mysql.Conn) MiddleRepository {
	return &MysqlMiddleRepository{id: id, db: db}
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

func (r *MysqlMiddleRepository) GetAvailableAppByType(ctx context.Context, t string) (pb.App, error) {
	var find pb.App
	err := r.db.WithContext(ctx).Where("type = ?", t).Where("token <> ?", "").First(&find).Error
	if err != nil {
		return pb.App{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) GetAppByType(ctx context.Context, t string) (pb.App, error) {
	var find pb.App
	err := r.db.WithContext(ctx).Where("type = ?", t).Last(&find).Error
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

func (r *MysqlMiddleRepository) ListTags(ctx context.Context) ([]*pb.Tag, error) {
	var items []*pb.Tag
	err := r.db.WithContext(ctx).Order("id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) GetOrCreateTag(ctx context.Context, tag *pb.Tag) (pb.Tag, error) {
	var find pb.Tag
	err := r.db.WithContext(ctx).Where("name = ?", tag.Name).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return pb.Tag{}, err
	}

	if find.Id <= 0 {
		tag.Id = r.id.Generate(ctx)
		tag.CreatedAt = time.Now().Unix()
		tag.UpdatedAt = time.Now().Unix()
		err = r.db.WithContext(ctx).Create(&tag).Error
		if err != nil {
			return pb.Tag{}, err
		}
	}

	return find, nil
}

func (r *MysqlMiddleRepository) GetOrCreateModelTag(ctx context.Context, model *pb.ModelTag) (pb.ModelTag, error) {
	var find pb.ModelTag
	err := r.db.WithContext(ctx).Where("service = ? AND model = ? AND model_id = ? AND tag_id = ?", model.Service, model.Model, model.ModelId, model.TagId).First(&find).Error
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
