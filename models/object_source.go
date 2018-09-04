package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"strconv"
)

func init() {
	orm.RegisterModel(new(ObjectSource),new(ObjectSource1),new(ObjectSource2),new(ObjectSource3),new(ObjectSource4),
		new(ObjectSource5),new(ObjectSource6),new(ObjectSource7),new(ObjectSource8),new(ObjectSource9),new(ObjectSource10))
}

type ObjectSource struct {
	Id int64
	Key string
	HostId int
	MiniType string
	Property string
	Secret   int
	Size     int
	Ext      string
}

type ObjectSource1 struct {
	ObjectSource
}
type ObjectSource2 struct {
	ObjectSource
}
type ObjectSource3 struct {
	ObjectSource
}
type ObjectSource4 struct {
	ObjectSource
}
type ObjectSource5 struct {
	ObjectSource
}
type ObjectSource6 struct {
	ObjectSource
}
type ObjectSource7 struct {
	ObjectSource
}
type ObjectSource8 struct {
	ObjectSource
}
type ObjectSource9 struct {
	ObjectSource
}
type ObjectSource10 struct {
	ObjectSource
}

/**
根据获取文件存储原文件列表
 */
func (os ObjectSource) GetListByKeys(keyList []string) []ObjectSource {
	// 获取文件列表
	var objectSource []ObjectSource
	o := orm.NewOrm()

	keys := strings.Replace(strings.Trim(fmt.Sprint(keyList), "[]"), " ", "','", -1)
	appendWhere := fmt.Sprintf("where `key` in ('%s')", keys)

	sql := os.getSql(appendWhere)
	o.Raw(sql).QueryRows(&objectSource)

	return objectSource
}

/**
根据key获取一个列表
 */
func (os ObjectSource) GetOneByKey(key string) ObjectSource {
	var objectSource ObjectSource
	o := orm.NewOrm()

	appendWhere := fmt.Sprintf("where `key` = '%s'", key)

	sql := os.getSql(appendWhere)
	o.Raw(sql).QueryRow(&objectSource)

	return objectSource
}

/**
编辑
 */
func (os * ObjectSource) Edit(ormer orm.Ormer, fields ...string) (int64, error){
	talbeName := os.getTableName(os.Key)

	// 对象转map
	m := StructToMap(os)

	return ormer.QueryTable(TableName(talbeName)).Filter("key", os.Key).Update(m)
}

/**
获取查询sql
 */
func (os ObjectSource) getSql(appendWhere string) string {
	var sql string

	for i := 1; i <= 10; i++ {
		sql += fmt.Sprintf("(select `id`, `key`, `host_id`, `mini_type`, `property`, `secret`, "+
			"`size`, `ext`, `created_at` from object_source%s %s)", strconv.Itoa(i), appendWhere)

		if i < 10 {
			sql += " UNION "
		}
	}
	return sql
}

func (os *ObjectSource) getTableName(key string) string {
	return "object_source3"
}
