package repository

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/util"
	"time"
)

type MiddleRepository interface {
	CreatePage(page pb.Page) (int64, error)
	GetPageByUUID(uuid string) (pb.Page, error)
	ListApps() ([]pb.App, error)
	GetAvailableAppByType(t string) (pb.App, error)
	GetAppByType(t string) (pb.App, error)
	UpdateAppByID(id int64, token, extra string) error
	CreateApp(app pb.App) (int64, error)
	GetCredentialByName(name string) (pb.Credential, error)
	GetCredentialByType(t string) (pb.Credential, error)
	ListCredentials() ([]pb.Credential, error)
	CreateCredential(credential pb.Credential) (int64, error)
}

type RqliteMiddleRepository struct {
	db *rqlite.Conn
}

func NewRqliteMiddleRepository(db *rqlite.Conn) MiddleRepository {
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
	rows, err := r.db.QueryOne("SELECT id, name, `type` FROM credentials WHERE name = '%s' LIMIT 1", name)
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
	rows, err := r.db.QueryOne("SELECT id, name, `type` FROM credentials WHERE type = '%s' LIMIT 1", t)
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
	rows, err := r.db.QueryOne("SELECT name, `type`, content, `created_at` FROM `credentials` ORDER BY `id` DESC")
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
