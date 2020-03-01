package main

import (
	"github.com/gin-gonic/gin"
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
func main() {
	e := gin.Default()
	udf(e)
	loadTemplates(e)
	demo1(e)
	demo2(e)
	demo3(e)
	demo4(e)
	e.Run(":9091")
}
