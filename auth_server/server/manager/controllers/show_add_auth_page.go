package controllers

import (
	"github.com/astaxie/beego"
)

type ShowAddAuthPageController struct {
	beego.Controller
}

func (c *ShowAddAuthPageController) Get() {
	// 选择list模板
	c.TplNames = "add_auth.tpl"
}
