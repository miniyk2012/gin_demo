package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/miniyk2012/gin_demo/tools"
	"html/template"
	"log"
	"net/http"
	"path"
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
	e.Static("/www", "../static")
	e.LoadHTMLGlob("../templates/**/*")
}

func jsonDemo(e *gin.Engine) {
	e.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "杨恺",
			"Age":  18,
		})
	})

	e.GET("json_more", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Name string `json:"name"`
			Age  int
		}{
			Name: "杨恺",
			Age:  19,
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
			Label:            &label,
			Reps:             reps,
			Type:             &type_,
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
			Label:            &label,
			Reps:             reps,
			Type:             &type_,
			XXX_unrecognized: []byte("123"),
		}
		c.JSON(http.StatusOK, data)
	})
}

func paramsDemo(e *gin.Engine) {
	e.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	e.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	e.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
}

func bindParam(e *gin.Engine) {
	// Binding from JSON
	type Login struct {
		User     string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"Passwords" binding:"required"`
	}

	e.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 绑定form表单示例 (user=q1mi&password=123456)
	e.POST("/login", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, login)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?username=q1mi&password=123456)
	e.GET("/loginQuery", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

}

func uploadFiles(e *gin.Engine) {
	e.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		dst := path.Join("..", "data", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	e.POST("/uploads", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["f1"]

		for _, file := range files {
			log.Println(file.Filename)
			dst := path.Join("..", "data", file.Filename)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.Redirect(http.StatusMovedPermanently, "/login") // 重定向
	})
}

func routeHandleContext(e *gin.Engine) {
	e.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "a",
		})
	})
	e.GET("/b", func(c *gin.Context) {
		c.Request.URL.Path = "/a"
		e.HandleContext(c)
	})
}

func routeDemo(e *gin.Engine) {
	e.Any("/test", func(c *gin.Context) {
		var ret = c.Request.Method
		c.JSON(http.StatusOK, gin.H{
			"message": ret,
		})
	})
	e.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}

func routeGroupDemo(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user/index",
			})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user/login",
			})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user/postlogin",
			})
		})
		subShop := userGroup.Group("xx")
		subShop.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "xx/oo",
			})
		})

	}
	shopGroup := e.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "shop/index",
			})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "shop/cart",
			})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "shop/checkout",
			})
		})
	}
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
	paramsDemo(e)
	bindParam(e)
	uploadFiles(e)
	routeHandleContext(e)
	routeDemo(e)
	routeGroupDemo(e)
	e.Run(":9091")
}
