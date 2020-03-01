package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("/index")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}
func m1(c *gin.Context) {
	fmt.Println("in m1")
	// 计时
	start := time.Now()
	c.Next() // 调用剩余处理程序
	//c.Abort() // 不调用该请求的剩余处理程序
	cost := time.Since(start)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "m1",
	//})
	fmt.Println(cost)
}

func m2(c *gin.Context) {
	fmt.Println("in m2")
	c.Set("name", "杨恺")
}

func main() {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(m1)
	e.GET("/test", m2, func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	e.GET("/index", indexHandler)
	e.Run()
}
