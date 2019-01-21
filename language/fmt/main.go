package main

import "fmt"

func main() {
	var e interface{} = 2.7182
	fmt.Printf("e = %v (%T)\n", e, e) // e = 2.7182 (float64)
	type Point struct {
		X int
		Y int
	}
	p := &Point{1, 2}
	fmt.Printf("%v %+v %#v \n", p, p, p) // &{1 2} &{X:1 Y:2} &main.Point{X:1, Y:2}
}
