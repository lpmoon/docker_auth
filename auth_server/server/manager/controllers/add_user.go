package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
	"github.com/golang/glog"
)

type AddUserController struct {
	beego.Controller
}

func (c *AddUserController) DoAddUser() {
	// 参数获取
	username := c.GetString("username")
	glog.Infof("username is %s \n", username)
	password := c.GetString("password")
	glog.Infof("password is %s \n", password)

	_, success := models.ACManager.AddUser(username, password)

	ret := map[string]interface{}{"success": success}
	c.Data["json"] = ret
	c.ServeJson() // 直接返回json数据

	glog.Flush()
}
