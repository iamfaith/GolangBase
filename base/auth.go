package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type CallBackResult map[string]interface{}
type CallBack func(*context.Context, CallBackResult)

type AuthStatus int

const (
	AUTH_ERR AuthStatus = iota
	AUTH_NO_LOGIN
	AUTH_NO_USER
	AUTH_NO_PERM
	AUTH_HAS_READ_PERM
	AUTH_HAS_WRITE_PERM
	AUTH_SUCCESS
)

type EnumType int

const (
	begin            EnumType = iota
	Status           EnumType = iota
	end              EnumType = iota
)

var enums = [...]string{"begin", "status", "end"}

func (a EnumType) String() string {
	if a <= begin || a >= end {
		return ""
	}
	return enums[a]
}

func ContainsEnum(modelType string) bool {
	for _, t := range enums {
		if t == begin.String() || t == end.String() {
			continue
		}
		if modelType == t {
			return true
		}
	}
	return false
}

var (
	defaultAllowHeaders = []string{"Origin", "Accept", "Content-Type", "Authorization"}
	// Regex patterns are generated from AllowOrigins. These are used and generated internally.
	allowOriginPatterns = []string{}
)

func auth(ctx *context.Context, args ...interface{}) CallBackResult {
	name := ctx.Request.URL.Query().Get("uname")
	var ret = make(CallBackResult)
	if name == "" {
		name = ctx.Input.Cookie("uname")
		if name == "" {
			ret[Status.String()] = AUTH_NO_LOGIN
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
