package requestProcess

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFormDataByPOST() {
	r := gin.Default()

	r.LoadHTMLFiles("html/upload.html")
	r.GET("/myform", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.POST("/myform_post", func(c *gin.Context) {
		//表单参数设置默认值
		ftype := c.DefaultPostForm("type", "alert")
		//获取表单参数
		user := c.PostForm("username")
		passwd := c.PostForm("password")
		//表单的多选框
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type is %s,username is %s,passwd is %s,爱好是%v", ftype, user, passwd, hobbys))
	})

	r.Run(":8080")
	return
}
