package cmd

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "apitest",
	Short: "apitest is a tester for api-server",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		glog.Errorln(err)
		os.Exit(-1)
	}
}

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})
	flag.Set("logtostderr", "true")
}
