package routers

import (
	"github.com/astaxie/beego"
	"lab004/quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//添加路由
	beego.Router("/user", &controllers.UserController{})
}
