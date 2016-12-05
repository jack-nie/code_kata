package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	if isKaprekarNumber(arg) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isKaprekarNumber(arg string) bool {
	num, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	square := num * num
	str := fmt.Sprintf("%d", square)
	var length int = len(str)
	for i := 0; i < length; i++ {
		num1, _ := strconv.Atoi(str[0:i])
		num2, _ := strconv.Atoi(str[i:length])

		if num1+num2 == num {
			return true
		}
	}
	return false
}
