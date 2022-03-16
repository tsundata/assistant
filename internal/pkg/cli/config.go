package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/util"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Import default configs",
	Run: func(cmd *cobra.Command, args []string) {
		// Get a new client
		kv, err := etcd.New()
		if err != nil {
			panic(err)
		}

		err = filepath.Walk("./configs", func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if ext := filepath.Ext(path); ext != ".yml" {
				return nil
			}

			path = filepath.Clean(path)
			data, err := ioutil.ReadFile(path) // #nosec
			if err != nil {
				return err
			}

			name := strings.ReplaceAll(path, ".yml", "")
			name = strings.ReplaceAll(name, "configs", "")
			name = strings.ReplaceAll(name, `\`, "")
			name = strings.ReplaceAll(name, `/`, "")
			name = strings.ReplaceAll(name, "_", "/")
			_, err = kv.Put(context.Background(), fmt.Sprintf("config/%s", name), util.ByteToString(data))
			if err != nil {
				panic(err)
			}

			return err
		})
		if err != nil {
			panic(err)
		}

		// info
		resp, err := kv.Get(context.Background(), "config/common")
		if err != nil {
			panic(err)
		}
		for _, ev := range resp.Kvs {
			fmt.Printf("config common: %s\n", ev.Value)
		}

		fmt.Println("Done")
	},
}
