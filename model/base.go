package model

import (
	"GolangBase/define"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

type UploadFile struct {
	Status   string `json:"status" bson:"status"`
	FilePath string `json:"-"`
}

type Link struct {
	Uid     string `json:"uid" bson:"uid"`
	Content string `json:"content" bson:"content"`
}

type Kv struct {
	Key       string
	Value     string
	TableName string
}

func (l *Link) InsertLink() int64 {
	o := orm.NewOrm()
	o.Using(DbName)
	id, err := o.Insert(l)
	if err != nil {
		logs.Error(err, l)
		return -1
	}
	return id
}

func InsertM(m map[string]interface{}) (int64, error) {
	tbl, ok := m["tbl"]
	if !ok {
		return 0, errors.New("param error")
	} else {
		delete(m, "tbl")
	}
	o := orm.NewOrm()
	o.Using(DbName)
	param := ""
	count := len(m)
	i := 0
	pval := make([]interface{}, count)
	for k, v := range m {
		if i == count-1 {
			param += fmt.Sprintf("%s", strings.ToLower(k))
		} else {
			param += fmt.Sprintf("%s, ", strings.ToLower(k))
		}
		pval[i] = v
		i++
	}
	sql := fmt.Sprintf("insert into %s (%s) values (?, ?)", tbl, param)
	result, err := o.Raw(sql, pval).Exec()
	if err != nil {
		logs.Error(err)
		return 0, err
	} else {
		num, _ := result.RowsAffected()
		return num, nil
	}
}

func Insert(t interface{}) (int64, error) {
	o := orm.NewOrm()
	o.Using(DbName)
	val := reflect.ValueOf(t)
	param := ""
	count := val.NumField()
	pval := make([]interface{}, count)
	for i := 0; i < count; i++ {
		if i == count-1 {
			param += fmt.Sprintf("%s", strings.ToLower(val.Type().Field(i).Name))
		} else {
			param += fmt.Sprintf("%s, ", strings.ToLower(val.Type().Field(i).Name))
		}
		pval[i] = val.Field(i).Interface()
	}
	sql := fmt.Sprintf("insert into %s (%s) values (?, ?)", define.LinkTbl, param)
	result, err := o.Raw(sql, pval).Exec()
	if err != nil {
		logs.Error(err)
		return 0, err
	} else {
		num, _ := result.RowsAffected()
		return num, nil
	}
}

func GetAll(kv interface{}) (interface{}, error) {
	val := reflect.ValueOf(kv)
	if val.Kind() != reflect.Struct || !val.IsValid() {
		return nil, errors.New("param error")
	}
	if val.NumField() != 2 {
		return nil, errors.New("must be key value")
	}
	o := orm.NewOrm()
	o.Using(DbName)
	sql := fmt.Sprintf("SELECT * from %s", strings.ToLower(reflect.TypeOf(kv).Name()))
	var u orm.Params
	key := strings.ToLower(val.Type().Field(0).Name)
	value := strings.ToLower(val.Type().Field(1).Name)
	_, err := o.Raw(sql).RowsToMap(&u, key, value)
	if err != nil {
		logs.Error(err)
		return nil, err

	}
	return u, nil
}

func GetAllByKv(kv Kv) (interface{}, error) {
	o := orm.NewOrm()
	o.Using(DbName)
	sql := fmt.Sprintf("SELECT * from %s", kv.TableName)
	var u orm.Params
	_, err := o.Raw(sql).RowsToMap(&u, kv.Key, kv.Value)
	if err != nil {
		logs.Error(err)
		return nil, err

	}
	return u, nil
}

func FindLinkByUid(uid string) (*Link, error) {
	//link := Link{UploadFile:UploadFile{Status:"aaa", FilePath:"bb"}}
	link := Link{}
	o := orm.NewOrm()
	o.Using(DbName)
	sql := fmt.Sprintf("SELECT * from %s WHERE uid=?", define.LinkTbl)
	err := o.Raw(sql, uid).QueryRow(&link)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &link, nil
}
