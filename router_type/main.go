package main

import "github.com/gin-gonic/gin"

/**
  各种路由方法的使用
*/
func main() {
	r := gin.Default()
	//get
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")

	})
	//post, curl -X POST "http://127.0.0.1:8080/post"
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")

	})
	//delete curl -X DELETE "http://127.0.0.1:8080/delete"
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	//any request url method support
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
