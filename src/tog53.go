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
