package main

import (
	_ "GolangBase/routers"
	"GolangBase/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
)

func CopyBuildTime() {
	logs.Info("begin to copy build time file")
	if util.Exist("/build_time") {
		if util.Exist("/data/build_time") {
			os.RemoveAll("/data/build_time")
		}
		if err := util.CopyDir("/build_time", "/data/build_time"); err != nil {
			logs.Error(err)
		} else {
			logs.Info("copy build time success")
		}
	} else {
		logs.Info("/build_time not exists")
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
