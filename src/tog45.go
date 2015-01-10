package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OSX")
		//自動でbreakするので下におろしたい場合はfallthroughを書く
		//fallthrough
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
}
