package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql <command>",
	Short: "mysql test data",
}

func init() {
	rootCmd.AddCommand(mysqlCmd)
}
