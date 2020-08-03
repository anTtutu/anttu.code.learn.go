package main

import "fmt"

/**
 * @Author: Anttu
 * @Date: 2020-02-11 22:49
 */
func main() {
	//写法1：
	switch i := 6; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数5")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("other", i)
	}

	//写法2：
	finger := 2
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效的数字")
	}

	//写法3：
	age := 30
	switch {
	case age < 24:
		fmt.Println("学习阶段")
	case age >= 24 && age < 60:
		fmt.Println("工作阶段")
	case age > 60:
		fmt.Println("退休阶段")
	default:
		fmt.Println("活着真好")
	}
}
