package main

import "github.com/gin-gonic/gin"

/**
  获取GET参数
*/
func main() {
	//获取一个first_name参数，和一个last_name,其中last_name要有默认值
	//curl -X GET "http://127.0.0.1:8080/test?first_name=wang"
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		//get first_name
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "default last name")
		c.String(200, "%s %s", firstName, lastName)

	})

	_ = r.Run()

}
