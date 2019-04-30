package tests

import (
	"encoding/json"
	"fmt"
	"testing"
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
