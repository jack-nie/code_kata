package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var buffer bytes.Buffer

	buffer.WriteString("heello")
	fmt.Println(buffer.String())
	fmt.Println(strings.Repeat(buffer.String(), 2))
	fmt.Println(strings.Replace("ssssss", "s", "t", -1))
}
