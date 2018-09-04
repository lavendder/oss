/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 00:18:02
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-16 17:26:48
***********************************************/

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

func Init() {
	dbhost := beego.AppConfig.String("database::db_host")
	dbport := beego.AppConfig.String("database::db_port")
	dbuser := beego.AppConfig.String("database::db_user")
	dbpassword := beego.AppConfig.String("database::db_password")
	dbname := beego.AppConfig.String("database::db_name")
	//timezone := beego.AppConfig.String("database::db_timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	//fmt.Println(dsn)

	//if timezone != "" {
	//	dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	//}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Analysis))
	//	new(Group), new(Env), new(Code), new(Api), new(ApiDetail), new(ApiParam), new(InfoList), new(InfoClass))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	// database will report error message like this: "invalid connection", in order to deal this problem
	// we have to ping database before query data
	db, err := orm.GetDB("default")
	if err == nil {
		db.Ping()
	}

	//return beego.AppConfig.String("database::db_prefix") + name
	return name
}
