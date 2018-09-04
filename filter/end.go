package filter

import (
	"github.com/astaxie/beego/context"
	. "oss_dfs/library/server"
	"time"
	."oss_dfs/library/log"
)

// set request id to SERVER param
var FilterEnd = func(ctx *context.Context) {
	beginTime := GetRequestTime()
	currentTime := time.Now()
	executeTime := currentTime.Sub(beginTime)
	uri := ctx.Request.RequestURI
	OssLogger.Info(" uri: " + uri + " - executeTime: " + executeTime.String())
}