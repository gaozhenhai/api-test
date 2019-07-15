package md5

import (
	"crypto/md5"
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func encryption(txt string) string {
	hexDigits := []byte{'!', '/', '|', '~', '`', '^', '\'', '*', '?', '/', '-', '_', '=', '+', '.', ','}

	mdInst := md5.New()
	if _, err := mdInst.Write([]byte(txt)); err != nil {
		glog.Errorln(err)
	}
	md := mdInst.Sum(nil)

	pos := 0
	mdLength := len(md)
	str := make([]byte, mdLength*2)

	for i := 0; i < mdLength; i++ {
		byte0 := md[i]
		str[pos] = hexDigits[byte0>>4&0xf]
		pos++
		str[pos] = hexDigits[byte0&0xf]
		pos++
	}

	return string(str)
}

func Md5(cmd *cobra.Command, args []string) {
	txt, _ := cmd.Flags().GetString("string")
	for i := 0; i < 32; i++ {
		txt = encryption(txt)
	}
	fmt.Println(txt)
}
