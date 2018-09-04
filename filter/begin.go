package filter

import (
	"github.com/astaxie/beego/context"
	"oss_dfs/library/constvar"
	. "oss_dfs/library/server"
	"net/url"
)

// set request id to SERVER param
var FilterBegin = func(ctx *context.Context) {
	url.ParseQuery(ctx.Request.URL.RawQuery)

	requestId := ctx.Request.Header.Get(constvar.XRequestId)
	SetRequestId(requestId)
	SetRequestTime()
}