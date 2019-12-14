package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" binding:"required"`
	Age     int    `form:"age" binding:"required,gt=10"`
	Address string `form:"address" binding:"required"`
}

/**
结构体验证
curl -X GET 'http://127.0.0.1:8080/testing?name=wang&age=19&address=shanghai'
*/
func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		//对数据进行绑定到结构体
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
		} else {
			c.String(200, "%v", person)
		}
	})

	_ = r.Run()

}
