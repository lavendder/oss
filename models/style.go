package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Style))
}

type Style struct {
	Id            int64
	Name          string
	Value          string
	Desc          string
	Status        int
	CreatedAt     time.Time
}

func StyleLists(page, pageSize int) (int64, []Style) {
	orm.Debug = true
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(TableName("style"))
	query = query.Filter("status", 1)
	count, _ := query.Count()

	var styles []Style
	query.Offset(offset).Limit(pageSize).OrderBy("id", "desc").All(&styles)

	return count, styles
}

func StyleAdd(style *Style) (int64, error) {
	return orm.NewOrm().Insert(style)
}

func StyleUpdate(style *Style, fields ...string) error {
	_, err := orm.NewOrm().Update(style, fields...)
	return err
}