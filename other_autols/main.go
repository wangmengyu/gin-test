package main

import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/autotls"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})

	//自动化证书配
	autotls.Run(r, "www.itpp.tk")

}
