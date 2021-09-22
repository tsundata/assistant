package mysql

import (
	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Conn struct {
	*gorm.DB
}

func New(c *config.AppConfig) (*Conn, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	})
	if err != nil {
		return nil, err
	}

	// migrate
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrate.Sets)
	if err := m.Migrate(); err != nil {
		panic(err)
	}

	return &Conn{db}, nil
}

var ProviderSet = wire.NewSet(New)
