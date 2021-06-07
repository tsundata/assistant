package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/internal/pkg/config"
)

func New(c *config.AppConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", c.Mysql.Url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

var ProviderSet = wire.NewSet(New)
