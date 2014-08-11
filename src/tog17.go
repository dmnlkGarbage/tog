/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:10
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}


