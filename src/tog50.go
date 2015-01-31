package main

import (
	"math"
	"fmt"
)
// 構造体定義
type Vertex struct {
	X, Y float64
}

// 上で定義した構造体に関数を実装する。
// 関数名の前がそれ
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
