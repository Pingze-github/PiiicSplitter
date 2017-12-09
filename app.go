package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server()
}

func server() {
	gin.SetMode(gin.DebugMode)
	port := 667
	router := gin.Default()
	router.StaticFile("/", "./public/index.html")
	router.POST("/upload", uploadHandler)
	// 启动服务
	fmt.Println(`Server start @`, port)
	router.Run(fmt.Sprintf(":%d", port))
}




