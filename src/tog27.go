/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:41
 * To change this template use File | Settings | File Templates.
 */
package main


import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}


func main() {
	p := Vertex{2, 10}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}
