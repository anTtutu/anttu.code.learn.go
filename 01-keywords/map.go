package main

import "fmt"

/**
 * @Author: Anttu
 * @Date: 2020-02-12 00:17
 */
func main() {
	//1、声明map
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	//2、遍历map
	for k, v := range scoreMap {
		fmt.Printf("key[%s] - value[%d]\n", k, v)
	}

	//3、遍历map只要key或者value
	for k := range scoreMap {
		fmt.Printf("key[%s]\n", k)
	}

	for _, v := range scoreMap {
		fmt.Printf("value[%d]\n", v)
	}

	//4、判断某个key是否存在  特殊写法：value, ok := map[key]
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//5、删除某个key
	//将小明:100从map中删除
	delete(scoreMap, "小明")
	for k,v := range scoreMap{
		fmt.Printf("key[%s] - value[%d]\n", k, v)
	}
}
