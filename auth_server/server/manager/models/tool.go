//package models
package models

import (
	"bytes"
	"github.com/golang/glog"
	"os/exec"
	"strings"
)

func EncrptByHtpasswd(username string, password string) string {
	// call htpasswd to encrpt password
	glog.Infof("password to be encrpt is %s", password)
	cmd := exec.Command("htpasswd", "-nbB", username, password)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		glog.Errorln(error.Error)
		return ""
	}

	tmp := out.String()
	return tmp[(strings.Index(tmp, ":") + 1):(len(tmp) - 2)]
}
