package main

import (
	"net/http"
	"fmt"
)

func main() {
	var s String = "aaaaa"
	http.ListenAndServe("localhost:4000", s)
}


type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}


type Struct struct {
	Greeting string
	Punct    string
	Who      string
}
