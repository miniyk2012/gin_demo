package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/miniyk2012/gin_demo/tools"
	"html/template"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
func demo1(e *gin.Engine) {
	e.GET("/hello", sayHello)

	e.GET("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "GET",
		})
	})
	e.POST("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "POST",
		})
	})
	e.PUT("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "PUT",
		})
	})
	e.DELETE("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "DELETE",
		})
	})
}

func demo2(e *gin.Engine) {
	e.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	e.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})
}

func demo3(e *gin.Engine) {

	e.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]string{
			"url":         "<a href='https://liwenzhou.com'>李文周的博客</a>",
			"currentPath": tools.GetCurrentPath(),
		})
	})
}

func demo4(e *gin.Engine) {
	e.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
}

func udf(e *gin.Engine) {
	e.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
}

func loadTemplates(e *gin.Engine) {
	e.Static("/www", "./static")
	e.LoadHTMLGlob("templates/**/*")
}

func jsonDemo(e *gin.Engine) {
	e.GET("/json", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "杨恺",
			"Age": 18,
		})
	})

	e.GET("json_more", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Name string `json:"name"`
			Age int
		}{
			Name: "杨恺",
			Age: 19,
		})
	})
	e.GET("/someXML", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	e.GET("/moreXML", func(c *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})

	e.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})

	e.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2), int64(5679988)}
		label := "test"
		type_ := int32(65)
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		var data = &protoexample.Test{
			Label: &label,
			Reps:  reps,
			Type:  &type_,
			XXX_unrecognized: []byte("123"),
		} // 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
	e.GET("/someProtoBufJson", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2), int64(5679988)}
		label := "test"
		type_ := int32(65)
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
			Type : &type_,
			XXX_unrecognized: []byte("123"),
		}
		c.JSON(http.StatusOK, data)
	})
}
func main() {
	e := gin.Default()
	udf(e)
	loadTemplates(e)
	demo1(e)
	demo2(e)
	demo3(e)
	demo4(e)
	jsonDemo(e)
	e.Run(":9091")
}
