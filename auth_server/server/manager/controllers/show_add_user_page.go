package controllers

import (
	"github.com/astaxie/beego"
)

type ShowAddUserPageController struct {
	beego.Controller
}

func (c *ShowAddUserPageController) Get() {
	// 选择list模板
	c.TplNames = "add_user.tpl"
}
