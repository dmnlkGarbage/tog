/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:47
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}

