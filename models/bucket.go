package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Bucket))
}

type Bucket struct {
	Id           int64
	Name         string
	Desc         string
	Secret       int
	Access       int
	IsReferer    int
	RefererValue string
	Token        string
	Status       int
	CreatedAt    string
	UpdatedAt    string
}

func (bucket *Bucket) GetLists(page, pageSize int) (int64, []Bucket) {
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(TableName("bucket"))
	query = query.Filter("status", 1)
	count, _ := query.Count()

	var buckets []Bucket
	query.Offset(offset).Limit(pageSize).OrderBy("id", "desc").All(&buckets)

	return count, buckets
}

func (bucket *Bucket) GetById(id int64) {
	orm.NewOrm().QueryTable(TableName("bucket")).Filter("id", id).Filter("status", 1).One(bucket)
	return
}


func GetByBucket(name string) (*Bucket, error) {
	b := new(Bucket)
	err := orm.NewOrm().QueryTable(TableName("bucket")).Filter("name", name).Filter("status", 1).One(b)
	if err != nil {
		return nil, err
	}
	return b, err
}

func (bucket *Bucket) Add () (int64, error) {
	return orm.NewOrm().Insert(bucket)
}

func (bucket *Bucket) Edit(fields ...string) (int64, error) {
	return orm.NewOrm().Update(bucket, fields...)
}
