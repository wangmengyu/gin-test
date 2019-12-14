package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/**
  POST请求，
  获取POST请求的request.body内容 curl -X POST 'http://127.0.0.1:8080/test' -d 'first_name=wang&last_name=mengyu'
*/
func main() {

	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body) //将request body内容全部读取到bytes
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		//回填刚才读取的bytes内容到c.Request.Body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//回填后才可以访问到Post来的数据
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default value")
		c.String(200, "%s, %s, %s", string(bodyBytes), firstName, lastName)

	})
	_ = r.Run()
}
