package redisop

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisHp struct {
	redisConn redis.Conn
}


const (
	HASH_SET_NOT_FOUND = "没有找哈希集"
	HASH_KEY_NOT_FOUND = "在哈希集中没有找到键值"
)

//从redis连接池中获取一个连接
func getConnFromPool(pool *redis.Pool) (err error, conn redis.Conn) {
	if redisHp.redisConn != nil {
		conn = redisHp.redisConn
		return
	}
	conn = pool.Get()

	if err = conn.Err(); err != nil {
		fmt.Println("从redis连接池获取连接失败!")
		conn.Close()
		return
	}

	return
}

func QueryHashKeyVal(hkey string, vkey string) (res string, err error) {
	err, conn := getConnFromPool(pool)
	if err != nil {
		return
	}

	res, err = redis.String(conn.Do("hget", hkey, vkey))
	if err != nil {
		fmt.Printf("在redis hash %s中没有找到key:%s,err=%s", hkey, vkey,err)
    err = errors.New(HASH_KEY_NOT_FOUND)
		return
	}
	fmt.Println(res)
	return
}

func SetHashKey(hkey, val string) (err error) {
	err, conn := getConnFromPool(pool)
	if err != nil {
		return
	}
	if _, err = conn.Do("hset", val); err != nil {
		return
	}

	return
}
