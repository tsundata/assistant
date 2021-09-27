package migrate

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
)

func Run(conn *mysql.Conn) {
	// migrate
	m := gormigrate.New(conn.DB, &gormigrate.Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}, []*gormigrate.Migration{
		m202109181651,
		m202109271035,
	})
	if err := m.Migrate(); err != nil {
		panic(err)
	}
}
