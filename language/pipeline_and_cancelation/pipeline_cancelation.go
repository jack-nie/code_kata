package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)
	out := sq(in)
	fmt.Println(<-out)
	fmt.Println(<-out)
	if n, ok := <-out; !ok {
		fmt.Println(n)
	}
	fmt.Println(<-out)
	in1 := gen(3, 4)
	c1 := sq(in1)
	c2 := sq(in1)
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
