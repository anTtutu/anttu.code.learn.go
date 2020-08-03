package main

import "fmt"

/**
 * @Author: Anttu
 * @Date: 2020-02-11 22:46
 */
func main() {
	//跳出循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 6 {
			break
		}
	}
}
