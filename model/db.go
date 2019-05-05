package model

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	DbName      = "base"
	MaxIdleConn = 50
	MaxOpenConn = 1000
)

func init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPWD := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	hostUrl := fmt.Sprintf("tcp(%s:%s)", mysqlHost, mysqlPort)
	if len(hostUrl) < 10 {
		fmt.Println("GolangBase:init:mysql env lost")
		os.Exit(0)
	}
	DriveUrl := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", mysqlUser, mysqlPWD, hostUrl, DbName)
	beego.Debug(DriveUrl)
	orm.RegisterDataBase("default", "mysql", DriveUrl, MaxIdleConn, MaxOpenConn)
	orm.RegisterDataBase(DbName, "mysql", DriveUrl, MaxIdleConn, MaxOpenConn)
}
