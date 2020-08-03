package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a(){
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Printf("A %d \n", i)
	}
}

func b(){
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		fmt.Printf("B %d \n", i)
	}
}

/**
 * @Author: Anttu
 * @Date: 6/3/2020 20:46
 */
func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
