package controllers

import "github.com/astaxie/beego"

type AutoController struct {
	beego.Controller
}

func (c *AutoController) Test1() {
	c.Ctx.Output.Body([]byte("autoController:Test1"))
}

func (c *AutoController) Test2() {
	c.Ctx.Output.Body([]byte("autoController:Test2"))
}
