//package models
package models

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func EncrptByHtpasswd(username string, password string) string {
	// call htpasswd to encrpt password
	cmd := exec.Command("htpasswd", "-nbB", username, password)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	tmp := out.String()
	return tmp[(strings.Index(tmp, ":") + 1):]
}
