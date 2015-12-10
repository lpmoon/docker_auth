package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
)

type ShowUserAuthJsonController struct {
	beego.Controller
}

func (c *ShowUserAuthJsonController) Get() {
	// c.TplNames = "index.tpl"

	user := c.GetString("user")
	detail := models.ACManager.QueryDetail(user)

	data := map[string]interface{}{"user": user, "detail": detail}

	c.Data["json"] = data
	c.ServeJson()
}
