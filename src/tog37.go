/**
 * Created with IntelliJ IDEA.
 * User: masaki
 * Date: 2014/08/18
 * Time: 0:44
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make (map[string]Vertex)
	m["Bell labs"] = Vertex{
		40.22221, -212.211,
	}
	fmt.Println(m["Bell labs"])
}



