package routers

import (
	"github.com/astaxie/beego"
	"lab008/quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
