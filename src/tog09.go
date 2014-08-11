/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 19:32
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hoge", "fuga")
	fmt.Println(a, ":", b)
}
