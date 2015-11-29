package routers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/controllers"
)

func init() {
	beego.Router("/", &controllers.ListController{})
	beego.Router("/query", &controllers.DetailController{})
	beego.Router("/modify", &controllers.ModifyController{})
}
