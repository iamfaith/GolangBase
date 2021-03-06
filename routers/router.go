package routers

import (
	"GolangBase/base"
	"GolangBase/define"
	"GolangBase/service/pprof"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"io/ioutil"
)

func init() {

	beego.Head("/heart", func(ctx *context.Context) {
		ctx.Output.Body([]byte("still alive"))
	})

	beego.Get("/alive", func(ctx *context.Context) {
		ctx.Output.Body([]byte("ok"))
	})

	beego.Get("/build", func(ctx *context.Context) {
		files, _ := ioutil.ReadDir("/data/build_time")
		names := ""
		for _, f := range files {
			names += f.Name()
		}
		ctx.Output.Body([]byte(names))
	})

	fileRouter := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/file", &base.BaseController{}, "POST:Upload"),
			beego.NSRouter("/?:method/?:id", &base.BaseController{}, "GET:GetByReflect"),
			beego.NSRouter("/?:method", &base.BaseController{}, "POST:PostByReflect"),
		),
	)
	beego.AddNamespace(fileRouter)

	cb := func(ctx *context.Context, ret base.CallBackResult) {
		switch ret[define.Status.String()] {
		case define.AuthNoLogin:
			beego.Debug("should redirect to login")
			ctx.Redirect(303, "http://www.faithio.cn")
			return
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:  []string{"*.faithio.cn"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.InsertFilter("/api/*", beego.BeforeRouter, base.Filter(cb))

	pprof.Monitor()
}
