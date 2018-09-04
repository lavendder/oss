package controllers
import (
	"time"
	."oss_dfs/library/log"
	"oss_dfs/models"
)
type HomeController struct {
	BaseController
}
func (c *HomeController) Index() {
	time.Sleep(time.Millisecond * 500)
	log1  := make(map[string]string)
	log1["name"] = "felix"
	log2 := JsonResponse{Code: 12, Message: "14"}
	OssLogger.Info("example of oss project log using ", log2, log1, "testtttt", "nihao")
	c.Data["PageTitle"] = "对象存储系统"
	c.TplName = "index.tpl"
}

func (c *HomeController) GetSummaryData() {
	/*type data struct {
		summary  []models.Analysis
		//list_data map[string][]models.Analysis
	}*/

	//d := data{summary:models.GetSummaryData(), list_data:models.GetLatestData()}
	//fmt.Println("xxxxxxxxxx")
	//fmt.Println(d)
	list := make(map[string]interface{})
	list["summary"] = models.GetSummaryData()
	list["list"]    = models.GetLatestData()

	c.SuccessData(list)
}