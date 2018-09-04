package routers

import (
	"oss_dfs/controllers"
	"github.com/astaxie/beego"
	"oss_dfs/filter"
)

func init() {

	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterBegin)
	beego.InsertFilter("/*", beego.FinishRouter, filter.FilterEnd, false)


	//系统介绍页面 & 数据统计总览
    beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/summary", &controllers.HomeController{}, "*:GetSummaryData")

	//文件上传 区分bucket
	beego.Router("/upload/:bucket", &controllers.UploadController{}, "post:Upload")

	//图片展示
	//beego.Router("/view/{bucket}/{file_name}/{style_name}.{jpg}", &controllers.ViewController{}, "*:View")


	//----------------------bucket管理------------------------------//
	//bucket列表
	beego.Router("/bucket/list", &controllers.BucketController{}, "*:List")
	//编辑bucket信息
	beego.Router("/bucket/edit/", &controllers.BucketController{}, "*:Edit")
	//添加bucket
	beego.Router("/bucket/add", &controllers.BucketController{}, "*:Add")
	//详情bucket
	beego.Router("/bucket/get", &controllers.BucketController{}, "*:Get")

	//----------------------文件管理------------------------------//
	//文件列表
	beego.Router("/object/list", &controllers.ObjectController{}, "*:List")
	//单个文件编辑
	beego.Router("/object/edit/", &controllers.ObjectController{}, "*:Edit")
	//单个文件信息查询
	beego.Router("/object/info", &controllers.ObjectController{}, "*:Detail")


	//----------------------样式规则管理------------------------------//
	//样式列表
	beego.Router("/style/list", &controllers.StyleController{}, "*:List")
	//样式编辑
	beego.Router("/style/edit/", &controllers.StyleController{}, "*:Edit")
	//样式添加
	beego.Router("/style/add", &controllers.StyleController{}, "*:Add")

}
