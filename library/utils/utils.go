package utils

import (
	"time"
	"github.com/teris-io/shortid"
	"github.com/astaxie/beego/context"
	"oss_dfs/library/constvar"
)

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}


func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetRequestId(ctx *context.Context) string {

	return ctx.Request.Header.Get(constvar.XRequestId)
}

// get first day of current week monday's date
func GetFistDayOfThisWeek() string {
	weekMap := weekMap()
	day := time.Now().Weekday().String()
	daysAgo := weekMap[day]
	return time.Now().AddDate(0, 0, -daysAgo).Format(constvar.FormatDate)
}

// get first day of current month's first day
func GetFirstDayOfThisMonth() string {
	_, _, day := time.Now().Date()
	day = day - 1
	return time.Now().AddDate(0, 0, -day).Format(constvar.FormatDate)
}

func GetFirstDayOfLastNMonths(n int) string {
	_, _, day := time.Now().Date()
	day = day - 1
	return  time.Now().AddDate(0, -n, -day).Format(constvar.FormatDate)
}

func GetFistDayOfLastNWeeks(n int) string {
	weekMap := weekMap()
	day := time.Now().Weekday().String()
	daysAgo := weekMap[day]
	return time.Now().AddDate(0, 0, -(daysAgo+n*7)).Format(constvar.FormatDate)
}

func weekMap() map[string]int {
	weekMap := make(map[string]int)
	weekMap["Monday"] = 0
	weekMap["Tuesday"] = 1
	weekMap["Wednesday"] = 2
	weekMap["Thursday"] = 3
	weekMap["Friday"] = 4
	weekMap["Saturday"] = 5
	weekMap["Sunday"] = 6
	return weekMap
}