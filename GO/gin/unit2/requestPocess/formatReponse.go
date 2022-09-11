package requestProcess

import (
	"github.com/gin-gonic/gin"
)

func ResponseAsFormat() {
	r := gin.Default()

	//json格式响应
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "200", "msg": "someJSON response"})
	})
	//结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}

		msg.Name = "root"
		msg.Message = "some message"
		msg.Number = 200

		c.JSON(200, msg)
	})
	//xml响应
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "some XML message"})
	})

	//yaml响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan", "age": "9000"})
	})

	//protobuf格式，谷歌开发的高效存储读取工具
	/*
		r.GET("/someProtoBuf", func(c *gin.Context) {
			resp := []int64{int64(1), int64(2)}
			label := "label"
			data := &protoexample.Test{
				Label: &label,
				Resp:  &resp,
			}

			c.ProtoBuf(200, data)
		})
	*/

	r.Run(":8080")
}
