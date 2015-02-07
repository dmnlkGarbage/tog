package main

import (
	"net/http"
	"fmt"
)

// 構造体
type Hello struct {}


// ListernAndServeにServeHTTPというメソッドを持つHandler interfaceを渡せる
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!")
}


func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
