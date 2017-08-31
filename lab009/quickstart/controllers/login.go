package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	v := c.GetSession("li")
	if v == nil {
		c.SetSession("li", int(1))
		c.Data["login"] = "登录完成"
	} else {
		c.Data["login"] = "已登录"
	}

	c.TplName = "login.tpl"
}
