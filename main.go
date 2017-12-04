package main

import (
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"image/color"
	"image/draw"
)

func brightness(c color.Color) float32 {
	r, g, b, _ := c.RGBA()
	return float32(0.3 * float32(r) + float32(g) * 0.59 + float32(b) * 0.11)
}

func main() {
	var gate_brightness_diff float32 = 5000 // 行平均亮度差阈值
	var gate_split_min_height_rate float32 = 0.1 // 切块最小高度和宽度比
	output_dir := "e:/testdata/splits" // 切块最小高度和宽度比
	fmt.Println("程序启动")
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
	for y := 1; y < height; y++ {
		var diff float32 = 0
		for x := 0; x < width; x++ {
			diff += brightness(piiic.At(x, y)) - brightness(piiic.At(x, y - 1))
		}
		diffAvg := diff / float32(width)
		if diffAvg > gate_brightness_diff {
			fmt.Printf("查找到分割位置 %d, 亮度差值 %f\n", y, diffAvg)
			splitYList = append(splitYList, y)
		}
	}
	splitYList = append(splitYList, height)
	index := 1
	for i := 1; i < len(splitYList); i++ {
		if splitYList[i] - splitYList[i-1] > gate_split_min_height {
			splitPath := output_dir + fmt.Sprintf("/%d.jpg", index)
			splitFile, err := os.Open(splitPath)
			if err != nil {
				panic(err)
			}
			defer splitFile.Close()
			split := image.NewRGBA(image.Rect(0, width, splitYList[i-1], splitYList[i]))
			draw.Draw(split, piiic.Bounds().Add(image.Pt(10, 10)), piiic, piiic.Bounds().Min, draw.Src) //截取图片的一部分
			jpeg.Encode(splitFile, split, nil)
			fmt.Printf("保存 分块%d 到 %s\n", index, splitPath)
			index ++
		}
	}
}
