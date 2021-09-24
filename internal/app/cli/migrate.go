package cli

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	mysql2 "github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/migrate"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration",
	Run: func(cmd *cobra.Command, args []string) {
		kv, err := etcd.New()
		if err != nil {
			panic(err)
		}
		resp, err := kv.Get(context.Background(), "config/common")
		if err != nil {
			panic(err)
		}
		var value []byte
		for _, ev := range resp.Kvs {
			value = ev.Value
		}
		var c config.AppConfig
		err = yaml.Unmarshal(value, &c)
		if err != nil {
			panic(err)
		}

		db, err := gorm.Open(mysql.Open(c.Mysql.Dsn))
		if err != nil {
			panic(err)
		}

		migrate.Run(&mysql2.Conn{DB: db})
	},
}
