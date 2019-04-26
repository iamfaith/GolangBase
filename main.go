package main

import (
	_ "GolangBase/routers"
	"GolangBase/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func CopyBuildTime() {
	if err := util.CopyDir("/build_time", "/data/build_time"); err != nil {
		logs.Error(err)
	}
}

func main() {
	go CopyBuildTime()
	beego.Run()
}
