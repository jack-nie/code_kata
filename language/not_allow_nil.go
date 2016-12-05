package main

import "fmt"

func main() {
	// this will not compile
	// fmt.Println(getString())

	s := getSomeStruct()
	fmt.Println(s.name)
}

// func getString() string {
// 	 return nil
// }

type SomeStruct struct {
	name string
}

func getSomeStruct() *SomeStruct {
	return nil
}
