package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

//正则路由
type Reg1Controller struct {
	beego.Controller
}

func (c *Reg1Controller) Get() {
	id := c.Ctx.Input.Param(":id")
	output := fmt.Sprintf("id=%s", id)
	c.Ctx.Output.Body([]byte(output))
}

type Reg2Controller struct {
	beego.Controller
}

func (c *Reg2Controller) Get() {
	id := c.Ctx.Input.Param(":id")
	output := fmt.Sprintf("id=%s", id)
	c.Ctx.Output.Body([]byte(output))
}
