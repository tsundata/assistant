package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rootCmd.AddCommand(passwordCmd)
}

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Password Hashing (bcrypt)",
	Run: func(cmd *cobra.Command, args []string) {
		hashing, err := bcrypt.GenerateFromPassword([]byte(args[0]), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(hashing))
	},
}
