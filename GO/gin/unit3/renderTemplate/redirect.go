package renderTemplate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect2Another() {
	r := gin.Default()

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	r.Run(":8080")
}
