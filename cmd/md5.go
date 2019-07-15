package cmd

import (
	"github.com/gaozhenhai/api-test/md5"
	"github.com/spf13/cobra"
)

var md5Cmd = &cobra.Command{
	Use:   "md5 <command>",
	Short: "clean md5 feature",
	Run:   md5.Md5,
}

func init() {
	rootCmd.AddCommand(md5Cmd)

	md5Cmd.Flags().String("string", "", "input a string")
	md5Cmd.MarkFlagRequired("string")
}
