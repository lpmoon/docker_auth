package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
)

type LoginController struct {
	beego.Controller
}

// do login
func (lc *LoginController) DoLogin() {
	username := lc.GetString("username")
	password := lc.GetString("password")
	sessionid := lc.Ctx.GetCookie("gsessionid")

	if sessionid == lc.GetSession(username) {
		// success
		lc.Data["json"] = map[string]interface{}{"success": true}
		lc.ServeJson() // 直接返回json数据
		return
	}

	_, success := models.ACManager.DoLogin(username, password)
	if success {
		// if success add session
		lc.Data["json"] = map[string]interface{}{"success": true}
		lc.ServeJson() // 直接返回json数据
	}
}
