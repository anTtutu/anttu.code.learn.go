package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	path := "toy\\testImages\\cs.jpg"

	var o convert.Options
	o.Ratio = 0.3
	o.Colored = true

	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString(path, &o))
}
