package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type Booking struct {
	//checkIn时间必须大于现在的时间，自定义验证方法实现 bookabledate
	CheckIn time.Time `form:"check_in" binding:"required,bookableDate" time_format:"2006-01-02"`

	//checkOut时间必须大于checkIn时间
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

/**
curl -X GET 'http://127.0.0.1:8080/booking?check_in=2019-12-10&check_out=2019-12-08'
*/
func main() {
	r := gin.Default()
	//注册自定义方法
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if errV := v.RegisterValidation("bookableDate", bookableDate); errV != nil {
			panic(errV)
		}
	}
	r.GET("/booking", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "ok!", "booking": b})
	})
	_ = r.Run()

}

/**
  可预约的日期，自定义验证方法
  在文档页面搜索：custom validate 搜索范例
*/

func bookableDate(fl validator.FieldLevel) bool {
	if checkInTime, ok := fl.Field().Interface().(time.Time); ok {
		if checkInTime.Unix() > time.Now().Unix() {
			//如果预定时间>现在的时间
			return true
		}
	}
	return false
}
