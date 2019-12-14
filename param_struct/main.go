package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

/**
  绑定请求的数据到struct结构体中
*/

type Person struct {
	Name     string    `form:"name"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
	Address  string    `form:"address"`
}

/**
  GET
  curl -X GET 'http://127.0.0.1:8080/testing?name=wang&address=shanghai&birthday=1989-01-24'
  POST
  curl -X POST 'http://127.0.0.1:8080/testing' -d 'name=wang&address=shanghai&birthday=1989-01-24'
  JSON
  curl -H "Content-Type:application/json" -X POST 'http://127.0.0.1:8080/testing' -d '{"name":"wang","address":"shanghai","birthday":"1989-01-24 13:33:33"}'

*/
func main() {
	//无论从GET或者POST都兼容绑定数据
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	_ = r.Run()
}

func testing(c *gin.Context) {
	//新建结构体对象
	var person Person
	//根据不用类型的CONTENT-TYPE来做不同的binding操作
	if err := c.ShouldBind(&person); err != nil {
		c.String(200, err.Error())
		c.Abort()
	} else {
		c.String(200, "%v", person)
	}

}
