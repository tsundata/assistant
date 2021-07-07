package cli

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
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
		client, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			panic(err)
		}

		// Get a handle to the KV API
		kv := client.KV()

		err = filepath.Walk("./configs", func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if ext := filepath.Ext(path); ext != ".yml" {
				return nil
			}

			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			name := strings.ReplaceAll(path, ".yml", "")
			name = strings.ReplaceAll(name, "configs", "")
			name = strings.ReplaceAll(name, `\`, "")
			name = strings.ReplaceAll(name, `/`, "")
			name = strings.ReplaceAll(name, "_", "/")
			p := &api.KVPair{Key: fmt.Sprintf("config/%s", name), Value: data}
			_, err = kv.Put(p, nil)
			if err != nil {
				panic(err)
			}

			return err
		})
		if err != nil {
			panic(err)
		}

		// info
		pair, _, err := kv.Get("config/common", nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("config common: %s\n", pair.Value)

		fmt.Println("Done")
	},
}
