package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
	"github.com/golang/glog"
)

type AddAuthController struct {
	beego.Controller
}

func (c *AddAuthController) DoAddAuth() {
	// 参数获取
	username := c.GetString("username")
	glog.Infof("username is %s \n", username)
	imagename := c.GetString("imagename")
	glog.Infof("imagename is %s \n", imagename)
	actions := c.GetStrings("actions[]")
	glog.Infof("actions is %s \n", actions)

	_, success := models.ACManager.Add(username, imagename, actions)

	ret := map[string]interface{}{"success": success}
	c.Data["json"] = ret
	c.ServeJson() // 直接返回json数据

	glog.Flush()
}
