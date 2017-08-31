package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	info := c.GetString("info")

	if info == "" {
		c.Ctx.Output.Body([]byte("info is empty"))
	} else {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("info:%s", info)))
	}
}
