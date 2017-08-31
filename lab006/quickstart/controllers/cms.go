package controllers

import "github.com/astaxie/beego"

type CMSController struct {
	beego.Controller
}

// @router /staticblock/:key [get]
func (c *CMSController) StaticBlock() {
	c.Ctx.Output.Body([]byte("StaticBlock()"))
}

// @router /all/:key [get]
func (c *CMSController) AllBlock() {
	c.Ctx.Output.Body([]byte("AllBlock"))
}
