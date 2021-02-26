package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	path := "toy\\testImages\\cs.jpg"

	var o convert.Options
	o.Ratio = 0.3
	//o.Colored = true

	converter := convert.NewImageConverter()

	// 目前输出到控制台
	//fmt.Print(converter.ImageFile2ASCIIString(path, &o))

	// 输出到txt，彩色需要研究下。输出txt暂时不支持彩色了，彩色需要用html
	str := converter.ImageFile2ASCIIString(path, &o)

	fileName := "test.txt"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	dstFile.WriteString(str)
}
