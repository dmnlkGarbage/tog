/**
 * Created with IntelliJ IDEA.
 * User: m00217-kawaguchi
 * Date: 2014/08/11
 * Time: 19:03
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to Play Ground")
	fmt.Println("The time is", time.Now())

	fmt.Println(os.Open("fileName"))

	fmt.Println(net.Dial("tcp", "google.com"))
}
