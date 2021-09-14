package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"quickstart/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user",&controllers.UserController{})
}
