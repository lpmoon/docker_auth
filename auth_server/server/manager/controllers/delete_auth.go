package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
	"github.com/golang/glog"
)

type DeleteAuthController struct {
	beego.Controller
}

func (c *DeleteAuthController) DoDeleteAuth() {
	// 参数获取
	username := c.GetString("user")
	glog.Infof("username is %s \n", username)
	imagename := c.GetString("name")
	glog.Infof("imagename is %s \n", imagename)

	_, success := models.ACManager.Delete(username, imagename)

	ret := map[string]interface{}{"success": success}
	c.Data["json"] = ret
	c.ServeJson() // 直接返回json数据

	glog.Flush()
}
