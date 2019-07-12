package cmd

import (
	"github.com/gaozhenhai/api-test/http/user"
	"github.com/spf13/cobra"
)

var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete tenxcloud users",
	Run:   user.DeleteTenxcloudUsers,
}

func init() {
	userCmd.AddCommand(userDeleteCmd)
	userDeleteCmd.Flags().Int("cnt", 1, "user total")
	userDeleteCmd.Flags().String("token", "", "admin token")
	userDeleteCmd.Flags().String("host", "", "apiserver address[eg: 192.168.1.1:8000]")

	userDeleteCmd.MarkFlagRequired("token")
	userDeleteCmd.MarkFlagRequired("host")
}
