package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"golang.org/x/sync/errgroup"
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
	fmt.Println("out m1")
}

func m2(c *gin.Context) {
	fmt.Println("in m2")
	c.Set("name", "杨恺")
}

func m3Closure(use bool) gin.HandlerFunc {
	return func (c *gin.Context) {
		if use {
			c.Next()
		} else {
			fmt.Println("in m3")
			c.Next()
			fmt.Println("out m3")
		}
	}
}

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(m1, m3Closure(true))
	e.GET("/test", m2, func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	e.GET("/index", indexHandler)
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})
	return e
}

func main() {
	server01 := &http.Server{
		Addr: ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := &http.Server{
		Addr: ":8082",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	g.Wait()
}
