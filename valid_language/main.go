package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type Person struct {
	Name    string `form:"name" validate:"required"`
	Age     int    `form:"age" validate:"required,gt=10"`
	Address string `form:"address" validate:"required"`
}

/**
  使用英文的验证器：
curl -X GET 'http://127.0.0.1:8080/testing?name=wang&address=shanghai&age=1&locale=en'
  使用中文的验证器
 curl -X GET 'http://127.0.0.1:8080/testing?name=wang&address=shanghai&age=1&locale=cn'
  或者省略语言设置也会使用中文的
 curl -X GET 'http://127.0.0.1:8080/testing?name=wang&address=shanghai&age=1'
*/
func main() {

	//创建验证器
	validate := validator.New()
	zh := zh2.New()
	en := en.New()
	//创建翻译器
	uni := ut.New(zh, en)

	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {

		//从GET请参数中获取当前的语言
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(validate, trans)
		}

		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}

		//需要用Validate.Struct来验证结构体
		if err := validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceError := []string{}
			for _, e := range errs {
				sliceError = append(sliceError, e.Translate(trans))
			}
			c.String(500, "%v", sliceError)
			c.Abort()
			return
		}

		c.String(200, "%v", person)
	})
	_ = r.Run()

}
