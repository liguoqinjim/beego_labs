package routers

import (
	"github.com/astaxie/beego"
	"lab007/quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/info", &controllers.InfoController{})
}
