package main

import (
	"fmt"
	"math"
)

// 構造体定義
type Vertex struct {
	X, Y float64
}

// ポインタへのメソッドだと中身を変更できるけど、値へのメソッドだと変更できない
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 値を読むだけなのでポインタだろうが関係ない
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
