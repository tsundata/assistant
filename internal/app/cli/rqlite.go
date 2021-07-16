package cli

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func init() {
	rootCmd.AddCommand(rqliteCmd)
}

var rqliteCmd = &cobra.Command{
	Use:   "rqlite",
	Short: "import sql to rqlite",
	Run: func(cmd *cobra.Command, args []string) {
		sql, err := ioutil.ReadFile("./scripts/sqlite.sql")
		if err != nil {
			fmt.Println(err)
			return
		}
		client := resty.New()
		resp, err := client.R().
			SetBody([]string{string(sql)}).
			Post("http://127.0.0.1:4001/db/execute?pretty&timings")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(resp.Body()))
	},
}
