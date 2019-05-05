package model

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UploadFile struct {
	Status   string `json:"status" bson:"status"`
	FilePath string `json:"-"`
}

type Link struct {
	Uid     string `json:"uid" bson:"uid"`
	Content string `json:"content" bson:"content"`
}

const (
	LinkTbl = "link"
)

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

func FindLinkByUid(uid string) (*Link, error) {
	//link := Link{UploadFile:UploadFile{Status:"aaa", FilePath:"bb"}}
	link := Link{}
	o := orm.NewOrm()
	o.Using(DbName)
	sql := fmt.Sprintf("SELECT * from %s WHERE uid=?", LinkTbl)
	err := o.Raw(sql, uid).QueryRow(&link)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &link, nil
}
