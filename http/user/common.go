package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

func verify(cmd *cobra.Command) (token, host string, cnt int, err error) {
	token, _ = cmd.Flags().GetString("token")
	if token == "" {
		err = fmt.Errorf("admin token is nil")
		return
	}

	host, _ = cmd.Flags().GetString("host")
	if host == "" {
		err = fmt.Errorf("host address is nil")
		return
	}

	cnt, _ = cmd.Flags().GetInt("cnt")
	return
}
