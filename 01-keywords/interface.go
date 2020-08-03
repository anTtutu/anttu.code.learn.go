package main

import "fmt"

type Dog interface {
	Category()
}

type Ha struct {
	Name string
}

func (h Ha) Category() {
	fmt.Println("狗子")
}

func main() {
	h := Ha{"二哈"}
	h.Category()
	test(h)
}

func test(a Dog) {
	fmt.Println("成功调用啦")
}
