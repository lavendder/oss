package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
)

func init() {
	orm.RegisterModel(new(ObjectList),new(ObjectList1),new(ObjectList2),new(ObjectList3),new(ObjectList4),new(ObjectList5),
		new(ObjectList6),new(ObjectList7),new(ObjectList8),new(ObjectList9),new(ObjectList10))
	//orm.RegisterModel("object_list2")
}

type ObjectList struct {
	Id int64
	Name string
	Path string
	Access int
	BucketId int
	ObjectSourceKey string
	From string
	ClientIp string
	CreatedAt string
	UpdatedAt string
}

type ObjectList1 struct {
	ObjectList
}
type ObjectList2 struct {
	ObjectList
}
type ObjectList3 struct {
	ObjectList
}
type ObjectList4 struct {
	ObjectList
}
type ObjectList5 struct {
	ObjectList
}
type ObjectList6 struct {
	ObjectList
}
type ObjectList7 struct {
	ObjectList
}
type ObjectList8 struct {
	ObjectList
}
type ObjectList9 struct {
	ObjectList
}
type ObjectList10 struct {
	ObjectList
}

type ListCount struct {
	Total int
}

/**
获取上传文件列表
 */
func (ol *ObjectList) GetList(bucketId, page, pageSize int) []ObjectList {
	offset := (page - 1) * pageSize

	// 获取文件列表
	var objectList []ObjectList
	o := orm.NewOrm()

	appendWhere := "where `bucket_id` = " + strconv.Itoa(bucketId)
	sql := fmt.Sprintf("%s limit %s, %s", ol.getSql(appendWhere), strconv.Itoa(offset), strconv.Itoa(pageSize))
	o.Raw(sql).QueryRows(&objectList)

	return objectList
}

/**
查询总数
 */
func (ol *ObjectList) GetTotal(bucketId int) int {
	o := orm.NewOrm()
	var listCount ListCount

	appendWhere := "where `bucket_id` = " + strconv.Itoa(bucketId)
	sql := fmt.Sprintf("select count(1) as total from (%s) as listCount", ol.getSql(appendWhere))
	o.Raw(sql).QueryRow(&listCount)
	return listCount.Total
}

func (ol * ObjectList) GetOneByName(name string) ObjectList {
	appendWhere := fmt.Sprintf("where `name` = '%s'", name)

	var objectList ObjectList
	o := orm.NewOrm()
	o.Raw(ol.getSql(appendWhere)).QueryRow(&objectList)

	return objectList
}

/**
编辑
 */
func (ol * ObjectList) Edit(ormer orm.Ormer, fields ...string) (int64, error){
	talbeName := ol.getTableName(ol.ObjectSourceKey)

	// 对象转map
	m := StructToMap(ol)

	return ormer.QueryTable(TableName(talbeName)).Filter("name", ol.Name).Update(m)
}

/**
获取查询sql
 */
func (ol *ObjectList) getSql(appendWhere string) string {
	var sql string
	for i := 1; i <= 10; i++ {
		sql += fmt.Sprintf("(select `id`, `name`, `path`, `access`, `bucket_id`, `object_source_key`, " +
			"`from`, `client_ip`, `created_at`, `updated_at` from object_list%s %s)", strconv.Itoa(i), appendWhere)

		if i < 10 {
			sql += " UNION "
		}
	}
	return sql
}

func (ol *ObjectList) getTableName(key string) string {
	return "object_list2"
}