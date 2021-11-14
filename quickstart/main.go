package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "quickstart/routers"
)

func main() {
	beego.SetStaticPath("/down1", "download1")
	//beego.SetStaticPath("/down2", "download2")
	beego.Run()
}

