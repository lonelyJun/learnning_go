package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	var size = 300
	//生成灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//纯白图片
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pic.SetGray(i, j, color.Gray{255})
		}
	}
	//描绘正弦函数
	for x := 0; x < size; x++ {
		angle := float64(x) * 2 * math.Pi / float64(size)
		y := math.Sin(angle)*float64(size)/2 + float64(size)/2
		//绘制轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	//写入图片文件
	file, err := os.Create("sin.png") //创建了一个sin.png 此时file为文件指针（猜测）

	if err != nil {
		log.Fatal(err) //如果出错，则记入日志
	}

	png.Encode(file, pic) //编码png，参数1为写入的文件指针，参数2为数据

	file.Close() //关闭文件

}
