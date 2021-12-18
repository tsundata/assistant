package migrate

import (
	_ "embed"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/tsundata/assistant/api/pb"
	"gorm.io/gorm"
	"strings"
)

//go:embed core.sql
var sqlCore string

var mCore = &gormigrate.Migration{
	ID: "core",
	Migrate: func(tx *gorm.DB) error {
		s := strings.Split(sqlCore, ";")
		for _, item := range s {
			item := strings.TrimSpace(item)
			if item == "" {
				continue
			}
			if err := tx.Exec(item).Error; err != nil {
				return err
			}
		}
		err := tx.AutoMigrate(&pb.Group{}, &pb.Device{}, &pb.Bot{}, &pb.Node{}) // fixme
		if err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
