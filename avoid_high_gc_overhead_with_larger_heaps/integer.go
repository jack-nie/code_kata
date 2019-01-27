package main

import (
	"fmt"
	"runtime"
	"time"
)

// GC took 323.951µs
// GC took 349.137µs
// GC took 287.04µs
// GC took 187.144µs
// GC took 210.253µs
// GC took 167.926µs
// GC took 148.807µs
// GC took 232.849µs
// GC took 470.34µs
func main() {
	a := make([]int, 1e9)

	for i := 0; i < 9; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}
