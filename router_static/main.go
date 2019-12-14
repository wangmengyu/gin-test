package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
设置静态文件路由
*/
func main() {
	r := gin.Default()

	//设置静态文件路由
	//  cd router_static
	//  go build -o router_static && ./router_static
	//  curl "http://127.0.0.1:8080/assets/a.html"
	r.Static("/assets", "./assets")

	//StaticFs的设定静态文件路由, 于Static一样，需要go build -o 执行
	//curl "http://127.0.0.1:8080/static/b.html"
	r.StaticFS("/static", http.Dir("static"))

	_ = r.Run()

}
