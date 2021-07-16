package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/log"
	"time"
)

type MiddleRepository interface {
	CreatePage(page pb.Page) (int64, error)
	GetPageByUUID(uuid string) (pb.Page, error)
	ListApps() ([]model.App, error)
	GetAvailableAppByType(t string) (model.App, error)
	GetAppByType(t string) (model.App, error)
	UpdateAppByID(id int64, token, extra string) error
	CreateApp(app model.App) (int64, error)
	GetCredentialByName(name string) (model.Credential, error)
	GetCredentialByType(t string) (model.Credential, error)
	ListCredentials() ([]model.Credential, error)
	CreateCredential(credential model.Credential) (int64, error)
}

type MysqlMiddleRepository struct {
	logger log.Logger
	db     *sqlx.DB
}

func NewMysqlMiddleRepository(logger log.Logger, db *sqlx.DB) MiddleRepository {
	return &MysqlMiddleRepository{logger: logger, db: db}
}

func (r *MysqlMiddleRepository) CreatePage(page pb.Page) (int64, error) {
	res, err := r.db.NamedExec("INSERT INTO `pages` (`uuid`, `type`, `title`, `content`) VALUES (:uuid, :type, :title, :content)", page)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MysqlMiddleRepository) GetPageByUUID(uuid string) (pb.Page, error) {
	var find pb.Page
	err := r.db.Get(&find, "SELECT uuid, `type`, title, content FROM `pages` WHERE `uuid` = ?", uuid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return pb.Page{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) ListApps() ([]model.App, error) {
	var apps []model.App
	err := r.db.Select(&apps, "SELECT name, `type`, token, extra, `created_at` FROM `apps` ORDER BY `created_at` DESC")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return apps, nil
}

func (r *MysqlMiddleRepository) GetAvailableAppByType(t string) (model.App, error) {
	var find model.App
	err := r.db.Get(&find, "SELECT id, name, `type`, token FROM apps WHERE `type` = ? AND `token` <> '' LIMIT 1", t)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.App{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) GetAppByType(t string) (model.App, error) {
	var app model.App
	err := r.db.Get(&app, "SELECT id FROM apps WHERE type = ? ORDER BY id DESC LIMIT 1", t)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.App{}, err
	}
	return app, nil
}

func (r *MysqlMiddleRepository) UpdateAppByID(id int64, token, extra string) error {
	_, err := r.db.Exec("UPDATE apps SET `token` = ?, `extra` = ?, `created_at` = ? WHERE id = ?", token, extra, time.Now(), id)
	return err
}

func (r *MysqlMiddleRepository) CreateApp(app model.App) (int64, error) {
	res, err := r.db.Exec("INSERT INTO `apps` (`name`, `type`, `token`, `extra`) VALUES (?, ?, ?, ?)",
		app.Name, app.Type, app.Token, app.Extra)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *MysqlMiddleRepository) GetCredentialByName(name string) (model.Credential, error) {
	var find model.Credential
	err := r.db.Get(&find, "SELECT id, name, `type` FROM credentials WHERE name = ? LIMIT 1", name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Credential{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) GetCredentialByType(t string) (model.Credential, error) {
	var find model.Credential
	err := r.db.Get(&find, "SELECT id, name, `type` FROM credentials WHERE type = ? LIMIT 1", t)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.Credential{}, err
	}
	return find, nil
}

func (r *MysqlMiddleRepository) ListCredentials() ([]model.Credential, error) {
	var items []model.Credential
	err := r.db.Select(&items, "SELECT name, `type`, content, `created_at` FROM `credentials` ORDER BY `id` DESC")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return items, nil
}

func (r *MysqlMiddleRepository) CreateCredential(credential model.Credential) (int64, error) {
	res, err := r.db.Exec("INSERT INTO `credentials` (`name`, `type`, `content`) VALUES (?, ?, ?)",
		credential.Name, credential.Type, credential.Content)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
