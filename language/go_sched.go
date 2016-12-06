package main

import (
	"fmt"
	"runtime"
	"sync"
)

func test(name string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(name)
	}
}

func main() {
	var wg sync.WaitGroup
	go test("jack")
	test("bone")
	wg.Wait()
}
