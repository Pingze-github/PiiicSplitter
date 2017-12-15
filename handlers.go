package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
)

func uploadHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["uploads"]
	var splitPaths []string
	for _, file := range files {
		fmt.Println("接收到上传文件",file.Filename)
		uploadPath := UPLOAD_DIR + "/" + file.Filename
		outDir := OUTPUT_DIR
		c.SaveUploadedFile(file, uploadPath)
		splitPaths = imgSplit(uploadPath, outDir)
	}
	for i, splitPath := range splitPaths {
		splitPaths[i] = strings.Replace(splitPath, OUTPUT_DIR, "/split", 1)
	}
	CReturn(c, RetBody{Data: splitPaths})
}
