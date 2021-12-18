package migrate

import (
	_ "embed"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"strings"
)

//go:embed bots.sql
var sqlBots string

var mBots = &gormigrate.Migration{
	ID: "bots",
	Migrate: func(tx *gorm.DB) error {
		s := strings.Split(sqlBots, ";")
		for _, item := range s {
			item := strings.TrimSpace(item)
			if item == "" {
				continue
			}
			if err := tx.Exec(item).Error; err != nil {
				return err
			}
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
