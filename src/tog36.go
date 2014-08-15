/**
 * Created with IntelliJ IDEA.
 * User: masaki
 * Date: 2014/08/16
 * Time: 0:24
 * To change this template use File | Settings | File Templates.
 */
package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
	}

	for y := 0; y < dy; y++ {
		for x :=0; x < dx; x++ {
			image[y][x] = uint8((x + y) / 2)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}

