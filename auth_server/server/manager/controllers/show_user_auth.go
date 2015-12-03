package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
)

type ShowUserAuthController struct {
	beego.Controller
}

func (c *ShowUserAuthController) Get() {
	// c.TplNames = "index.tpl"

	user := c.GetString("user")
	detail := models.ACManager.QueryDetail(user)
	c.Data["user"] = user
	c.Data["detail"] = detail
	c.Data["idx"] = 1
	c.TplNames = "show_user_auth.tpl"
}
