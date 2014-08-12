/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/12
 * Time: 14:01
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

var pow = []int {1,2,4,8,16,32,64,128}

func main() {
	for i, v:=range pow {
		fmt.Printf("2**%d = %d ¥n", i, v)
	}
}

