package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	data := "page not found"
	c.Ctx.WriteString(data)
}

func (c *ErrorController) Error500() {
	data :=  "Internal server error"
	c.Ctx.WriteString(data)
}

func (c *ErrorController) Error501() {
	data := "server error"
	c.Ctx.WriteString(data)
}


func (c *ErrorController) ErrorDb() {
	data := "database is now down"
	c.Ctx.WriteString(data)
}