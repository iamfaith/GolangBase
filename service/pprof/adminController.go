package pprof

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

type adminController struct {
	beego.Controller
}

func checkPrivilege(ctx *context.Context) error {
	adminKey := beego.AppConfig.String("adminKey")
	headKey := ctx.Input.Header("adminKey")
	if headKey != adminKey {
		return errors.New("not the administrator")
	}
	return nil
}

func (c *adminController) Auth(ctx *context.Context) {
	err := checkPrivilege(ctx)
	if err != nil {
		logs.Error(err)
		ctx.Output.Body([]byte("Not Authorized"))
		ctx.Output.SetStatus(403)
		c.StopRun()
	}
}
