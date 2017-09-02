package main

import (
	"github.com/astaxie/beego"
	_ "lab010/quickstart/routers"
)

func main() {
	beego.SetStaticPath("/download1", "down1")

	beego.Run()
}
