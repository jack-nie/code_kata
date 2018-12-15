package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strconv"
)

func countGoroutine() int {
	return runtime.NumGoroutine()
}

func getGoroutinesCountHandler(w http.ResponseWriter, r *http.Request) {
	count := countGoroutine()
	w.Write([]byte(strconv.Itoa(count)))
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func sumConcurrent(w http.ResponseWriter, r *http.Request) {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)
	go sum(s[:len(s)/2], c1)
	go sum(s[len(s)/2:], c2)
	x := <-c1
	fmt.Fprintf(w, strconv.Itoa(x))
}

func getStackTraceHandler(w http.ResponseWriter, r *http.Request) {
	stack := debug.Stack()
	w.Write(stack)
	pprof.Lookup("goroutine").WriteTo(w, 2)
}

func main() {
	http.HandleFunc("/_count", getGoroutinesCountHandler)
	http.HandleFunc("/sum", sumConcurrent)
	http.HandleFunc("/_stack", getStackTraceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
