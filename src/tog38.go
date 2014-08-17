/**
 * Created with IntelliJ IDEA.
 * User: masaki
 * Date: 2014/08/18
 * Time: 1:02
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell":Vertex{
		11.11,22.32,
	},
	"Log":Vertex{
		333.11,44.22,
	},
}

func main() {
	fmt.Println(m)
}
