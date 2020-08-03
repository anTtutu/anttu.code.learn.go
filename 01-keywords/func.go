package main

import "fmt"

/**
 * @Author: Anttu
 * @Date: 2020-02-11 23:10
 */
func main() {
}

//1、带返回值的函数
func add(a, b int) int {
	return a + b
}

//2、多个返回值的函数
func vals() (int, int) {
	return 3, 7
}

//3、没返回值的函数
func query () {
	fmt.Println("query success.")
}
