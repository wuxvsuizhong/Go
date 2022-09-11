package requestProcess

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义接收数据的结构体
type info struct {
	// binding:"required" 这个修饰是限定User和Password是必须字段，如果字段为空则会报错
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func ParseAsJSON() {
	r := gin.Default()

	/*
		使用curl发送请求：
		win命令：curl http://127.0.0.1:8080/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\",\"password\":\"admin\"}" -X POST
		linux命令：curl http://127.0.0.1:8080/loginJSON -H 'content-type:application/json' -d {"user":"root","password":"admin"}" -X POST
	*/

	r.POST("loginJSON", func(c *gin.Context) {
		var jsoninfo info
		//按照json格式把request的bosy数据，解析到结构体中
		if err := c.ShouldBindJSON(&jsoninfo); err != nil {
			//返回错误信息
			//gin.H 封装了生成json的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//对获取到的参数做判断，如果不满足那么返回
		if jsoninfo.User != "root" || jsoninfo.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})

	})

	r.Run(":8080")
}

func ParseFormDataAsJson() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/loginJSON", func(c *gin.Context) {
		var forminfo info
		//Bind 默认解析并绑定forminfo 格式
		//根据请求头中的content-type 自动判断
		if err := c.Bind(&forminfo); err != nil {
			//Bind失败则返回
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}

		if forminfo.User != "root" || forminfo.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304", "msg": "用户名或者密码不符合要求！"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run(":8080")
}

func ParseUriParamsAsJson() {
	//解析请求中的参数项
	r := gin.Default()

	/*
		使用curl发送请求：
		win命令：curl http://127.0.0.1:8080/login/root/admin
	*/

	r.GET("/login/:user/:password", func(c *gin.Context) {
		var jsoninfo info
		if err := c.ShouldBindUri(&jsoninfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		if jsoninfo.User != "root" || jsoninfo.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304", "msg": "请求uri不符合要求!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})

	})

	r.Run(":8080")
}
