package controllers

import (
	"oss_dfs/models"
	"oss_dfs/library/constvar"
)

type ObjectController struct {
	BaseController
}

/**
列表
 */
func (c *ObjectController) List() {
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 20)
	bucketId, _ := c.GetInt("bucketId")

	object := new(models.Object)
	data := make(map[string]interface{})

	data["total"] = object.GetTotal(bucketId)
	data["list"] = object.GetList(bucketId, page, pageSize)

	c.SuccessData(data)
}

/**
详情
 */
func (c *ObjectController) Detail() {
	name := c.GetString("name")

	object := new(models.Object)
	data := object.GetDetail(name)

	c.SuccessData(data)
}

/**
编辑
 */
func (c * ObjectController) Edit() {
	name := c.GetString("name")

	clientIp := c.GetString("clientIp")
	miniType := c.GetString("miniType")

	object := new(models.Object)
	data, ok := object.Edit(name, clientIp, miniType)

	if ok {
		c.SuccessData(data)
	} else {
		c.ErrorData(data, constvar.CodeError, "")
	}
}