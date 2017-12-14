package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func uploadHandler(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["uploads"]

	for _, file := range files {
		fmt.Println("接收到上传文件",file.Filename)

		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))


	// CReturn(c, RetBody{Data: "123"})
}
