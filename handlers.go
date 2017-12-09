package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"log"
	"net/http"
)

func uploadHandler(c *gin.Context) {

	name := c.PostForm("uploads")
	log.Println("name",name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		warn(err)
		c.String(http.StatusBadRequest, "Interal Server Error")
		return
	}
	filename := header.Filename

	log.Println(file, err, filename)

	out, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
	}
	c.String(http.StatusCreated, "upload successful")


	CReturn(c, RetBody{Data: "123"})
}
