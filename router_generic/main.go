package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//所有请求/user/开头的都打到一个回调函数里
	//curl -X GET "http://127.0.0.1:8080/user/xxx"
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	_ = r.Run()

}
