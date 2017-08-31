package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"lab006/quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//基本路由
	beego.Get("/baseGet", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})

	//固定路由
	beego.Router("/user", &controllers.UserController{})

	//正则表达式
	beego.Router("/reg1/?:id", &controllers.Reg1Controller{})
	beego.Router("/reg2/:id", &controllers.Reg2Controller{})

	//自定义方法及 RESTful 规则
	beego.Router("/rest/list", &controllers.RestController{}, "*:ListFood")

	//自动匹配
	beego.AutoRouter(&controllers.AutoController{})

	//注解路由
	beego.Include(&controllers.CMSController{})
}
