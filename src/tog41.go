package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	//map作る
	m := make(map[string] int)
	//まず空白区切りでバラす
	val := strings.Fields(s)
	//forで回して、既にある単語ならインクリメント
	for _, v := range val {
		val, ok := m[v]
		if ok {
			m[v] = val +1
		} else {
			m[v] = 1
		}

	}
	return m
}

func main() {
	wc.Test(WordCount)
}
