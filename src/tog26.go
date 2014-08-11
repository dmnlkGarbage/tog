/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:38
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

type V3etex struct {
	X int
	Y int
}

func main() {
	v := Vertex{2,3}
	v.X = 4
	fmt.Print(v.X)
}
