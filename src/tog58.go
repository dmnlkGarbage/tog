package main

import (
	"net/http"
	"fmt"
)

func main() {
	var s String = "aaaaa"

	http.Handle("/struct", &Struct{"hello", ":", "me"})
	http.Handle("/str", s)
	http.ListenAndServe("localhost:4000", nil)
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

func (s Struct) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, s.Greeting + s.Punct + s.Who)
}
