package routers

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/server/manager/controllers"
)

func init() {
	beego.Router("/", &controllers.ListAllUserController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/query", &controllers.ShowUserAuthController{})
	beego.Router("/modify", &controllers.ModifyUserAuthController{}, "get,post:DoModify")
	beego.Router("/showadduser", &controllers.ShowAddUserPageController{})
	beego.Router("/adduser", &controllers.AddUserController{}, "get,post:DoAddUser")
	beego.Router("/deleteuser", &controllers.DeleteUserController{}, "get,post:DoDeleteUser")
	beego.Router("/showaddauth", &controllers.ShowAddAuthPageController{})
	beego.Router("/addauth", &controllers.AddAuthController{}, "get,post:DoAddAuth")
	beego.Router("/deleteauth", &controllers.DeleteAuthController{}, "get,post:DoDeleteAuth")
	beego.Router("/showallauth", &controllers.ShowAllAuthController{})
}
