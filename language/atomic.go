package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var s int32

	flag := atomic.CompareAndSwapInt32(&s, 0, 100)
	fmt.Println(flag)
	fmt.Println(s)

	atomic.StoreInt32(&s, 12)
	fmt.Println(s)
	fmt.Println(atomic.LoadInt32(&s))
	var config atomic.Value
	type Map map[string]string
	config.Store(make(Map))
	m := config.Load()
	for k, v := range m.(Map) {
		fmt.Println(k, v)
	}
}
