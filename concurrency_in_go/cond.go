package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var tmpwg sync.WaitGroup
		tmpwg.Add(1)

		go func() {
			tmpwg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		tmpwg.Wait()
	}

	var wg sync.WaitGroup
	wg.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("exiting.....")
		wg.Done()
	})

	button.Clicked.Broadcast()

	wg.Wait()

}
