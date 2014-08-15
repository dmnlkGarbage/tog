/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:22
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
			pow(3,2,10),
			pow(3,3,20),
			pow(3,4,30),
	)
}
