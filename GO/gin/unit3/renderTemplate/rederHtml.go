package renderTemplate

import (
	"github.com/gin-gonic/gin"
)

func LoadTemplate() {
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		//根据模板文件名渲染
		//渲染时json会把模板渲染，此处会把title替换
		c.HTML(200, "index.tmpl", gin.H{"title": "我的标题"})
	})

	r.Run(":8080")
}
