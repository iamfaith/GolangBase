package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"GolangBase/base"
)

func init() {

	beego.Head("/heart", func(ctx *context.Context) {
		ctx.Output.Body([]byte("still alive"))
	})

	beego.Get("/alive", func(ctx *context.Context) {
		ctx.Output.Body([]byte("ok"))
	})

	fileRouter := beego.NewNamespace("/v1",
		beego.NSRouter("/file", &base.BaseController{}, "POST:Upload"),
	)
	beego.AddNamespace(fileRouter)

	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowOrigins:  []string{"*.wps.cn"},
	//	AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:  []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
	//	ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
	//}))
}
