package main

import (
	"fmt"
	"time"
)

func main() {
	running := true
	f := func() {
		for running {
			fmt.Println("running inside go ")
			time.Sleep(1 * time.Second)
		}
		fmt.Println(".")
		fmt.Println("running outside go")
	}

	go f()
	go f()
	go f()
	time.Sleep(2 * time.Second)
	running = false
	time.Sleep(1 * time.Second)
	fmt.Println("go exit")
}
