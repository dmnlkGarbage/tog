package main

import (
	"math"
	"fmt"
)

// interface宣言
// Absを実装してることが必須
type Abser interface {
	Abs() float64
}

func main() {

	// Absを実装してるのでAbserとして扱われる
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// Abser型の変数に入る
	a = f

	//Abserが実装されてないと判断されてる
	v := Vertex{3, 4}
	//a = v

	// ポインタを入れればコンパイル通る
	a = &v

	fmt.Println(a.Abs())
}

// 適当な型
type MyFloat float64

// Absを実装する
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 適当な構造体
type Vertex struct {
	X,Y float64
}

// Absを実装してるように見える
// *Vertexに対してのAbs。*を外すとコンパイル通る
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
