package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type InfoController struct {
	beego.Controller
}

type info struct {
	Id    string `-`
	Name  string `form:"username"`
	Age   int    `form:"age"`
	Email string `form:"email"`
}

func (c *InfoController) Post() {
	info := info{}
	if err := c.ParseForm(&info); err != nil {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("parseForm error:%v", err)))
	} else {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("info:%+v", info)))
	}
}
