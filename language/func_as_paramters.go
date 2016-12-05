package main

import "fmt"

type convert func(int) string

func value(x int) string {
	return fmt.Sprintf("%v", x)
}

func quote123(fn convert) string {
	return fmt.Sprintf("%q", fn(123))
}

func main() {
	var result string
	result = value(123)
	fmt.Println(result)

	result = quote123(value)
	fmt.Println(result)

	result = quote123(func(x int) string { return fmt.Sprintf("%b", x) })
	fmt.Println(result)

	foo := func(x int) string { return "foo" }
	result = quote123(foo)
	fmt.Println(result)
}
