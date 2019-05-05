package model

import (
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
	var link Link
	o := orm.NewOrm()
	o.Using(DbName)
	sql := "SELECT * from ? WHERE uid=?"
	_, err := o.Raw(sql, LinkTbl, uid).QueryRows(&link)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &link, nil
}
