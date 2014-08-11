/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:29
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64{
	z := float64(1)
	prevZ := z + 1
	for math.Abs(prevZ - z) > 0.0000001 {
		prevZ = z
		z = z - ((z*z - x) / 2 * z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}

