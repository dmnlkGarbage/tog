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
	hoge := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		fuga := make([]uint8, dx)
		for j := dx; j < dx; j++ {
			fuga[j] = uint8(i * j)
		}
		hoge[i] = fuga
	}
	return hoge
}

func main() {
	pic.Show(Pic)
}

