package main

import "github.com/gin-gonic/gin"

/**
curl -X GET 'http://127.0.0.1:8080/test'
*/
func main() {
	r := gin.Default()
	//需要模仿Logger(), Recovery() 等系统内部的中间件的实现格式
	r.Use(IpAuth())

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})

	_ = r.Run()

}

/**
IP 验证方法
*/
func IpAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.2",
		}
		flag := false

		//当前访问者的IP
		clientIp := c.ClientIP()

		//检查客户端IP是否在白名单内
		for _, ip := range ipList {
			if clientIp == ip {
				flag = true
				break
			}
		}

		//不在白名单内，进行报错
		if flag == false {
			c.String(401, "%s not in iplist", clientIp)
			c.Abort()
		}

	}

}
