package main

import (
	"github.com/astaxie/beego"
	_ "lab011/quickstart01/routers"
	"log"
)

func main() {
	//读取mysql配置
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	database := beego.AppConfig.String("mysqldb")

	log.Println(user, pass, url, database)

	beego.Run()
}
