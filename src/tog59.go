package main

import (
	"image"
	"fmt"
)

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(img.Bounds())
	fmt.Println(img.At(0,0).RGBA())
}

