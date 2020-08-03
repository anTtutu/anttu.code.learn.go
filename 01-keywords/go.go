package main

import (
	"fmt"
	"time"
)

/**
 * @Author: Anttu
 * @Date: 2020-02-11 23:26
 */
func main() {
	go testName()
	fmt.Println("main method")
	time.Sleep(time.Second)
}

func testName () {
	fmt.Println("test method")
}
