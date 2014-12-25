package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string] int)
	val := strings.Fields(s)

	for _, v := range val {
		m[v] = 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
