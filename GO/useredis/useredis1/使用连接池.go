package useredis

import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
  pool = &redis.Pool{
    MaxIdle:     5,   //最大空闲连接数
    MaxActive:   0,   //和数据的最大连接数，设置为0时表示不限制
    IdleTimeout: 100, //最大开空闲时间
    Dial: func() (conn redis.Conn, err error) { //连接那个redis
      return redis.Dial("tcp", "192.168.0.120:6379")
    },
  }
}

func UseRedisPool() {
  conn := pool.Get() //从连接池里面获取一个连接
  defer conn.Close()

  _, err := conn.Do("set", "name", "tom2")
  if err != nil {
    fmt.Println("使用连接池设置key失败,err:", err)
    return
  }

  r, err := redis.String(conn.Do("get", "name"))
  if err != nil {
    fmt.Println("使用连接池获取key失败,err:", err)
    return
  }
  fmt.Println("使用连接池获取key:", r)

  pool.Close() //连接池关闭后，就不能再使用了

  /*
  	conn2 := pool.Get()
  	_, err = conn2.Do("set", "name", "jerry")
  	if err != nil {
  		fmt.Println("使用连接池设置key失败,err:", err)
  		return
  	}
  */
}
