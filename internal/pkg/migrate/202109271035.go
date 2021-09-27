package migrate

import (
	_ "embed"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/tsundata/assistant/api/pb"
	"gorm.io/gorm"
	"strings"
)

//go:embed 202109271035.sql
var sql202109271035 string

var m202109271035 = &gormigrate.Migration{
	ID: "202109271035",
	Migrate: func(tx *gorm.DB) error {
		s := strings.Split(sql202109271035, ";")
		for _, item := range s {
			item := strings.TrimSpace(item)
			if item == "" {
				continue
			}
			if err := tx.Exec(item).Error; err != nil {
				return err
			}
		}
		err := tx.AutoMigrate(&pb.Group{}, &pb.Device{}, &pb.Bot{})// fixme
		if err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
