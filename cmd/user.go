package cmd

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <command>",
	Short: "create tenxcloud users",
}

func init() {
	rootCmd.AddCommand(userCmd)
}
