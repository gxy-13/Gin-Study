package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 处理multipart form 提交文件时默认内存限制时32mb
	// 可以这样修改  r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1")
		// 将文件保存在本地
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSONP(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

		// 多个文件
		//form , _ := c.MultipartForm()
		//files := form.File["file"]
		//
		//for index , file := range files {
		//	log.Println(file.Filename)
		//	dst := fmt.Sprintf("./tmp/%s_%d",file.Filename,index)
		//	c.SaveUploadedFile(file,dst)
		//}
		//c.JSONP(http.StatusOK,gin.H{
		//	"message" : fmt.Sprintf("%d files uploaded!", len(files)),
		//})

	})
	r.Run()
}
