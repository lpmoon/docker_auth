package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
)

type ShowAllAuthController struct {
	beego.Controller
}

func (c *ShowAllAuthController) Get() {
	// c.TplNames = "index.tpl"

	detail := models.ACManager.QueryAllAuth()
	c.Data["detail"] = detail
	c.TplNames = "show_all_auth.tpl"
}
