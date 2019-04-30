package base

import (
	"GolangBase/model"
	"GolangBase/service/redis_cluster"
	"GolangBase/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
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

var funcs = util.NewFuncs(2)

func init() {
	funcs.Bind("ListAll", redis_cluster.ListAll)
	funcs.Bind("GetValue", redis_cluster.GetValue)
}

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

func (this *BaseController) Redis() {
	method := this.Ctx.Input.Param(":method")
	id := this.Ctx.Input.Param(":id")
	if strings.Contains(id, "*") {
		this.Fail(CodeBadParam, "bad param")
	}
	//var err error
	//var ret interface{}
	//switch method {
	//case "ListAll":
	//	ret, err = redis_cluster.ListAll(id)
	//	break
	//case "GetValue":
	//	ret, err = redis_cluster.GetValue(id)
	//	break
	//default:
	//	err = errors.New("bad params")
	//}
	//if err != nil {
	//	logs.Error(err)
	//	this.Fail(CodeBadParam, err.Error())
	//} else {
	//	this.Success("ok", ret)
	//}
	if val, err := funcs.Call(method, id); err != nil || len(val) == 0 {
		logs.Error("Call %s: %s %v", method, id, err)
		this.Fail(CodeBadParam, fmt.Sprintf("bad param: %v", err.Error()))
	} else {
		ret := val[0]
		logs.Info(ret.Kind())
		switch ret.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result := strconv.FormatInt(ret.Int(), 10)
			this.Success("ok", result)
		case reflect.String:
			result := ret.String()
			if len(result) > 0 && result[0] == '[' && result[len(result)-1] == ']' {
				if strings.Contains(result, "\"") {
					var arrObj []map[string]interface{}
					json.Unmarshal([]byte(result), &arrObj)
					this.Success("ok", arrObj)
				} else {
					var arr []int
					json.Unmarshal([]byte(result), &arr)
					this.Success("ok", arr)
				}
			} else if len(result) > 0 && result[0] == '{' && result[len(result)-1] == '}' {
				var retObj map[string]interface{}
				json.Unmarshal([]byte(result), &retObj)
				this.Success("ok", retObj)
			}
			this.Success("ok", result)
		case reflect.Slice:
			count := ret.Len()
			var arrays []string
			for i := 0; i < count; i++ {
				arrays = append(arrays, ret.Index(i).String())
			}
			this.Success("ok", arrays)
		}
		logs.Info("call %s: %s: %v %v %v", method, id, val, ret)
		this.Success("ok", "")
	}

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
			path := beego.AppConfig.String("upload_path")
			if !utils.FileExists(path) {
				if err := os.MkdirAll(path, os.ModePerm); err != nil {
					path = "./"
				}
			}
			sha1, _ := util.HashFileSha1(f)
			path = fmt.Sprintf("%s/%s", path, sha1)
			if !utils.FileExists(path) {
				os.MkdirAll(path, os.ModePerm)
			}
			path = filepath.Join(path, fileName)
			logs.Info("upload file at ", path)
			defer f.Close()
			if !utils.FileExists(path) {
				this.SaveToFile("file", path)
			}
			fileObj, _ := json.Marshal(model.UploadFile{Status: NotProcess.String(), FilePath: path})

			if err := redis_cluster.SetExValue(UploadFile+sha1, string(fileObj), 600); err != nil {
				logs.Error(sha1, path, err)
				this.Fail(CodeBadParam, err.Error())
			} else {
				redis_cluster.LPush(UploadFile, sha1)
				this.Success("ok", sha1)
			}
		}
	}
}
