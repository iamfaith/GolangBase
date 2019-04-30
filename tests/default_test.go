package tests

import (
	"GolangBase/service/redis_cluster"
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

	ret, _ := json.Marshal(user[1])

	var retObj []map[string]interface{}
	json.Unmarshal([]byte(ret), &retObj)
	fmt.Print(retObj)
	redis_cluster.SetValue("c", ret, -1)
}
