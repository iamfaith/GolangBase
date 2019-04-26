package redis_cluster

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestInt64(t *testing.T) {
	//c1 := []byte{'1','0','4'}
	c1 := []byte("104")
	c2 := int64(34)
	err := errors.New("查询错误")
	result, e := Int64(c1, nil)
	log.Println(result)
	assert.NoError(t, e)
	assert.Equal(t, int64(104), result)

	result, e = Int64(c2, nil)
	log.Println(result)
	assert.NoError(t, e)
	assert.Equal(t, int64(34), result)

	result, e = Int64(err, nil)
	log.Println(result)
	assert.Error(t, e)
	assert.Equal(t, int64(0), result)

	result, e = Int64(nil, nil)
	log.Println(result)
	assert.Error(t, e)
	assert.Equal(t, int64(0), result)

	result, e = Int64(c1, err)
	log.Println(result)
	assert.Error(t, e)
	assert.Equal(t, int64(0), result)

	result, e = Int64([]byte("ok"), nil)
	log.Println(result)
	assert.Error(t, e)
	assert.Equal(t, int64(0), result)

}

func TestValues(t *testing.T) {
	i := []interface{}{"string", []byte{0, 3}, int64(0), int64(45)}
	err := errors.New("查询错误")

	result, e := Values(i, nil)
	log.Println(result)
	assert.NoError(t, e)
	log.Println(i)
	assert.Equal(t, i, result)

	result, e = Values(i, err)
	//注意nil要进行强转
	assert.Equal(t, ([]interface{})(nil), result)
	assert.Error(t, e)

	result, e = Values(err, err)
	assert.Equal(t, ([]interface{})(nil), result)
	assert.Error(t, e)

	result, e = Values(nil, nil)
	assert.Equal(t, ([]interface{})(nil), result)
	assert.Error(t, e)

}

func TestInt64Map(t *testing.T) {
	i := []interface{}{[]byte("key1"), int64(11), []byte("key2"), int64(111)}
	res, e := Int64Map(i, nil)
	assert.NoError(t, e)
	j := make(map[string]int64, len(i)/2)
	j["key1"] = int64(11)
	j["key2"] = int64(111)
	assert.Equal(t, j, res)
	log.Println(res)

	i = []interface{}{[]byte("key1"), int64(11), []byte("key2"), int64(111), []byte("key2")}
	res, e = Int64Map(i, nil)
	assert.Error(t, e)
	assert.Equal(t, (map[string]int64)(nil), res)
	log.Println(res)

	i = []interface{}{"key1", int64(11), "key2", int64(111)}
	res, e = Int64Map(i, nil)
	assert.Error(t, e)
	assert.Equal(t, (map[string]int64)(nil), res)
	log.Println(res)

}

func TestInts(t *testing.T) {
	i := []interface{}{int64(1), int64(1), int64(1), []byte("123"), []byte("12")}
	res, e := Ints(i, nil)
	j := []int{1, 1, 1, 123, 12}
	assert.Equal(t, j, res)
	assert.NoError(t, e)

	i = []interface{}{int64(1), int64(1), int64(1), []byte("123"), "12", "123"}
	res, e = Ints(i, nil)
	j = []int{1, 1, 1, 123, 0, 0}
	assert.Equal(t, j, res)
	assert.Error(t, e)
}

func TestError(t *testing.T) {
	var err Error = ""
	err.Error()
}
