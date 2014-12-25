package main

import "fmt"



func main() {
	//Map作る
	m := make(map[string] int)
	//代入はputとかじゃなくていい
	m["answer"] = 42
	fmt.Println("the value:", m["answer"])
	//書き換えできる
	m["answer"] = 48
	fmt.Println("the value:", m["answer"])
	//要素の削除はdelete文
	delete(m, "answer")
	fmt.Println("the value:", m["answer"])
	//取得結果をタプルで受けると1個めがvalueで2個めがcontainsKeyのboolean。なければ0とfalseがかえる
	v, ok := m["answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
