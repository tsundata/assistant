package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
	"time"
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

type RqliteMiddleRepository struct {
	db *rqlite.Conn
}

func NewRqliteMiddleRepository(db *rqlite.Conn) *RqliteMiddleRepository {
	return &RqliteMiddleRepository{db: db}
}

func (r *RqliteMiddleRepository) CreatePage(page pb.Page) (int64, error) {
	res, err := r.db.WriteOne("INSERT INTO `pages` (`uuid`, `type`, `title`, `content`) VALUES ('%s', '%s', '%s', '%s')", page.Uuid, page.Type, page.Title, page.Content)
	if err != nil {
		return 0, err
	}
	return res.LastInsertID, nil
}

func (r *RqliteMiddleRepository) GetPageByUUID(uuid string) (pb.Page, error) {
	rows, err := r.db.QueryOne("SELECT uuid, `type`, title, content FROM `pages` WHERE `uuid` = '%s'", uuid)
	if err != nil {
		return pb.Page{}, nil
	}

	var find pb.Page
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Page{}, err
		}
		util.Inject(&find, m)
	}

	return find, nil
}

func (r *RqliteMiddleRepository) ListApps() ([]pb.App, error) {
	rows, err := r.db.QueryOne("SELECT name, `type`, token, extra, `created_at` FROM `apps` ORDER BY `created_at` DESC")
	if err != nil {
		return nil, err
	}

	var apps []pb.App
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.App
		util.Inject(&item, m)
		apps = append(apps, item)
	}

	return apps, nil
}

func (r *RqliteMiddleRepository) GetAvailableAppByType(t string) (pb.App, error) {
	rows, err := r.db.QueryOne("SELECT id, name, `type`, token FROM apps WHERE `type` = '%s' AND `token` <> '' LIMIT 1", t)
	if err != nil {
		return pb.App{}, err
	}

	var find pb.App
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.App{}, err
		}
		util.Inject(&find, m)
	}

	return find, nil
}

func (r *RqliteMiddleRepository) GetAppByType(t string) (pb.App, error) {
	rows, err := r.db.QueryOne("SELECT id FROM apps WHERE type = '%s' ORDER BY id DESC LIMIT 1", t)
	if err != nil {
		return pb.App{}, nil
	}

	var find pb.App
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.App{}, err
		}
		util.Inject(&find, m)
	}

	return find, nil
}

func (r *RqliteMiddleRepository) UpdateAppByID(id int64, token, extra string) error {
	_, err := r.db.WriteOne("UPDATE apps SET `token` = '%s', `extra` = '%s', `created_at` = '%s' WHERE id = '%d'", token, extra, time.Now(), id)
	return err
}

func (r *RqliteMiddleRepository) CreateApp(app pb.App) (int64, error) {
	res, err := r.db.WriteOne("INSERT INTO `apps` (`name`, `type`, `token`, `extra`) VALUES ('%s', '%s', '%s', '%s')",
		app.Name, app.Type, app.Token, app.Extra)
	if err != nil {
		return 0, err
	}
	return res.LastInsertID, err
}

func (r *RqliteMiddleRepository) GetCredentialByName(name string) (pb.Credential, error) {
	rows, err := r.db.QueryOne("SELECT * FROM credentials WHERE name = '%s' LIMIT 1", name)
	if err != nil {
		return pb.Credential{}, nil
	}

	var find pb.Credential
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Credential{}, err
		}
		util.Inject(&find, m)
	}

	return find, nil
}

func (r *RqliteMiddleRepository) GetCredentialByType(t string) (pb.Credential, error) {
	rows, err := r.db.QueryOne("SELECT * FROM credentials WHERE type = '%s' LIMIT 1", t)
	if err != nil {
		return pb.Credential{}, nil
	}

	var find pb.Credential
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Credential{}, err
		}
		util.Inject(&find, m)
	}

	return find, nil
}

func (r *RqliteMiddleRepository) ListCredentials() ([]pb.Credential, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `credentials` ORDER BY `id` DESC")
	if err != nil {
		return nil, err
	}

	var items []pb.Credential
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Credential
		util.Inject(&item, m)
		items = append(items, item)
	}

	return items, nil
}

func (r *RqliteMiddleRepository) CreateCredential(credential pb.Credential) (int64, error) {
	res, err := r.db.WriteOne("INSERT INTO `credentials` (`name`, `type`, `content`) VALUES ('%s', '%s', '%s')",
		credential.Name, credential.Type, credential.Content)
	if err != nil {
		return 0, err
	}
	return res.LastInsertID, nil
}

func (r *RqliteMiddleRepository) ListTags() ([]pb.Tag, error) {
	rows, err := r.db.QueryOne("SELECT * FROM `tags` ORDER BY `id` DESC")
	if err != nil {
		return nil, err
	}

	var items []pb.Tag
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return nil, err
		}
		var item pb.Tag
		util.Inject(&item, m)
		items = append(items, item)
	}

	return items, nil
}

func (r *RqliteMiddleRepository) GetOrCreateTag(tag pb.Tag) (pb.Tag, error) {
	rows, err := r.db.QueryOne("SELECT * FROM tags WHERE name = '%s' LIMIT 1", tag.Name)
	if err != nil {
		return pb.Tag{}, nil
	}

	var find pb.Tag
	for rows.Next() {
		m, err := rows.Map()
		if err != nil {
			return pb.Tag{}, err
		}
		util.Inject(&find, m)
	}

	if find.Id <= 0 {
		now := util.Now()
		res, err := r.db.WriteOne("INSERT INTO `tags` (`name`, `created_at`) VALUES ('%s', '%s')", tag.Name, now)
		if err != nil {
			return pb.Tag{}, err
		}
		tag.Id = res.LastInsertID
		tag.CreatedAt = time.Now().Unix()
		return tag, nil
	}

	return find, nil
}
