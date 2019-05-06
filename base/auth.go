package base

import (
	"GolangBase/define"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type CallBackResult map[string]interface{}
type CallBack func(*context.Context, CallBackResult)

func auth(ctx *context.Context, args ...interface{}) CallBackResult {
	name := ctx.Request.URL.Query().Get("uname")
	var ret = make(CallBackResult)
	if name == "" {
		name = ctx.Input.Cookie("uname")
		if name == "" {
			ret[define.Status.String()] = define.AuthNoLogin
		}
	}
	return ret
}

func Filter(cb CallBack, args ...interface{}) beego.FilterFunc {
	filter := func(ctx *context.Context) {
		authRes := auth(ctx, args)
		cb(ctx, authRes)
	}
	return filter
}
