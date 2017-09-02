package routers

import (
	"github.com/astaxie/beego"
	"lab012/quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/test", &controllers.TestController{})
}
