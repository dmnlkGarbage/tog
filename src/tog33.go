/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/12
 * Time: 13:58
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("slice is nil")
	}
}

