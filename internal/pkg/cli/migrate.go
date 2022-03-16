package cli

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	migratePkg "github.com/tsundata/assistant/internal/pkg/migrate"
	"gopkg.in/yaml.v2"
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

		d, err := iofs.New(migratePkg.Fs, "migrations")
		if err != nil {
			panic(err)
		}

		m, err := migrate.NewWithSourceInstance("iofs", d, fmt.Sprintf("mysql://%s", c.Mysql.Dsn))
		if err != nil {
			panic(err)
		}

		err = m.Up()
		if err != nil {
			panic(err)
		}
	},
}
