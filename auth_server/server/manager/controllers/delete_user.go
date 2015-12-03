package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
	"github.com/golang/glog"
)

type DeleteUserController struct {
	beego.Controller
}

func (c *DeleteUserController) DoDeleteUser() {
	// 参数获取
	username := c.GetString("username")
	glog.Infof("username is %s \n", username)

	_, success := models.ACManager.DeleteUser(username)

	ret := map[string]interface{}{"success": success}
	c.Data["json"] = ret
	c.ServeJson() // 直接返回json数据

	glog.Flush()
}
