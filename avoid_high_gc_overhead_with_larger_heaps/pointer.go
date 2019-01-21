package main

import (
	"fmt"
	"runtime"
	"time"
)

// GC took 17.906192214s s
// GC took 10.01669487s s
// GC took 13.297150481s s
// GC took 10.164840762s s
// GC took 10.329762246s s
// GC took 9.791182076s s
// GC took 11.136856031s s
// GC took 10.971730415s s
// GC took 9.255500277s s
func main() {
	a := make([]*int, 1e9)

	for i := 0; i < 9; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
	}
	runtime.KeepAlive(a)
}
