package main

import (
	"oss_dfs/models"
	_ "oss_dfs/routers"
	_ "oss_dfs/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"oss_dfs/jobs"
)

func init() {
	models.Init()
	jobs.InitJobs()
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	//beego.ErrorController(&controllers.ErrorController{})
	beego.SetStaticPath("/index.html", "index.html")
	beego.Run()
}
