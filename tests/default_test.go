package tests

import (
	"GolangBase/define"
	_ "GolangBase/init_config"
	"GolangBase/service/redis_cluster"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type User struct {
	Status int64  `json:"status" bson:"status"`
	Name   string `json:"name"`
}

func TestUser(t *testing.T) {
	user := make([]User, 2)

	user[0] = User{Status: 1, Name: "jin"}
	user[1] = User{Status: 0, Name: "f"}

	aa := []int{0, 1}
	ret, _ := json.Marshal(aa)

	var retObj []map[string]interface{}
	json.Unmarshal([]byte(ret), &retObj)
	fmt.Print(retObj)

	var nums []int
	json.Unmarshal([]byte("[0,1]"), &nums)
	fmt.Print(nums)
	//redis_cluster.SetValue("d", ret, -1)
}

func TestMysql(t *testing.T) {
	//fmt.Println(model.GetAllByKv(model.Kv{Key: "uid", Value: "content", TableName: model.LinkTbl}))
	//fmt.Println(model.Insert(model.Link{Uid: "70", Content: "aa"}))
	//fmt.Println(unsafe.Sizeof(model.Link{}))
	//m := make(map[string]interface{})
	//m["uid"] = "99"
	//m["Content"] = "11"
	//m["tbl"] = "link"
	//fmt.Println(model.InsertM(m))
	//fmt.Println(model.GetAll(model.Link{}))
	//
	//ty := reflect.TypeOf(redis_cluster.SetValue)
	//fmt.Println(ty, ty.In(0))

	time.Sleep(1 * time.Second)
	err := redis_cluster.LPush(define.UploadFile, "1111")
	fmt.Println(err)
	err = redis_cluster.LPush(define.UploadFile, "2222")
	fmt.Println(err)
	err = redis_cluster.LPush(define.UploadFile, "3333")
	fmt.Println(err)
	go ConsumeRedisTask()
	time.Sleep(1000 * time.Second)
}

func ConsumeRedisTask() {
	for true {
		if ret, err := redis_cluster.RPop(define.UploadFile); err != nil {
			fmt.Println(ret)
		} else {
			fmt.Println("wait...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
