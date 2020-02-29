package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
func demo1(r *gin.Engine) {
	r.GET("/hello", sayHello)

	r.GET("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "GET",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "POST",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"method": "DELETE",
		})
	})
}

func demo2(r *gin.Engine) {
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})
}

func demo3(router *gin.Engine) {
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles("./index.tmpl")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
}

func main() {
	r := gin.Default()
	demo1(r)
	demo2(r)
	demo3(r)
	r.Run(":9091")
}
