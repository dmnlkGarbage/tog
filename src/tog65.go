package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// バッファ以上に投げるとdeadlock起きて死ぬ
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)
}

