package main

import (
	"time"
	"fmt"
)

// 共通の例外クラスをつくる感じ

// 自分でerror型をつくる
type MyError struct {
	When time.Time
	What string
}

// errorというinterfaceは Error()関数を実装していれば良い
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// errorを返す関数
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		// fmt.Printlnはerror型を渡すとError()関数を実行するらしい
		fmt.Println(err)
	}
}
