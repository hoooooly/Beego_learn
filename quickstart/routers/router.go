package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"quickstart/controllers"
)

func init() {
	// 执行路由映射
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user",&controllers.UserController{})
}
