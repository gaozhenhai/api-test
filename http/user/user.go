package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func createOneUser(token, host string, cnt, n int, waitGroup *sync.WaitGroup, quick bool) {
	if quick {
		defer waitGroup.Done()
	}

	request := fmt.Sprintf("{\n    \"userName\":\"t%04d\",\n    \"password\":\"tenxcloud001\",\n    \"email\":\"%04d@qq.com\",\n    \"phone\":\"1361234%04d\",\n    \"role\":4,\n    \"authority\":[\n        \"RID-Ezeg3KPhm3mS\",\n        \"RID-XwPiLfrBYjqd\"\n    ]\n}", n, n, n)

	url := fmt.Sprintf("http://%s/api/v2/users", host)
	req, err := http.NewRequest("POST", url, strings.NewReader(request))
	if err != nil {
		glog.Errorln(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	req.Header.Add("username", "admin")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Errorf("error create user %04d %v", n, err)
	}
	defer res.Body.Close()

	glog.Infof("==> create user t%04d is done status %d", n, res.StatusCode)
}

func CreateTenxcloudUsers(cmd *cobra.Command, args []string) {
	token, host, cnt, err := verify(cmd)
	if err != nil {
		cmd.Usage()
		glog.Errorln(err)
		return
	}

	glog.Infof("host: %v token: %v cnt: %d", host, token, cnt)

	waitGroup := &sync.WaitGroup{}
	quick, _ := cmd.Flags().GetBool("quick")

	for n := 1; n < cnt+1; n++ {
		if quick {
			waitGroup.Add(1)
			go createOneUser(token, host, cnt, n, waitGroup, quick)
			glog.Infof("==> create user t%04d is running...", n)
		} else {
			createOneUser(token, host, cnt, n, waitGroup, quick)
		}
	}

	if quick {
		waitGroup.Wait()
	}
}

func DeleteTenxcloudUsers(cmd *cobra.Command, args []string) {
	token, host, cnt, err := verify(cmd)
	if err != nil {
		cmd.Usage()
		glog.Errorln(err)
		return
	}

	glog.Infof("host: %v token: %v cnt: %d", host, token, cnt)

	detail := struct {
		Users []struct {
			UserID      int    `json:"userID"`
			UserName    string `json:"userName"`
			Namespace   string `json:"namespace"`
			DisplayName string `json:"displayName"`
		} `json:"users"`
	}{}

	for n := 1; n < cnt+1; n++ {
		detailUrl := fmt.Sprintf("http://%s/api/v2/users?filter=userName,t%04d", host, n)
		req, err := http.NewRequest("GET", detailUrl, nil)
		if err != nil {
			glog.Errorln(err)
			continue
		}

		req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
		req.Header.Add("username", "admin")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			glog.Errorf("error get user detail %04d %v", n, err)
			continue
		}
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(body, &detail); err != nil {
			glog.Errorln(err)
			continue
		}

		if len(detail.Users) > 0 && detail.Users[0].UserName == detail.Users[0].Namespace && detail.Users[0].UserName == detail.Users[0].DisplayName {
			deletelUrl := fmt.Sprintf("http://%s/api/v2/users/%d", host, detail.Users[0].UserID)
			req, err := http.NewRequest("DELETE", deletelUrl, nil)
			if err != nil {
				glog.Errorln(err)
				continue
			}

			req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
			req.Header.Add("username", "admin")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				glog.Errorf("error delete user %04d %v", n, err)
				continue
			}
			defer res.Body.Close()
		}
		glog.Infof("==> delete user t%04d is done status %d", n, res.StatusCode)
	}
}
