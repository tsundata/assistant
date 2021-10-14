package migrate

import (
	_ "embed"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"strings"
)

//go:embed 202109181651.sql
var sql202109181651 string

var m202109181651 = &gormigrate.Migration{
	ID: "202109181651",
	Migrate: func(tx *gorm.DB) error {
		s := strings.Split(sql202109181651, ";")
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