package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	//まず空白区切りでバラす
	splittedStrings := strings.Fields(s)
	//map作る
	m = make(map[string] int, len(splittedStrings))

	//forで回して、既にある単語ならインクリメント
	for _, v := range splittedStrings {
		val, ok := m[v]
		if ok {
			m[v] = val +1
		} else {
			m[v] = 1
		}

	}
	return
}

func main() {
	wc.Test(WordCount)
}
