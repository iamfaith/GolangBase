package main

import (
	_ "GolangBase/routers"
	"GolangBase/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func CopyBuildTime() {
	logs.Info("begin to copy build time file")
	if err := util.CopyDir("/build_time", "/data/build_time"); err != nil {
		logs.Error(err)
	} else {
		logs.Info("copy build time success")
	}
}

func main() {
	go CopyBuildTime()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
