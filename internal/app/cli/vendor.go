package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(vendorCmd)
}

var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Remove vendor",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.RemoveAll("vendor")
		if err != nil {
			fmt.Println(err)
		}
	},
}
