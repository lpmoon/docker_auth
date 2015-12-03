package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
	"github.com/golang/glog"
)

type ModifyUserAuthController struct {
	beego.Controller
}

func (c *ModifyUserAuthController) DoModify() {
	// 参数获取
	glog.Infoln("--modify request start--")
	user := c.GetString("user")
	glog.Infof("accout is %s \n", user)
	name := c.GetString("name")
	glog.Infof("image name is %s \n", name)
	mtype, err := c.GetInt("mtype")
	if err != nil {
		ret := map[string]interface{}{"success": false, "msg": "mtype参数错误"}
		c.Data["json"] = ret
		glog.Errorln("mtype 错误!!")
		c.ServeJson()
		return
	}
	glog.Infof("mtype is %s", mtype)
	ispull, err := c.GetBool("ispull")
	if err != nil {
		ret := map[string]interface{}{"success": false, "msg": "ispull参数错误"}
		c.Data["json"] = ret
		glog.Errorln("ispull 错误!!")
		c.ServeJson()
		return
	}
	glog.Infof("ispull is %s \n", ispull)

	success := models.ACManager.Update(user, &name, mtype, ispull)

	ret := map[string]interface{}{"success": success}
	c.Data["json"] = ret
	c.ServeJson() // 直接返回json数据

	glog.Flush()
}
