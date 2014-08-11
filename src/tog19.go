/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:14
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func main() {
	sum	:= 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

