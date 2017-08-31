package main

import (
	"github.com/astaxie/beego"
	_ "lab008/quickstart/routers"
)

func main() {
	//开启session，这行一定要在beego.Run()前面
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
