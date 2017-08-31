package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	v := c.GetSession("li")
	if v == nil {
		c.SetSession("li", int(1))
		c.Data["num"] = 0
	} else {
		c.SetSession("li", v.(int)+1)
		c.Data["num"] = v.(int)
	}

	c.TplName = "user.tpl"
}
