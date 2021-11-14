package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller // 控制器
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *UserController) Get() {
	c.Data["Website"] = "UserController.web"
	c.Data["Email"] = "espholychan@outlook.com"
	c.TplName = "index.html"
}