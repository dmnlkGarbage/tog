package main

import (
	"math"
	"fmt"
)
// 組み込みパッケージに対してメソッドを生やすために型化
type MyFloat float64

// 拡張メソッド
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func main() {
	v := MyFloat(-math.Sqrt2)
	fmt.Println(v.Abs())
}
