package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"os"
	"path/filepath"
	"GolangBase/util"
	"GolangBase/model"
	"GolangBase/service/redis_cluster"
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

const (
	UploadFile = "upload_file"
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
		logs.Error(err)
		this.Fail(CodeBadParam, err.Error())
	} else {
		if f == nil {
			this.Fail(CodeBadParam, "file is null")
		} else {
			fileName := h.Filename
			path := "/data/upload/"
			if !utils.FileExists(path) {
				os.MkdirAll(path, os.ModePerm)
			}
			path = filepath.Join(path, fileName)
			logs.Info("upload file at ", path)
			defer f.Close()
			this.SaveToFile("file", path)
			sha1, _ := util.HashFileSha1(path)
			if err := redis_cluster.SetExValue(UploadFile+sha1, model.UploadFile{Status: NotProcess.String(), FilePath: path}, 600); err != nil {
				this.Fail(CodeBadParam, "redis gone wrong")
			} else {
				redis_cluster.LPush(UploadFile, sha1)
				this.Success("ok", sha1)
			}
		}
	}
}
