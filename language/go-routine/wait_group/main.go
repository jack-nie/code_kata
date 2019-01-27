package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	stop := make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(stop <-chan bool) {
			defer wg.Done()
			consumer(stop)
		}(stop)
	}
	waitForSignal()
	close(stop)
	fmt.Println("Waiting.....")
	wg.Wait()
	fmt.Println("All stopped")
}

func consumer(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Println("running....")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func waitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}
