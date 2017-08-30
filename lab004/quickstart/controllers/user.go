package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["title1"] = "hello"
	c.Data["title2"] = "world"
	c.TplName = "test.tpl"
}
