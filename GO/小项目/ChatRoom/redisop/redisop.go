package redisop

import (
	"time"
	"github.com/gomodule/redigo/redis"
)
// redis连接池全局变量指针
var pool *redis.Pool

func InitPool(addr string, maxiIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxiIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", addr)
		},
	}
}

