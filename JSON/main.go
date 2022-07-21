package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() // 返回默认的路由引擎

	r.GET("/json", func(c *gin.Context) {
		// 方法1 使用map
		//data := map[string]interface{}{
		//	"name":    "gxy",
		//	"message": "hello,world",
		//	"age":     12,
		//}
		// Gin 将map[string]interface{} 设置了成了自定义类型H
		data := gin.H{
			"name":    "gxy",
			"message": "hello,world",
			"age":     12,
		}

		c.JSONP(http.StatusOK, data)
	})

	// 方法2 结构体 灵活使用tag对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"User"`
		Message string
		Age     int
	}
	r.GET("/struct_json", func(c *gin.Context) {
		data := msg{
			"gxy",
			"hello,golang!",
			15,
		}
		c.JSONP(http.StatusOK, data)
	})

	// 启动服务
	r.Run(":8081")
}
