package main

import "github.com/gin-gonic/gin"

/**
参数作为URL
*/
func main() {
	r := gin.Default()
	//指定一个获取name参数和一个id参数的URI
	//测试：curl -X GET "http://127.0.0.1:8080/x/1"
	r.GET("/:name/:id", func(c *gin.Context) {
		//返回GET中获得的name和id参数值
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})
	_ = r.Run()
}
