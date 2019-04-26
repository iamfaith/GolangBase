package redis_cluster

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	prefix = "__go__test_1:"
	lkey   = prefix + "list"
	hkey   = prefix + "hash"
	allkey = prefix + "*"
	setkey = prefix + "set"
	zkey   = prefix + "z"
)

func TestGetValue(t *testing.T) {
	e := SetValue(hkey, "ttt", -1)
	assert.NoError(t, e)
	res, e := GetValue(hkey)
	assert.NoError(t, e)
	res, e = GetValue(hkey)
	assert.NoError(t, e)
	assert.Equal(t, "ttt", res)
}

func TestLPush(t *testing.T) {
	err := LPush(lkey, "lpush step1")
	if err != nil {
		t.Error(err)
		return
	}
	err = LPush(lkey, "lpush step2")
	if err != nil {
		t.Error(err)
	}

	//TestLRange(t)
	res, err := LPop(lkey)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "lpush step2", res)
}

func TestKeys(t *testing.T) {
	// https://github.com/go-redis/redis/issues/678
	// 在集群里面不管用，要自己去每个节点执行keys。。。
	Del(lkey)
	Del(hkey)
	LPush(lkey, "ttt")
	RPush(hkey, "ttt")
	keys, err := Keys(allkey)
	if err != nil {
		t.Error(err)
		return
	}
	if len(keys) == 0 {
		t.Error("Keys() empty return")
		return
	}
	log.Println("got keys: ", keys)
}

//添加到链表尾部
func TestRPush(t *testing.T) {
	Del(lkey)
	err := RPush(lkey, "rpush step3")
	if err != nil {
		t.Error(err)
	}
	err = RPush(lkey, "rpush step4")
	if err != nil {
		t.Error(err)
	}
	//TestLRange(t)
	re, e := LRange(lkey, 0, -1)
	if e != nil {
		t.Error(e)
	}
	log.Println("got vals:", re)
	assert.Equal(t, []string{"rpush step3", "rpush step4"}, re)
}

func TestLPop(t *testing.T) {
	LPush(lkey, "ttt")
	poped, err := LPop(lkey)
	if err != nil {
		t.Error(err)
		return
	}
	//log.Println("poped: ", poped)
	assert.Equal(t, "ttt", poped)
}

func TestLRange(t *testing.T) {
	Del(lkey)
	LPush(lkey, "rpush step5")
	LPush(lkey, "rpush step6")
	vals, err := LRange(lkey, 0, -1)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, []string{"rpush step6", "rpush step5"}, vals)
	log.Println("got vals: ", vals)
}

func TestLInsert(t *testing.T) {
	Del(lkey)
	LPush(lkey, "lpush step1")
	LPush(lkey, "lpush step2")
	LPush(lkey, "lpush step8")
	LPush(lkey, "lpush step9")
	err := LInsert(lkey, "lpush step8", "inserted before lpush step8")
	if err != nil {
		t.Error(err)
		return
	}
	res, e := LRange(lkey, 0, -1)
	if e != nil {
		t.Error(e)
		return
	}
	log.Println(res)
	var x, y string = " ", " "
	for _, val := range res {
		x = y
		y = val
		if y == "lpush step8" {
			break
		}

	}
	assert.Equal(t, "inserted before lpush step8", x)
}

func TestLRem(t *testing.T) {
	LPush(lkey, "inserted before lpush step8")
	LPush(lkey, "ttt")
	LPush(lkey, "tt")
	LPush(lkey, "t")
	err := LRem(lkey, "inserted before lpush step8")
	if err != nil {
		t.Error(err)
		return
	}
	res, e := LRange(lkey, 0, -1)
	if e != nil {
		t.Error(e)
		return
	}
	for _, val := range res {
		assert.True(t, val != "inserted before lpush step8")
	}
	//TestLRange(t)
}

func TestMGet(t *testing.T) {
	_, err := MGet([]string{lkey, hkey})
	assert.Error(t, err, errors.New("MGet don't allow use"))
	//禁止使用
	// cluster不支持MGet命令，除非keys在同一个slot，所以CROSSSLOT是正常返回
}

func TestDel(t *testing.T) {
	SetValue(hkey, "2", -1)
	r, e := Exists(hkey)
	assert.NoError(t, e)
	assert.Equal(t, true, r)
	e = Del(hkey)
	assert.NoError(t, e)
	r, e = Exists(hkey)
	assert.NoError(t, e)
	assert.Equal(t, false, r)
}

func TestHashIncr(t *testing.T) {
	e := Del(hkey)
	assert.NoError(t, e)
	e = HSet(hkey, "tt", 1)
	assert.NoError(t, e)
	v, e := HGet(hkey, "tt")
	assert.Equal(t, "1", v)
	assert.NoError(t, e)
	i, e := HashIncr(hkey, "tt")
	assert.Equal(t, int64(2), i)
	assert.NoError(t, e)
	v, _ = HGet(hkey, "tt")
	assert.Equal(t, "2", v)

	r, _ := HDel(hkey, "tt")
	assert.Equal(t, int64(1), r)
}

func TestHashIncrByFloat(t *testing.T) {
	e := Del(hkey)
	assert.NoError(t, e)
	v, e := HashIncrByFloat(hkey, "tt", 1.1)
	assert.NoError(t, e)
	assert.Equal(t, float64(1.1), v)
	v, e = HashIncrByFloat(hkey, "tt", 1.1)
	assert.NoError(t, e)
	assert.Equal(t, float64(2.2), v)
}

func TestIncr(t *testing.T) {
	SetValue(hkey, 1, -1)
	v, e := Incr(hkey)
	assert.NoError(t, e)
	assert.Equal(t, int64(2), v)
}

func TestSet(t *testing.T) {
	v, e := SAdd(setkey, "1", "2", "3")
	assert.NoError(t, e)
	assert.Equal(t, int64(3), v)

	b, e := SIsMember(setkey, "1")
	assert.NoError(t, e)
	assert.Equal(t, true, b)

	b, e = SIsMember(setkey, "3")
	assert.NoError(t, e)
	assert.Equal(t, true, b)

	r, e := SMembers(setkey)
	assert.NoError(t, e)
	assert.Equal(t, 3, len(r))

	v, e = SRem(setkey, r)
	assert.NoError(t, e)
	assert.Equal(t, int64(3), v)

	b, e = SIsMember(setkey, "3")
	assert.NoError(t, e)
	assert.Equal(t, false, b)

	r, e = SMembers(setkey)
	assert.NoError(t, e)
	assert.Equal(t, 0, len(r))
}

func TestClearKey(t *testing.T) {
	// keys := []string{lkey, hkey}
	keys, err := Keys(allkey)
	if err != nil {
		log.Println("failed to get keys")
		t.Error(err)
		return
	}
	for _, k := range keys {
		log.Println("clear key ", k)
		if err := ClearKey(k); err != nil {
			t.Error(err)
			return
		}
	}
}

func TestZ(t *testing.T) {
	re, e := ZRange(zkey, 0, -1)
	assert.NoError(t, e)

	if len(re) != 0 {
		r, e := ZRem(zkey, re...)
		assert.NoError(t, e)
		assert.Equal(t, int64(1), r)
	}

	v := make(map[string]float64)
	v["a"] = 0.3
	v["c"] = 1.9
	v["d"] = 9.7
	r, e := ZAdd(zkey, v)
	assert.NoError(t, e)
	assert.Equal(t, int64(3), r)

	re, e = ZRangeByScore(zkey, "1", "10", 0, 10)
	assert.NoError(t, e)
	assert.Equal(t, 2, len(re))
	assert.Equal(t, "c", re[0])
	assert.Equal(t, "d", re[1])

	re, e = ZRange(zkey, 0, 1)
	assert.NoError(t, e)
	assert.Equal(t, 2, len(re))
	assert.Equal(t, "a", re[0])
	assert.Equal(t, "c", re[1])

	r, e = ZRem(zkey, re...)
	assert.NoError(t, e)
	assert.Equal(t, int64(2), r)

	re, e = ZRange(zkey, 0, 100)
	assert.NoError(t, e)
	assert.Equal(t, 1, len(re))
	assert.Equal(t, "d", re[0])

	r, e = ZRem(zkey, re...)
	assert.NoError(t, e)
	assert.Equal(t, int64(1), r)
}

func TestTTL(t *testing.T) {
	Del(hkey)
	SetValue(hkey, "wawa", 2)
	r, e := TTL(hkey)
	assert.NoError(t, e)
	if r > 0 && r <= 2 {
		assert.Equal(t, 1, 1)
	} else {
		assert.Equal(t, 1, r)
	}

	time.Sleep(2 * time.Second)

	_, e = GetValue(hkey)
	assert.Error(t, errors.New("redis: nil"), e)
}

func TestLLen(t *testing.T) {
	type args struct {
		listName string
	}
	listName := args{listName: "testList"}
	ClearKey(listName.listName)
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "LLen",
			args:    listName,
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LLen(tt.args.listName)
			assert.NoError(t, err, "LLen() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
	SetValue(listName.listName, "1", 60)
}

func TestRPushWithInx(t *testing.T) {
	type args struct {
		listName string
		value    string
	}
	para := args{listName: "testList", value: "10"}
	ClearKey(para.listName)
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "RPushWithInx",
			args:    para,
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RPushWithInx(tt.args.listName, tt.args.value)
			assert.NoError(t, err, "RPushWithInx() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestListAll(t *testing.T) {
	type args struct {
		listName string
	}
	listName := args{listName: "testList"}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "ListAll",
			args:    listName,
			want:    []string{"10"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAll(tt.args.listName)
			assert.NoError(t, err, "ListAll() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLRemWithDir(t *testing.T) {
	type args struct {
		listName  string
		direction int64
		value     string
	}
	param := args{listName: "testList", direction: 1, value: "123"}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "LRemWithDir",
			args:    param,
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LRemWithDir(tt.args.listName, tt.args.direction, tt.args.value)
			assert.NoError(t, err, "LRemWithDir() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPing(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Ping",
			want:    "PONG",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ping()
			assert.NoError(t, err, "Ping() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExistAndGet(t *testing.T) {
	type args struct {
		key string
	}
	ClearKey("test")
	key := args{key: "test"}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "ExistAndGet",
			args:    key,
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExistAndGet(tt.args.key)
			assert.NoError(t, err, "ExistAndGet() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
	SetValue("test", "1", 60)
}

func TestHExists(t *testing.T) {
	type args struct {
		key   string
		field string
	}
	key := args{key: "test123", field: "123"}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "HGetAll",
			args:    key,
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HExists(tt.args.key, tt.args.field)
			assert.NoError(t, err, "HExists() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetExValue(t *testing.T) {
	type args struct {
		key     string
		value   interface{}
		timeout int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "RPushWithInx",
			args:    args{"testabc", "1", 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetExValue(tt.args.key, tt.args.value, tt.args.timeout)
			assert.NoError(t, err, "SetExValue() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func TestHashIncrBy(t *testing.T) {
	HashIncrBy("hkey", "fieldkey", 2)
}

func TestHGetAll(t *testing.T) {
	HGetAll("hkey")
}

func TestExpire(t *testing.T) {
	Expire("hkey", 1000)
}

func TestPublish(t *testing.T) {
	Publish("channel", "interface data")
}

func TestSubscribe(t *testing.T) {
	Subscribe("channel", "interface data")
}
