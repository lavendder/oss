package controllers

import (
	"oss_dfs/models"
	"strings"
	"oss_dfs/library/constvar"
	"github.com/astaxie/beego"
	"time"
)

var (
	bucket *models.Bucket
)

type BucketController struct {
	BaseController
}

func init() {
	bucket = &models.Bucket{}
}

// 列表
func (c *BucketController) List() {
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 20)
	total, list := bucket.GetLists(page, pageSize)

	data := make(map[string]interface{})
	data["total"] = total
	data["list"] = list

	c.SuccessData(data)
}

// 编辑
func (c *BucketController) Edit() {
	id, _ := c.GetInt64("id")
	bucket.GetById(id)
	if bucket.Id == 0 {
		c.ErrorData(bucket, constvar.CodeError, "")
	}
	bucket.Name = strings.TrimSpace(c.GetString("Name"))
	bucket.Desc = strings.TrimSpace(c.GetString("Desc"))
	bucket.Secret, _ = c.GetInt("Secret")
	bucket.Access, _ = c.GetInt("Access")
	bucket.IsReferer, _ = c.GetInt("IsReferer")
	bucket.RefererValue = strings.TrimSpace(c.GetString("RefererValue"))
	bucket.Token = strings.TrimSpace(c.GetString("Token"))
	bucket.UpdatedAt = beego.DateFormat(time.Now(), constvar.FormatDateTime)

	if _, err := bucket.Edit(); err != nil {
		c.SuccessData(bucket)
	}
	c.ErrorData(bucket, constvar.CodeError, "")
}

// 添加
func (c *BucketController) Add() {
	bucket.Name = strings.TrimSpace(c.GetString("Name"))
	bucket.Desc = strings.TrimSpace(c.GetString("Desc"))
	bucket.Secret, _ = c.GetInt("Secret")
	bucket.Access, _ = c.GetInt("Access")
	bucket.IsReferer, _ = c.GetInt("IsReferer")
	bucket.RefererValue = strings.TrimSpace(c.GetString("RefererValue"))
	bucket.Token = strings.TrimSpace(c.GetString("Token"))
	bucket.Status = constvar.STATUS_VALID
	bucket.CreatedAt = beego.DateFormat(time.Now(), constvar.FormatDateTime)
	bucket.UpdatedAt = beego.DateFormat(time.Now(), constvar.FormatDateTime)
	id, err := bucket.Add()
	if err != nil {
		c.ErrorData(bucket, constvar.CodeError, "")
	}
	bucket.Id = id
	c.SuccessData(bucket)
}

// 详情
func (c *BucketController) Get() {
	id, _ := c.GetInt("id", 0)
	bucket.GetById(int64(id))
	data := make(map[string]interface{})
	data["data"] = bucket
	c.SuccessData(data)
}
