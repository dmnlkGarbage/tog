/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/12
 * Time: 14:11
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func main() {
	pow := make([]int , 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d¥n ", value)
	}

}

