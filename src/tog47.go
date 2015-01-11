package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	// == switch true
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
	
}

