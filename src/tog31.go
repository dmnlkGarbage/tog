/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 20:53
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	fmt.Println("p[:3] ==", p[:3])

	fmt.Println("p[4:] ==", p[4:])
}

