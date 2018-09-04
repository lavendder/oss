package models

import "github.com/astaxie/beego/orm"

type Object struct {
	ObjectListId int64
	ObjectList
	ObjectSourceId int64
	ObjectSource
}

/**
获取文件列表
 */
func (o *Object) GetList(bucketId, page, pageSize int) []Object {
	// 获取上传列表文件
	objectList := new(ObjectList)
	objectListData := objectList.GetList(bucketId, page, pageSize)

	// 获取文件信息和key的关系
	objects := make(map[string]Object)
	var keys []string
	var object Object
	for _, info := range objectListData {
		object.ObjectListId = info.Id
		object.Name = info.Name
		object.Path = info.Path
		object.Access = info.Access
		object.BucketId = info.BucketId
		object.ObjectSourceKey = info.ObjectSourceKey
		object.From = info.From
		object.ClientIp = info.ClientIp
		object.CreatedAt = info.CreatedAt
		object.UpdatedAt = info.UpdatedAt

		objects[info.ObjectSourceKey] = object
		keys = append(keys, info.ObjectSourceKey)
	}

	// 获取文件存储原文件
	objectSource := new(ObjectSource)
	objectSourceList := objectSource.GetListByKeys(keys)

	var list []Object
	for _, info := range objectSourceList {
		object = objects[info.Key]
		object.ObjectSourceId = info.Id
		object.Key = info.Key
		object.HostId = info.HostId
		object.MiniType = info.MiniType
		object.Property = info.Property
		object.Secret = info.Secret
		object.Size = info.Size
		object.Ext = info.Ext
		list = append(list, object)
	}

	return list
}

/**
获取总数
 */
func (o *Object) GetTotal(bucketId int) int {
	objectList := new(ObjectList)
	return objectList.GetTotal(bucketId)
}

/**
获取文件详情
 */
func (o * Object) GetDetail(name string) Object {
	objectListObj := new(ObjectList)
	objectList := objectListObj.GetOneByName(name)

	objectSourceObj := new(ObjectSource)
	objectSource := objectSourceObj.GetOneByKey(objectList.ObjectSourceKey)

	// 组装数据
	var object Object
	object.ObjectListId = objectList.Id
	object.Name = objectList.Name
	object.Path = objectList.Path
	object.Access = objectList.Access
	object.BucketId = objectList.BucketId
	object.ObjectSourceKey = objectList.ObjectSourceKey
	object.From = objectList.From
	object.ClientIp = objectList.ClientIp
	object.CreatedAt = objectList.CreatedAt
	object.UpdatedAt = objectList.UpdatedAt

	object.ObjectSourceId = objectSource.Id
	object.Key = objectSource.Key
	object.HostId = objectSource.HostId
	object.MiniType = objectSource.MiniType
	object.Property = objectSource.Property
	object.Secret = objectSource.Secret
	object.Size = objectSource.Size
	object.Ext = objectSource.Ext

	return object
}

/**
编辑
 */
func (o *Object) Edit(name string, clientIp string, miniType string) (Object, bool) {
	data := o.GetDetail(name)

	// 如果不存在
	if data.Name == "" {
		return data, false
	}

	data.ClientIp = clientIp
	data.MiniType = miniType
	objectList, objectSource := o.resolve(data)

	ormer := orm.NewOrm()
	ormer.Begin()
	// 修改上传文件列表
	if _, err := objectList.Edit(ormer); err != nil {
		ormer.Rollback()
		return data, false
	}

	// 修改文件存储原文件
	if _, err := objectSource.Edit(ormer); err != nil {
		ormer.Rollback()
		return data, false
	}

	ormer.Commit()
	return data, true
}

func (o *Object) resolve(data Object) (ObjectList, ObjectSource) {
	var objectList ObjectList
	var objectSource ObjectSource

	objectList.Id = data.ObjectListId
	objectList.Name = data.Name
	objectList.Path = data.Path
	objectList.Access = data.Access
	objectList.BucketId = data.BucketId
	objectList.ObjectSourceKey = data.ObjectSourceKey
	objectList.From = data.From
	objectList.ClientIp = data.ClientIp
	objectList.CreatedAt = data.CreatedAt
	objectList.UpdatedAt = data.UpdatedAt

	objectSource.Id = data.ObjectSourceId
	objectSource.Key = data.Key
	objectSource.HostId = data.HostId
	objectSource.MiniType = data.MiniType
	objectSource.Property = data.Property
	objectSource.Secret = data.Secret
	objectSource.Size = data.Size
	objectSource.Ext = data.Ext

	return objectList, objectSource
}