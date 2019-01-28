package main

import (
	"fmt"
	"sync"
	"bytes"
)

func main() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}

	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("received %d\n", result)
		}
		fmt.Println("Done receive!")
	}

	results := chanOwner()
	consumer(results)

	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer

		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}

		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	byteData := []byte("golang")
	go printData(&wg, byteData[:3])
	go printData(&wg, byteData[3:])
	wg.Wait()
}
