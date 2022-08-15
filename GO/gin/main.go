package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main1() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context 封装了Request和Response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	// 3.监听端口
	r.Run(":9000")
}

func main() {
	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello"+name)
	})
	r.Run(":9000")
}
