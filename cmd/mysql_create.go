package cmd

import (
	"github.com/gaozhenhai/api-test/mysql"
	"github.com/spf13/cobra"
)

var mysqlCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create many test data",
	Run:   mysql.CreateTestData,
}

func init() {
	mysqlCmd.AddCommand(mysqlCreateCmd)
	mysqlCreateCmd.Flags().String("dsn", "", "mysql address [e.g: username:password@tcp(ip:port)/dbname]")
	mysqlCreateCmd.Flags().Int("connection", 1, "mysql connection")
	mysqlCreateCmd.Flags().Int("total", 0, "data total installed")

	mysqlCreateCmd.MarkFlagRequired("dsn")
	mysqlCreateCmd.MarkFlagRequired("total")
}
