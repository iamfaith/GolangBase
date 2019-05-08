package tests

import (
	_ "GolangBase/init_config"
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

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func TestMysql(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}
