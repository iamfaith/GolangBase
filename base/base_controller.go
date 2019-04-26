package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"os"
	"path/filepath"
	"wps.ai/yun/ytalk/common/log"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	CodeOk        = 0
	CodeBadParam  = -1
	CodeServerErr = -2
)

func NewFailResponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}

func NewSuccessReponse(msg string, data interface{}) *Response {
	return &Response{
		Code: CodeOk,
		Msg:  msg,
		Data: data,
	}
}

func (this *BaseController) Fail(code int, msg string) {
	res := NewFailResponse(code, msg)
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) Success(msg string, data interface{}) {
	res := NewSuccessReponse(msg, data)
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) Upload() {
	if f, h, err := this.GetFile("file"); err != nil {
		log.Error(err)
		this.Fail(CodeBadParam, err.Error())
	} else {
		if f == nil {
			this.Fail(CodeBadParam, "file is null")
		} else {
			path := "upload/" + h.Filename
			if !utils.FileExists(path) {
				workPath, _ := os.Getwd()
				path = filepath.Join(filepath.Dir(workPath), path)
			}
			log.Info("upload file at ", path)
			defer f.Close()
			this.SaveToFile("file", path)
			this.Success("ok", nil)
		}
	}
}
