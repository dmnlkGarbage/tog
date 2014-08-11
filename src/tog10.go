/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 19:35
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4
	y = sum - x
	return
}

func main() {
	fmt.Println(split(10))
}
