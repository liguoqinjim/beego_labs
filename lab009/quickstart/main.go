package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "lab009/quickstart/routers"
)

func main() {
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	//过滤器
	var FilterLogin = func(ctx *context.Context) {
		//判断session中是否有li这个值
		v := ctx.Input.Session("li")
		if v == nil {
			ctx.Redirect(302, "/")
		} else {
			if _, ok := v.(int); !ok {
				ctx.Redirect(302, "/")
			}
		}
	}

	//插入过滤器
	beego.InsertFilter("/user", beego.BeforeRouter, FilterLogin)

	beego.Run()
}
