package requestProcess

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile() {
	r := gin.Default()
	r.LoadHTMLFiles("html/upload.html")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.POST("/upload_post", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err:%s\n", err.Error()))
			return
		}
		c.SaveUploadedFile(file, filepath.Join("store/", file.Filename))

		c.String(http.StatusOK, fmt.Sprintf("upload %s success", file.Filename))
	})

	//上传多个文件，表单限制大小为8M,默认为32M
	r.MaxMultipartMemory = 8 << 20
	r.POST("/uploads_post", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err:%s\n", err.Error()))
			return
		}
		//获取所有图片
		filelist := form.File["files"]
		//遍历所有图片
		for _, f := range filelist {
			if err := c.SaveUploadedFile(f, filepath.Join("store/", f.Filename)); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err:%s", err.Error()))
				return
			}
		}
		//保存文件后返回前端
		c.String(http.StatusOK, fmt.Sprintf("upload %d files success!", len(filelist)))
	})
	r.Run(":8080")
	return
}
