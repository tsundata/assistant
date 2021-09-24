package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"gorm.io/gorm"
)

type MiddleRepository interface {
	CreatePage(ctx context.Context, page *pb.Page) (int64, error)
	GetPageByUUID(ctx context.Context, uuid string) (*pb.Page, error)
	ListApps(ctx context.Context, ) ([]*pb.App, error)
	GetAvailableAppByType(ctx context.Context, t string) (*pb.App, error)
	GetAppByType(ctx context.Context, t string) (*pb.App, error)
	UpdateAppByID(ctx context.Context, id int64, token, extra string) error
	CreateApp(ctx context.Context, app *pb.App) (int64, error)
	GetCredentialByName(ctx context.Context, name string) (*pb.Credential, error)
	GetCredentialByType(ctx context.Context, t string) (*pb.Credential, error)
	ListCredentials(ctx context.Context, ) ([]*pb.Credential, error)
	CreateCredential(ctx context.Context, credential *pb.Credential) (int64, error)
	ListTags(ctx context.Context, ) ([]*pb.Tag, error)
	GetOrCreateTag(ctx context.Context, tag *pb.Tag) (*pb.Tag, error)
}

type MysqlMiddleRepository struct {
	db *mysql.Conn
}

func NewMysqlMiddleRepository(db *mysql.Conn) MiddleRepository {
	return &MysqlMiddleRepository{db: db}
}

func (r *MysqlMiddleRepository) CreatePage(ctx context.Context, page *pb.Page) (int64, error) {
	err := r.db.WithContext(ctx).Create(&page).Error
	if err != nil {
		return 0, err
	}
	return page.Id, nil
}

func (r *MysqlMiddleRepository) GetPageByUUID(ctx context.Context, uuid string) (*pb.Page, error) {
	var find pb.Page
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlMiddleRepository) ListApps(ctx context.Context) ([]*pb.App, error) {
	var items []*pb.App
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) GetAvailableAppByType(ctx context.Context, t string) (*pb.App, error) {
	var find pb.App
	err := r.db.WithContext(ctx).Where("type = ?", t).Where("token <> ?", "").First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlMiddleRepository) GetAppByType(ctx context.Context, t string) (*pb.App, error) {
	var find pb.App
	err := r.db.WithContext(ctx).Where("type = ?", t).Last(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlMiddleRepository) UpdateAppByID(ctx context.Context, id int64, token, extra string) error {
	return r.db.WithContext(ctx).Model(&pb.App{}).Where("id = ?", id).Update("token", token).
		Update("extra", extra).Error
}

func (r *MysqlMiddleRepository) CreateApp(ctx context.Context, app *pb.App) (int64, error) {
	err := r.db.WithContext(ctx).Create(&app).Error
	if err != nil {
		return 0, err
	}
	return app.Id, nil
}

func (r *MysqlMiddleRepository) GetCredentialByName(ctx context.Context, name string) (*pb.Credential, error) {
	var find pb.Credential
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlMiddleRepository) GetCredentialByType(ctx context.Context, t string) (*pb.Credential, error) {
	var find pb.Credential
	err := r.db.WithContext(ctx).Where("type = ?", t).First(&find).Error
	if err != nil {
		return nil, err
	}
	return &find, nil
}

func (r *MysqlMiddleRepository) ListCredentials(ctx context.Context) ([]*pb.Credential, error) {
	var items []*pb.Credential
	err := r.db.WithContext(ctx).Order("id DESC").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) CreateCredential(ctx context.Context, credential *pb.Credential) (int64, error) {
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

func (r *MysqlMiddleRepository) GetOrCreateTag(ctx context.Context, tag *pb.Tag) (*pb.Tag, error) {
	var find pb.Tag
	err := r.db.WithContext(ctx).Where("name = ?", tag.Name).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if find.Id <= 0 {
		err = r.db.WithContext(ctx).Create(&tag).Error
		if err != nil {
			return nil, err
		}
	}

	return &find, nil
}
