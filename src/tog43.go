package main

import "fmt"
//関数を戻り値に持つ関数
func adder() func(int) int {
	//ずっと保持してる
	sum := 0
	//返す関数。外側の変数を内側から参照できる？
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2 * i),
		)
	}
}

