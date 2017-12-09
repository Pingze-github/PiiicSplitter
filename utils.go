package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"errors"
	"io/ioutil"
	"net/http"
	"github.com/fatih/color"
)

type RetBody struct {
	Status int
	String string
	Html string
	Raw []byte
	ContentType string
	Code int
	Msg string
	Data interface{}
}

func CReturn(c *gin.Context, ret RetBody) {
	if ret.Status == 0 {
		ret.Status = http.StatusOK
	}
	if ret.String != "" {
		c.Data(ret.Status, "text/plain", []byte(ret.String))
	} else if ret.Html != "" {
		c.Data(ret.Status, "text/html", []byte(ret.Html))
	} else if ret.Raw != nil {
		if ret.ContentType == "" {
			ret.ContentType = "text/plain"
		}
		c.Data(ret.Status, ret.ContentType, []byte(ret.Raw))
	} else {
		type JSONRet struct {
			Code int `json:"Code"`
			Msg string `json:"Msg"`
			Data interface{}
		}
		if ret.Msg == "" {ret.Msg = "请求成功"}
		if ret.Data == "" {ret.Data = gin.H{}}
		jsonRet := JSONRet{Code: ret.Code, Msg: ret.Msg, Data: ret.Data}
		c.JSON(ret.Status, jsonRet)
	}
}

func request(url string) (content string, bytes []byte, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil);
	req.Header.Set("Referer", "http://www.bilibili.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	res, errRes := client.Do(req)
	if errRes != nil {
		fmt.Println(errRes)
		err = errors.New("http get failed")
		return
	}
	defer res.Body.Close()
	bytes, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		fmt.Println(errRead)
		err = errors.New("res body read failed")
		return
	}
	content = string(bytes)
	return
}

func warn(things interface{}) {
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset()
	fmt.Println(things)
}
