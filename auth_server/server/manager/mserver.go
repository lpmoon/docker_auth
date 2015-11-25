package manager

import (
	"github.com/astaxie/beego"
	"github.com/cesanta/docker_auth/auth_server/authn"
	"github.com/cesanta/docker_auth/auth_server/authz"
	"github.com/cesanta/docker_auth/auth_server/server/config"
	"github.com/cesanta/docker_auth/auth_server/server/manager/models"
	_ "github.com/cesanta/docker_auth/auth_server/server/manager/routers"
)

type MServer struct {
}

// 初始化配置，以及新建服务器
func NewMServer(config *config.Config, authenticators []authn.Authenticator, authorizers []authz.Authorizer) *MServer {
	models.InitAuthConfigManager(config, authenticators, authorizers)
	return &MServer{}
}

// 代理函数，实际调用beego的运行方法
func (ms *MServer) RunManagerServer() {
	beego.SetStaticPath("/img", "img")
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/js", "js")
	beego.Run()
}
