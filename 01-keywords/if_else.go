package main

import "fmt"

/**
 * @Author: Anttu
 * @Date: 2020-02-12 16:52
 */
func main() {
	//1、if条件判断基本写法
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//2、if条件判断特殊写法
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
