package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/cesanta/docker_auth/auth_server/server"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	// 选择list模板
	c.TplNames = "list.tpl"
	c.Data["names"] = models.ACManager.QueryAllUser()
}
