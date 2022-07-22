package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// gin 中间件 实现功能类似java aop 实现的功能，权限，日志
// Gin的中间件必须是一个gin.HandlerFunc 类型，
//type HandlerFunc func(*Context)
func indexHandler(c *gin.Context) {
	name := c.MustGet("name")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "index",
		"name": name,
	})
}

// 定义一个中间件m1 统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1ing.......")
	c.Set("name", "gxy")
	// go funcXX(c.Copy()) 如果要在中间件中使用goroutine 只能传递c的拷贝
	// 计时
	start := time.Now()
	c.Next() // 调用后续处理函数
	// c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\v", cost)
}

//func authMiddleware(c *gin.Context) {
//	// 是否登录判断
//
//	// if是登录用户
//	// c.Next()
//	// else
//	// c.Abort()
//}
func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 链接数据库
	// 或者其它准备工作

	return func(c *gin.Context) {
		// 存放具体的逻辑
		if doCheck {
			// 是否登录判断
			//
			//	// if是登录用户
			//	// c.Next()
			//	// else
			//	// c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default() // 默认使用logger()和recovery()的中间件
	// gin.New()

	r.Use(m1) // 全局注册中间件函数m1

	//GET(relativePath string, handlers ...HandlerFunc)
	r.GET("/index", m1, indexHandler)

	// 给路由组注册中间件方法1
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}

	// 给路由组注册中间件方法2
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	}

	r.Run()
}
