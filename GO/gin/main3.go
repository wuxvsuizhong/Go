package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		// 如果在get的承诺书中获取不到name那么使用默认值zhww
		name := c.DefaultQuery("name", "zhww")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	// http://127.0.0.1:9000/welcome?name=李四
	//能获取到get请求参数name的值
	// http://127.0.0.1:9000/welcome
	//获取不到get请求参数name的值，用默认值zhww返回
	r.Run(":9000")
}
