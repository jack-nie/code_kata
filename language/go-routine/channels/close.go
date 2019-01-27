package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Helo channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)
	close(stringStream)
	salutation, ok = <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)

}
