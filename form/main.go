package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取form表单提交的参数
type user struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	// 加载登录页面
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 处理login
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		data := user{
			username,
			password,
		}
		c.JSONP(http.StatusOK, data)
	})
	r.Run(":8090")
}
