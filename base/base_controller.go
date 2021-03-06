package base

import (
	"GolangBase/define"
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

var funcs = util.NewFuncs(2)
var typeRegistry = make(map[string]reflect.Type)

func init() {
	funcs.Bind("ListAll", redis_cluster.ListAll)
	funcs.Bind("GetValue", redis_cluster.GetValue)
	funcs.Bind("FindLinkByUid", model.FindLinkByUid)
	funcs.Bind("GetAll", model.GetAll)
	funcs.Bind("Insert", model.InsertM)

	typeRegistry["Link"] = reflect.TypeOf(model.Link{})
}

func makeInstance(name string) interface{} {
	reflectT := typeRegistry[name]
	if reflectT == nil {
		return nil
	}
	v := reflect.New(reflectT).Elem()
	// Maybe fill in fields here if necessary
	return v.Interface()
}

func NewFailResponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}

func NewSuccessReponse(msg string, data interface{}) *Response {
	return &Response{
		Code: define.CodeOk,
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

func (this *BaseController) PostByReflect() {
	method := this.Ctx.Input.Param(":method")
	var m map[string]interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &m); err != nil {
		logs.Error(err)
		this.Fail(define.CodeBadParam, "param error")
	}
	this.callFunc(method, m)
}

func (this *BaseController) GetByReflect() {
	method := this.Ctx.Input.Param(":method")
	id := this.Ctx.Input.Param(":id")
	t := this.GetString("t")
	if strings.Contains(id, "*") {
		this.Fail(define.CodeBadParam, "bad param")
	}
	if t == "all" {
		if reflectType := makeInstance(id); reflectType == nil {
			this.Fail(define.CodeBadParam, "param error")
		} else {
			this.callFunc(method, reflectType)
		}

	}
	this.callFunc(method, id)

}

func (this *BaseController) callFunc(method string, param interface{}) {
	if val, err := funcs.Call(method, param); err != nil || len(val) == 0 {
		logs.Error("Call %s: %s %v", method, param, err)
		this.Fail(define.CodeBadParam, fmt.Sprintf("bad param: %v", err.Error()))
	} else {
		ret := val[0]
		logs.Info(ret.Kind(), ret)
		logs.Info("call %s: %s: %v %v", method, param, val, ret)
		if ret.IsValid() {
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
			case reflect.Ptr:
				if !ret.IsNil() {
					fields := ret.Elem()
					if fields.Kind() == reflect.Map {
						this.handleReflectMap(fields)
					} else {
						logs.Info(fields.Kind())
					}
					m := make(map[string]interface{})
					for i := 0; i < fields.NumField(); i++ {
						valueField := fields.Field(i)
						typeField := fields.Type().Field(i)
						if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
							elm := valueField.Elem()
							if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
								valueField = elm
							}
						}
						if valueField.Kind() == reflect.Ptr {
							valueField = valueField.Elem()

						}
						m[strings.ToLower(typeField.Name)] = valueField.Interface()
						//address:= "not-addressable"
						//if valueField.CanAddr() {
						//	address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
						//}
						//fmt.Printf("Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n", typeField.Name,
						//	valueField.Interface(), address, typeField.Type, valueField.Kind())
						//if valueField.Kind() == reflect.Struct {
						//}
					}
					this.Success("ok", m)
				}
			case reflect.Interface:
				if !ret.IsNil() {
					this.Success("ok", ret.Elem().Interface())
				}
			case reflect.Map:
				this.handleReflectMap(ret)
			}
		}
		this.Success("ok", "")
	}
}

func (this *BaseController) handleReflectMap(ret reflect.Value) {
	if !ret.IsNil() {
		m := make(map[string]interface{})
		for _, key := range ret.MapKeys() {
			v := ret.MapIndex(key)
			m[key.String()] = v.Interface()
		}
		this.Success("ok", m)
	}
}

func (this *BaseController) Upload() {
	if f, h, err := this.GetFile("file"); err != nil {
		logs.Error(err)
		this.Fail(define.CodeBadParam, err.Error())
	} else {
		if f == nil {
			this.Fail(define.CodeBadParam, "file is null")
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
			fileObj, _ := json.Marshal(model.UploadFile{Status: define.NotProcess.String(), FilePath: path})

			if err := redis_cluster.SetValue(define.UploadFile+sha1, string(fileObj), -1); err != nil {
				logs.Error(sha1, path, err)
				this.Fail(define.CodeBadParam, err.Error())
			} else {
				redis_cluster.LPush(define.UploadFile, sha1)
				this.Success("ok", sha1)
			}
		}
	}
}
