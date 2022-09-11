package requestProcess

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UseRouteGroup() {
	r := gin.Default()
	// r.LoadHTMLFiles("html/*")
	r.LoadHTMLGlob("html/*")
	//创建路由组v1 ,用来专门处理GET请求
	v1 := r.Group("/v1")
	v1.GET("/myform", showmyform)
	v1.GET("/upload", showupload)

	//创建路由组v2,来专门处理POST请求

	v2 := r.Group("v2")
	v2.POST("/myform_post", process_myform_post)
	v2.POST("/upload_post", process_upload_post)
	r.MaxMultipartMemory = 8 << 20
	v2.POST("/uploads_post", process_uploads_post)

	r.Run(":8080")
}

// 处理各种请求的子函数
func showmyform(c *gin.Context) {
	c.HTML(http.StatusOK, "routeGroup_myform.html", nil)
}

func showupload(c *gin.Context) {
	c.HTML(http.StatusOK, "routeGroup_upload.html", nil)
}

func process_myform_post(c *gin.Context) {
	ftype := c.DefaultPostForm("type", "alert")
	user := c.PostForm("username")
	password := c.PostForm("password")
	hobbys := c.PostFormArray("hobby")
	c.String(http.StatusOK, fmt.Sprintf("type :%s, user %s, password: %s, 爱好是：%v\n", ftype, user, password, hobbys))
}

func process_upload_post(c *gin.Context) {
	file, err := c.FormFile("file")
	// fmt.Printf("%v\n", file)
	if file == nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("没有获取到任何文件！"))
		return
	}
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err:%s", err.Error()))
		return
	}
	c.SaveUploadedFile(file, filepath.Join("store/", file.Filename))
	c.String(http.StatusOK, fmt.Sprintf("upload %s success!", file.Filename))
}

func process_uploads_post(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err:%s", err.Error()))
		return
	}

	filelist := form.File["files"]
	if len(filelist) == 0 {
		c.String(http.StatusBadRequest, fmt.Sprintf("没有获取到任何文件！"))
		return
	}
	for _, f := range filelist {
		if err := c.SaveUploadedFile(f, filepath.Join("store/", f.Filename)); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err:%s", err.Error()))
			return
		}
	}

	c.String(http.StatusOK, fmt.Sprintf("upload %d files success!", len(filelist)))
}
