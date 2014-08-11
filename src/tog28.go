/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:44
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}
	q = &Vertex{1, 3}
	r = Vertex{X: 1}
	s = Vertex{}
)

func main() {
	fmt.Println(p, q, r, s)
}

