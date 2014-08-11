/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:00
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

const (
	Pi = 3.14
)

func main () {
	const World = "world"
	fmt.Println("hello", World)
	fmt.Println("happy", Pi, "day")

	const True = true
	fmt.Println("Go rules?", True)
}

