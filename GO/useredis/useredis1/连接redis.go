package useredis

import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)

func connectRedis() *redis.Conn {
  conn, err := redis.Dial("tcp", "192.168.0.120:6379")
  if err != nil {
    fmt.Println("连接redis失败，err=", err)
  }
  fmt.Println("连接redis成功!conn=", conn)
  return &conn
}

func AddKeyVal() {
  c := connectRedis()
  conn := *c
  defer conn.Close()

  _, err := conn.Do("set", "name", "tom猫")
  if err != nil {
    fmt.Println("set err:", err)
    return
  }

  r, err := redis.String(conn.Do("get", "name"))
  if err != nil {
    fmt.Println("get err:", err)
    return
  }
  fmt.Println("r:", r)
}

func AddHash() {
  c := connectRedis()
  conn := *c
  defer conn.Close()

  conn.Do("hset", "user01", "name", "张三")
  conn.Do("hset", "user01", "age", "27")

  //使用hash意思设置对象的多个field和值
  conn.Do("hset", "user01", "name", "jerry", "age", 22)

  //返回的是单个的字符串，使用redis.string即可
  r, err := redis.String(conn.Do("hget", "user01", "name"))
  if err != nil {
    fmt.Println("get err:", err)
    return
  }
  fmt.Println("r:", r)

  //返回结果是多个字符串，需要使用redios.strings
  r2, err := redis.Strings(conn.Do("hgetall", "user01"))
  if err != nil {
    fmt.Println("get err:", err)
    return
  }
  fmt.Println("r:", r2)

  //一次获取某个hash的多个值
  r3, err := redis.Strings(conn.Do("hmget", "user01", "name", "age"))
  if err != nil {
    fmt.Println("get err:", err)
    return
  }
  fmt.Println("r:", r3)

}
