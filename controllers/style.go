package controllers

import (
	"github.com/astaxie/beego"
	"oss_dfs/models"
)

type StyleController struct {
	beego.Controller
}

// 列表
func (c *StyleController) List() {
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 20)
	total, list := models.StyleLists(page, pageSize)
	c.Data["total"] = total
	c.Data["list"] = list
	c.TplName = "style/list.tpl"
}

// 编辑
func (c *StyleController) Edit() {
	c.Data["title"] = "style title"
	c.Data["desc"] = "style desc"
	c.TplName = "style/edit.tpl"
}

// 添加
func (c *StyleController) Add() {
	c.Data["title"] = "style title"
	c.Data["desc"] = "style desc"
	c.TplName = "style/add.tpl"
}
