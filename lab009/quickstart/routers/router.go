package routers

import (
	"github.com/astaxie/beego"
	"lab009/quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/user", &controllers.UserController{})
}
