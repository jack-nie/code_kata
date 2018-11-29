package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type LargeStruct struct {
	Data [1024]byte
}

var wg = sync.WaitGroup{}

func copy(ch chan *LargeStruct) {
	data := make([]*LargeStruct, 0)
	for i := 0; i < 100*100*100; i++ {
		data = append(data, <-ch)
	}
	wg.Done()
}

func main() {

	ch := make(chan *LargeStruct, 1)

	wg.Add(1)
	go copy(ch)
	PrintMemUsage()
	now := time.Now()
	for i := 0; i < 100*100*100; i++ {
		ch <- &LargeStruct{}
	}
	wg.Wait()
	fmt.Println(time.Since(now).Nanoseconds() / 1000000)
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
