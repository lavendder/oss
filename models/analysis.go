package models

import (
	"github.com/astaxie/beego/orm"
	"oss_dfs/library/constvar"
	"oss_dfs/library/utils"
	"time"
	"fmt"
	"strconv"
	"strings"
)

var bucketName = [3]string{"all", "image", "txt"}

type Analysis struct {
	Id     int    `json:"id"`
	Bucket string `json:"bucket"`
	Count  int    `json:"count"`
	Size   int    `json:"size"`
	Type   string `json:"type"`
	Day    string    `json:"day"`
}

func tableName() string {
	return TableName("analysis")
}

func (a *Analysis) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Analysis) Insert(fields ...string) error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

func AnalysisAdd(analysis *Analysis) (int64, error) {
	return orm.NewOrm().Insert(analysis)
}

func AnalysisGetById(id int) (*Analysis, error) {
	a := new(Analysis)
	err := orm.NewOrm().QueryTable(tableName()).Filter("id", id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AnalysisGetByBucket(bucket string) (*Analysis, error) {
	a := new(Analysis)
	err := orm.NewOrm().QueryTable(tableName()).Filter("bucket", bucket).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AnalysisGetByBucketAndDay(bucket string, day int) (*Analysis, error) {
	a := new(Analysis)
	err := orm.NewOrm().QueryTable(tableName()).Filter("bucket", bucket).Filter("day", day).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AnalysisUpdate(analysis *Analysis, fields ...string) error {
	_, err := orm.NewOrm().Update(analysis, fields...)
	return err
}

// get oss summary data
func GetSummaryData1() []Analysis {
	var summaryList []Analysis
	_, err := orm.NewOrm().QueryTable(tableName()).Filter("bucket", "all").All(&summaryList)
	if err == nil {
		return summaryList
	}
	var nullAnalysis []Analysis
	return nullAnalysis
}

// get oss summary data
func GetSummaryData() []Analysis {
	var summaryList []Analysis

	// all
	analysis, err := getSummaryDataByType(constvar.TypeALL, constvar.TypeALL, "")
	if err == nil && analysis.Count > 0 {
		summaryList = append(summaryList, *analysis)
	}

	// week
	analysis, err = getSummaryDataByType(constvar.TypeALL, constvar.TypeWeek, utils.GetFistDayOfThisWeek())
	if err == nil && analysis.Count > 0 {
		summaryList = append(summaryList, *analysis)
	}

	// month
	analysis, err = getSummaryDataByType(constvar.TypeALL, constvar.TypeMonth, utils.GetFirstDayOfThisMonth())
	if err == nil && analysis.Count > 0 {
		summaryList = append(summaryList, *analysis)
	}

	// day
	analysis, err = getSummaryDataByType(constvar.TypeALL, constvar.TypeDay, time.Now().Format(constvar.FormatDate))
	if err == nil && analysis.Count > 0 {
		summaryList = append(summaryList, *analysis)
	}
	return summaryList
}

func getSummaryDataByType(bucket string, dataType string, day string) (*Analysis, error) {
	sum := new(Analysis)
	var err error
	if day == "" {
		err = orm.NewOrm().QueryTable(tableName()).
			Filter("bucket", bucket).
			Filter("type", dataType).
			One(sum)
	} else {
		err = orm.NewOrm().QueryTable(tableName()).
			Filter("bucket", bucket).
			Filter("type", dataType).
			Filter("day", day).One(sum)
	}

	if err != nil {
		return sum, err
	}
	return sum, nil
}

func SyncAnalysisDay() bool {
	date := time.Now().Format("2006-01-02")
	start := fmt.Sprintf("%s 00:00:00", date)
	end := fmt.Sprintf("%s 23:59:59", date)
	dateInt := strings.Replace(date, "-", "", -1)

	for _, bucket := range bucketName {
		count := 0
		size := 0
		var maps []orm.Params
		for i := 1; i <= 10; i++ {
			tableName := fmt.Sprintf("object_source%s", strconv.Itoa(i))
			data, err := orm.NewOrm().Raw(fmt.Sprintf("SELECT COUNT(`id`) AS sourceCount,ifnull(SUM(`size`),0) AS sizeSum FROM `%s` WHERE created_at >= '%s' AND created_at <= '%s'", tableName, start, end)).Values(&maps)
			if err == nil && data > 0 {
				sourceCount, _ := strconv.Atoi(maps[0]["sourceCount"].(string))
				sizeSum, _ := strconv.Atoi(maps[0]["sizeSum"].(string))
				count += sourceCount
				size += sizeSum
			}
		}
		analysis := new(Analysis)
		analysis.Type = "day"
		analysis.Count = count
		analysis.Bucket = bucket
		analysis.Day = dateInt
		analysis.Size = size

		orm.NewOrm().InsertOrUpdate(analysis)
	}

	return true
}

func SyncAnalysisWeek() bool {
	fmt.Println("week")
	for _, bucket := range bucketName {
		start := utils.GetFistDayOfThisWeek()
		end := time.Now().Format(constvar.FormatDate)
		var maps []orm.Params
		orm.NewOrm().Raw(fmt.Sprintf("SELECT ifnull(SUM(`count`),0) AS sourceCount,ifnull(SUM(`size`),0) AS sizeSum FROM `analysis` WHERE day >= '%s' AND day <= '%s' AND bucket='%s'", start, end, bucket)).Values(&maps)

		analysis := new(Analysis)
		analysis.Type = "week"
		analysis.Count = maps[0]["sourceCount"].(int)
		analysis.Bucket = bucket
		analysis.Day = start
		analysis.Size = maps[0]["sizeSum"].(int)

		orm.NewOrm().InsertOrUpdate(analysis)
	}
	return true

}

func SyncAnalysisAll() bool {
	var maps []orm.Params
	orm.NewOrm().Raw("SELECT ifnull(SUM(`count`),0) AS sourceCount,ifnull(SUM(`size`),0) AS sizeSum FROM `analysis` type='month' AND bucket='all'").Values(&maps)

	analysis := new(Analysis)
	analysis.Type = "all"
	analysis.Count = maps[0]["sourceCount"].(int)
	analysis.Bucket = "all"
	analysis.Size = maps[0]["sizeSum"].(int)

	orm.NewOrm().InsertOrUpdate(analysis)
	return true
}

func SyncAnalysisMonth() bool {
	fmt.Println("week")
	for _, bucket := range bucketName {
		start := utils.GetFirstDayOfThisMonth()
		end := time.Now().Format(constvar.FormatDate)
		var maps []orm.Params
		orm.NewOrm().Raw(fmt.Sprintf("SELECT ifnull(SUM(`count`),0) AS sourceCount,ifnull(SUM(`size`),0) AS sizeSum FROM `analysis` WHERE day >= '%s' AND day <= '%s' AND bucket='%s'", start, end, bucket)).Values(&maps)

		analysis := new(Analysis)
		analysis.Type = "month"
		analysis.Count = maps[0]["sourceCount"].(int)
		analysis.Bucket = bucket
		analysis.Day = start
		analysis.Size = maps[0]["sizeSum"].(int)

		orm.NewOrm().InsertOrUpdate(analysis)
	}
	return true
}

func GetLatestData() map[string][]Analysis {
	data := make(map[string][]Analysis)
	data[constvar.TypeMonth] = getLatestDataByNum(constvar.TypeALL, constvar.TypeMonth, 5)
	data[constvar.TypeWeek] = getLatestDataByNum(constvar.TypeALL, constvar.TypeWeek, 5)
	return data
}

// 取出最近的几条数据
func getLatestDataByNum(bucket, dataType string, num int) []Analysis{
	var summaryList []Analysis
	_, err := orm.NewOrm().QueryTable(tableName()).
		Filter("bucket", bucket).
		Filter("type", dataType).
		OrderBy("day").
		Limit(num).
		All(&summaryList)
	if err == nil {
		return summaryList
	}
	var nullAnalysis []Analysis
	return nullAnalysis
}
