package main

import (
	"fmt"
	"time"
)

func work() {
	exitChan := make(chan int)
	go task1(exitChan)
	go task2(exitChan)
	time.Sleep(5 * time.Second)
	close(exitChan)
	time.Sleep(5 * time.Second)
}
func task1(exitChan chan int) {
	<-exitChan
	fmt.Println("task1 exiting")
}

func task2(exitChan chan int) {
	<-exitChan
	fmt.Println("task2 exiting")
}

func main() {
	work()
}
