package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := userInfo{
		//	username,
		//	password,
		//}
		var u userInfo // 声明一个userinfo类型的变量u
		// go 中的方法都是值传递，所以需要传递地址
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSONP(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		var u userInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSONP(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
	r.Run()
}
