package cmd

import (
	"github.com/gaozhenhai/api-test/http/user"
	"github.com/spf13/cobra"
)

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create tenxcloud users",
	Run:   user.CreateTenxcloudUsers,
}

func init() {
	userCmd.AddCommand(userCreateCmd)
	userCreateCmd.Flags().Bool("quick", false, "concurrent creation users")
	userCreateCmd.Flags().Int("cnt", 1, "user total")
	userCreateCmd.Flags().String("token", "", "admin token")
	userCreateCmd.Flags().String("host", "", "apiserver address[eg: 192.168.1.1:8000]")

	userCreateCmd.MarkFlagRequired("token")
	userCreateCmd.MarkFlagRequired("host")
}
