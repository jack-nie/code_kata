package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("sent messages: ", msg)
	default:
		fmt.Println("no messages received!")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent messages: ", msg)
	default:
		fmt.Println("no messages sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received messages: ", msg)
	case sig := <-signals:
		fmt.Println("received signal: ", sig)
	default:
		fmt.Println("no activity")
	}
}
