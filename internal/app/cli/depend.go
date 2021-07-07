package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(dependCmd)
}

var dependCmd = &cobra.Command{
	Use:   "depend",
	Short: "App service dependency chain",
	Run: func(cmd *cobra.Command, args []string) {
		dirs, _ := filepath.Glob("./internal/app/*")
		for _, dir := range dirs {
			app := strings.ReplaceAll(dir, "internal/app/", "")
			app = strings.ReplaceAll(app, `internal\app\`, "")

			subItems, _ := filepath.Glob(fmt.Sprintf("%s/rpcclient/*.go", dir))
			if len(subItems) == 0 {
				continue
			}

			for _, item := range subItems {
				service := strings.ReplaceAll(item, fmt.Sprintf("internal/app/%s/rpcclient/", app), "")
				service = strings.ReplaceAll(service, fmt.Sprintf(`internal\app\%s\rpcclient\`, app), "")
				service = strings.ReplaceAll(service, `.go`, "")
				if service == "client" {
					continue
				}
				fmt.Println(app, service)
			}
		}
	},
}
