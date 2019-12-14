package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//指定LOG的文件位置
	f, _ := os.Create("gin.log")
	//重定向默认输出和错误输出到指定的文件
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	//创建gin实例
	r := gin.New()
	//使用logger中间件
	r.Use(gin.Logger())

	//使用recovery中间件，如果期间出现panic,不会中断Main进程，否则会中断
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default name")
		//panic("test panic")
		c.String(200, "%s", name)
	})

	_ = r.Run()

}
