/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:50
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func main() {
	p := [] int{2, 3, 5, 7 ,11 ,13}
	fmt.Println("p == ", p)



	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
}

