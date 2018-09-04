package jobs

import (
	//"oss_dfs/models"
	//"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/toolbox"
	"fmt"
	//"oss_dfs/models"
	//"os"
)

func registerTask() []map[string]interface{} {
	taskList := []map[string]interface{}{
		{"tname": "SyncAnalysisMonth", "spec": "0/6 * * * * *", "method": func() {
			//models.SyncAnalysisMonth()
			fmt.Println("month")
		}},
		{"tname": "SyncAnalysisAll", "spec": "0/3 * * * * *", "method": func() {
			//models.SyncAnalysisAll()
			fmt.Println("all")
		}},
	}
	return taskList
}

func InitJobs() {
	//log := logs.NewLogger(10000)
	//log.SetLogger("file", `{"filename":"./logs/task.log"}`)
	//res := models.SyncAnalysisWeek()
	//fmt.Println(res)
	//os.Exit(3)
	//taskList := registerTask()

	//for _, value := range taskList {
	//	toolbox.AddTask(value["tname"].(string),
	//		toolbox.NewTask(value["tname"].(string), value["spec"].(string), func(method interface{}) func() error {
	//			return func() error {
	//				method.(func())()
	//				return nil
	//			}
	//		}(value["method"])))
	//}

	//toolbox.StartTask()
}
