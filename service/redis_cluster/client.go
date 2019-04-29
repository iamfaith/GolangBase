package redis_cluster

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	// v5 有连接问题
	// "gopkg.in/redis.v5"
	"github.com/go-redis/redis"
	"os"
	"strings"
	"time"
)

const (
	REDIS_CLUSTER_URL  = "redis_cluster_url"
	REDIS_POOL_MAX_CON = "redis_pool_max_con"
)

var (
	client *redis.ClusterClient
)

func init() {
	addrs := os.Getenv("REDIS_CLUSTER")
	logs.Debug("redis cluster from env:" + addrs)
	if len(addrs) == 0 {
		addrs = beego.AppConfig.String(REDIS_CLUSTER_URL)
		logs.Debug("redis cluster from conf:" + addrs)
		if len(addrs) == 0 {
			logs.Debug("init: REDIS_CLUSTER env lost:" + addrs)
			os.Exit(0)
		}
	}
	urls := strings.Split(addrs, ",")
	logs.Debug("urls", urls)
	client = newClient(urls)
}

//添加到链表头部
func LPush(key string, value string) error {
	cmd := client.LPush(key, value)
	return cmd.Err()
}

//添加到链表尾部
func RPush(key string, value string) error {
	cmd := client.RPush(key, value)
	return cmd.Err()
}

func LPop(key string) (string, error) {
	re := client.LPop(key)
	if re.Err() != nil {
		return "", re.Err()
	}
	return re.Result()
}

func RPop(key string) (string, error) {
	re := client.RPop(key)
	if re.Err() != nil {
		return "", re.Err()
	}
	return re.Result()
}

//如果要返回所有，start=0 end=-1
func LRange(key string, start, end int64) ([]string, error) {
	re := client.LRange(key, start, end)
	if re.Err() != nil {
		return []string{}, re.Err()
	}
	return re.Result()
}

//在pivot 之前插入value
func LInsert(key string, pivot string, value string) error {
	cmd := client.LInsert(key, "BEFORE", pivot, value)
	return cmd.Err()
}

//删除list中的某个值
func LRem(key string, value string) error {
	cmd := client.LRem(key, 0, value)
	return cmd.Err()
}

func LLen(listName string) (int64, error) {
	cmd := client.LLen(listName)
	if cmd.Err() != nil {
		return -1, cmd.Err()
	}
	return cmd.Val(), nil
}

func RPushWithInx(listName, value string) (int64, error) {
	cmd := client.RPush(listName, value)
	if cmd.Err() != nil {
		return -1, cmd.Err()
	}
	return cmd.Val(), nil
}

func ListAll(listName string) ([]string, error) {
	return LRange(listName, 0, -1)
}

func LRemWithDir(listName string, direction int64, value string) (int64, error) {
	cmd := client.LRem(listName, direction, value)
	if cmd.Err() != nil {
		return -1, cmd.Err()
	}
	return cmd.Val(), nil
}

func Ping() (string, error) {
	cmd := client.Ping()
	return cmd.Val(), cmd.Err()
}

func ExistAndGet(key string) (string, error) {
	if ret, err := Exists(key); err != nil {
		return "", err
	} else {
		if ret {
			return GetValue(key)
		} else {
			return "", nil
		}
	}
}

func SetExValue(key string, value interface{}, timeout int) error {
	cmd := client.Set(key, value, time.Duration(timeout)*time.Second)
	return cmd.Err()
}

// timeout 单位s  , -1 表示不设置
func SetValue(key string, value interface{}, timeout int) error {
	cmd := client.Set(key, value, time.Duration(timeout)*time.Second)
	return cmd.Err()
}

// cluster不支持MGet命令，会返回CROSSSLOT错误
func MGet(key []string) ([]interface{}, error) {
	return nil, errors.New("MGet don't allow use")
}

func GetValue(key string) (string, error) {
	cmd := client.Get(key)
	return cmd.Result()
}

func ClearKey(key string) error {
	cmd := client.Del(key)
	return cmd.Err()
}

func Incr(key string) (int64, error) {
	cmd := client.Incr(key)
	return cmd.Result()
}

func HashIncr(hkey, fieldkey string) (int64, error) {
	cmd := client.HIncrBy(hkey, fieldkey, 1)
	return cmd.Result()
}

func HashIncrBy(hkey, fieldkey string, incr int64) (int64, error) {
	cmd := client.HIncrBy(hkey, fieldkey, incr)
	return cmd.Result()
}

func HashIncrByFloat(hkey, fieldkey string, incr float64) (float64, error) {
	cmd := client.HIncrByFloat(hkey, fieldkey, incr)
	return cmd.Result()
}

func HSet(key, field string, value interface{}) error {
	cmd := client.HSet(key, field, value)
	return cmd.Err()
}

func HGet(key string, field string) (string, error) {
	cmd := client.HGet(key, field)
	return cmd.Result()
}

func HDel(key string, field ...string) (int64, error) {
	cmd := client.HDel(key, field...)
	return cmd.Result()
}

func HGetAll(key string) (map[string]string, error) {
	cmd := client.HGetAll(key)
	return cmd.Result()
}

func HExists(key, field string) (bool, error) {
	cmd := client.HExists(key, field)
	return cmd.Result()
}

func Expire(key string, exp int) (bool, error) {
	cmd := client.Expire(key, time.Duration(exp)*time.Second)
	return cmd.Result()
}

func SMembers(key string) ([]string, error) {
	cmd := client.SMembers(key)
	return cmd.Result()
}

func SRem(key string, members []string) (int64, error) {
	cmd := client.SRem(key, members)
	return cmd.Result()
}

func SAdd(key string, value ...string) (int64, error) {
	cmd := client.SAdd(key, value)
	return cmd.Result()
}

func SIsMember(key string, member string) (bool, error) {
	cmd := client.SIsMember(key, member)
	return cmd.Result()
}

func ZRange(key string, start, stop int64) ([]string, error) {
	cmd := client.ZRange(key, start, stop)
	return cmd.Result()
}

func ZRangeByScore(key, min, max string, offset, count int64) ([]string, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	cmd := client.ZRangeByScore(key, opt)
	return cmd.Result()
}

func ZRem(key string, mem ...string) (int64, error) {
	cmd := client.ZRem(key, mem)
	return cmd.Result()
}

func ZAdd(key string, mem map[string]float64) (int64, error) {
	zv := []redis.Z{}
	for k, v := range mem {
		z := redis.Z{Score: v, Member: k}
		zv = append(zv, z)
	}
	cmd := client.ZAdd(key, zv...)
	return cmd.Result()
}

func TTL(key string) (time.Duration, error) {
	cmd := client.TTL(key)
	return cmd.Result()
}

func Publish(channel string, message interface{}) (int64, error) {
	cmd := client.Publish(channel, message)
	return cmd.Result()
}

func Subscribe(channel ...string) *redis.PubSub {
	cmd := client.Subscribe(channel...)
	return cmd
}

func Del(key ...string) error {
	cmd := client.Del(key...)
	return cmd.Err()
}

// 判断key是否存在，存在返回true，不存在返回false
func Exists(key string) (bool, error) {
	cmd := client.Exists(key)
	r, e := cmd.Result()
	return r == 1, e
}

func Keys(key string) ([]string, error) {
	var keys []string
	nodeKey := func(client *redis.Client) error {
		// fmt.Printf("%+v\n", client)
		cmd := client.Keys(key)
		ks, err := cmd.Result()
		if err != nil {
			return err
		}
		keys = append(keys, ks...)
		return nil
	}
	// cluster不支持keys命令，自己遍历一下所有master
	if err := client.ForEachMaster(nodeKey); err != nil {
		return keys, err
	}
	return keys, nil
}

func newClient(addrs []string) *redis.ClusterClient {
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
		FixIPRemapping: true,
	})

	cmd := client.Ping()
	if cmd.Err() != nil {
		logs.Error("redis cluster err:", cmd.Err())
		os.Exit(-1)
	}

	return client
}
