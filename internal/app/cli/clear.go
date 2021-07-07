package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(clearCmd)
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear api/pb/*.pb.go",
	Run: func(cmd *cobra.Command, args []string) {
		matches, err := filepath.Glob("api/pb/*.pb.go")
		if err != nil {
			panic(err)
		}
		for _, i := range matches {
			err = os.Remove(i)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("remove %s\n", i)
		}
	},
}
