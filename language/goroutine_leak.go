package main

import (
	"fmt"
	"runtime"
)

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}

func countGoroutine() int {
	return runtime.NumGoroutine()
}

func main() {
	fmt.Printf("Count goroutines: %d", countGoroutine())
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Printf("Count goroutines: %d", countGoroutine())
}
