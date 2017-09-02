package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

//func (c *TestController) Get() {
//	c.Data["name"] = "test1"
//	c.TplName = "test1.tpl"
//}

func (c *TestController) Post() {
	c.Data["name"] = "test2"
	c.TplName = "test2.tpl"
}
