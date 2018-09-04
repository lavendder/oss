package controllers

import (
	"github.com/astaxie/beego"
	"oss_dfs/library/constvar"
	"reflect"
)

type BaseController struct {
	beego.Controller
}

type JsonResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

// 成功返回的数据
func (base *BaseController) SuccessData(data interface{}) {
	base.Data["json"] = JsonResponse{constvar.CodeSuccess, "success", data}
	base.ServeJSON()
}

// 失败返回的数据
func (base *BaseController) ErrorData(data interface{}, code interface{}, message interface{} ) {
	coder := constvar.CodeError
	if reflect.TypeOf(code).String() == "int" {
		coder = code.(int)
	}
	messager := "error"
	if reflect.TypeOf(message).String() == "string" && message.(string) != "" {
		messager = message.(string)
	}
	base.Data["json"] = JsonResponse{coder, messager, data}
	base.ServeJSON()
}