package controllers

import "github.com/astaxie/beego"

type RestController struct {
	beego.Controller
}

func (c *RestController) ListFood() {
	c.Ctx.Output.Body([]byte("apple;milk;noodle"))
}
