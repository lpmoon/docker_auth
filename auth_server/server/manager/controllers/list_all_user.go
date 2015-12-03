package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/cesanta/docker_auth/auth_server/server"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
)

type ListAllUserController struct {
	beego.Controller
}

func (c *ListAllUserController) Get() {
	// 选择list模板
	c.TplNames = "list_all_user.tpl"
	c.Data["names"] = models.ACManager.QueryAllUser()
}
