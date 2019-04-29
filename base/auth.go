package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type CallBackResult map[string]interface{}
type CallBack func(*context.Context, CallBackResult)

type AuthStatus int

const (
	AuthErr          AuthStatus = iota
	AuthNoLogin
	AuthNoUser
	AuthNoPerm
	AuthHasReadPerm
	AuthHasWritePerm
	AuthSuccess
)

func auth(ctx *context.Context, args ...interface{}) CallBackResult {
	name := ctx.Request.URL.Query().Get("uname")
	var ret = make(CallBackResult)
	if name == "" {
		name = ctx.Input.Cookie("uname")
		if name == "" {
			ret[Status.String()] = AuthNoLogin
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
