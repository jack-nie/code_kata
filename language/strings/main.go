package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// stringptr returns a pointer to the string data.
func stringptr(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

func main() {
	s1 := "1234"
	s2 := s1[:2]
	fmt.Println(stringptr(s1) == stringptr(s2))

	s1 = "12"
	s2 = "1" + "2"
	fmt.Println(stringptr(s1) == stringptr(s2)) // true

	// strings generated at runtime are not interned
	s1 = "12"
	s2 = strconv.Itoa(12)
	fmt.Println(stringptr(s1) == stringptr(s2)) // false
}
