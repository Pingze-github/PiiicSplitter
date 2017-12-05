package main

import (
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"image/color"
	"image/draw"
	"io"
	"time"
)

func ads(num float32) float32 {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func brightness(c color.Color) float32 {
	r, g, b, _  := c.RGBA()
	R := float32(r) / float32(255)
	G := float32(g) / float32(255)
	B := float32(b) / float32(255)
	return R * 0.3 + G * 0.59 + B * 0.11
}

func main() {
	var gate_brightness_diff float32 = 50 // 行平均亮度差阈值
	var gate_split_min_height_rate float32 = 0.1 // 切块最小高度和宽度比
	var sample_step int = 10
	output_dir := "e:/testdata/splits" // 切块最小高度和宽度比
	start := time.Now().UnixNano()
	fmt.Println("程序启动...")
	path := "e:/testdata/testPiiic.jpg"
	piiicFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer piiicFile.Close()
	piiic, _ := jpeg.Decode(piiicFile)
	width := piiic.Bounds().Max.X
	height:= piiic.Bounds().Max.Y
	gate_split_min_height := int(gate_split_min_height_rate * float32(width))
	fmt.Printf("读取到图片 %s (%dx%d)\n", path, width, height)
	splitYList := [] int {0}
	fmt.Println("正在查询分割位置...")
	sampleNum := width / sample_step
	for y := 1; y < height; y++ {
		var diff float32 = 0
		for x := 0; x < width; x += sample_step {
			diff += ads(brightness(piiic.At(x, y)) - brightness(piiic.At(x, y - 1)))
		}
		diffAvg := diff / float32(sampleNum)
		if diffAvg > gate_brightness_diff {
			fmt.Printf("查找到分割位置 %d, 亮度差值 %f\n", y, diffAvg)
			splitYList = append(splitYList, y)
		}
	}
	splitYList = append(splitYList, height)
	index := 1

	channel := make(chan int, 10)
	quit := make(chan int)

	finished := 0
	jpegSave := func (splitFile io.Writer, split image.Image, num int, splitPath string) {
		time.Sleep(1e9)
		jpeg.Encode(splitFile, split, &jpeg.Options{Quality: 100})
		fmt.Printf("保存 分块%d 到 %s\n", num, splitPath)
		<- channel
		finished ++
		if finished == index - 1 {
			fmt.Printf("程序执行完毕，耗时 %d ms \n", (time.Now().UnixNano() - start) / 1e6)
			<- quit
		}
	}

	for i := 1; i < len(splitYList); i++ {
		if splitYList[i] - splitYList[i-1] > gate_split_min_height {
			splitPath := output_dir + fmt.Sprintf("/%d.jpg", index)
			splitFile, err := os.OpenFile(splitPath, os.O_RDWR|os.O_CREATE, 0664)
			if err != nil {
				panic(err)
			}
			defer splitFile.Close()
			rect := image.Rect(0, splitYList[i-1], width, splitYList[i])
			split := image.NewRGBA(rect)
			draw.Draw(split, rect, piiic, rect.Min, draw.Src) //截取图片的一部分
			channel <- 1
			go jpegSave(splitFile, split, index, splitPath)
			index ++
		}
	}
	quit <- 1
}


