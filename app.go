package main

import (
	"fmt"
	"runtime"
	"github.com/gin-gonic/gin"
)

func server() {
	gin.SetMode(gin.DebugMode)
	port := 667
	router := gin.Default()
	router.StaticFile("/", "./public/index.html")
	router.POST("/upload", uploadHandler)
	// 启动服务
	fmt.Println(`Server start @`, port)
	fmt.Println(`OS:`, OS)
	router.Run(fmt.Sprintf(":%d", port))
}

func main() {
	server()
}

var (
	OS string
	UPLOAD_DIR string
	OUTPUT_DIR string
)
func init() {
	OS = runtime.GOOS
	if (OS == "windows") {
		UPLOAD_DIR = "E:/raid/piiic/upload"
		OUTPUT_DIR = "E:/raid/piiic/output"
	} else {
		UPLOAD_DIR = "/var/raid/piiic/upload"
		OUTPUT_DIR = "/var/raid/piiic/output"
	}
}


